package kubeclient

import (
	ocappsclient "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	ocrouteclient "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubeClient struct {
	CoreClient    kubernetes.Interface
	OcAppsClient  ocappsclient.AppsV1Interface
	OcRouteClient ocrouteclient.RouteV1Interface
}

func NewKubeClient() *KubeClient {
	var err error
	kc := new(KubeClient)
	config := getKubeConfig()
	kc.CoreClient, err = kubernetes.NewForConfig(&config)
	if err != nil {
		panic(err)
	}
	kc.OcAppsClient, err = ocappsclient.NewForConfig(&config)
	if err != nil {
		panic(err)
	}

	kc.OcRouteClient, err = ocrouteclient.NewForConfig(&config)
	if err != nil {
		panic(err)
	}

	return kc
}

func getKubeConfig() rest.Config {
	host := "https://api.tkurian22.devcluster.openshift.com:6443"
	bearerToken := "X6ppoX-V1M5kZaVe1Yw8PZaIzDEwUxx2qEp7q3tIySs"

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
