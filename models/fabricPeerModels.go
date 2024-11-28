package models

type CreatePeerOptions struct {
	StateDB     string `json:"statedb"`
	PeerImage   string `json:"peerImage"`
	PeerVersion string `json:"peerVersion"`
	SCName      string `json:"scName"`
	EnrollID    string `json:"enrollId"`
	MSPID       string `json:"mspid"`
	EnrollPW    string `json:"enrollPw"`
	Capacity    string `json:"capacity"`
	Name        string `json:"name"`
	CAName      string `json:"caName"`
	Hosts       string `json:"hosts"`
	IstioPort   int    `json:"istioPort"`
}

type CheckPeerOptions struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type FabricPeerResponse struct {
	Status string `json:"status"`
}
