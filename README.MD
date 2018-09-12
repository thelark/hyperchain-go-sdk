### Hyperchain Golang SDK

#### Import(引入)
```go
import "github.com/thelark/hyperchain-go-sdk/api"
```
#### Usage(使用)
```go
const (
    phone     = "<phone>"
    password  = "<password>"
    apiKey    = "<apiKey>"
    apiSecret = "<apiSecret>"
)
api := api.New(phone, password, apiKey, apiSecret)
api.GetApiToken()
fmt.Println(api.GetAccessToken())
```
#### API(接口)
##### Auth(授权)
###### 1. getApiToken(获取Token)
```go
resp, err := api.GetApiToken()
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
###### 2. refreshApiToken(刷新Token)
```go
resp, err := api.RefreshApiToken()
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
##### Account(账户)
###### 1. createAccount(新建区块链账户)
```go
resp, err := api.CreateAccount()
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
##### Block(区块)
###### 1. queryBlock(查询区块信息)
```go
resp, err := api.QueryBlock("number", "latest")
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
###### 2. queryBlocks(分页获取区块信息)
```go
resp, err := api.QueryBlocks(1, 3)
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
###### 3. queryBlocksByRange(查询某范围内区块信息)
```go
resp, err := api.QueryBlocksByRange(2, "4")
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
##### Transaction(交易)
###### 1. queryTransactionCount(获取交易总数)
```go
resp, err := api.QueryTransactionCount()
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
###### 2. queryTransactionByHash(通过hash查询某交易)
```go
resp, err := api.QueryTransactionByHash("0xed70377c261bfdc7dd7f4fc15c8961c145f9457186d6ff95f60907e9fb63d827")
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
###### 3. queryTransactionReceipt(查询交易回执)
```go
resp, err := api.QueryTransactionReceipt("0xed70377c261bfdc7dd7f4fc15c8961c145f9457186d6ff95f60907e9fb63d827")
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
###### 4. queryDiscardTransaction(查询无效交易)
```go
resp, err := api.QueryDiscardTransaction(1515140350903604865, 1515140350903604867)
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
##### Contract(合约)
###### 1. compileContract(编译合约)
```go
resp, err := api.CompileContract(`contract test{}`)
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
###### 2. deployContract&deployContractSync(部署合约)
```go
resp1, err := api.DeployContract("0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820b4c36b8b61723f302432d246407a061599017f8607ed26f1c053b5ecc63a54200029", "0x3713c3d01ae09cf32787c9c66c9c0781cf4b613d")
if err != nil {
    log.Println(err)
    return
}
jb1, _ := json.Marshal(resp1)
log.Println(string(jb1))
resp2, err := api.DeployContractSync("0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820b4c36b8b61723f302432d246407a061599017f8607ed26f1c053b5ecc63a54200029", "0x3713c3d01ae09cf32787c9c66c9c0781cf4b613d")
if err != nil {
    log.Println(err)
    return
}
jb2, _ := json.Marshal(resp2)
log.Println(string(jb2))
```
###### 3. getPayload(获取调用方法Payload)
```go
var(
    abi = "[{\"constant\":false,\"inputs\":[{\"name\":\"num1\",\"type\":\"uint32\"},{\"name\":\"num2\",\"type\":\"uint32\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getSum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"payable\":false,\"type\":\"function\"}]"
    args = []string{"1","2"}
    fn = "add"
)
resp, err := api.GetPayload(abi, fn, args)
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
###### 4. invokeContract&invokeContractSync(调用合约方法)
```go
var (
    _const =  false
    _from = "0x3713c3d01ae09cf32787c9c66c9c0781cf4b613d"
    _payload = "34141"
    _to = "0x8255340c2c4a1aec4010d2b6fdbb98727c65523d"
)
resp1, err := api.InvokeContract(_const, _from, _to, _payload)
if err != nil {
    log.Println(err)
    return
}
jb1, _ := json.Marshal(resp1)
log.Println(string(jb1))
resp2, err := api.InvokeContractSync(_const, _from, _to, _payload)
if err != nil {
    log.Println(err)
    return
}
jb2, _ := json.Marshal(resp2)
log.Println(string(jb2))    
```
###### 5. maintainContract(合约维护)
```go
var (
    _from = "0x19a170a0413096a9b18f2ca4066faa127f4d6f4a"
    _operation = 1
    _payload = "0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820b4c36b8b61723f302432d246407a061599017f8607ed26f1c053b5ecc63a54200029"
    _to = "0xd3a7bdd391f6aa13b28a72690e19d2ab3d845ac8"
)
resp, err := api.MaintainContract(_from, _to, _operation, _payload)
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
log.Println(string(jb))
```
###### 6. queryContractStatus(查询合约状态)
```go
resp, err := api.QueryContractStatus("0xd3a7bdd391f6aa13b28a72690e19d2ab3d845ac8")
if err != nil {
    log.Println(err)
    return
}
jb, _ := json.Marshal(resp)
```
##### cmd/main.go 仅供参考
