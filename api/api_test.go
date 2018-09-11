package api

import "testing"

const (
	phone     = "<phone>"
	password  = "<password>"
	apiKey    = "<apiKey>"
	apiSecret = "<apiSecret>"
)
const (
	contractCode = `
	pragma solidity ^0.4.10;
	contract Token {
		address issuer;
		mapping(address => bool) whiteList; // white list
		mapping(address => uint) balances;
	
		// status code
		uint CODE_SUCCESS = 3300; // success
		uint CODE_INSUFFICIENT_BALANCE = 3401; // balance is not enough
		uint CODE_INSUFFICIENT_PERMISSION = 3402; // permission denied
		uint CODE_INVALID_PARAMS = 3600; // invalid param
	
		function Token() {
			issuer = msg.sender;
		}
		// issue assets
		function issue(address account, uint amount) returns(uint) {
			if (msg.sender != issuer) return CODE_INSUFFICIENT_PERMISSION;
			balances[account] += amount;
			return CODE_SUCCESS;
		}
		// transfer assets
		function transfer(address to, uint amount) returns(uint) {
			if (balances[msg.sender] < amount) return CODE_INSUFFICIENT_BALANCE;
			balances[msg.sender] -= amount;
			balances[to] += amount;
			return CODE_SUCCESS;
		}
		// get user's balance
		function getBalance(address account) constant returns(uint, uint) {
			if (!whiteList[msg.sender]) return (CODE_INSUFFICIENT_PERMISSION, 0);
			return (CODE_SUCCESS, balances[account]);
		}
		// update white list
		function updateWhiteList(address account, uint opt) returns(uint) {
			if (msg.sender != issuer) return CODE_INSUFFICIENT_PERMISSION;
			if (opt == 3501) {
				// add whiteList user
				whiteList[account] = true;
			} else if (opt == 3502) {
				// remove whiteList user
				whiteList[account] = false;
			} else {
				return CODE_INVALID_PARAMS;
			}
			return CODE_SUCCESS;
		}
	}
	` // 合约源码
	contractBin  = "0x6060604052610ce4600355610d49600455610d4a600555610e10600655341561002757600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b6102f5806100546000396000f300606060405263ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663867904b4811461005e578063a386cfb914610092578063a9059cbb146100c6578063f8b2cb4f146100fa575b600080fd5b341561006957600080fd5b610080600160a060020a0360043516602435610131565b60405190815260200160405180910390f35b341561009d57600080fd5b610080600160a060020a036004351660243561017a565b60405190815260200160405180910390f35b34156100d157600080fd5b610080600160a060020a0360043516602435610213565b60405190815260200160405180910390f35b341561010557600080fd5b610119600160a060020a0360043516610275565b60405191825260208201526040908101905180910390f35b6000805433600160a060020a039081169116146101515750600554610174565b50600160a060020a03821660009081526002602052604090208054820190556003545b92915050565b6000805433600160a060020a0390811691161461019a5750600554610174565b81610dad14156101cf57600160a060020a0383166000908152600160208190526040909120805460ff19169091179055610207565b81610dae14156101fe57600160a060020a0383166000908152600160205260409020805460ff19169055610207565b50600654610174565b5b506003545b92915050565b600160a060020a0333166000908152600260205260408120548290101561023d5750600454610174565b50600160a060020a03338116600090815260026020526040808220805485900390559184168152208054820190556003545b92915050565b600160a060020a033316600090815260016020526040812054819060ff1615156102a557505060055460006102c4565b5050600354600160a060020a0382166000908152600260205260409020545b9150915600a165627a7a723058204e53bb53714ee98211364a7225d67f118a88f5a0edc47482f007110aa63c070c0029" // 合约BIN
	contractAbi  = `[{"constant":false,"inputs":[{"name":"account","type":"address"},{"name":"amount","type":"uint256"}],"name":"issue","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"account","type":"address"},{"name":"opt","type":"uint256"}],"name":"updateWhiteList","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"account","type":"address"}],"name":"getBalance","outputs":[{"name":"","type":"uint256"},{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[],"payable":false,"type":"constructor"}]` // 合约ABI
	address      = "" // 合约地址
)
const (
	from      = ""
	to        = ""
	payload   = ""
	operation = 1
)
const (
	transactionHash   = "" // 交易Hash
	transactionTxHash = "" // 交易Hash
	transactionStart  = 1504250000000000000  // 交易查询开始时间
	transactionEnd    = 1504368000000000000  // 交易查询结束时间
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
