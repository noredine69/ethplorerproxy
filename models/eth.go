package models

type LastBlock struct {
	LastBlockId uint `json:"lastBlock"`
}

type ProxyEthRequest struct {
	JsonRpcVersion string   `json:"jsonrpc"`
	Method         string   `json:"method"`
	Params         []string `json:"params"`
	Id             uint     `json:"id"`
}

type ProxyEthResponse struct {
	JsonRpcVersion string `json:"jsonrpc"`
	Id             uint   `json:"id"`
	Result         string `json:"result"`
}
