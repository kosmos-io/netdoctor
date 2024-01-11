package utils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
)

var (
	defaultKubeConfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
)

// RestConfig creates a rest config from the context and kubeconfig.
func RestConfig(kubeconfigPath, context string) (*rest.Config, error) {
	rawConfig, err := loadKubeconfig(kubeconfigPath, context)
	if err != nil {
		return nil, err
	}

	restConfig, err := clientcmd.NewDefaultClientConfig(*rawConfig, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, err
	}

	return restConfig, nil
}

func loadKubeconfig(kubeconfigPath, context string) (*clientcmdapi.Config, error) {
	if kubeconfigPath == "" {
		kubeconfigPath = GetEnvString("KUBECONFIG", defaultKubeConfig)
	}

	if _, err := os.Stat(kubeconfigPath); err != nil {
		return nil, fmt.Errorf("kubeconfig path %s does not exist", kubeconfigPath)
	}

	pathOptions := clientcmd.NewDefaultPathOptions()
	loadingRules := *pathOptions.LoadingRules
	loadingRules.ExplicitPath = kubeconfigPath
	overrides := &clientcmd.ConfigOverrides{
		CurrentContext: context,
	}

	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(&loadingRules, overrides)
	rawConfig, err := clientConfig.RawConfig()
	if err != nil {
		return nil, err
	}

	return &rawConfig, nil
}

// GetEnvString returns the env variable,if the env is not set,return the defaultValue
func GetEnvString(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}

func GenerateDeployment(deployTemplate string, obj interface{}) (*appsv1.Deployment, error) {
	deployBytes, err := parseTemplate(deployTemplate, obj)
	if err != nil {
		return nil, fmt.Errorf("kosmosctl parsing Deployment template exception, error: %v", err)
	} else if deployBytes == nil {
		return nil, fmt.Errorf("kosmosctl get Deployment template exception, value is empty")
	}

	deployStruct := &appsv1.Deployment{}

	if err = runtime.DecodeInto(scheme.Codecs.UniversalDecoder(), deployBytes, deployStruct); err != nil {
		return nil, fmt.Errorf("kosmosctl decode deployBytes error: %v", err)
	}

	return deployStruct, nil
}

func GenerateDaemonSet(dsTemplate string, obj interface{}) (*appsv1.DaemonSet, error) {
	dsBytes, err := parseTemplate(dsTemplate, obj)
	if err != nil {
		return nil, fmt.Errorf("kosmosctl parsing DaemonSet template exception, error: %v", err)
	} else if dsBytes == nil {
		return nil, fmt.Errorf("kosmosctl get DaemonSet template exception, value is empty")
	}

	dsStruct := &appsv1.DaemonSet{}

	if err = runtime.DecodeInto(scheme.Codecs.UniversalDecoder(), dsBytes, dsStruct); err != nil {
		return nil, fmt.Errorf("kosmosctl decode dsBytes error: %v", err)
	}

	return dsStruct, nil
}

func GenerateServiceAccount(saTemplate string, obj interface{}) (*corev1.ServiceAccount, error) {
	saBytes, err := parseTemplate(saTemplate, obj)
	if err != nil {
		return nil, fmt.Errorf("kosmosctl parsing ServiceAccount template exception, error: %v", err)
	} else if saBytes == nil {
		return nil, fmt.Errorf("kosmosctl get ServiceAccount template exception, value is empty")
	}

	saStruct := &corev1.ServiceAccount{}

	if err = runtime.DecodeInto(scheme.Codecs.UniversalDecoder(), saBytes, saStruct); err != nil {
		return nil, fmt.Errorf("kosmosctl decode saBytes error: %v", err)
	}

	return saStruct, nil
}

func GenerateClusterRole(crTemplate string, obj interface{}) (*rbacv1.ClusterRole, error) {
	crBytes, err := parseTemplate(crTemplate, obj)
	if err != nil {
		return nil, fmt.Errorf("kosmosctl parsing ClusterRole template exception, error: %v", err)
	} else if crBytes == nil {
		return nil, fmt.Errorf("kosmosctl get ClusterRole template exception, value is empty")
	}

	crStruct := &rbacv1.ClusterRole{}

	if err = runtime.DecodeInto(scheme.Codecs.UniversalDecoder(), crBytes, crStruct); err != nil {
		return nil, fmt.Errorf("kosmosctl decode crBytes error: %v", err)
	}

	return crStruct, nil
}

func GenerateClusterRoleBinding(crbTemplate string, obj interface{}) (*rbacv1.ClusterRoleBinding, error) {
	crbBytes, err := parseTemplate(crbTemplate, obj)
	if err != nil {
		return nil, fmt.Errorf("kosmosctl parsing ClusterRoleBinding template exception, error: %v", err)
	} else if crbBytes == nil {
		return nil, fmt.Errorf("kosmosctl get ClusterRoleBinding template exception, value is empty")
	}

	crbStruct := &rbacv1.ClusterRoleBinding{}

	if err = runtime.DecodeInto(scheme.Codecs.UniversalDecoder(), crbBytes, crbStruct); err != nil {
		return nil, fmt.Errorf("kosmosctl decode crbBytes error: %v", err)
	}

	return crbStruct, nil
}

func parseTemplate(strTmpl string, obj interface{}) ([]byte, error) {
	var buf bytes.Buffer
	tmpl, err := template.New("template").Parse(strTmpl)
	if err != nil {
		return nil, fmt.Errorf("error when parsing template: %v", err)
	}
	err = tmpl.Execute(&buf, obj)
	if err != nil {
		return nil, fmt.Errorf("error when executing template: %v", err)
	}
	return buf.Bytes(), nil
}
