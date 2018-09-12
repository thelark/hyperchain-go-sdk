package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type api struct {
	phone     string // phone number (username)
	password  string // password
	apiKey    string // apiKey
	apiSecret string // apiSecret
	// ----------- token
	accessToken  string // access token
	refreshToken string // refresh token
}

const (
	apiUrl = "https://api.hyperchain.cn/v1"
)

var client = &http.Client{
	// Timeout: time.Second,
}

// New - Create new api
func New(_phone, _password string, _apiKey, _apiSecret string) *api {
	return &api{phone: _phone, password: _password, apiKey: _apiKey, apiSecret: _apiSecret}
}

// CheckAccessToken -  Check access token
func (a *api) checkAccessToken() bool {
	return a.accessToken != ""
}

// CheckRefreshToken -  Check refresh token
func (a *api) checkRefreshToken() bool {
	return a.refreshToken != ""
}

// GetAccessToken - Get access token
func (a *api) GetAccessToken() (string, error) {
	if !a.checkAccessToken() {
		return "", ERR_NOACCESSTOKEN
	}
	return a.accessToken, nil
}

// GetRefreshToken - Get refresh token
func (a *api) GetRefreshToken() (string, error) {
	if !a.checkRefreshToken() {
		return "", ERR_NOREFRESHTOKEN
	}
	return a.refreshToken, nil
}

// -------------------------------------------------------------

// GetApiToken - Get the auth token(获取API接入码)
func (a *api) GetApiToken() (*commonTokenResponse, error) {
	uri := fmt.Sprintf("%s/token/gtoken/", apiUrl)
	body := strings.NewReader(fmt.Sprintf("phone=%s&password=%s&client_id=%s&client_secret=%s", a.phone, a.password, a.apiKey, a.apiSecret))
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept":       "text/plain",
	}
	buffer, err := doRequest(uri, "POST", header, body)
	if err != nil {
		return nil, err
	}
	tmpResp := new(commonResponse)
	err = json.Unmarshal(buffer, &tmpResp)
	if err != nil {
		return nil, err
	}
	if tmpResp.Code != 0 {
		return nil, fmt.Errorf(tmpResp.Status)
	}
	resp := new(commonTokenResponse)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	a.accessToken = resp.AccessToken
	a.refreshToken = resp.RefreshToken
	return resp, nil
}

// RefreshApiToken - Refresh the auth token(刷新API接入码)
func (a *api) RefreshApiToken() (*commonTokenResponse, error) {
	if !a.checkRefreshToken() {
		return nil, ERR_NOREFRESHTOKEN
	}
	uri := fmt.Sprintf("%s/token/rtoken/", apiUrl)
	body := strings.NewReader(fmt.Sprintf("refresh_token=%s&client_id=%s&client_secret=%s", a.refreshToken, a.apiKey, a.apiSecret))
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept":       "text/plain",
	}
	buffer, err := doRequest(uri, "POST", header, body)
	if err != nil {
		return nil, err
	}
	tmpResp := new(commonResponse)
	err = json.Unmarshal(buffer, &tmpResp)
	if err != nil {
		return nil, err
	}
	if tmpResp.Code != 0 {
		return nil, fmt.Errorf(tmpResp.Status)
	}
	errorResp := new(commonErrorResponse)
	err = json.Unmarshal(buffer, &errorResp)
	if err != nil {
		return nil, err
	}
	if errorResp.Error != "" {
		return nil, fmt.Errorf(errorResp.ErrorDescription)
	}
	resp := new(commonTokenResponse)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	a.accessToken = resp.AccessToken
	a.refreshToken = resp.RefreshToken
	return resp, nil
}

