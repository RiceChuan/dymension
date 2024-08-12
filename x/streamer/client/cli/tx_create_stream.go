package cli

import (
	"errors"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/spf13/cobra"

	"github.com/dymensionxyz/dymension/v3/utils"
	"github.com/dymensionxyz/dymension/v3/x/streamer/types"
)

const FlagSponsored = "sponsored"

// NewCmdSubmitCreateStreamProposal broadcasts a CreateStream message.
func NewCmdSubmitCreateStreamProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-stream-proposal gaugeIds weights reward [flags]",
		Short: "proposal to create a stream of incentives rewards over a period of time",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposal, deposit, err := utils.ParseProposal(cmd)
			if err != nil {
				return err
			}

			sponsoredInt, err := cmd.Flags().GetCount(FlagSponsored)
			if err != nil {
				return err
			}
			sponsored := sponsoredInt > 0

			var records []types.DistrRecord
			// Ignore records is the stream is sponsored
			if !sponsored {
				records, err = parseRecords(args[0], args[1])
				if err != nil {
					return err
				}
			}

			coins, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			var startTime time.Time
			timeStr, err := cmd.Flags().GetString(FlagStartTime)
			if err != nil {
				return err
			}
			if timeStr == "" { // empty start time
				startTime = time.Unix(0, 0)
			} else if timeUnix, err := strconv.ParseInt(timeStr, 10, 64); err == nil { // unix time
				startTime = time.Unix(timeUnix, 0)
			} else if timeRFC, err := time.Parse(time.RFC3339, timeStr); err == nil { // RFC time
				startTime = timeRFC
			} else { // invalid input
				return errors.New("invalid start time format")
			}

			epochIdentifier, err := cmd.Flags().GetString(FlagEpochIdentifier)
			if err != nil {
				return err
			}

			epochs, err := cmd.Flags().GetUint64(FlagEpochs)
			if err != nil {
				return err
			}

			content := types.NewCreateStreamProposal(proposal.Title, proposal.Description, coins, records, startTime, epochIdentifier, epochs, sponsored)
			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			txfCli, err := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			txf := txfCli.WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)
			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "The proposal title")
	cmd.Flags().String(govcli.FlagDescription, "", "The proposal description")
	cmd.Flags().String(govcli.FlagDeposit, "", "The proposal deposit")
	cmd.Flags().Count(FlagSponsored, "The stream is based on the sponsorship distribution")

	cmd.Flags().AddFlagSet(FlagSetCreateStream())
	return cmd
}
