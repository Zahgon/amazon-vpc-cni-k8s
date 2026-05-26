package common

import (
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/aws/amazon-vpc-cni-k8s/test/agent/pkg/input"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	coreV1 "k8s.io/api/core/v1"

	"github.com/aws/amazon-vpc-cni-k8s/test/framework"
	"github.com/aws/amazon-vpc-cni-k8s/test/framework/resources/agent"
	"github.com/aws/amazon-vpc-cni-k8s/test/framework/resources/k8s/manifest"
)

type TestType int

var (
	// The Pod labels for client and server in order to retrieve the
	// client and server Pods belonging to a Deployment/Jobs
	labelKey          = "app"
	serverPodLabelVal = "server-pod"
	clientPodLabelVal = "client-pod"
)

const (
	NetworkingTearDownSucceeds TestType = iota
	NetworkingTearDownFails
	NetworkingSetupSucceeds
	NetworkingSetupFails
)

type InterfaceTypeToPodList struct {
	PodsOnPrimaryENI   []coreV1.Pod
	PodsOnSecondaryENI []coreV1.Pod
}

func GetPodNetworkingValidationInput(interfaceTypeToPodList InterfaceTypeToPodList, vpcCIDRs []string) input.PodNetworkingValidationInput {
	_ = "STUB: not implemented"
	return *new(input.PodNetworkingValidationInput)
}

// Validate host networking for the list of pods supplied
func ValidateHostNetworking(testType TestType, podValidationInputString string, nodeName string, f *framework.Framework) {
	_ = "STUB: not implemented"
	return
}

// GetPodsOnPrimaryAndSecondaryInterface returns the list of Pods on Primary Networking
// Interface and Secondary Network Interface on a given Node
func GetPodsOnPrimaryAndSecondaryInterface(node coreV1.Node,
	podLabelKey string, podLabelVal string, f *framework.Framework) InterfaceTypeToPodList {
	_ = "STUB: not implemented"
	return *new(InterfaceTypeToPodList)
}

func GetTrafficTestConfig(f *framework.Framework, protocol string, serverDeploymentBuilder *manifest.DeploymentBuilder, clientCount int, serverCount int) agent.TrafficTest {
	_ = "STUB: not implemented"
	return *new(agent.TrafficTest)
}

func IsPrimaryENI(nwInterface ec2types.InstanceNetworkInterface, instanceIPAddr *string) bool {
	_ = "STUB: not implemented"
	return false
}

func ApplyCNIManifest(filepath string) { _ = "STUB: not implemented"; return }

func ValidateTraffic(f *framework.Framework, serverDeploymentBuilder *manifest.DeploymentBuilder, succesRate float64, protocol string) {
	_ = "STUB: not implemented"
	return
}
