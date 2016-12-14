package xssh

func MakeXSSHConfig() *ClientConfig {
	ret := &ClientConfig{
		DontAuthenticate: true, // IOT scan ethically, never attempt to authenticate
		ClientVersion:    pkgConfig.ClientID,
	}

	return ret
}
