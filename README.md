# ethquery

`Ethquery` is a command to make easy the task of request and debug `Ethereum JSON RPC` commands on command line, it supports multiple configurations and JSON RPC commands without the needing of long JSON RPC strings using `CURL` command

## Comparing CURL vs Ethquery

```bash
# Curl way get block by number
curl -X POST --data '{"id":0,"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0",true]}' http://localhost:8545

# Curl way get block number
curl -X POST --data '{"id":0,"jsonrpc":"2.0","method":"eth_blockNumber","params":[]}' http://localhost:8545
```

```bash
# Ethquery way get block by number
ethquery blockbynumber 0x0

# Or you can use decimal values here
ethquery blockbynumber 47 # auto converted to 0x2f

# Ethquery way get block number
ethquery blocknumber
```

## Install

```bash
go get -u github.com/eduardonunesp/ethquery
```

## Configuring

The first use requires to add at least one configuration which will save the configurations on home directory using the path `~/.ethquery/config`

```bash
# Adding ganache and loom configurations
ethquery configuration new ganache http://localhost:8545
ethquery configuration new loom http://localhost:46658
```

```bash
# Listing available configurations
ethquery configuration list

+---------+------------------------+---------+
|  NAME   |          URL           | CURRENT |
+---------+------------------------+---------+
| ganache | http://localhost:8545  |         |
| loom    | http://localhost:46658 | *       |
+---------+------------------------+---------+
```

> The CURRENT means the configuration used at the moment

```bash
# Switching current configuration
ethquery configuration current ganache
```

## Commands

```bash
# List commands available by type ethquery [Enter]

ethquery
Usage:
  ethquery [command]

Available Commands:
  blockbyhash        Get current block by hash
  blockbynumber      Get current block by number
  blocknumber        Get current block height
  code               Get hex code for a contract
  configuration      Manage configurations
  help               Help about any command
  transactionbyhash  Get transaction information
  transactionreceipt Get transaction receipt
```

You can combine commands with the tool [jq](https://github.com/stedolan/jq) which is a great command line JSON processor

```bash
# Command get block by hash with jq
ethquery blockbynumber 0x0 | jq .result.hash
"0x7b86a1ff2c4aa712bbfb51c07be05ab82a1a0d9aeaeab3b1d02c3695708a56af"
```

## Commands available

- [x] eth_accounts
- [x] eth_blockNumber
- [ ] eth_call
- [ ] eth_estimateGas
- [x] eth_gasPrice
- [x] eth_getBlockByHash
- [x] eth_getBlockByNumber
- [x] eth_getCode
- [ ] eth_getFilterChanges
- [ ] eth_getLogs
- [x] eth_getTransactionByHash
- [x] eth_getTransactionReceipt
- [ ] eth_newBlockFilter
- [ ] eth_newFilter
- [ ] eth_newPendingTransactionFilter
- [ ] eth_sendTransaction
- [ ] eth_subscribe
- [ ] eth_uninstallFilter
- [ ] eth_unsubscribe
- [x] net_version
- [ ] net_peerCount (Returns the current ethereum protocol version)
- [ ] eth_syncing (Returns an object with data about the sync status or false)
- [ ] eth_coinbase (Returns the client coinbase address)
- [ ] web3_clientVersion (Returns the current client version)
- [x] eth_getTransactionCount (Returns the number of transactions sent from an address)
- [ ] eth_getBlockTransactionCountByHash (Returns the number of transactions in a block from a block matching the given block hash)
- [ ] eth_getBlockTransactionCountByNumber (Returns the number of transactions in a block matching the given block number)

## LICENSE
Copyright (c) 2015, Eduardo Nunes Pereira
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of sslb nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.