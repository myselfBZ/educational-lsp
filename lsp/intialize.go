package lsp

type InitializeRequest struct {
	Request
    Params InitParams `json:"params"`
}

type InitParams struct {
	ClientInfo ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
    Name    string `json:"name"`
	Version string `json:"version"`
}
