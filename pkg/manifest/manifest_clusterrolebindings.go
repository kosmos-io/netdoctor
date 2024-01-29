package manifest

const (
	ClusterlinkFloaterClusterRoleBinding = `
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: clusterlink-floater
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: clusterlink-floater
subjects:
  - kind: ServiceAccount
    name: clusterlink-floater
    namespace: {{ .Namespace }}
`
)

type ClusterRoleBindingReplace struct {
	Namespace string
}
