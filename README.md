# Distributed Accounting System


Go client application for the DAS (Distributed Accounting System)

#### Building the Go client

Clone the repository, change the current directory to src and build the project using the Go compiler

```
cd src
go build das.go
```

#### Console application 

#Charge another walet

```
./das charge --wallet <WALLET_NAME> --amount <CHARGE_AMOUNT>
```
Here:
<WALLET_NAME> - name of the wallet you want to charge
<CHARGE_AMOUNT> - amount to charge

#Get balance

```
./das balance --my <MY_WALLET_NAME> --their <THEIR_WALLET_NAMES>
```
Here:
<MY_WALLET_NAME> - name of the wallet we are checking balance for
<THEIR_WALLET_NAMES> - one or more wallet names (comma separated). Provide those to get wallet-to-wallet balances

#Get history

```
./das history --wallet <WALLET_NAMES> --from <START_TIME> --to <END_TIME>
```
Here:
<WALLET_NAMES> - one or more wallet names (comma separated). Provide those to filter the output
<START_TIME>, <END_TIME> - time period to filter the output
