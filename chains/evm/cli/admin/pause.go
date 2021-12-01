package admin

import (
	"fmt"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/bridge"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/transactor"
	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/contracts"

	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/flags"
	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause deposits and proposals",
	Long:  "Pause deposits and proposals",
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.LoggerMetadata(cmd.Name(), cmd.Flags())
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		bridgeContract, err := contracts.InitializeBridgeContract(
			url, gasLimit, gasPrice, senderKeyPair, bridgeAddr,
		)
		if err != nil {
			return err
		}
		return PauseCmd(cmd, args, bridgeContract)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		err := ValidatePauseCmdFlags(cmd, args)
		if err != nil {
			return err
		}

		ProcessPauseCmdFlags(cmd, args)

		return nil
	},
}

func BindPauseCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&Bridge, "bridge", "", "bridge contract address")
	flags.MarkFlagsAsRequired(cmd, "bridge")
}

func init() {
	BindPauseCmdFlags(pauseCmd)
}

func ValidatePauseCmdFlags(cmd *cobra.Command, args []string) error {
	if !common.IsHexAddress(Bridge) {
		return fmt.Errorf("invalid bridge address: %s", Bridge)
	}
	return nil
}

func ProcessPauseCmdFlags(cmd *cobra.Command, args []string) {
	bridgeAddr = common.HexToAddress(Bridge)
}

func PauseCmd(cmd *cobra.Command, args []string, contract *bridge.BridgeContract) error {
	hash, err := contract.Pause(transactor.TransactOptions{})
	if err != nil {
		log.Error().Err(fmt.Errorf("admin pause error: %v", err))
		return err
	}

	log.Info().Msgf("successfully paused bridge: %s; tx hash: %s", Bridge, hash.Hex())
	return nil
}
