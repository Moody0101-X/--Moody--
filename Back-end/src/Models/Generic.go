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