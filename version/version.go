package version

import (
	"runtime"
)

// const set at build time
const (
	UTCBuildTime = "undefined"
	GitCommit    = "undefined"
	OS           = runtime.GOOS
	Arch         = runtime.GOARCH
	GoVersion    = "undefined"
)

const (
	Version = "1.2.2"
)
