package manifest

const (
	ClusterlinkFloaterServiceAccount = `
apiVersion: v1
kind: ServiceAccount
metadata:
  name:  clusterlink-floater
  namespace: {{ .Namespace }}
`
)

type ServiceAccountReplace struct {
	Namespace string
}
