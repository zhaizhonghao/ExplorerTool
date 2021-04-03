# A Tool for configuring the Hyperledger explorer
## Usage

```shell
go build
```

This command will generate a executable file **explorerTool**

We need to move the executable file **explorerToool** into the `BlockchainExplorer` folder, then start the **explorerTool** service, by following command:

```
./explorerTool
```

## What can explorerTool service do 

Once the explorerTool service is started, it will listen on the port 8282 to wait for event from the [HyperledgerHelper-front-end](https://github.com/zhaizhonghao/HyperledgerHelper-front-end). Triggered by the event, the explorerTool service will  

* generate the connection profile(here named `first-network_2.2_bak.json`) in the [BlockchainExplorer/connection-profile](https://github.com/zhaizhonghao/BlockchainExplorer/tree/main/connection-profile) folder
* copy the `crypto-config` folder from the [BasicNetwork-2.0/artifacts/channel](https://github.com/zhaizhonghao/BasicNetwork-2.0/tree/main/artifacts/channel) into the [BlockchainExplorer](https://github.com/zhaizhonghao/BlockchainExplorer) folder
* start the explorer docker container in the background and listen on 8080 port

