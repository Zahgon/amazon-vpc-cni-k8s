package cp

func TouchFile(filePath string) error { _ = "STUB: not implemented"; return nil }

func cp(src, dst string) error { _ = "STUB: not implemented"; return nil }

func CopyFile(src, dst string) (err error) { _ = "STUB: not implemented"; return nil }

func InstallBinaries(pluginBins []string, hostCNIBinPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func InstallBinariesFromDir(readDir string, hostCNIBinPath string, excludeBins map[string]bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Only copy files

// Exclude binaries in deny-list
