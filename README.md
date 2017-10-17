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
./das server --port 8089
DAS Server started on port 8089
```

Now you should be able to run the same commands as above using your favorite browser or simply a curl client

1. Charge another wallet

```
curl -d "wallet=1F7B169C846F218A&amount=23.45" -X POST http://localhost:8089/das/charge
{"Result":"OK","Message":"Charged"}
```

2. Get balance

```
curl -X GET http://localhost:8089/das/balance?wallet=1F7B169C846F218A
{"Result":"OK","Message":"","Wallet":"1F7B169C846F218A","Balance":100.50}
```

3. Get history

```
curl -X GET http://localhost:8089/das/history?wallet=1F7B169C846F218A
[{"Wallet":"1F7B169C846F218A","Amount":-43.39317,"Message":"SPENT","Time":"09.12.2016 09:12:59"},
{"Wallet":"1F7B169C846F218A","Amount":95.048325,"Message":"INCOME","Time":"17.12.2016 17:58:11"},
{"Wallet":"1F7B169C846F218A","Amount":78.872345,"Message":"INCOME","Time":"18.12.2016 07:20:56"},
{"Wallet":"1F7B169C846F218A","Amount":-81.83254,"Message":"SPENT","Time":"13.01.2017 13:32:03"},
{"Wallet":"1F7B169C846F218A","Amount":-82.8959,"Message":"SPENT","Time":"13.01.2017 13:53:25"},
{"Wallet":"1F7B169C846F218A","Amount":-99.89724,"Message":"SPENT","Time":"14.01.2017 02:26:35"},
{"Wallet":"1F7B169C846F218A","Amount":-50.06136,"Message":"SPENT","Time":"14.01.2017 14:40:55"},
{"Wallet":"1F7B169C846F218A","Amount":-0.32113647,"Message":"SPENT","Time":"15.01.2017 15:06:43"},
{"Wallet":"1F7B169C846F218A","Amount":47.965164,"Message":"INCOME","Time":"17.01.2017 17:57:56"},
{"Wallet":"1F7B169C846F218A","Amount":43.27368,"Message":"INCOME","Time":"19.01.2017 19:25:32"}]

```

