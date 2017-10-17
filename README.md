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

1) {Wallet:1F7B169C846F218A Amount:-43.39317 Message:SPENT Time:06.08.2015 06:00:31}
2) {Wallet:52FA82FBF86758BF Amount:95.048325 Message:INCOME Time:16.01.1985 16:43:09}
3) {Wallet:97D2D2A313E4F959 Amount:78.872345 Message:INCOME Time:21.08.1981 21:07:17}
4) {Wallet:818A7B3EDCA492F2 Amount:-81.83254 Message:SPENT Time:02.07.1976 02:05:43}
5) {Wallet:A67697C4F91D9B93 Amount:-82.8959 Message:SPENT Time:07.10.1979 07:48:38}
6) {Wallet:E8234783DE17BD7A Amount:-99.89724 Message:SPENT Time:22.10.1986 22:19:03}
7) {Wallet:E0A9F6813976EADF Amount:-50.06136 Message:SPENT Time:21.08.1986 21:16:04}
8) {Wallet:DEB5475EB5820F83 Amount:-0.32113647 Message:SPENT Time:09.03.1981 09:11:32}
9) {Wallet:0FCABC87CC1F1A22 Amount:47.965164 Message:INCOME Time:07.10.1998 07:35:46}
10) {Wallet:AAE7E0F0EE788A1F Amount:43.27368 Message:INCOME Time:03.05.2013 03:58:48}
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

