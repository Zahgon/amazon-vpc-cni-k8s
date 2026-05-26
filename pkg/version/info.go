package version

import (
	"runtime"
)

// Build information. Populated at build-time.
var (
	Version   string
	GoVersion = runtime.Version()
)

func RegisterMetric() { _ = "STUB: not implemented"; return }
