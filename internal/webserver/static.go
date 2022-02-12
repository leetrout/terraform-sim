package webserver

import "embed"

// NB: Go 1.18 required

//go:embed all:static
var f embed.FS
