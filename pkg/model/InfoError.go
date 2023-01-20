package model

type ErrorInfo struct {
	Code          string `json:"code"`
	Usermessage   string `json:"usermessage"`
	Systemmessage string `json:"systemmessage"`
}
