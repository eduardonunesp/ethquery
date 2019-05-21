# ethquery

Ethquery is a command to make easy the task of request and debug Ethereum JSON RPC commands on command line, it supports multiple configurations and JSON RPC commands without the needing of long JSON RPC strings using CURL command

# Install

```bash
go get -u github.com/eduardonunesp/ethquery
```

# Configuring

The first use requires to add at least one configuration which will save the configurations on home directory using the path `~/.ethquery/config`

```bash
# Adding ganache configuration
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

> The CURRENT means the the configuration used at the moment

```bash
# Switching current configuration
ethquery configuration current ganache
```

# Commands

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