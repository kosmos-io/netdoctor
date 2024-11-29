package manifest

const (
	NetDoctorFloaterClusterRoleBinding = `
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: netdr-floater
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: netdr-floater
subjects:
  - kind: ServiceAccount
    name: netdr-floater
    namespace: {{ .Namespace }}
`
)

type ClusterRoleBindingReplace struct {
	Namespace string
}
