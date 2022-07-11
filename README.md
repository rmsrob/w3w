[![Go Reference](https://pkg.go.dev/badge/github.com/rrobrms/w3w.svg)](https://pkg.go.dev/github.com/rrobrms/w3w)
[![Go Report Card](https://goreportcard.com/badge/github.com/rrobrms/w3w)](https://goreportcard.com/report/github.com/rrobrms/w3w)
[![Coverage Status](https://coveralls.io/repos/github/rrobrms/w3w/badge.svg?branch=main)](https://coveralls.io/github/rrobrms/w3w?branch=main)
<!-- [![Latest Release](https://img.shields.io/github/v/release/rrobrms/w3w)](https://github.com/rrobrms/w3w/releases) -->

# W3W - eth wallet CLI

> Simple CLI to generate ether wallets from `go-ethereum`

## Install

#### Packages
Direct downloads are available through the [releases page](https://github.com/rrobrms/w3w/releases/latest).

#### GO
> `go install github.com/rrobrms/w3w@latest`

#### Manual

> You'll need to [install Go](https://golang.org/doc/install)
>```sh
> git clone https://github.com/rrobrms/w3w.git
> cd w3w
> go install
>```

## usage

> w3w 

```sh
W3W CLI is a tool that gives you a soom basic insight into the web3 ecosystem.

Usage:
  w3w [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  key         Get private key tools
  tx          Get transaction hash info

Flags:
      --config string   config file (default is $HOME/.w3w.yaml)
  -h, --help            help for w3w
  -t, --toggle          Help message for toggle

Use "w3w [command] --help" for more information about a command.
```

> Generate new wallet key pair
```sh
$ w3w key

$ ╭─────────────┬──────────────────────────────────────────────────────────────────╮
$ │ private key │ 23ed07206ce891e433a0193879caa12b6f50b69qbe374ecrd5fed1c7658a90b6 │
$ ├─────────────├──────────────────────────────────────────────────────────────────┤
$ │ address     │ 0xC3c7902A11f79184D5138B047FD9a993633b642d                       │
$ ╰─────────────┴──────────────────────────────────────────────────────────────────╯
$   valid   https://etherscan.io/address/0xC3c7902A11f79184D5138B047FD9a993633b642d
```

> Derive address from private key

```sh
$ w3w key --priv="23ed07206ce891e433a0193879caa12b6f50b69qbe374ecrd5fed1c7658a90b6"
$ Derivated address: 0xC3c7902A11f79184D5138B047FD9a993633b642d
```

> w3w tx work best with etherscan api key/ infura key **wip**

```sh
# ommit the flag `--chainid="rinkeby"` for mainnet lookup
$ w3w tx --chainid="rinkeby" 0x5c18e23444ba147a08ba2b07061c4224896805a6a9b609003da25760cfe4f0d6

tx isPending: false
Status: 1
Contract Implementation of: ConcaveNFTMarketplace
BlockNumber: 10863431 | Chain ID: 4
From: 0x8522093305253EfB2685241dc0C587CDD9B10e4B
To: 0x6Cc7F744f720e01834F0b497519e22e921FDF161 (is a contract: true)
Value: 0
Transaction Fee: 0.000207512677167615 Ether
Gas Price: 1.500073569e-09 Gwei
Gas Limit/Gas Used: 138335 | 136686
Type: 2 | Nonce: 360 | tx index: 21
Transaction Data in hex: d15758dd00000000000000000000000000000000000000000000000000000000000000970000000000000000000000004a8b871784a8e6344126f47d48283a87ea987f270000000000000000000000000000000000000000000000008ac7230489e80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000010000000000000000000000008522093305253efb2685241dc0c587cdd9b10e4b00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000002710
map[_buyNowPrice:10000000000000000000 _erc20Token:0x4A8b871784A8e6344126F47d48283a87Ea987f27 _feePercentages:[10000] _feeRecipients:[0x8522093305253EfB2685241dc0C587CDD9B10e4B] _tokenId:151 _whitelistedBuyer:0x0000000000000000000000000000000000000000]
Method Name: createSale
Method inputs: map[_buyNowPrice:10000000000000000000 _erc20Token:0x4A8b871784A8e6344126F47d48283a87Ea987f27 _feePercentages:[10000] _feeRecipients:[0x8522093305253EfB2685241dc0C587CDD9B10e4B] _tokenId:151 _whitelistedBuyer:0x0000000000000000000000000000000000000000]
Event Name: SaleCreated
Log Data in Hex: 000000000000000000000000b9e431fc34152246bb28453b6ce117829e8a5b0c00000000000000000000000000000000000000000000000000000000000000970000000000000000000000008522093305253efb2685241dc0c587cdd9b10e4b0000000000000000000000004a8b871784a8e6344126f47d48283a87ea987f270000000000000000000000000000000000000000000000008ac7230489e8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000010000000000000000000000008522093305253efb2685241dc0c587cdd9b10e4b00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000002710
Event outputs: map[buyNowPrice:10000000000000000000 erc20Token:0x4A8b871784A8e6344126F47d48283a87Ea987f27 feePercentages:[10000] feeRecipients:[0x8522093305253EfB2685241dc0C587CDD9B10e4B] nftContractAddress:0xB9E431Fc34152246BB28453b6ce117829E8A5B0C nftSeller:0x8522093305253EfB2685241dc0C587CDD9B10e4B tokenId:151 whitelistedBuyer:0x0000000000000000000000000000000000000000]

```