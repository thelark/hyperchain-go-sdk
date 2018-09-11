package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"../api"
	"strings"
)

const (
	phone     = "<phone>"
	password  = "<password>"
	apiKey    = "<apiKey>"
	apiSecret = "<apiSecret>"
)

var (
	// contract info
	DEFAULT_CONTRACT_SOURCE = "contract test{}"
	DEFAULT_CONTRACT_BIN    = "0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a723058207799045e48fe5a1bd53859a58b9f2b52388e4cb227e7792eb4b894b366bd0bd40029"
	DEFAULT_CONTRACT_ABI    = "[]"
	// account info
	DEFAULT_ACCOUNT_ADDRESS = "0x3be60875d005800671e5fbfda15b0f49f1727494"
)

var _type = flag.String("type", "number", "查询类型,按区块高度查询:number,按区块哈希值查询:hash")
var _value = flag.String("value", "latest", "查询值,按类型填入哈希值或者区块号,注意可用字符串latest表示最新生成的区块号")
var _index = flag.Int("index", 1, "页码")
var _pageSize = flag.Int("pageSize", 2, "每页区块数量")
var _from = flag.Int("from", 1, "起始区块高度")
var _to = flag.String("to", "2", "终点区块高度,注意可用字符串latest代表最新区块号")
var _contractCode = flag.String("contractCode", "", "合约源码字符串,注意不能包含换行符,引号等特殊符号需转义")
var _bin = flag.String("bin", "", "合约BIN")
var _address = flag.String("address", "", "发起者地址|合约地址")
var _abi = flag.String("abi", "", "合约ABI")
var _args = flag.String("args", "", "方法参数列表,用','号隔开")
var _func = flag.String("func", "", "方法名")
var _const = flag.Bool("const", false, "表示交易不走共识，false表示走共识，默认为false")
var _fromAddress = flag.String("fAddress", "", "合约调用者地址")
var _toAddress = flag.String("tAddress", "", "合约地址")
var _operation = flag.Int("operation", 0, "执行操作,1：升级，2：冻结，3：解冻")
var _payload = flag.String("payload", "", `有多重含义:
1.方法名和方法参数经过编码后的input字节码
2.修改后的合约BIN`)
var _hash = flag.String("hash", "", "交易哈希值")
var _txhash = flag.String("txhash", "", "交易哈希值")
var _start = flag.Int64("start", 0, "起始时间戳(单位:ns)")
var _end = flag.Int64("end", 0, "终点时间戳(单位:ns)")

func init() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
func main() {
	api := api.New(phone, password, apiKey, apiSecret)
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("please input func name")
		return
	}
	switch args[0] {
	case "getApiToken":
		resp, err := api.GetApiToken()
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "refreshApiToken":
		api.GetApiToken()
		resp, err := api.RefreshApiToken()
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "createAccount":
		api.GetApiToken()
		resp, err := api.CreateAccount()
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "queryBlock":
		if *_type == "" || *_value == "" {
			fmt.Println("eg. queryBlock -type 'number' -value 'latest'")
			return
		}
		api.GetApiToken()
		resp, err := api.QueryBlock(*_type, *_value)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "queryBlocks":
		api.GetApiToken()
		resp, err := api.QueryBlocks(*_index, *_pageSize)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "queryBlocksByRange":
		api.GetApiToken()
		resp, err := api.QueryBlocksByRange(*_from, *_to)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "compileContract":
		code := *_contractCode
		if code == "" {
			code = DEFAULT_CONTRACT_SOURCE
		}
		api.GetApiToken()
		resp, err := api.CompileContract(code)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "deployContract":
		bin := *_bin
		address := *_address
		if bin == "" || address == "" {
			bin = DEFAULT_CONTRACT_BIN
			address = DEFAULT_ACCOUNT_ADDRESS
		}
		api.GetApiToken()
		resp, err := api.DeployContract(bin, address)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "deployContractSync":
		bin := *_bin
		address := *_address
		if bin == "" || address == "" {
			bin = DEFAULT_CONTRACT_BIN
			address = DEFAULT_ACCOUNT_ADDRESS
		}
		api.GetApiToken()
		resp, err := api.DeployContractSync(bin, address)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "getPayload":
		abi := *_abi
		args := strings.Split(*_args, ",")
		fn := *_func
		if abi == "" || fn == "" {
			fmt.Println("abi and func must input")
			return
		}
		api.GetApiToken()
		resp, err := api.GetPayload(abi, fn, args)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "invokeContract":
		from := *_fromAddress
		to := *_toAddress
		payload := *_payload
		if from == "" || to == "" || payload == "" {
			fmt.Println("fAddress, tAddress and payload is must input")
			return
		}
		api.GetApiToken()
		resp, err := api.InvokeContract(*_const, from, to, payload)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "invokeContractSync":
		from := *_fromAddress
		to := *_toAddress
		payload := *_payload
		if from == "" || to == "" || payload == "" {
			fmt.Println("fAddress, tAddress and payload is must input")
			return
		}
		api.GetApiToken()
		resp, err := api.InvokeContractSync(*_const, from, to, payload)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "maintainContract":
		operation := *_operation
		from := *_fromAddress
		to := *_toAddress
		payload := *_payload
		if operation == 0 || from == "" || to == "" || payload == "" {
			fmt.Println("operation, fAddress, tAddress and payload is must input")
			return
		}
		api.GetApiToken()
		resp, err := api.MaintainContract(from, to, operation, payload)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "queryContractStatus":
		address := *_address
		if address == "" {
			fmt.Println("address must input")
			return
		}
		api.GetApiToken()
		resp, err := api.QueryContractStatus(address)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "queryTransactionCount":
		api.GetApiToken()
		resp, err := api.QueryTransactionCount()
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "queryTransactionByHash":
		hash := *_hash
		if hash == "" {
			fmt.Println("hash must input")
			return
		}
		api.GetApiToken()
		resp, err := api.QueryTransactionByHash(hash)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "queryTransactionReceipt":
		txhash := *_txhash
		if txhash == "" {
			fmt.Println("hash must input")
			return
		}
		api.GetApiToken()
		resp, err := api.QueryTransactionReceipt(txhash)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	case "queryDiscardTransaction":
		start := *_start
		end := *_end
		if start == 0 || end == 0 || start > end {
			fmt.Println("start, end must input and end must gt end")
			return
		}
		api.GetApiToken()
		resp, err := api.QueryDiscardTransaction(start, end)
		if err != nil {
			log.Println(err)
			return
		}
		jb, _ := json.Marshal(resp)
		log.Println(string(jb))
	}
	return
}
