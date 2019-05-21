package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/eduardonunesp/ethquery/config"
	"github.com/spf13/cobra"
)

var blockNumberCmd = &cobra.Command{
	Use:   "blocknumber",
	Short: "Get current block height",
	Run: func(cmd *cobra.Command, args []string) {
		configurationList := config.Load()
		currentConfiguration := configurationList.GetCurrent()
		postRequest(currentConfiguration.URL, "eth_blockNumber", []string{})
	},
}

var blockByNumberCmd = &cobra.Command{
	Use:   "blockbynumber <block number hex|block number decimal>",
	Short: "Get current block by number",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		blockNumber := args[0]
		var full string

		if full == "" || len(args) < 1 {
			full = "true"
		}

		configurationList := config.Load()
		currentConfiguration := configurationList.GetCurrent()
		postRequest(
			currentConfiguration.URL,
			"eth_getBlockByNumber",
			[]string{hexify(blockNumber), full},
		)
	},
}

var blockByHashCmd = &cobra.Command{
	Use:   "blockbyhash <block hash>",
	Short: "Get current block by hash",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		blockHash := args[0]
		var full string

		if full == "" || len(args) < 1 {
			full = "true"
		}

		configurationList := config.Load()
		currentConfiguration := configurationList.GetCurrent()
		postRequest(
			currentConfiguration.URL,
			"eth_getBlockByHash",
			[]string{blockHash, full},
		)
	},
}

var codeCmd = &cobra.Command{
	Use:   "code <contract address>",
	Short: "Get hex code for a contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		contractAddress := args[0]
		configurationList := config.Load()
		currentConfiguration := configurationList.GetCurrent()
		postRequest(
			currentConfiguration.URL,
			"eth_getCode",
			[]string{contractAddress},
		)
	},
}

var tranactionByHashCmd = &cobra.Command{
	Use:   "transactionbyhash <transaction hash>",
	Short: "Get transaction information",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		transactionHash := args[0]
		configurationList := config.Load()
		currentConfiguration := configurationList.GetCurrent()
		postRequest(
			currentConfiguration.URL,
			"eth_getTransactionByHash",
			[]string{transactionHash},
		)
	},
}

var transactionReceipt = &cobra.Command{
	Use:   "transactionreceipt <transaction hash>",
	Short: "Get transaction receipt",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		transactionHash := args[0]
		configurationList := config.Load()
		currentConfiguration := configurationList.GetCurrent()
		postRequest(
			currentConfiguration.URL,
			"eth_getTransactionReceipt",
			[]string{transactionHash},
		)
	},
}

func init() {
	rootCmd.AddCommand(
		blockNumberCmd,
		blockByHashCmd,
		blockByNumberCmd,
		codeCmd,
		tranactionByHashCmd,
		transactionReceipt,
	)
}

func postRequest(url string, method string, params []string) {
	var paramsString []string

	for _, param := range params {
		var quoted string

		if param != "true" && param != "false" {
			quoted = fmt.Sprintf(`"%s"`, param)
		} else {
			quoted = param
		}

		paramsString = append(paramsString, quoted)
	}

	jsonStr := []byte(fmt.Sprintf(
		`{"id": 0, "jsonrpc":"2.0", "method": "%s", "params": [%s]}`,
		method, strings.Join(paramsString, ","),
	))

	// fmt.Println(string(jsonStr))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func hexify(numericalString string) string {
	if len(numericalString) >= 3 && numericalString[:2] == "0x" {
		return numericalString
	}

	u64, err := strconv.ParseUint(numericalString, 10, 32)
	if err != nil {
		panic(err)
	}

	if len(numericalString) < 3 {
		return fmt.Sprintf("0x%x", u64)
	}

	if numericalString[:2] != "0x" {
		return fmt.Sprintf("0x%x", u64)
	}

	return numericalString
}
