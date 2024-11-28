package models

type CreateCAOptions struct {
	Image        string   `json:"image"`
	Version      string   `json:"version"`
	StorageClass string   `json:"storageclass"`
	Capacity     string   `json:"capacity"`
	Name         string   `json:"name"`
	EnrollID     string   `json:"enroll_id"`
	EnrollPW     string   `json:"enroll_pw"`
	Hosts        []string `json:"hosts"`
	IstioPort    int      `json:"istio_port"`
}

type RegisterCAOptions struct {
	Name         string `json:"name"`
	User         string `json:"user"`
	Secret       string `json:"secret"`
	Type         string `json:"type"`
	EnrollID     string `json:"enrollId"`
	EnrollSecret string `json:"enrollSecret"`
	MSPID        string `json:"mspid"`
}

type EnrollCAOptions struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	User      string `json:"user"`
	Secret    string `json:"secret"`
	MSPID     string `json:"mspid"`
	CAName    string `json:"ca_name"`
	Output    string `json:"output"`
}

type RegisterUserCAOptions struct {
	Name         string `json:"name"`
	User         string `json:"user"`
	Secret       string `json:"secret"`
	Type         string `json:"type"`
	EnrollID     string `json:"enrollId"`
	EnrollSecret string `json:"enrollSecret"`
	MSPID        string `json:"mspid"`
}

type CheckCAOptions struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
