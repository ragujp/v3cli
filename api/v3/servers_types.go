package v3

import "gorm.io/gorm"

type LibrespeedTestServer struct {
	gorm.Model                          // id, created_at, updated_at, deleted_at
	TypeName   LibrespeedTestServerType `json:"typeName"`
	Name       string                   `json:"name"`
	Server     string                   `json:"server"`
	DlUrl      string                   `json:"dlURL"`
	UlUrl      string                   `json:"ulURL"`
	PingUrl    string                   `json:"pingURL"`
	GetIpUrl   string                   `json:"getIpURL"` // IP is special word, but need to refer librespeed design.
}

type LibrespeedTestServerType string

const (
	LibrespeedTestServerTypeIPv4 LibrespeedTestServerType = "ipv4"
	LibrespeedTestServerTypeIPv6 LibrespeedTestServerType = "ipv6"
)

type OneshotTestServer struct {
	gorm.Model                         // id, created_at, updated_at, deleted_at
	TypeName     OneshotTestServerType `json:"typeName"`
	HTTPEndpoint string                `json:"httpEndpoint"`
}

type OneshotTestServerType string

const (
	OneshotTestServerTypeQUIC OneshotTestServerType = "quic"
	OneshotTestServerTypeMSS  OneshotTestServerType = "mss"
)

// TestServerの一覧を種別別に返す用
type TestServersList struct {
	Librespeed []LibrespeedTestServer `json:"librespeed"`
	Oneshot    []OneshotTestServer    `json:"oneshot"`
}
