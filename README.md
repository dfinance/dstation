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
 
 
 # User Migration v1.0.0
 
 ## Introduction
 
 We are happy to announce a new version of the Dfinance node - [dstation](https://github.com/dfinance/dstation). This version is based on the **Cosmos SDK v0.42.6**, and **Tendermint v0.34.11,** introduces fixes for the critical HDD space usage issue and opens the doors to continue support of the Dfinance blockchain: implementing new features and functionality. The changes are significant therefore validator nodes will be required to migrate to the new version - meaning all validators will need to be re-created and delegators will be required to redelegate. 
 
 ## Migration
 
 Step 1 - Navigate to your server with dfinance node and stop it:
 
 ```bash
 cd dfinance-bootstrap
 docker-compose down
 cd ../
 ```
 
 Step 2 - Clone the new version of [docker-compose](https://github.com/dfinance/bootstrap) from Github:
 
 ```bash
 git clone https://github.com/dfinance/bootstrap.git dstation-bootstrap
 cd dstation-bootstrap
 cp .env.mainnet .env
 docker-compose pull
 ```
 
 Step 3 - Launch node:
 
 ```bash
 docker-compose up -d # Up node in background
 docker-compose logs -f # Check logs
 ```
 
 Now you can remove *dfinance-bootstrap* folder and all blockchain data there.
 
 Step 4 - Recreate your validator / delegations.
 
 **Create validator:**
 
 ```bash
 cd dstation-bootstrap
 docker-compose exec node sh # Login into docker container
 ./dstation tendermint show-validator # Show validator key (copy it)
 exit # Exit from contaier
 ```
 
 Now install dstation from [Github](https://github.com/dfinance/dstation/releases/tag/v1.0.1) Releases.
 
 For Linux: 
 
 ```bash
 wget https://github.com/dfinance/dstation/releases/download/v1.0.1/dstation-v1.0.1-82f51b96214efa654b318484fbb218437cd2f773-linux-amd64
 mv dstation-v1.0.1-82f51b96214efa654b318484fbb218437cd2f773-linux-amd64 ./dstation
 sudo chmod +x ./dstation
 mv dstation /usr/local/bin/dstation
 ```
 
 For Mac OS (download dstation from [Github](https://github.com/dfinance/dstation/releases/tag/v1.0.1)), and run  following commands in terminal:
 
 ```bash
 mv ./dstation-v1.0.1-82f51b96214efa654b318484fbb218437cd2f773-darwin-10.12-amd64 ./dstation
 sudo chmod +x ./dstation
 mv ./dstation /usr/local/bin/
 ```
 
 For Windows:
 
 1. Download binary from [Github](https://github.com/dfinance/dstation/releases/tag/v1.0.1).
 2. Go to **"Program Files"** directory.
 3. Create there **"dn"** directory.
 4. Rename the downloaded file to **"dstation"** and put it into **"dn"** directory.
 
 Now run **"cmd"** and execute:
 
 ```
 setx path "%path%;%ProgramFiles%\dn"
 ```
 
 Now restart **"cmd"**.
 
 Configure dstation:
 
 ```bash
 dstation config chain-id dn-alpha-mainnet
 dstation config output text
 dstation config node https://rpc.dfinance.co:443
 dstation config keyring-backend file
 ```
 
 Add your keys to dstation:
 
 ```bash
 dstation keys add <name> --recover # Add with existing mnemonic
 dstation keys add <name> --recover --index <index> # Add with existing mnemonic and indexx
 dstation keys add <name> --ledger # Add with ledger
 
 # See list of added keys
 dstation keys list
 
 # Query your account balance
 dstation q bank balances <address> 
 ```
 
 Send transaction to create validator:
 
 ```bash
 dstation tx staking create-validator \
    --amount=100000000000000000000000xfi \
    --pubkey=<your pub key> \
    --moniker=<moniker> \
    --commission-rate="0.10" \
    --commission-max-rate="0.20" \
    --commission-max-change-rate="0.01" \
    --min-self-delegation="1000000000000000000000" \ 
    --from <account>
 ```
 
 Where:
 
 - `amount` - XFI amount to self-stake, at least 1 XFI.
 - `pubkey` - validator consensus public key received during `dstation tendermint show-validator` command.
 - `moniker` - name of your validator.
 - `commission-rate` - how much your validator is going to take a commission from received rewards/fees, currently 10% by default.
 - `commission-max-rate` - maximum that validator can take as commission, 20% by default.
 - `commission-max-change-rate` - how percent per day validator can change commission, currently 1% per day.
 - `from` - an account that is going to send transaction and will self-stake coins for your validator, also, you can use this account to manage your validator later.
 
 Almost all commands similar to previous dnode, so you can look at our previous [docs](https://docs.dfinance.co) (will be updated in the near time) and dstation help:
 
 ```bash
 dstation help
 ```
 
 *dstation-bootstrap/config/.dstation/priv_validator_key.json* - private key of validator, backup it and don't miss, as it's the only way to access your validator.
 
 ### Changes
 
 **Staking**
 
 In parallel with the development of the new version we also researched the Cosmos [Gravity Bridge](https://github.com/cosmos/gravity-bridge). Our main goal is to enable transfers of assets between the Ethereum and Dfinance chains. While preparations for the Gravity Bridge have been implemented, the Gravity Bridge itself is not yet production-ready, therefore for the time being we continue to use our [Staking Portal](https://stake.dfinance.co) until the stable release of the Gravity Bridge. 
 
 The [Staking Portal](https://stake.dfinance.co) can be used to stake your XFI tokens.  The new version also enables the transfer of XFI tokens between accounts within the Dfinance network. Note that only addresses connected through the Staking Portal will remain eligible for rewards. If you unstake your XFI tokens using Staking Portal without withdrawing, your XFI will remain within the Dfinance network and your account would be marked as disabled, however, you will not be eligible to receive any rewards. 
 
 **XFI Staking**
 
 The new version also contains inflation changes. The new inflation model continues to allow to stake XFI and earn rewards while the inflation and block rewards have been updated to float between 20% till 7% per year. The staking goal is 67% of the total XFI which would reduce the inflation reduces to 7% p.a. Lower amounts will increase the inflation amount up to 20% p.a.
 
 Please note that all inflation parameters are not final and might be subject to changes through governance.
 
 **LP Staking**
 
 LP staking is now disabled within the Dfinance node and will not be available for the time being. LP staking will be later re-enabled using an out-of-blockchain approach. 
 
 **IBC**
 
 Lastly, the new version implements the latest version of IBC as well, however remains disabled as IBC is not being production-ready and is currently not yet enabled in the Cosmos network. Once enabled, IBC will allow the transfer of assets between Cosmos-based chains.
 
 # User Migration v1.1.0
 
 We released a new version of dstation v1.1.0 that fixed an issue with node synchronization that happened during v1.0.0 usage.
 
 During the update, at 2021-07-19 10:00:00 UTC time the network stopped at block 211948 because of a successful upgrade proposal.
 
 Here is instruction for validators how to update their nodes:
 
 **First of all, backup your validator private key:**
 
 ```bash
 cd dstation-bootstrap
 cp ./config/.dstation/priv_validator_key.json ../v1.0.0_priv_validator_key.json
```

**Check if it copied correctly (don't share content of your private key with anyone!):**

```bash
cat ../v1.0.0_priv_validator_key.json
```

**Stop your node:**

```bash
docker-compose down -v
```

**Pull the latest version of bootstrap:**

```bash
git pull origin master
docker-compose pull
cp .env.mainnet .env

 # If you want to change moniker or open p2p port edit .env content 
nano .env
```

**Remove data folder:**

```bash
rm -rf ./data
```

**Start your validator:**

```bash
docker-compose up -d
```

**Install the latest dstation:**

```bash
wget https://github.com/dfinance/dstation/releases/download/v1.1.0/dstation-v1.1.0-571475329ddacd76e50e3428755db63e87130d79-linux-amd64
mv dstation-v1.1.0-571475329ddacd76e50e3428755db63e87130d79-linux-amd64 ./dstation
chmod +x ./dstation
mv ./dstation /usr/local/bin/
```

**Change chain-id:**

```bash
dstation config chain-id dn-alpha-mainnet-v1-1-0
```

Because of migration almost all validators got jailed, so network could be migrated without risks of stuck.
You need to unjail your validator, so it's become active.

**Send unjail transaction**

```bash
dstation tx slashing unjail --from <account> # Account you used during validator creation.
```