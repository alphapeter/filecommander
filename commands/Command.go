package commands

type Command struct {
	RpcVersion string   `josn:"jsonrpc"`
	Method     string   `json:"method"`
	Params     []string `json:"params"`
	Id         string   `json:"id"`
}
