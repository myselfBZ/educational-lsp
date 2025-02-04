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

type InitializeResponse struct{
    Response
    Result InitializeResult `json:"result"`
}

type InitializeResult struct{
    Capabalities ServerCapabalities `json:"capabalities"`
    ServerInfo  ServerInfo          `json:"serverInfo"`
}

type ServerCapabalities struct{

}

type ServerInfo struct{
    Version string `json:"version"`
    Name string `json:"name"`
}

func NewResponse(id int) *InitializeResponse {
    return &InitializeResponse{
        Response: Response{
            RPC: "2.0",
            ID:id,
        },
        Result: InitializeResult{
            Capabalities: ServerCapabalities{},
            ServerInfo: ServerInfo{
                Name: "monkey",
                Version: "1.0",
            },
        },
    }
}
