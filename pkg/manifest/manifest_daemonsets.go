package manifest

const (
	NetDoctorFloaterDaemonSet = `
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: netdr-floater
  namespace: {{ .Namespace }}
  labels:
    app: netdr-floater
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{ .Name }}
  template:
    metadata:
      labels:
        app: {{ .Name }}
    spec:
      hostNetwork: {{ .EnableHostNetwork }}
      serviceAccountName: netdr-floater
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: kosmos.io/exclude
                operator: DoesNotExist
      containers:
      - name: floater
        image: {{ .ImageRepository }}/netdr-floater:{{ .Version }}
        imagePullPolicy: IfNotPresent
        command:
          - netdr-floater
        securityContext:
          privileged: true
        env: 
          - name: "PORT"
            value: "{{ .Port }}"
          - name: "ENABLE_ANALYSIS"
            value: "{{ .EnableAnalysis }}"
      hostPID: true
      tolerations:
      - effect: NoSchedule
        operator: Exists
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoExecute
        operator: Exists
`
)

type DaemonSetReplace struct {
	Namespace       string
	Name            string
	ImageRepository string
	Version         string
	Port            string

	EnableHostNetwork bool `default:"false"`
	EnableAnalysis    bool `default:"false"`
}
