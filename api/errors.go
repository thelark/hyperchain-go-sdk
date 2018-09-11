package api

import "errors"

var (
	ERR_NOACCESSTOKEN  = errors.New("Don't have access token, please get api token first.")
	ERR_NOREFRESHTOKEN = errors.New("Don't have refresh token, please get api token first.")
)

/*
1008 	授权未通过
1009 	非法参数
1010 	查询异常
1011 	appkey不存在
1012 	不支持的查询类型
1013 	合约编译异常
1014 	账户私钥异常
1015 	合约部署异常
1016 	合约调用异常
1017 	合约维护操作异常
*/
