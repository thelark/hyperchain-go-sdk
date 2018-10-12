package main

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/thelark/hyperchain-go-sdk/api"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	api := api.New("", "", "", "")
	getApiTokenCmd := &cli.Command{
		Name:    "getApiToken",
		Aliases: []string{},
		Usage:   "获取指定用户的API接入授权码",
		Flags:   []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			resp, err := api.GetApiToken()
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	refreshApiTokenCmd := &cli.Command{
		Name:    "refreshApiToken",
		Aliases: []string{},
		Usage:   "刷新API授权码",
		Flags:   []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.RefreshApiToken()
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	createAccountCmd := &cli.Command{
		Name:    "createAccount",
		Aliases: []string{},
		Usage:   "新建区块链账户,返回区块链地址",
		Flags:   []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.CreateAccount()
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	queryBlockCmd := &cli.Command{
		Name:    "queryBlock",
		Aliases: []string{},
		Usage:   "查询指定区块的详细信息",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "type",
				Value:   "number",
				Usage:   "查询类型,按区块高度查询:number,按区块哈希值查询:hash",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "value",
				Value:   "latest",
				Usage:   "查询值,按类型填入哈希值或者区块号,注意可用字符串latest表示最新生成的区块号",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.QueryBlock(ctx.String("type"), ctx.String("value"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	queryBlocksCmd := &cli.Command{
		Name:    "queryBlocks",
		Aliases: []string{},
		Usage:   "区块信息分页查询",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Aliases: []string{},
				Name:    "index",
				Value:   1,
				Usage:   "页码",
			},
			&cli.IntFlag{
				Aliases: []string{},
				Name:    "pageSize",
				Value:   2,
				Usage:   "每页区块数量",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.QueryBlocks(ctx.Int("index"), ctx.Int("pageSize"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	queryBlocksByRangeCmd := &cli.Command{
		Name:    "queryBlocksByRange",
		Aliases: []string{},
		Usage:   "查询指定高度区间的区块列表",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Aliases: []string{},
				Name:    "from",
				Value:   1,
				Usage:   "起始区块高度",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "to",
				Value:   "2",
				Usage:   "终点区块高度,注意可用字符串latest代表最新区块号",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.QueryBlocksByRange(ctx.Int("from"), ctx.String("to"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	compileContractCmd := &cli.Command{
		Name:    "compileContract",
		Aliases: []string{},
		Usage:   "编译智能合约",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "code",
				Value:   "contract test{}",
				Usage:   "合约源码字符串,注意不能包含换行符,引号等特殊符号需转义",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.CompileContract(ctx.String("code"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	deployContractCmd := &cli.Command{
		Name:    "deployContract",
		Aliases: []string{},
		Usage:   "部署智能合约",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "bin",
				Value:   "0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a723058207799045e48fe5a1bd53859a58b9f2b52388e4cb227e7792eb4b894b366bd0bd40029",
				Usage:   "合约BIN",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "address",
				Value:   "0x3be60875d005800671e5fbfda15b0f49f1727494",
				Usage:   "发起者地址|合约地址",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.DeployContract(ctx.String("bin"), ctx.String("address"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	deployContractSyncCmd := &cli.Command{
		Name:    "deployContractSync",
		Aliases: []string{},
		Usage:   "同步部署智能合约,直接返回部署结果,包含合约地址",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "bin",
				Value:   "0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a723058207799045e48fe5a1bd53859a58b9f2b52388e4cb227e7792eb4b894b366bd0bd40029",
				Usage:   "合约BIN",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "address",
				Value:   "0x3be60875d005800671e5fbfda15b0f49f1727494",
				Usage:   "发起者地址|合约地址",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.DeployContractSync(ctx.String("bin"), ctx.String("address"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	getPayloadCmd := &cli.Command{
		Name:    "getPayload",
		Aliases: []string{},
		Usage:   "调用合约需要转入合约方法与合约参数编码后的input字节码,该接口可为用户返回Payload",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "abi",
				Value:   "[]",
				Usage:   "合约ABI",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "func",
				Value:   "",
				Usage:   "方法名",
			},
			&cli.StringSliceFlag{
				Aliases: []string{},
				Name:    "args",
				Value:   &cli.StringSlice{},
				Usage:   "方法参数列表,用','号隔开",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.GetPayload(ctx.String("abi"), ctx.String("func"), ctx.StringSlice("args"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	invokeContractCmd := &cli.Command{
		Name:    "invokeContract",
		Aliases: []string{},
		Usage:   "调用合约",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Aliases: []string{},
				Name:    "const",
				Value:   false,
				Usage:   "表示交易不走共识，false表示走共识，默认为false",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "from",
				Value:   "",
				Usage:   "合约调用者地址",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "to",
				Value:   "",
				Usage:   "合约地址",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "payload",
				Value:   "",
				Usage:   "方法名和方法参数经过编码后的input字节码",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.InvokeContract(ctx.Bool("const"), ctx.String("from"), ctx.String("to"), ctx.String("payload"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	invokeContractSyncCmd := &cli.Command{
		Name:    "invokeContractSync",
		Aliases: []string{},
		Usage:   "同步调用合约",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Aliases: []string{},
				Name:    "const",
				Value:   false,
				Usage:   "表示交易不走共识，false表示走共识，默认为false",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "from",
				Value:   "",
				Usage:   "合约调用者地址",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "to",
				Value:   "",
				Usage:   "合约地址",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "payload",
				Value:   "",
				Usage:   "方法名和方法参数经过编码后的input字节码",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.InvokeContractSync(ctx.Bool("const"), ctx.String("from"), ctx.String("to"), ctx.String("payload"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	maintainContractCmd := &cli.Command{
		Name:    "maintainContract",
		Aliases: []string{},
		Usage:   "合约维护,包括合约的升级,冻结和解冻",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "from",
				Value:   "",
				Usage:   "合约调用者地址",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "to",
				Value:   "",
				Usage:   "合约地址",
			},
			&cli.IntFlag{
				Aliases: []string{},
				Name:    "operate",
				Value:   0,
				Usage:   "执行操作,1：升级，2：冻结，3：解冻",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "payload",
				Value:   "",
				Usage:   "修改后的合约BIN",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.MaintainContract(ctx.String("from"), ctx.String("to"), ctx.Int("operate"), ctx.String("payload"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	queryContractStatusCmd := &cli.Command{
		Name:    "queryContractStatus",
		Aliases: []string{},
		Usage:   "查询合约当前状态",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "address",
				Value:   "",
				Usage:   "合约地址",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.QueryContractStatus(ctx.String("address"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	queryTransactionCountCmd := &cli.Command{
		Name:    "queryTransactionCount",
		Aliases: []string{},
		Usage:   "查询联盟链上的交易总数",
		Flags:   []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.QueryTransactionCount()
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	queryTransactionByHashCmd := &cli.Command{
		Name:    "queryTransactionByHash",
		Aliases: []string{},
		Usage:   "查询指定交易hash的交易详情",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "hash",
				Value:   "",
				Usage:   "交易哈希值",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.QueryTransactionByHash(ctx.String("hash"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	queryTransactionReceiptCmd := &cli.Command{
		Name:    "queryTransactionReceipt",
		Aliases: []string{},
		Usage:   "查询指定交易hash的交易回执",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "txhash",
				Value:   "",
				Usage:   "交易哈希值",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.QueryTransactionReceipt(ctx.String("txhash"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	queryDiscardTransactionCmd := &cli.Command{
		Name:    "queryDiscardTransaction",
		Aliases: []string{},
		Usage:   "查询指定时间区间内的非法交易",
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Aliases: []string{},
				Name:    "start",
				Value:   0,
				Usage:   "起始时间戳(单位:ns)",
			},
			&cli.Int64Flag{
				Aliases: []string{},
				Name:    "end",
				Value:   0,
				Usage:   "终点时间戳(单位:ns)",
			},
		},
		Action: func(ctx *cli.Context) error {
			_, err := api.GetApiToken()
			if err != nil {
				return err
			}
			resp, err := api.QueryDiscardTransaction(ctx.Int64("start"), ctx.Int64("end"))
			if err != nil {
				return err
			}
			jb, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Println(string(jb))
			return nil
		},
	}
	app := &cli.App{
		Name:        "hyperchain-cli",
		Usage:       "hyperchain-cli <method> <args>",
		Version:     "1.0.0",
		Description: "hyperchain sdk cli",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "phone",
				Value:   "",
				Usage:   "注册手机号",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "password",
				Value:   "",
				Usage:   "密码",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "apiKey",
				Value:   "",
				Usage:   "api key",
			},
			&cli.StringFlag{
				Aliases: []string{},
				Name:    "apiSecret",
				Value:   "",
				Usage:   "api secret",
			},
		},
		Commands: []*cli.Command{
			getApiTokenCmd,
			refreshApiTokenCmd,
			createAccountCmd,
			queryBlockCmd,
			queryBlocksCmd,
			queryBlocksByRangeCmd,
			compileContractCmd,
			deployContractCmd,
			deployContractSyncCmd,
			getPayloadCmd,
			invokeContractCmd,
			invokeContractSyncCmd,
			maintainContractCmd,
			queryContractStatusCmd,
			queryTransactionCountCmd,
			queryTransactionByHashCmd,
			queryTransactionReceiptCmd,
			queryDiscardTransactionCmd,
		},
		Before: func(ctx *cli.Context) error {
			api.SetPhone(ctx.String("phone"))
			api.SetPassword(ctx.String("password"))
			api.SetApiKey(ctx.String("apiKey"))
			api.SetApiSecret(ctx.String("apiSecret"))
			return nil
		},
	}
	app.Run(os.Args)
}
