package common

type ValueMsg struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Value interface{} `json:"value"`
}

type NormalMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
