package ibctesting_test

import (
	"testing"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
	"github.com/cosmos/ibc-go/v7/testing/simapp"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
	"github.com/stretchr/testify/suite"

	"github.com/dymensionxyz/dymension/v3/app/apptesting"
)

type transfersEnabledSuite struct {
	utilSuite

	path *ibctesting.Path
}

func TestTransfersEnabledTestSuite(t *testing.T) {
	suite.Run(t, new(transfersEnabledSuite))
}

func (s *transfersEnabledSuite) SetupTest() {
	s.utilSuite.SetupTest()
	path := s.newTransferPath(s.hubChain(), s.rollappChain())
	s.coordinator.Setup(path)
	s.createRollapp(false, nil)
	s.registerSequencer()
	s.path = path

	// manually set the rollapp to have transfers disabled by default
	// (rollapp is setup correctly, meaning transfer channel is canonical)
	ra := s.hubApp().RollappKeeper.MustGetRollapp(s.hubCtx(), rollappChainID())
	ra.GenesisState.TransferProofHeight = 0
	ra.ChannelId = s.path.EndpointA.ChannelID
	s.hubApp().RollappKeeper.SetRollapp(s.hubCtx(), ra)
	s.hubApp().LightClientKeeper.SetCanonicalClient(s.hubCtx(), rollappChainID(), s.path.EndpointA.ClientID)
}

// Regular (non genesis) transfers (RA->Hub) and Hub->RA should both be blocked when the bridge is not open
func (s *transfersEnabledSuite) TestHubToRollappDisabled() {
	amt := math.NewIntFromUint64(10000000000000000000)
	denom := "foo"
	tokens := sdk.NewCoin(denom, amt)

	timeoutHeight := clienttypes.NewHeight(100, 110)

	msg := types.NewMsgTransfer(
		s.path.EndpointA.ChannelConfig.PortID,
		s.path.EndpointA.ChannelID,
		tokens,
		s.hubChain().SenderAccount.GetAddress().String(),
		s.rollappChain().SenderAccount.GetAddress().String(),
		timeoutHeight,
		0,
		"",
	)

	shouldFail := true

	for i := range 2 {

		apptesting.FundAccount(s.hubApp(), s.hubCtx(), s.hubChain().SenderAccount.GetAddress(), sdk.Coins{msg.Token})

		_, _, err := simapp.SignAndDeliver(
			s.hubChain().T,
			s.hubChain().TxConfig,
			s.hubApp().GetBaseApp(),
			s.hubCtx().BlockHeader(),
			[]sdk.Msg{msg},
			hubChainID(),
			[]uint64{s.hubChain().SenderAccount.GetAccountNumber()},
			// TODO: not sure why, but the simapp doesn't seem to properly update the sequence after a failed tx
			[]uint64{s.hubChain().SenderAccount.GetSequence() + uint64(i)},
			true,
			!shouldFail,
			s.hubChain().SenderPrivKey,
		)

		if shouldFail {
			shouldFail = false
			s.Require().True(errorsmod.IsOf(err, gerrc.ErrFailedPrecondition))
			ra := s.hubApp().RollappKeeper.MustGetRollapp(s.hubCtx(), rollappChainID())
			ra.ChannelId = s.path.EndpointA.ChannelID
			ra.GenesisState.TransferProofHeight = 1 // enable
			s.hubApp().RollappKeeper.SetRollapp(s.hubCtx(), ra)
		} else {
			s.Require().NoError(err)
		}
	}
}
