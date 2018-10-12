package api

import "testing"

const (
	phone     = "<phone>"
	password  = "<password>"
	apiKey    = "<apiKey>"
	apiSecret = "<apiSecret>"
)
const (
	contractCode = `` // 合约源码
	contractBin  = "" // 合约BIN
	contractAbi  = "" // 合约ABI
	address      = "" // 合约地址
)
const (
	from      = ""
	to        = ""
	payload   = ""
	operation = 1
)
const (
	transactionHash   = ""                  // 交易Hash
	transactionTxHash = ""                  // 交易Hash
	transactionStart  = 1504250000000000000 // 交易查询开始时间
	transactionEnd    = 1504368000000000000 // 交易查询结束时间
)

var (
	fn   = ""
	args = []string{""}
)

func TestNew(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	if api.phone == "" || api.password == "" || api.apiKey == "" || api.apiSecret == "" {
		t.Errorf("The corresponding parameter is not introduced.")
	}
}
func TestApi_GetApiToken(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	apiToken, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	if apiToken.AccessToken == "" || apiToken.RefreshToken == "" {
		t.Errorf("Request api failed.")
		return
	}
}
func TestApi_GetAccessToken(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	accessToken, err := api.GetAccessToken()
	if err != nil {
		t.Error(err)
		return
	}
	if accessToken == "" {
		t.Errorf("Request api failed.")
		return
	}
}
func TestApi_GetRefreshToken(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	refreshToken, err := api.GetRefreshToken()
	if err != nil {
		t.Error(err)
		return
	}
	if refreshToken == "" {
		t.Errorf("Request api failed.")
		return
	}
}
func TestApi_RefreshApiToken(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	refreshApiToken, err := api.RefreshApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	if refreshApiToken.RefreshToken == "" || refreshApiToken.AccessToken == "" {
		t.Errorf("Request api failed.")
		return
	}
}
func TestApi_CreateAccount(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	account, err := api.CreateAccount()
	if err != nil {
		t.Error(err)
		return
	}
	if account.Address == "" {
		t.Errorf("Request api failed.")
		return
	}
}
func TestApi_QueryBlock(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	block, err := api.QueryBlock("number", "latest")
	if err != nil {
		t.Error(err)
		return
	}
	if block.Block.Hash == "" {
		t.Errorf("Block is empty.")
		return
	}
}
func TestApi_QueryBlocks(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	blocks, err := api.QueryBlocks(1, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if len(blocks.List) == 0 {
		t.Errorf("Blocks is empty.")
		return
	}
}
func TestApi_QueryBlocksByRange(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	blocks, err := api.QueryBlocksByRange(1, "2")
	if err != nil {
		t.Error(err)
		return
	}
	if len(blocks.Blocks) == 0 {
		t.Errorf("Blocks is empty.")
		return
	}
}
func TestApi_CompileContract(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	result, err := api.CompileContract(contractCode)
	if err != nil {
		t.Error(err)
		return
	}
	if result.Cts[0].Abi == "" || result.Cts[0].Bin == "" {
		t.Errorf("Compile contract failed.")
		return
	}
}
func TestApi_DeployContract(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	result, err := api.DeployContract(contractBin, address)
	if err != nil {
		t.Error(err)
		return
	}
	if result.TxHash == "" {
		t.Errorf("Deploy contract failed.")
		return
	}
}
func TestApi_DeployContractSync(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	result, err := api.DeployContractSync(contractBin, address)
	if err != nil {
		t.Error(err)
		return
	}
	if result.TxHash == "" {
		t.Errorf("Deploy contract failed.")
		return
	}
}
func TestApi_GetPayload(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	result, err := api.GetPayload(contractAbi, fn, args)
	if err != nil {
		t.Error(err)
		return
	}
	if result == "" {
		t.Errorf("Get payload failed.")
		return
	}
}
func TestApi_InvokeContract(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	result, err := api.InvokeContract(false, from, to, payload)
	if err != nil {
		t.Error(err)
		return
	}
	if result.TxHash == "" {
		t.Errorf("Invoke contract failed.")
		return
	}
}
func TestApi_InvokeContractSync(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	result, err := api.InvokeContractSync(false, from, to, payload)
	if err != nil {
		t.Error(err)
		return
	}
	if result.TxHash == "" {
		t.Errorf("Invoke contract failed.")
		return
	}
}
func TestApi_MaintainContract(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	result, err := api.MaintainContract(from, to, operation, payload)
	if err != nil {
		t.Error(err)
		return
	}
	if result.TxHash == "" {
		t.Errorf("Invoke contract failed.")
		return
	}
}
func TestApi_QueryContractStatus(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	result, err := api.QueryContractStatus(address)
	if err != nil {
		t.Error(err)
		return
	}
	if result.CtStatus == "" {
		t.Errorf("Query contract status failed.")
		return
	}
}
func TestApi_QueryTransactionCount(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	result, err := api.QueryTransactionCount()
	if err != nil {
		t.Error(err)
		return
	}
	if result.Timestamp == 0 {
		t.Errorf("Query transaction count failed.")
		return
	}
}
func TestApi_QueryTransactionByHash(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	result, err := api.QueryTransactionByHash(transactionHash)
	if err != nil {
		t.Error(err)
		return
	}
	if result.Transaction.Timestamp == 0 {
		t.Errorf("Query transaction by hash failed.")
		return
	}
}
func TestApi_QueryTransactionReceipt(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	result, err := api.QueryTransactionReceipt(transactionTxHash)
	if err != nil {
		t.Error(err)
		return
	}
	if result.TxHash == "" {
		t.Errorf("Query transaction receipt failed.")
		return
	}
}
func TestApi_QueryDiscardTransaction(t *testing.T) {
	api := New(phone, password, apiKey, apiSecret)
	_, err := api.GetApiToken()
	if err != nil {
		t.Error(err)
		return
	}
	_, err = api.QueryDiscardTransaction(transactionStart, transactionEnd)
	if err != nil {
		t.Error(err)
		return
	}
}
