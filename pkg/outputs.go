package pkg

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/gitlabkubernetes"
	"github.com/plantoncloud/stack-job-runner-golang-sdk/pkg/automationapi/autoapistackoutput"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

const (
	NamespaceOutputName          = "namespace"
	ServiceOutputName            = "service"
	PortForwardCommandOutputName = "port-forward-command"
	KubeEndpointOutputName       = "kube-endpoint"
)

func PulumiOutputToStackOutputsConverter(pulumiOutputs auto.OutputMap,
	input *gitlabkubernetes.GitlabKubernetesStackInput) *gitlabkubernetes.GitlabKubernetesStackOutputs {
	return &gitlabkubernetes.GitlabKubernetesStackOutputs{
		Namespace:          autoapistackoutput.GetVal(pulumiOutputs, NamespaceOutputName),
		Service:            autoapistackoutput.GetVal(pulumiOutputs, ServiceOutputName),
		PortForwardCommand: autoapistackoutput.GetVal(pulumiOutputs, PortForwardCommandOutputName),
		KubeEndpoint:       autoapistackoutput.GetVal(pulumiOutputs, KubeEndpointOutputName),
	}
}
