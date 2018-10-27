package module

type Error struct {
	ErrCode int64
	ErrMsg  error
}

type ApiResp struct {
	ErrNo  int64       `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data,omitempty"`
}
