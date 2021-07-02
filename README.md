# DStation

`dstation` app is a Cosmos SDK based blockchain service for Dfinance chain.
The project is a stripped down [dnode](https://github.com/dfinance/dnode) which is moved to a newer SDK version v0.42.X.

Refer to [Cosmos SDK documentation](https://docs.cosmos.network) for base functionality details and usage tips.

The project has a few custom Dfinance modules:

- `x/oracle`
    Module is used to store asset direct and reverse exchange rates (*btcusdt* for example) and requires an external service as a data source.
  Dfinance maintains an external service which captures data from the [Binance](https://www.binance.com/en) stock.
  
- `x/staker`
    Module is used to bridge tokens from other chains to the Dfinance chain (the Ethereum chain for example) and requires an external backend alongside deployed smart contracts.
  Dfinance maintains an external service as well.
  This is a temporary solution as we're waiting for the [Gravity bridge](https://github.com/cosmos/gravity-bridge) release.
  
- `x/vm`
    Module provides smart contracts functionality using Libra Move language.
  
## CLI

The `dstation` binary combines server and client functionality.
CLI has a help build in for every command it has:

```bash
dstation -h

# Transaction commands
dstation tx -h

# Query commands
dstation q -h
```

To communicate with a specific chain you have to define its parameters:

```bash
dstation --chain-id dn-testnet-v3 --node https://rpc.demo1.dfinance.co:443 tx wallet1unv2d9vw2f3jdnwqnm5qfl2kkm8tz2zuw2c3xu 1xfi --from user1 -y
```

To avoid typing in `chain-id` and `node` credentials for every request, you can store them into local config (`${HOME}/.dstation/config/client.toml`):

```bash
dstation config chain-id dn-testnet-v3
dstation config node https://rpc.demo1.dfinance.co:443
```

## Events

Each module can emit events on transactions.

### `x/oracle` module

* New asset added

  Type: `oracle.add_asset`

  Attributes:
    - `asset_code` - new asset assetCode [string];

* Price updated for assetCode

  Type: `oracle.price`

  Attributes:
    - `asset_code` - assetCode [string];
    - `price` - updated price [int];
    - `received_at` - price received UNIX timestamp (in seconds) by oracles system [int];

### `x/vm` module

Depending on VM execution status, module emits multiple events per Tx with variadic number of attributes.

* VM execution status `keep` received

    * Contract status event ("keep" event)

      Type: `vm.contract_status`

      Attributes:
        - `status` - `keep` [string const];

    * VM events (number of events depends on execution status)

      Type: `vm.contract_events`

      Attributes:
        - `sender_address` - VM event sender address [`0x1` for stdlib / Bech32 string for account resource];
        - `source` - VM event source [`script` for script source / `{moduleAddress}::{moduleName}` for module source];
        - `type` - VM event type string representation in Move format [string];
        - `data` - HEX string VM event data representation [string];

* VM execution status `keep` received (failed with an error)

    1. "keep" event

       Type: `vm.contract_status`

       Attributes:
        - `status` - `keep` [string const];

    2. "error" event

       Type: `vm.contract_status`

       Attributes:
        - `status` - `error` [string const];
        - `major_status` - error majorStatus [uint];
        - `sub_status` - error subStatus [uint];
        - `message` - error message [string];

* VM execution status `discard` received

  Type: `vm.contract_status`

  Attributes:
    - `status` - `discard` [string const];

* VM execution status `discard` received (failed with an error)

  Type: `vm.contract_status`

  Attributes:
    - `status` - `discard` [string const];
    - `major_status` - error majorStatus [uint];
    - `sub_status` - error subStatus [uint];
    - `message` - error message [string];
 