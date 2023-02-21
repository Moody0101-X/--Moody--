package models

func IsEmpty(s string) bool { return len(s) == 0 }

type ServerResult struct {
	Ok 	 	bool
	Desc 	string
	Code    int
}

func NewServerResult(ok bool, text string, code int) ServerResult {

	return ServerResult{
		Ok: ok,
		Desc: text,
		Code: code,
	}

}

func MakeServerRespFromResult(res ServerResult) ServerResponse {
	return ServerResponse{
		Data: res.Desc,
		Code: res.Code,
	}
}

func MakeServerResp(code int, data interface{}) ServerResponse {
	return ServerResponse{
		Data: data,
		Code: code,	
	}
}

type ServerResponse struct {
	Data interface{}	`json:"data"`
	Code int         	`json:"code"`
}