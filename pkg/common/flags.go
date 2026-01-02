package common

import (
	"github.com/spf13/pflag"
	"k8s.io/client-go/tools/clientcmd"
)

func AddFlags(flags pflag.FlagSet, loadingRules *clientcmd.ClientConfigLoadingRules, overrides *clientcmd.ConfigOverrides) {
	flags.StringVar(&loadingRules.ExplicitPath, "kubeconfig", loadingRules.ExplicitPath, "Path to the kubeconfig file to use")
	flags.StringVar(&overrides.CurrentContext, "context", overrides.CurrentContext, "The name of the kubeconfig context to use")
	flags.StringVar(&overrides.Context.AuthInfo, "user", overrides.Context.AuthInfo, "The name of the kubeconfig user to use")
	flags.StringVar(&overrides.Context.Cluster, "cluster", overrides.Context.Cluster, "The name of the kubeconfig cluster to use")
	flags.StringVarP(&overrides.Context.Namespace, "namespace", "n", overrides.Context.Namespace, "The name of the Kubernetes Namespace to work in (NOT optional)")
}
