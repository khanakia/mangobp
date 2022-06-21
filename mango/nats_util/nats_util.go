package nats_util

type ErrorResp struct {
	Msg        string `json:"msg,omitempty"`
	Code       string `json:"code,omitempty"`       // auth_login, token_error
	StatusCode string `json:"statusCode,omitempty"` // 400|401
}

type Resp struct {
	Data  interface{} `json:"data"`
	Error *ErrorResp  `json:"error"`
}

func CreateRespWithErr(msg, code, statusCode string) Resp {
	return Resp{
		Error: &ErrorResp{
			Msg:        msg,
			Code:       code,
			StatusCode: statusCode,
		},
	}
}

func CreateRespWithData(data interface{}) Resp {
	return Resp{
		Data: data,
	}
}
