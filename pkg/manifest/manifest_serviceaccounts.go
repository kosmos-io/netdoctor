package manifest

const (
	NetDoctorFloaterServiceAccount = `
apiVersion: v1
kind: ServiceAccount
metadata:
  name:  netdr-floater
  namespace: {{ .Namespace }}
`
)

type ServiceAccountReplace struct {
	Namespace string
}
