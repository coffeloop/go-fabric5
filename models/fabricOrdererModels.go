package models

type OrdererOptions struct {
	Image        string `json:"ordererImage"`
	Version      string `json:"ordererVersion"`
	StorageClass string `json:"storageClass"`
	EnrollID     string `json:"enrollId"`
	MSPID        string `json:"mspid"`
	EnrollPW     string `json:"enrollPw"`
	Capacity     string `json:"capacity"`
	Name         string `json:"name"`
	CAName       string `json:"caName"`
	Hosts        string `json:"hosts"`
	IstioPort    int    `json:"istioPort"`
}

type CheckOrdererOptions struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
