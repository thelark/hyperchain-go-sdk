### Hyperchain Golang CLI

#### Usage(使用)
##### Auth(授权)
###### 1. 获取指定用户的API接入授权码
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> getApiToken
```
###### 2. 刷新API授权码
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> refreshApiToken
```
##### Account(账户)
###### 1. 新建区块链账户,返回区块链地址
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> createAccount
```
##### Block(区块)
###### 1. 查询指定区块的详细信息
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> queryBlock --type number --value latest
```
###### 2. 区块信息分页查询
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> queryBlocks --index 1 --pageSize 2
```
###### 3. 查询指定高度区间的区块列表
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> queryBlocksByRange --from 1 --to 2
```
##### Transaction(交易)
###### 1. 查询联盟链上的交易总数
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> queryTransactionCount
```
###### 2. 查询指定交易hash的交易详情
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> queryTransactionByHash --hash <hash>
```
###### 3. 查询联盟链上的交易总数
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> queryTransactionReceipt --txhash <txhash>
```
###### 4. 查询指定时间区间内的非法交易
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> queryDiscardTransaction --start 1539330293000000 --end 1539330294000000
```
##### Contract(合约)
###### 1. 编译智能合约
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> compileContract --code 'contract test{}'
```
###### 2. 部署智能合约
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> deployContract --bin '0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a723058207799045e48fe5a1bd53859a58b9f2b52388e4cb227e7792eb4b894b366bd0bd40029' --address '0x3be60875d005800671e5fbfda15b0f49f1727494'
```
###### 3. 同步部署智能合约,直接返回部署结果,包含合约地址
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> deployContractSync --bin '0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a723058207799045e48fe5a1bd53859a58b9f2b52388e4cb227e7792eb4b894b366bd0bd40029' --address '0x3be60875d005800671e5fbfda15b0f49f1727494'
```
###### 4. 调用合约需要转入合约方法与合约参数编码后的input字节码,该接口可为用户返回Payload
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> getPayload --abi '[]' --func '<func>' --args '<arg1>' --args '<arg2>' --args '<arg3>'
```
###### 5. 调用合约
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> invokeContract --const --from '<from>' --to '<to>' --payload '<payload>'
```
###### 6. 同步调用合约
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> invokeContractSync --const --from '<from>' --to '<to>' --payload '<payload>'
```
###### 7. 合约维护,包括合约的升级,冻结和解冻
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> maintainContract --from '<from>' --to '<to>' --operate 1 --payload '<payload>'
```
###### 8. 查询合约当前状态
```shell
hyperchain-cli --phone <phone> --password <password> --apiKey <apiKey> --apiSecret <apiSecret> queryContractStatus --address '<address>'
```