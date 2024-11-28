package models

type ConnectionChain struct {
	Output string `json:"output"`
	MSPID  string `json:"mspid"`
}

type AddUserToConnectionChainOptions struct {
	UserPath string `json:"userPath"`
	Config   string `json:"config"`
	Username string `json:"username"`
	MSPID    string `json:"mspid"`
}
