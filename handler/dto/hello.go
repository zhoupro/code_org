package dto

type ReqHello struct {
	Name    string `json:"name"`
	ReqBase `json:"req_base"`
}

type ResHelloData struct {
	Echo string `json:"echo"`
}

type ResHello struct {
	Data []ResHelloData `json:"data"`
	ResBase
}
