package log

import "strings"

type Infra string

const (
	InfraLocal    Infra = "local"
	InfraCloudGcp Infra = "gcp"
	InfraCloudAws Infra = "aws"
)

var AllInfras = []Infra{
	InfraLocal,
	InfraCloudGcp,
	InfraCloudAws,
}

func ParseInfra(infra string) Infra {
	switch strings.ToLower(strings.Trim(infra, "")) {
	case "local":
		return InfraLocal
	case "gcp":
		return InfraCloudGcp
	case "aws":
		return InfraCloudAws
	default:
		Warnf("unrecognized infra type = %s, falling back to InfraLocal, available Infra are %v", infra, AllInfras)
		return InfraLocal
	}
}
