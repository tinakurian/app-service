package kubeclient

import (
	ocappsclient "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubeClient struct {
	CoreClient   kubernetes.Interface
	OcAppsClient ocappsclient.AppsV1Interface
	//OcRouteClient ocrouteclient.RouteV1Interface
}

func NewKubeClient(config *rest.Config) *KubeClient {
	var err error
	kc := new(KubeClient)
	kc.CoreClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	kc.OcAppsClient, err = ocappsclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return kc
}

func getKubeConfig() rest.Config {
	host := "https://api.tkurian16.devcluster.openshift.com:6443"
	bearerToken := "7DdF7VrdYl2F9MrR95J_v0Z0pJj1qh6tMZrSzbn_Uno"

	return getOpenshiftAPIConfig(host, bearerToken)
}

func getOpenshiftAPIConfig(host string, bearerToken string) rest.Config {
	return rest.Config{
		Host:        host,
		BearerToken: bearerToken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
}