// CreateAccount - Create account(新建账号)
func (a *api) CreateAccount() (*respCreateAccount, error) {
	/**
		curl -X GET \
	    --header 'Accept: application/json' \
	    --header 'Authorization: EQJN4SFONAMGLLLADMH2MG' \
	    'https://api.hyperchain.cn/v1/dev/account/create'
	*/
	// {"Code":0,"Status":"","id":4850,"address":"0x825b5e9fc3546e086b8b9832abf68ba9f464d970","time":"2018-09-09 14:46:36","isDisabled":false,"appName":"学习"}
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/account/create", apiUrl)
	header := map[string]string{
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	buffer, err := doRequest(uri, "GET", header, nil)
	if err != nil {
		return nil, err
	}
	resp := new(respCreateAccount)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// QueryBlock - Query block(查询区块信息)
// @param _type		string 查询类型,按区块高度查询:number,按区块哈希值查询:hash
// @param _value	string 查询值,按类型填入哈希值或者区块号,注意可用字符串latest表示最新生成的区块号
func (a *api) QueryBlock(_type, _value string) (*respQueryBlock, error) {
	/**
		curl -X GET \
	    --header 'Accept: application/json' \
	    --header 'Authorization: JRKTO3PDOT-U1JTHO5VS7A' \
		'https://api.hyperchain.cn/v1/dev/block/query?type=number&value=latest'
	*/
	// {"Code":0,"Status":"ok","block":{"Number":385088,"Hash":"0x709c3f66c3bab5b5a08d2ff247ea544edc0deab3f75403998c5eb328070e6f91","ParentHash":"0x456cc045659198137486a5c7a518cd7c0c5fa480ab8bc580596a7ff0669b8a51","WriteTime":1536466493503394800,"AvgTime":11,"Txcounts":1,"MerkleRoot":"0x524b329f6e3729af91b44915f70ac22d9df5c868f4ec186b521d0d25b67d769f","Transactions":[{"Version":"1.2","Hash":"0xee83c1f1c91c2400edba5ae942744c0feaf3aaa85c28ce9ab7ff4244762298db","BlockNumber":385088,"BlockHash":"0x709c3f66c3bab5b5a08d2ff247ea544edc0deab3f75403998c5eb328070e6f91","TxIndex":0,"From":"0x4ec56dab780f1d35ba740e5af2c08db0312ef43b","To":"0x0b9b8f244ca9e63eb6c2013e9ceaacc8bf4af689","Amount":0,"Timestamp":1536466497151322600,"Nonce":646033713864111700,"ExecuteTime":11,"Payload":"0xf8b2cb4f0000000000000000000000004ec56dab780f1d35ba740e5af2c08db0312ef43b","Invalid":false,"InvalidMsg":""}]}}
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/block/query?type=%s&value=%s", apiUrl, _type, _value)
	header := map[string]string{
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	buffer, err := doRequest(uri, "GET", header, nil)
	if err != nil {
		return nil, err
	}
	resp := new(respQueryBlock)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// QueryBlocks - Query blocks(分页查询区块信息)
// @param _index	number 页码
// @param _pageSize	number 每页区块数量
func (a *api) QueryBlocks(_index, _pageSize int) (*respQueryBlocks, error) {
	/**
		curl -X GET \
	    --header 'Accept: application/json' \
	    --header 'Authorization: 1SL1AUWHM-S-BYZZVYXYFG' \
	    'https://api.hyperchain.cn/v1/dev/blocks/page?index=1&pageSize=2'
	*/
	// {"Code":0,"Status":"ok","List":[{"Number":385088,"Hash":"0x709c3f66c3bab5b5a08d2ff247ea544edc0deab3f75403998c5eb328070e6f91","ParentHash":"0x456cc045659198137486a5c7a518cd7c0c5fa480ab8bc580596a7ff0669b8a51","WriteTime":1536466493503394800,"AvgTime":11,"Txcounts":1,"MerkleRoot":"0x524b329f6e3729af91b44915f70ac22d9df5c868f4ec186b521d0d25b67d769f","Transactions":[{"Version":"1.2","Hash":"0xee83c1f1c91c2400edba5ae942744c0feaf3aaa85c28ce9ab7ff4244762298db","BlockNumber":385088,"BlockHash":"0x709c3f66c3bab5b5a08d2ff247ea544edc0deab3f75403998c5eb328070e6f91","TxIndex":0,"From":"0x4ec56dab780f1d35ba740e5af2c08db0312ef43b","To":"0x0b9b8f244ca9e63eb6c2013e9ceaacc8bf4af689","Amount":0,"Timestamp":1536466497151322600,"Nonce":646033713864111700,"ExecuteTime":11,"Payload":"0xf8b2cb4f0000000000000000000000004ec56dab780f1d35ba740e5af2c08db0312ef43b","Invalid":false,"InvalidMsg":""}]},{"Number":385087,"Hash":"0x456cc045659198137486a5c7a518cd7c0c5fa480ab8bc580596a7ff0669b8a51","ParentHash":"0xdf70b2e4e201c8a04a12bfa67fe1bddd3dc36691c74a7064326a94241a658011","WriteTime":1536464464186828000,"AvgTime":9,"Txcounts":1,"MerkleRoot":"0x524b329f6e3729af91b44915f70ac22d9df5c868f4ec186b521d0d25b67d769f","Transactions":[{"Version":"1.2","Hash":"0xb2921194691e60da395769d41197d2dce60f7ed62b40a9609de3a4ee258db2b9","BlockNumber":385087,"BlockHash":"0x456cc045659198137486a5c7a518cd7c0c5fa480ab8bc580596a7ff0669b8a51","TxIndex":0,"From":"0x0f8ee0dad301718bb145d4baa47190c48c15d41b","To":"0xc9c49fcfbed7c0b30310370ee677e648654fae88","Amount":0,"Timestamp":1536464467837475800,"Nonce":8660119264438445000,"ExecuteTime":9,"Payload":"0xf93471306f30526a74774d6750515f5f6f2d775f2d536b6d4a33716c71533641000000000000000000000000000000000000000000000000000000000000000000000002","Invalid":false,"InvalidMsg":""}]}],"Count":385088}
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/blocks/page?index=%d&pageSize=%d", apiUrl, _index, _pageSize)
	header := map[string]string{
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	buffer, err := doRequest(uri, "GET", header, nil)
	if err != nil {
		return nil, err
	}
	resp := new(respQueryBlocks)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// QueryBlocksByRange - Query blocks by range(根据区块号范围查询区块信息)
// @param _from	number 起始区块高度
// @param _to	string 终点区块高度,注意可用字符串latest代表最新区块号
func (a *api) QueryBlocksByRange(_from int, _to string) (*respQueryBlocksByRange, error) {
	/**
		curl -X GET \
	    --header 'Accept: application/json' \
	    --header 'Authorization: 1SL1AUWHM-S-BYZZVYXYFG' \
	    'https://api.hyperchain.cn/v1/dev/blocks/range?from=1&to=2'
	*/
	// {"Code":0,"Status":"ok","Blocks":[{"Number":2,"Hash":"0x913cdeb66c57ff2be994f66bdbbb2ccc2c3a845ce25bbb5a6b37fc6a0702d23c","ParentHash":"0x5ad3c9b9f156c3ea2b018f330c0663221ccfdbc6dc872120ae36123ed0d35515","WriteTime":1503399161444970000,"AvgTime":20,"Txcounts":1,"MerkleRoot":"0x3b7801bc24bb6ada08bb3cfb4add51791f51335b480853fee7ac77bf85a3e97b","Transactions":[{"Version":"1.2","Hash":"0x0c4aae58494fada572de4f8d0d15fa3b9f922a154ab0016c6159c3d2d196a00a","BlockNumber":2,"BlockHash":"0x913cdeb66c57ff2be994f66bdbbb2ccc2c3a845ce25bbb5a6b37fc6a0702d23c","TxIndex":0,"From":"0x27cb880403feb4e7f1fdba3410334e607becb7a5","To":"0x0000000000000000000000000000000000000000","Amount":0,"Timestamp":1503399160979970300,"Nonce":4831389563158288000,"ExecuteTime":20,"Payload":"0x60606040523415600b57fe5b604051604080606b8339810160405280516020909101515b8181015b5050505b60338060386000396000f30060606040525bfe00a165627a7a723058200cb906fa6d1f24809c4cac4552d2b7f8f1d6677613d1405aa42f82120a1a88210029","Invalid":false,"InvalidMsg":""}]},{"Number":1,"Hash":"0x5ad3c9b9f156c3ea2b018f330c0663221ccfdbc6dc872120ae36123ed0d35515","ParentHash":"0x0000000000000000000000000000000000000000000000000000000000000000","WriteTime":1503394536661740800,"AvgTime":23,"Txcounts":1,"MerkleRoot":"0x239b95dcb13681c9da7cea8edbcff419afb0fdd0838a682bcb35ae4cc730a300","Transactions":[{"Version":"1.2","Hash":"0x81d0a2409ac4769de713bcb8a4bc2ea6a324e37c378a705c8c29d096ebc3a1d8","BlockNumber":1,"BlockHash":"0x5ad3c9b9f156c3ea2b018f330c0663221ccfdbc6dc872120ae36123ed0d35515","TxIndex":0,"From":"0x3d508bc18baa6cd651a809540f19322f25ca30ea","To":"0x0000000000000000000000000000000000000000","Amount":0,"Timestamp":1503394399375738400,"Nonce":4125532999287192000,"ExecuteTime":23,"Payload":"0x60a0604052600b60608190527f68656c6c6f20776f726c64000000000000000000000000000000000000000000608090815261003e916000919061004c565b50341561004757fe5b6100ec565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061008d57805160ff19168380011785556100ba565b828001600101855582156100ba579182015b828111156100ba57825182559160200191906001019061009f565b5b506100c79291506100cb565b5090565b6100e991905b808211156100c757600081556001016100d1565b5090565b90565b610188806100fb6000396000f300606060405263ffffffff60e060020a6000350416638da9b7728114610021575bfe5b341561002957fe5b6100316100b1565b604080516020808252835181830152835191928392908301918501908083838215610077575b80518252602083111561007757601f199092019160209182019101610057565b505050905090810190601f1680156100a35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100b961014a565b6000805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561013f5780601f106101145761010080835404028352916020019161013f565b820191906000526020600020905b81548152906001019060200180831161012257829003601f168201915b505050505090505b90565b604080516020810190915260008152905600a165627a7a72305820ff8a0b89ef40366cc75ff0dafe9442fab2283671d34cac35d7aa90ad1ee623330029","Invalid":false,"InvalidMsg":""}]}]}
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/blocks/range?from=%d&to=%s", apiUrl, _from, _to)
	header := map[string]string{
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	buffer, err := doRequest(uri, "GET", header, nil)
	if err != nil {
		return nil, err
	}
	resp := new(respQueryBlocksByRange)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// CompileContract - Compile contract(编译智能合约)
// @param _contractCode	string 合约源码字符串,注意不能包含换行符,引号等特殊符号需转义
func (a *api) CompileContract(_contractCode string) (*respCompileContract, error) {
	/**
		curl -X POST \
	    --header 'Content-Type: application/json' \
	    --header 'Accept: application/json' \
	    --header 'Authorization: 1SL1AUWHM-S-BYZZVYXYFG' \
	    -d '{"CTCode": "contract test{}"}' \
	    'https://api.hyperchain.cn/v1/dev/contract/compile'
	*/
	// {"Code":0,"Status":"ok","Cts":[{"Code":0,"Status":"","Id":0,"Bin":"0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a723058207799045e48fe5a1bd53859a58b9f2b52388e4cb227e7792eb4b894b366bd0bd40029","Abi":"[]","Name":"test","OK":true}]}
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/contract/compile", apiUrl)
	header := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	body := strings.NewReader(fmt.Sprintf(`{"CTCode":"%s"}`, _contractCode))
	buffer, err := doRequest(uri, "POST", header, body)
	if err != nil {
		return nil, err
	}
	resp := new(respCompileContract)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// DeployContract - Deploy contract(部署智能合约)
// @param _bin	string 合约BIN
// @param _from string 发起者地址
func (a *api) DeployContract(_bin, _from string) (*respDeployContract, error) {
	/**
		curl -X POST \
	    --header 'Content-Type: application/json' \
	    --header 'Accept: application/json' \
	    --header 'Authorization: 1SL1AUWHM-S-BYZZVYXYFG' \
	    -d '{ \
	        "Bin": "0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820b4c36b8b61723f302432d246407a061599017f8607ed26f1c053b5ecc63a54200029", \
	        "From": "0x3713c3d01ae09cf32787c9c66c9c0781cf4b613d" \
	        }' \
	   'https://api.hyperchain.cn/v1/dev/contract/deploy'
	*/
	// {"Code":0,"Status":"ok","TxHash":"0xd6271905df2ae9ad31358999c50afcb18dad804fd2f94487b805bf7fe4ee02a5"}
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/contract/deploy", apiUrl)
	header := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	body := strings.NewReader(fmt.Sprintf(`{"Bin":"%s","From":"%s"}`, _bin, _from))
	buffer, err := doRequest(uri, "POST", header, body)
	if err != nil {
		return nil, err
	}
	resp := new(respDeployContract)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// DeployContractSync - Deploy contract sync(同步部署智能合约,部署完成之后返回回执信息)
// @param _bin	string 合约BIN
// @param _from string 发起者地址
func (a *api) DeployContractSync(_bin, _from string) (*respDeployContractSync, error) {
	/**
		curl -X POST \
	    --header 'Content-Type: application/json' \
	    --header 'Accept: application/json' \
	    --header 'Authorization: 1SL1AUWHM-S-BYZZVYXYFG' \
	    -d '{ \
	        "Bin": "0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820b4c36b8b61723f302432d246407a061599017f8607ed26f1c053b5ecc63a54200029", \
	        "From": "0x3713c3d01ae09cf32787c9c66c9c0781cf4b613d" \
	        }'\
	    'https://api.hyperchain.cn/v1/dev/contract/deploysync'
	*/
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/contract/deploysync", apiUrl)
	header := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	body := strings.NewReader(fmt.Sprintf(`{"Bin":"%s","From":"%s"}`, _bin, _from))
	buffer, err := doRequest(uri, "POST", header, body)
	if err != nil {
		return nil, err
	}
	resp := new(respDeployContractSync)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// GetPayload - Get payload(获取Payload)
// @param _abi	string		合约ABI
// @param _args	[]string	方法参数列表
// @param _func string 		方法名
func (a *api) GetPayload(_abi string, _func string, _args []string) (string, error) {
	/**
		curl -X POST \
	    --header 'Content-Type: application/json' \
	    --header 'Accept: application/json' \
	    --header 'Authorization: EQJN4SFONAMGLLLADMH2MG' \
	    -d '{ \
	        "Abi": "[{\"constant\":false,\"inputs\":[{\"name\":\"num1\",\"type\":\"uint32\"},{\"name\":\"num2\",\"type\":\"uint32\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getSum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"payable\":false,\"type\":\"function\"}]", \
	        "Args": ["1", "2"], \
	        "Func": "add" \
	        }' \
	    'https://api.hyperchain.cn/v1/dev/payload'
	*/
	if !a.checkAccessToken() {
		return "", ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/payload", apiUrl)
	header := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	jsonBuffer, err := json.Marshal(_args)
	if err != nil {
		return "", err
	}
	body := strings.NewReader(fmt.Sprintf(`{"Abi":"%s","Args":%s,"Func":"%s"}`, _abi, string(jsonBuffer), _func))
	buffer, err := doRequest(uri, "POST", header, body)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

// InvokeContract - Invoke contract(调用合约)
// @param _const	boolean	表示交易不走共识，false表示走共识，默认为false
// @param _from		string	合约调用者地址
// @param _to		string	合约地址
// @param _payload	string	方法名和方法参数经过编码后的input字节码
func (a *api) InvokeContract(_const bool, _from, _to string, _payload string) (*respInvokeContract, error) {
	/**
		curl -X POST \
	    --header 'Content-Type: application/json' \
	    --header 'Accept: application/json' \
	    --header 'Authorization: 1SL1AUWHM-S-BYZZVYXYFG' \
	    -d '{ \
	        "Const": false, \
	        "From": "0x3713c3d01ae09cf32787c9c66c9c0781cf4b613d", \
	        "Payload": "34141", \
	        "To": "0x8255340c2c4a1aec4010d2b6fdbb98727c65523d" \
	        }' \
	    'https://api.hyperchain.cn/v1/dev/contract/invoke'
	*/
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/contract/invoke", apiUrl)
	header := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	body := strings.NewReader(fmt.Sprintf(`{"Const":%t,"From":"%s","Payload":"%s","To":"%s"}`, _const, _from, _payload, _to))
	buffer, err := doRequest(uri, "POST", header, body)
	if err != nil {
		return nil, err
	}
	resp := new(respInvokeContract)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// InvokeContractSync - Invoke contract sync(同步调用合约,等待共识完成之后返回回执信息)
// @param _const	boolean	表示交易不走共识，false表示走共识，默认为false
// @param _from		string	合约调用者地址
// @param _to		string	合约地址
// @param _payload	string	方法名和方法参数经过编码后的input字节码
func (a *api) InvokeContractSync(_const bool, _from, _to string, _payload string) (*respInvokeContractSync, error) {
	/**
		curl -X POST \
	    --header 'Content-Type: application/json' \
	    --header 'Accept: application/json' \
	    --header 'Authorization: 3C7_ZNDAPB-QWVGSK3R3DG' \
	    -d '{ \
	        "Const": false, \
	        "From": "0x23b4f4aa3be625ef5f629523dc7e06ed73f161f7", \
	        "Payload": "f39581ad", \
	        "To": "0x8bdc64dce18a743294ca480e86ea13a43a1bf255" \
	        }' \
	    'https://api.hyperchain.cn/v1/dev/contract/invokesync'
	*/
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/contract/invokesync", apiUrl)
	header := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	body := strings.NewReader(fmt.Sprintf(`{"Const":%t,"From":"%s","Payload":"%s","To":"%s"}`, _const, _from, _payload, _to))
	buffer, err := doRequest(uri, "POST", header, body)
	if err != nil {
		return nil, err
	}
	resp := new(respInvokeContractSync)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// MaintainContract - Maintain contract(维护智能合约,可对部署后的合约进行升级/冻结/解冻)
// @param _operation	number	执行操作,1：升级，2：冻结，3：解冻
// @param _from			string	合约调用者地址
// @param _to			string	合约地址
// @param _payload		string	修改后的合约BIN
func (a *api) MaintainContract(_from, _to string, _operation int, _payload string) (*respMaintainContract, error) {
	/**
		curl -X POST \
	    --header 'Content-Type: application/json' \
	    --header 'Accept: application/json' \
	    --header 'Authorization: EQJN4SFONAMGLLLADMH2MG' \
	    -d '{ \
	        "From": "0x19a170a0413096a9b18f2ca4066faa127f4d6f4a", \
	        "Operation": 1, \
	        "Payload": "0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820b4c36b8b61723f302432d246407a061599017f8607ed26f1c053b5ecc63a54200029", \
	        "To": "0xd3a7bdd391f6aa13b28a72690e19d2ab3d845ac8" \
	        }' \
	    'https://api.hyperchain.cn/v1/dev/contract/maintain'
	*/
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/contract/maintain", apiUrl)
	header := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	body := strings.NewReader(fmt.Sprintf(`{"From":"%s","Operation":%d,"Payload":"%s","To":"%s"}`, _from, _operation, _payload, _to))
	buffer, err := doRequest(uri, "POST", header, body)
	if err != nil {
		return nil, err
	}
	resp := new(respMaintainContract)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// QueryContractStatus - Query contract status(查询智能合约状态)
// @param _address string 合约地址
func (a *api) QueryContractStatus(_address string) (*respQueryContractStatus, error) {
	/**
		curl -X GET \
	    --header 'Accept: application/json' \
	    --header 'Authorization: EQJN4SFONAMGLLLADMH2MG' \
	    'https://api.hyperchain.cn/v1/dev/contract/status?address=0xd3a7bdd391f6aa13b28a72690e19d2ab3d845ac8' // 需要查询的合约地址
	*/
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/contract/status?address=%s", apiUrl, _address)
	header := map[string]string{
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	buffer, err := doRequest(uri, "GET", header, nil)
	if err != nil {
		return nil, err
	}
	resp := new(respQueryContractStatus)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// QueryTransactionCount - Query transaction count(查询区块链交易总数)
func (a *api) QueryTransactionCount() (*respQueryTransactionCount, error) {
	/**
		curl -X GET \
	    --header 'Accept: application/json' \
	    --header 'Authorization: JRKTO3PDOT-U1JTHO5VS7A' \
	    'https://api.hyperchain.cn/v1/dev/transaction/count'
	*/
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/transaction/count", apiUrl)
	header := map[string]string{
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	buffer, err := doRequest(uri, "GET", header, nil)
	if err != nil {
		return nil, err
	}
	resp := new(respQueryTransactionCount)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// QueryTransactionByHash - Query transaction by hash(通过交易哈希值查询交易信息)
// @param _hash	string 交易哈希值
func (a *api) QueryTransactionByHash(_hash string) (*respQueryTransactionByHash, error) {
	/**
		curl -X GET \
	    --header 'Accept: application/json' \
	    --header 'Authorization: JRKTO3PDOT-U1JTHO5VS7A' \
	    'https://api.hyperchain.cn/v1/dev/transaction/query?hash=0xed70377c261bfdc7dd7f4fc15c8961c145f9457186d6ff95f60907e9fb63d827'
	*/
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/transaction/query?hash=%s", apiUrl, _hash)
	header := map[string]string{
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	buffer, err := doRequest(uri, "GET", header, nil)
	if err != nil {
		return nil, err
	}
	resp := new(respQueryTransactionByHash)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// QueryTransactionReceipt - Query transaction receipt(查询交易回执信息)
// @param _txhash string 交易哈希值
func (a *api) QueryTransactionReceipt(_txhash string) (*respQueryTransactionReceipt, error) {
	/**
		curl -X GET \
	    --header 'Accept: application/json' \
	    --header 'Authorization: JRKTO3PDOT-U1JTHO5VS7A' \
	    'https://api.hyperchain.cn/v1/dev/transaction/txreceipt?txhash=0xed70377c261bfdc7dd7f4fc15c8961c145f9457186d6ff95f60907e9fb63d827'
	*/
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/transaction/txreceipt?txhash=%s", apiUrl, _txhash)
	header := map[string]string{
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	buffer, err := doRequest(uri, "GET", header, nil)
	if err != nil {
		return nil, err
	}
	resp := new(respQueryTransactionReceipt)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// QueryDiscardTransaction - Query discard transaction(根据时间戳范围查询非法交易)
// @param _start	string	起始时间戳(单位:ns)
// @param _end		string	终点时间戳(单位:ns)
func (a *api) QueryDiscardTransaction(_start int64, _end int64) (*respQueryDiscardTransaction, error) {
	/**
		curl -X GET \
	    --header 'Accept: application/json' \
	    --header 'Authorization: JRKTO3PDOT-U1JTHO5VS7A' \
	    'https://api.hyperchain.cn/v1/dev/transactions/discard?start=1504250000000000000&end=1504368000000000000'
	*/
	if !a.checkAccessToken() {
		return nil, ERR_NOACCESSTOKEN
	}
	uri := fmt.Sprintf("%s/dev/transactions/discard?start=%d&end=%d", apiUrl, _start, _end)
	header := map[string]string{
		"Accept":        "application/json",
		"Authorization": a.accessToken,
	}
	buffer, err := doRequest(uri, "GET", header, nil)
	if err != nil {
		return nil, err
	}
	resp := new(respQueryDiscardTransaction)
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// utils method
func doRequest(uri, method string, header map[string]string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, uri, body)
	for k, v := range header {
		req.Header.Add(k, v)
	}
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
