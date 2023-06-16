package dto

type ReqBase struct {
	LogId string            `json:"logId"`
	Extra map[string]string `json:"extra"`
}

type ResBase struct {
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Extra   map[string]string `json:"extra,omitempty"`
}
