# Cross Subnet Virtual Machine (XSVM)

Cross Subnet Asset Transfers

## Avalanche Subnets and Custom VMs

Avalanche is a network composed of multiple sub-networks (called [subnets][Subnet]) that each contain any number of blockchains. Each blockchain is an instance of a [Virtual Machine (VM)](https://docs.avax.network/learn/platform-overview#virtual-machines), much like an object in an object-oriented language is an instance of a class. That is, the VM defines the behavior of the blockchain where it is instantiated. For example, [Coreth (EVM)][Coreth] is a VM that is instantiated by the [C-Chain]. Likewise, one could deploy another instance of the EVM as their own blockchain (to take this to its logical conclusion).

## AvalancheGo Compatibility

```
[v1.0.2] AvalancheGo@v1.9.6-v1.9.7
[v1.0.1] AvalancheGo@v1.9.5
[v1.0.0] AvalancheGo@v1.9.5
```

## Introduction

Just as [Coreth] powers the [C-Chain], XSVM can be used to power its own blockchain in an Avalanche [Subnet]. Instead of providing a place to execute Solidity smart contracts, however, XSVM enables asset transfers for assets originating on it's own chain or other XSVM chains on other subnets.

## How it Works

XSVM utilizes AvalancheGo's [teleporter] package to create and authenticate Subnet Messages.

### Transfer

If you want to send an asset to someone, you can use a `tx.Transfer` to send to any address.

### Export

If you want to send this chain's native asset to a different subnet, you can use a `tx.Export` to send to any address on a destination chain. You may also use a `tx.Export` to return the destination chain's native asset.

### Import

To receive assets from another chain's `tx.Export`, you must issue a `tx.Import`. Note that, similarly to a bridge, the security of the other chain's native asset is tied to the other chain. The security of all other assets on this chain are unrelated to the other chain.

### Fees

Currently there are no fees enforced in the XSVM.

### xsvm

#### Install

```bash
git clone https://github.com/ava-labs/xsvm.git;
cd xsvm;
go install -v ./cmd/xsvm;
```

#### Usage

```
Runs an XSVM plugin

Usage:
  xsvm [flags]
  xsvm [command]

Available Commands:
  account     Displays the state of the requested account
  chain       Manages XS chains
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  issue       Issues transactions
  version     Prints out the version

Flags:
  -h, --help   help for xsvm

Use "xsvm [command] --help" for more information about a command.
```

### [Golang SDK](https://github.com/ava-labs/xsvm/blob/master/client/client.go)

```golang
// Client defines xsvm client operations.
type Client interface {
  Network(
    ctx context.Context,
    options ...rpc.Option,
  ) (uint32, ids.ID, ids.ID, error)
  Genesis(
    ctx context.Context,
    options ...rpc.Option,
  ) (*genesis.Genesis, error)
  Nonce(
    ctx context.Context,
    address ids.ShortID,
    options ...rpc.Option,
  ) (uint64, error)
  Balance(
    ctx context.Context,
    address ids.ShortID,
    assetID ids.ID,
    options ...rpc.Option,
  ) (uint64, error)
  Loan(
    ctx context.Context,
    chainID ids.ID,
    options ...rpc.Option,
  ) (uint64, error)
  IssueTx(
    ctx context.Context,
    tx *tx.Tx,
    options ...rpc.Option,
  ) (ids.ID, error)
  LastAccepted(
    ctx context.Context,
    options ...rpc.Option,
  ) (ids.ID, *block.Stateless, error)
  Block(
    ctx context.Context,
    blkID ids.ID,
    options ...rpc.Option,
   (*block.Stateless, error)
  Message(
    ctx context.Context,
    txID ids.ID,
    options ...rpc.Option,
  ) (*teleporter.UnsignedMessage, []byte, error)
}
```

### Public Endpoints

#### xsvm.network

```
<<< POST
{
  "jsonrpc": "2.0",
  "method": "xsvm.network",
  "params":{},
  "id": 1
}
>>> {"networkID":<uint32>, "subnetID":<ID>, "chainID":<ID>}
```

For example:

```bash
curl --location --request POST 'http://34.235.54.228:9650/ext/bc/28iioW2fYMBnKv24VG5nw9ifY2PsFuwuhxhyzxZB5MmxDd3rnT' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "xsvm.network",
    "params":{},
    "id": 1
}'
```

> `{"jsonrpc":"2.0","result":{"networkID":1000000,"subnetID":"2gToFoYXURMQ6y4ZApFuRZN1HurGcDkwmtvkcMHNHcYarvsJN1","chainID":"28iioW2fYMBnKv24VG5nw9ifY2PsFuwuhxhyzxZB5MmxDd3rnT"},"id":1}`

#### xsvm.genesis

```
<<< POST
{
  "jsonrpc": "2.0",
  "method": "xsvm.genesis",
  "params":{},
  "id": 1
}
>>> {"genesis":<genesis file>}
```

#### xsvm.nonce

```
<<< POST
{
  "jsonrpc": "2.0",
  "method": "xsvm.nonce",
  "params":{
    "address":<cb58 encoded>
  },
  "id": 1
}
>>> {"nonce":<uint64>}
```

#### xsvm.balance

```
<<< POST
{
  "jsonrpc": "2.0",
  "method": "xsvm.balance",
  "params":{
    "address":<cb58 encoded>,
    "assetID":<cb58 encoded>
  },
  "id": 1
}
>>> {"balance":<uint64>}
```

#### xsvm.loan

```
<<< POST
{
  "jsonrpc": "2.0",
  "method": "xsvm.loan",
  "params":{
    "chainID":<cb58 encoded>
  },
  "id": 1
}
>>> {"amount":<uint64>}
```

#### xsvm.issueTx

```
<<< POST
{
  "jsonrpc": "2.0",
  "method": "xsvm.issueTx",
  "params":{
    "tx":<bytes>
  },
  "id": 1
}
>>> {"txID":<cb58 encoded>}
```

#### xsvm.lastAccepted

```
<<< POST
{
  "jsonrpc": "2.0",
  "method": "xsvm.lastAccepted",
  "params":{},
  "id": 1
}
>>> {"blockID":<cb58 encoded>, "block":<json>}
```

#### xsvm.block

```
<<< POST
{
  "jsonrpc": "2.0",
  "method": "xsvm.block",
  "params":{
    "blockID":<cb58 encoded>
  },
  "id": 1
}
>>> {"block":<json>}
```

#### xsvm.message

```
<<< POST
{
  "jsonrpc": "2.0",
  "method": "xsvm.message",
  "params":{
    "txID":<cb58 encoded>
  },
  "id": 1
}
>>> {"message":<json>, "signature":<bytes>}
```

## Running the VM

To build the VM, run `./scripts/build.sh`.

### Deploying Your Own Network

Anyone can deploy their own instance of the XSVM as a subnet on Avalanche. All you need to do is compile it, create a genesis, and send a few txs to the
P-Chain.

You can do this by following the [subnet tutorial] or by using the [subnet-cli].

[teleporter]: https://github.com/ava-labs/avalanchego/tree/master/vms/platformvm/teleporter
[subnet tutorial]: https://docs.avax.network/build/tutorials/platform/subnets/create-a-subnet
[subnet-cli]: https://github.com/ava-labs/subnet-cli
[Coreth]: https://github.com/ava-labs/coreth
[C-Chain]: https://docs.avax.network/learn/platform-overview/#contract-chain-c-chain
[Subnet]: https://docs.avax.network/learn/platform-overview/#subnets

# Cross Subnet Transaction Example

The following example shows how to interact with XSVM to send and receive native assets across subnets.

### Overview of Steps
 1. Create & deploy Subnet A
 2. Create  & deploy Subnet B
 3. Issue an **export** TX on Subnet A
 4. Issue an **import** TX on Subnet B
 5. Confirm TXs processed correctly

> **Note:**  This demo requires [avalanche-cl](https://github.com/ava-labs/avalanche-cli)i version > 1.0.5, [xsvm](https://github.com/ava-labs/xsvm) version > 1.0.2 and [avalanche-network-runner](https://github.com/ava-labs/avalanche-network-runner) v1.3.5. 

