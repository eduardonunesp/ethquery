package cmd

import (
	"fmt"

	"github.com/eduardonunesp/ethquery/config"
	"github.com/spf13/cobra"
)

var blockNumberCmd = &cobra.Command{
	Use:     "blocknumber",
	Aliases: []string{"n"},
	Short:   "Returns the number of most recent block",
	RunE: func(cmd *cobra.Command, args []string) error {
		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)

		if err != nil {
			return err
		}

		postRequest(
			currentConfiguration.URL,
			"eth_blockNumber",
			[]string{},
		)

		return nil
	},
}

var blockByNumberCmd = &cobra.Command{
	Use:     "blockbynumber <block number hex|block number decimal>",
	Aliases: []string{"bn"},
	Short:   "Returns information about a block by block number",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		blockNumber := args[0]
		var full string

		if len(args) < 2 {
			full = "true"
		} else {
			full = args[1]
		}

		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)
		if err != nil {
			return err
		}
		postRequest(
			currentConfiguration.URL,
			"eth_getBlockByNumber",
			[]string{hexify(blockNumber), full},
		)

		return nil
	},
}

var blockByHashCmd = &cobra.Command{
	Use:     "blockbyhash <block hash>",
	Aliases: []string{"bh"},
	Short:   "Returns the information about a transaction requested by transaction hash",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		blockHash := args[0]
		var full string

		if len(args) < 2 {
			full = "true"
		} else {
			full = args[1]
		}

		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)

		if err != nil {
			return err
		}

		postRequest(
			currentConfiguration.URL,
			"eth_getBlockByHash",
			[]string{blockHash, full},
		)

		return nil
	},
}

var codeCmd = &cobra.Command{
	Use:     "code <contract address>",
	Aliases: []string{"co"},
	Short:   "Returns code at a given address",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		contractAddress := args[0]
		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)

		if err != nil {
			return err
		}

		postRequest(
			currentConfiguration.URL,
			"eth_getCode",
			[]string{contractAddress},
		)

		return nil
	},
}

var tranactionByHashCmd = &cobra.Command{
	Use:     "transactionbyhash <transaction hash>",
	Aliases: []string{"th"},
	Short:   "Returns the information about a transaction requested by transaction hash",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		transactionHash := args[0]
		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)

		if err != nil {
			return err
		}

		postRequest(
			currentConfiguration.URL,
			"eth_getTransactionByHash",
			[]string{transactionHash},
		)

		return nil
	},
}

var transactionReceiptCmd = &cobra.Command{
	Use:     "transactionreceipt <transaction hash>",
	Aliases: []string{"tr"},
	Short:   "Returns the receipt of a transaction by transaction hash",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		transactionHash := args[0]
		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)

		if err != nil {
			return err
		}

		postRequest(
			currentConfiguration.URL,
			"eth_getTransactionReceipt",
			[]string{transactionHash},
		)

		return nil
	},
}

var accountsCmd = &cobra.Command{
	Use:     "accounts",
	Aliases: []string{"ac"},
	Short:   "Returns a list of addresses owned by client",
	Args:    cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)

		if err != nil {
			return err
		}

		postRequest(
			currentConfiguration.URL,
			"eth_accounts",
			[]string{},
		)

		return nil
	},
}

var netVersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"vr"},
	Short:   "Returns the current network id",
	Args:    cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)

		if err != nil {
			return err
		}

		postRequest(
			currentConfiguration.URL,
			"net_version",
			[]string{},
		)

		return nil
	},
}

var gasPriceCmd = &cobra.Command{
	Use:     "gasprice",
	Aliases: []string{"gp"},
	Short:   "Returns the current price per gas in wei",
	Args:    cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)

		if err != nil {
			return err
		}

		postRequest(
			currentConfiguration.URL,
			"eth_gasPrice",
			[]string{},
		)

		return nil
	},
}

var transactionCountCmd = &cobra.Command{
	Use:     "transactioncount <address>",
	Aliases: []string{"tc"},
	Short:   "Returns the number of transactions sent from an address",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		address := args[0]
		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)

		if err != nil {
			return err
		}

		postRequest(
			currentConfiguration.URL,
			"eth_getTransactionCount",
			[]string{address},
		)

		return nil
	},
}

var balanceCmd = &cobra.Command{
	Use:     "balance <address> <tag>",
	Aliases: []string{"ba"},
	Short:   "Returns the balance for the given address",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		address := args[0]
		tag := "latest"

		if len(args) == 2 {
			tag = args[1]
		}

		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)

		if err != nil {
			return err
		}

		postRequest(
			currentConfiguration.URL,
			"eth_getBalance",
			[]string{address, tag},
		)

		return nil
	},
}

var dataFlag string
var toAddressFlag string
var fromAddressFlag string
var callCmd = &cobra.Command{
	Use:     "call",
	Aliases: []string{"ca"},
	Short:   "Static call a function on contract",
	RunE: func(cmd *cobra.Command, args []string) error {
		configurationList := config.Load()
		currentConfiguration, err := configurationList.GetCurrent(configurationFlag)

		if err != nil {
			return err
		}

		var params string

		if len(fromAddressFlag) > 0 {
			params = fmt.Sprintf(`{"from": "%s", "to":"%s", "data": "%s"}`, fromAddressFlag, toAddressFlag, dataFlag)
		} else {
			params = fmt.Sprintf(`{"to":"%s", "data": "%s"}`, toAddressFlag, dataFlag)
		}

		postRequest(
			currentConfiguration.URL,
			"eth_call",
			[]string{params},
		)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(
		blockNumberCmd,
		blockByHashCmd,
		blockByNumberCmd,
		codeCmd,
		tranactionByHashCmd,
		transactionReceiptCmd,
		accountsCmd,
		netVersionCmd,
		gasPriceCmd,
		transactionCountCmd,
		balanceCmd,
		callCmd,
	)

	callCmd.Flags().StringVarP(&fromAddressFlag, "from", "f", "", "from address")
	callCmd.Flags().StringVarP(&toAddressFlag, "to", "t", "", "to address")
	callCmd.Flags().StringVarP(&dataFlag, "data", "d", "", "hex data representation")
}
