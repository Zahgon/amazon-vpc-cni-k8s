package networkutils

// GeneratePodHostVethName generates the name for Pod's host-side veth device.
// The veth name is generated in a way that aligns with the value expected by Calico for NetworkPolicy enforcement.
func GeneratePodHostVethName(prefix string, podNamespace string, podName string, index int) string {
	_ = "STUB: not implemented"
	return ""
}

// GeneratePodHostVethNameSuffix generates the name suffix for Pod's hostVeth.
func GeneratePodHostVethNameSuffix(podNamespace string, podName string) string {
	_ = "STUB: not implemented"
	return ""
}

// Generates the interface name inside the pod namespace
func GenerateContainerVethName(defaultIfName string, prefix string, index int) string {
	_ = "STUB: not implemented"
	return ""
}
