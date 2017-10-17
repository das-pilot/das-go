# Distributed Accounting System


Go client application for the DAS (Distributed Accounting System)

#### Building the Go client

Clone the repository, change the current directory to src and build the project using the Go compiler

```
cd src
go build das.go
```

#### Console application 

1. Charge another walet

```
./das charge --wallet <WALLET_NAME> --amount <CHARGE_AMOUNT>
```
Here:

<WALLET_NAME> - name of the wallet you want to charge
<CHARGE_AMOUNT> - amount to charge

2. Get balance

```
./das balance --my <MY_WALLET_NAME> --their <THEIR_WALLET_NAMES>
```
Here:

<MY_WALLET_NAME> - name of the wallet we are checking balance for
<THEIR_WALLET_NAMES> - one or more wallet names (comma separated). Provide those to get wallet-to-wallet balances

3. Get history

```
./das history --wallet <WALLET_NAMES> --from <START_TIME> --to <END_TIME>
```
Here:

<WALLET_NAMES> - one or more wallet names (comma separated). Provide those to filter the output
<START_TIME>, <END_TIME> - time period to filter the output

#### Start REST server

You can start the REST server from the command line:

```
./das server --port <PORT_NUMBER>
```

Now you should be able to run the same commands as above using your favorite browser or simply a curl client

1. Charge another wallet

2. Get balance

3. Get history


