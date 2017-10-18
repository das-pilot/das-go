# Distributed Accounting System


Go client application for the DAS (Distributed Accounting System)

#### Building the Go client

Clone the repository, change the current directory to src and build the project using the Go compiler

```
go build das.go
```

#### Console application 

1. Charge another walet

```
./das charge --wallet <WALLET_NAME> --amount <CHARGE_AMOUNT>


./das charge --wallet 1F7B169C846F218A --amount 23.45

{Result:OK Message:Charged}
```
Here:

<WALLET_NAME> - name of the wallet you want to charge
<CHARGE_AMOUNT> - amount to charge

2. Get balance

```
./das balance --wallet <WALLET_NAMES>


./das balance --wallet 1F7B169C846F218A

{Result:OK Message: Wallet:1F7B169C846F218A Balance:100.50}
```
Here:

<WALLET_NAMES> - one or more wallet names (comma separated). Provide those to get wallet-to-wallet balances

3. Get history

```
./das history --wallet <WALLET_NAMES> --from <START_TIME> --to <END_TIME>


./das history --wallet 1F7B169C846F218A

1) {Wallet:1F7B169C846F218A Amount:-43.39 Message:SPENT  Time:16.08.2017 06:00:31}
2) {Wallet:1F7B169C846F218A Amount: 95.04 Message:INCOME Time:16.08.2017 16:43:09}
3) {Wallet:1F7B169C846F218A Amount: 78.87 Message:INCOME Time:16.08.2017 21:07:17}
4) {Wallet:1F7B169C846F218A Amount:-81.83 Message:SPENT  Time:17.08.2017 02:05:43}
5) {Wallet:1F7B169C846F218A Amount:-82.89 Message:SPENT  Time:17.08.2017 07:48:38}
6) {Wallet:1F7B169C846F218A Amount:-99.89 Message:SPENT  Time:17.08.2017 22:19:03}
7) {Wallet:1F7B169C846F218A Amount:-50.06 Message:SPENT  Time:17.08.2017 23:16:04}
8) {Wallet:1F7B169C846F218A Amount: -0.32 Message:SPENT  Time:18.08.2017 09:11:32}
9) {Wallet:1F7B169C846F218A Amount: 47.96 Message:INCOME Time:18.08.2017 09:35:46}
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

Now you should be able to run the same commands as above using your favorite browser or simply a Curl client

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

[{Wallet:1F7B169C846F218A Amount:-43.39 Message:SPENT  Time:16.08.2017 06:00:31},
 {Wallet:1F7B169C846F218A Amount: 95.04 Message:INCOME Time:16.08.2017 16:43:09},
 {Wallet:1F7B169C846F218A Amount: 78.87 Message:INCOME Time:16.08.2017 21:07:17},
 {Wallet:1F7B169C846F218A Amount:-81.83 Message:SPENT  Time:17.08.2017 02:05:43},
 {Wallet:1F7B169C846F218A Amount:-82.89 Message:SPENT  Time:17.08.2017 07:48:38},
 {Wallet:1F7B169C846F218A Amount:-99.89 Message:SPENT  Time:17.08.2017 22:19:03},
 {Wallet:1F7B169C846F218A Amount:-50.06 Message:SPENT  Time:17.08.2017 23:16:04},
 {Wallet:1F7B169C846F218A Amount: -0.32 Message:SPENT  Time:18.08.2017 09:11:32},
 {Wallet:1F7B169C846F218A Amount: 47.96 Message:INCOME Time:18.08.2017 09:35:46}]
```

