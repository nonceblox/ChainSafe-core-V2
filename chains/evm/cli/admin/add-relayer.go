package admin

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var addRelayerCmd = &cobra.Command{
	Use:   "add-relayer",
	Short: "Add a new relayer",
	Long:  "Add a new relayer",
	Run:   addRelayer,
}

func init() {
	addRelayerCmd.Flags().String("relayer", "", "address to add")
	addRelayerCmd.Flags().String("bridge", "", "bridge contract address")
}

func addRelayer(cmd *cobra.Command, args []string) {
	relayerAddress := cmd.Flag("relayer").Value
	bridgeAddress := cmd.Flag("bridge").Value
	log.Debug().Msgf(`
Adding relayer 
Relayer address: %s
Bridge address: %s`, relayerAddress, bridgeAddress)
}

/*
func addRelayer(cctx *cli.Context) error {
	url := cctx.String("url")
	gasLimit := cctx.Uint64("gasLimit")
	gasPrice := cctx.Uint64("gasPrice")
	sender, err := cliutils.DefineSender(cctx)
	if err != nil {
		return err
	}
	bridgeAddress, err := cliutils.DefineBridgeAddress(cctx)
	if err != nil {
		return err
	}
	relayer := cctx.String("relayer")
	if !common.IsHexAddress(relayer) {
		return fmt.Errorf("invalid relayer address %s", relayer)
	}
	relayerAddress := common.HexToAddress(relayer)
	ethClient, err := client.NewClient(url, false, sender, big.NewInt(0).SetUint64(gasLimit), big.NewInt(0).SetUint64(gasPrice), big.NewFloat(1))
	if err != nil {
		return err
	}
	err = utils.AdminAddRelyaer(ethClient, bridgeAddress, relayerAddress)
	if err != nil {
		return err
	}
	log.Info().Msgf("Address %s is relayer now", relayerAddress.String())
	return nil
}
*/