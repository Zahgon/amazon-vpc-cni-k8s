package sgpp

const vlanInterfacePrefix = "vlan"

// BuildHostVethNamePrefix computes the name prefix for host-side veth pairs for SGPP pods
// for the "standard" mode, we use the same hostVethNamePrefix as normal pods, which is "eni" by default, but can be overwritten as well.
// for the "strict" mode, we use dedicated "vlan" hostVethNamePrefix, which is to opt-out SNAT support and opt-out calico's workload management.
func BuildHostVethNamePrefix(hostVethNamePrefix string, podSGEnforcingMode EnforcingMode) string {
	_ = "STUB: not implemented"
	return ""
}

// LoadEnforcingModeFromEnv tries to load the enforcing mode from environment variable and fall-back to DefaultEnforcingMode.
func LoadEnforcingModeFromEnv() EnforcingMode {
	_ = "STUB: not implemented"
	return *new(EnforcingMode)
}
