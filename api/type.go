package api

type commonResponse struct {
	Code   int    `json:"Code"`
	Status string `json:"Status"`
}
type commonErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type commonTokenResponse struct {
	AccessToken  string `json:"access_token"`  // API授权码
	ExpiresIn    int    `json:"expires_in"`    // 7200s后授权码失效
	RefreshToken string `json:"refresh_token"` // 刷新码,可用于刷新授权码，注意:不能作为授权码使用
	Scope        string `json:"scope"`         // 同请求参数里的scope
	TokenType    string `json:"token_type"`    // 授权码类型
}
type (
	respCreateAccount struct {
		Code       int    `json:"Code"`       // 错误码
		Status     string `json:"Status"`     // 状态信息
		Id         int    `json:"id"`         // 账号ID
		Address    string `json:"address"`    // 区块链地址
		Time       string `json:"time"`       // 创建时间
		IsDisabled bool   `json:"isDisabled"` // 是否禁用
		AppName    string `json:"appName"`    // 所属应用的名称
	}
	respQueryBlock struct {
		Code   int    `json:"Code"`   // 错误码
		Status string `json:"Status"` // 状态信息
		Block  struct {
			Number       int64  `json:"Number"`     // 区块高度
			Hash         string `json:"Hash"`       // 交易哈希值,32字节的十六进制字符串
			ParentHash   string `json:"ParentHash"` // 父区块哈希值,32字节的十六进制字符串
			WriteTime    int64  `json:"WriteTime"`  // 区块的生成时间(单位:ns)
			AvgTime      int    `json:"AvgTime"`    // 当前区块中，交易的平均处理时间(单位ms)
			Txcounts     int    `json:"Txcounts"`   // 当前区块中打包的交易数量
			MerkleRoot   string `json:"MerkleRoot"` // Merkle树的根哈希
			Transactions []struct {
				Version     string `json:"Version"`
				Hash        string `json:"Hash"`
				BlockNumber int    `json:"BlockNumber"`
				BlockHash   string `json:"BlockHash"`
				TxIndex     int    `json:"TxIndex"`
				From        string `json:"From"`
				To          string `json:"To"`
				Amount      int    `json:"Amount"`
				Timestamp   int64  `json:"Timestamp"`
				Nonce       int64  `json:"Nonce"`
				ExecuteTime int    `json:"ExecuteTime"`
				Payload     string `json:"Payload"`
				Invalid     bool   `json:"Invalid"`
				InvalidMsg  string `json:"InvalidMsg"`
			} `json:"Transactions"` // 区块中的交易列表
		} `json:"block"` // 区块信息
	}
	respQueryBlocks struct {
		Code   int    `json:"Code"`   // 错误码
		Status string `json:"Status"` // 状态信息
		List   []struct {
			Number       int64  `json:"Number"`
			Hash         string `json:"Hash"`
			ParentHash   string `json:"ParentHash"`
			WriteTime    int64  `json:"WriteTime"`
			AvgTime      int    `json:"AvgTime"`
			Txcounts     int    `json:"Txcounts"`
			MerkleRoot   string `json:"MerkleRoot"`
			Transactions []struct {
				Version     string `json:"Version"`
				Hash        string `json:"Hash"`
				BlockNumber int    `json:"BlockNumber"`
				BlockHash   string `json:"BlockHash"`
				TxIndex     int    `json:"TxIndex"`
				From        string `json:"From"`
				To          string `json:"To"`
				Amount      int    `json:"Amount"`
				Timestamp   int64  `json:"Timestamp"`
				Nonce       int64  `json:"Nonce"`
				ExecuteTime int    `json:"ExecuteTime"`
				Payload     string `json:"Payload"`
				Invalid     bool   `json:"Invalid"`
				InvalidMsg  string `json:"InvalidMsg"`
			} `json:"Transactions"`
		} `json:"List"` // 区块列表
		Count int64 `json:"Count"` // 区块总数
	}
	respQueryBlocksByRange struct {
		Code   int    `json:"Code"`   // 错误码
		Status string `json:"Status"` // 状态信息
		Blocks []struct {
			Number       int64  `json:"Number"`
			Hash         string `json:"Hash"`
			ParentHash   string `json:"ParentHash"`
			WriteTime    int64  `json:"WriteTime"`
			AvgTime      int    `json:"AvgTime"`
			Txcounts     int    `json:"Txcounts"`
			MerkleRoot   string `json:"MerkleRoot"`
			Transactions []struct {
				Version     string `json:"Version"`
				Hash        string `json:"Hash"`
				BlockNumber int    `json:"BlockNumber"`
				BlockHash   string `json:"BlockHash"`
				TxIndex     int    `json:"TxIndex"`
				From        string `json:"From"`
				To          string `json:"To"`
				Amount      int    `json:"Amount"`
				Timestamp   int64  `json:"Timestamp"`
				Nonce       int64  `json:"Nonce"`
				ExecuteTime int    `json:"ExecuteTime"`
				Payload     string `json:"Payload"`
				Invalid     bool   `json:"Invalid"`
				InvalidMsg  string `json:"InvalidMsg"`
			} `json:"Transactions"`
		} `json:"Blocks"` // 区块列表
	}
	respCompileContract struct {
		Code   int    `json:"Code"`   // 错误码
		Status string `json:"Status"` // 状态信息
		Cts    []struct {
			Code   int    `json:"Code"`
			Status string `json:"Status"` // 编译状态描述
			Id     int    `json:"Id"`     // 编译项id
			Bin    string `json:"Bin"`    // 编译Bin结果
			Abi    string `json:"Abi"`    // 编译Abi结果
			Name   string `json:"Name"`   //
			OK     bool   `json:"OK"`     // 是否编译成功
		} `json:"Cts"` // 编译结果项列表
	}
	respDeployContract struct {
		Code   int    `json:"Code"`   // 错误码
		Status string `json:"Status"` // 状态信息
		TxHash string `json:"TxHash"` // 本次交易的哈希值
	}
	respDeployContractSync struct {
		Code            int    `json:"Code"`           // 错误码
		Status          string `json:"Status"`         // 状态信息
		TxHash          string `json:"TxHash"`         // 本次交易的哈希值
		PostState       string `json:"PostState"`      //
		ContractAddress string `json:"ContractAddress` // 部署后的合约地址
		Ret             string `json:"Ret"`            // 返回值
	}
	respInvokeContract struct {
		Code   int    `json:"Code"`   // 错误码
		Status string `json:"Status"` // 状态信息
		TxHash string `json:"TxHash"` // 本次交易的哈希值
	}
	respInvokeContractSync struct {
		Code            int    `json:"Code"`           // 错误码
		Status          string `json:"Status"`         // 状态信息
		TxHash          string `json:"TxHash"`         // 本次交易的哈希值
		PostState       string `json:"PostState"`      //
		ContractAddress string `json:"ContractAddress` // 合约地址
		Ret             string `json:"Ret"`            // 调用返回值
	}
	respMaintainContract struct {
		Code   int    `json:"Code"`   // 错误码
		Status string `json:"Status"` // 状态信息
		TxHash string `json:"TxHash"` // 本次交易的哈希值
	}
	respQueryContractStatus struct {
		Code     int    `json:"Code"`     // 错误码
		Status   string `json:"Status"`   // 状态信息
		CtStatus string `json:"ctStatus"` // 合约状态
	}
	respQueryTransactionCount struct {
		Code      int    `json:"Code"`      // 错误码
		Status    string `json:"Status"`    // 状态信息
		Count     int    `json:"Count"`     // 交易总数
		Timestamp int64  `json:"Timestamp"` // 时间戳
	}
	respQueryTransactionByHash struct {
		Code        int    `json:"Code"`   // 错误码
		Status      string `json:"Status"` // 状态信息
		Transaction struct {
			Version     string `json:"Version"`     // hyperchain平台版本
			Hash        string `json:"Hash"`        // 交易哈希值,32字节的十六进制字符串,交易的唯一标识
			BlockNumber int    `json:"BlockNumber"` // 交易所在区块的区块号
			BlockHash   string `json:"BlockHash"`   // 交易所在区块的区块哈希值
			TxIndex     int    `json:"TxIndex"`     // 交易在区块中的交易列表的位置，起始数字为0
			From        string `json:"From"`        // 交易发送方地址,20字节的十六进制字符串
			To          string `json:"To"`          // 交易接收方地址,若该交易为合约调用,则To为合约地址,20字节的十六进制字符串
			Amount      int    `json:"Amount"`      // 交易量（预留字段）
			Timestamp   int64  `json:"Timestamp"`   // 交易发生时间(单位:ns)
			Nonce       int64  `json:"Nonce"`       // 16位随机数，用于确保交易的唯一性
			ExecuteTime int    `json:"ExecuteTime"` // 交易的处理时间(单位:ms)
			Payload     string `json:"Payload"`     // 部署合约与调用合约的时候才有这个值，可以通过这个值追溯到合约调用的方法以及调用传入的参数
			Invalid     bool   `json:"Invalid"`     // 交易是否不合法
			InvalidMsg  string `json:"InvalidMsg"`  // 交易的不合法信息
		} `json:"Transaction"` // 交易详细信息
	}
	respQueryTransactionReceipt struct {
		Code            int    `json:"Code"`           // 错误码
		Status          string `json:"Status"`         // 状态信息
		TxHash          string `json:"TxHash"`         // 本次交易的哈希值
		PostState       string `json:"PostState"`      //
		ContractAddress string `json:"ContractAddress` // 合约地址
		Ret             string `json:"Ret"`            // 执行的结果
	}
	respQueryDiscardTransaction struct {
		Code         int    `json:"Code"`   // 错误码
		Status       string `json:"Status"` // 状态信息
		Transactions []struct {
			Version     string `json:"Version"`
			Hash        string `json:"Hash"`
			BlockNumber int    `json:"BlockNumber"`
			BlockHash   string `json:"BlockHash"`
			TxIndex     int    `json:"TxIndex"`
			From        string `json:"From"`
			To          string `json:"To"`
			Amount      int    `json:"Amount"`
			Timestamp   int64  `json:"Timestamp"`
			Nonce       int64  `json:"Nonce"`
			ExecuteTime int    `json:"ExecuteTime"`
			Payload     string `json:"Payload`
			Invalid     bool   `json:"Invalid"`
			InvalidMsg  string `json:"InvalidMsg"`
		} `json:"Transactions"` // 区间内所有非法交易列表
	}
)
