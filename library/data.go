package library

type UserData struct {
	Name     string `json:"name,omitempty" ddb:"name,omitempty"`
	Sex      int64  `json:"sex,omitempty" ddb:"sex,omitempty"`
	PassWard string `json:"pass,omitempty" ddb:"pass,omitempty"`
}

type ResData struct {
	Result  *ResResult  `json:"result"`
	Content *ResContent `json:"content"`
}

type ResResult struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
}

type ResContent struct {
	User interface{} `json:"user,omitempty"`
}

func ResJson(code int, msg string, data *ResContent) *ResData {
	return &ResData{
		Result: &ResResult{
			Errno: code,
			Msg:   msg,
		},
		Content: data,
	}
}
