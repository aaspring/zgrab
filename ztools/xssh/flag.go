package xssh

import (
	"flag"
)

type XSSHConfig struct {
	ClientID string
}

var pkgConfig XSSHConfig

func init() {
	flag.StringVar(&pkgConfig.ClientID, "xssh-client-id", packageVersion, "Specify the client ID string to use")
}
