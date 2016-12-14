package xssh

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

var pkgConfig XSSHConfig

type XSSHConfig struct {
	ClientID          string
	HostKeyAlgorithms HostKeyAlgorithmsList
}

type HostKeyAlgorithmsList struct {
	IsSet      bool
	Algorithms []string
}

func (hkaList *HostKeyAlgorithmsList) String() string {
	return "BROKEN HostKeyAlgorithmsList.String()"
}

func (hkaList *HostKeyAlgorithmsList) Set(value string) error {
	hkaList.IsSet = true
	for _, alg := range strings.Split(value, ",") {
		isValid := false
		for _, val := range supportedHostKeyAlgos {
			if val == alg {
				isValid = true
				break
			}
		}

		if !isValid {
			return errors.New(fmt.Sprintf(`Can not support host key algorithm : "%s"`, alg))
		}

		hkaList.Algorithms = append(hkaList.Algorithms, alg)
	}
	return nil
}

func (hkaList *HostKeyAlgorithmsList) GetStringSlice() []string {
	if !hkaList.IsSet {
		return supportedHostKeyAlgos
	} else {
		return hkaList.Algorithms
	}
}

func init() {
	flag.StringVar(&pkgConfig.ClientID, "xssh-client-id", packageVersion, "Specify the client ID string to use")

	hostKeyAlgUsage := fmt.Sprintf("Which host key algorithms to support (default \"%s\")", strings.Join(supportedHostKeyAlgos, ","))
	flag.Var(&pkgConfig.HostKeyAlgorithms, "xssh-host-key-algs", hostKeyAlgUsage)
}
