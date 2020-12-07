package app

import (
	"runtime"
)

// Application version information overriden during compilation.
var (
	Name    = "app"
	Version = "0.0.0"
	Commit  = "0000000"
	Date    = "1970-01-01 00:00:00"

	// Never overriden - will always use runtime information.
	Golang = runtime.Version()
	Arch   = runtime.GOOS + "/" + runtime.GOARCH
)
