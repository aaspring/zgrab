package xssh

func MakeXSSHConfig() *ClientConfig {
	ret := &ClientConfig{
		DontAuthenticate:  true, // IOT scan ethically, never attempt to authenticate
		ClientVersion:     pkgConfig.ClientID,
		HostKeyAlgorithms: pkgConfig.HostKeyAlgorithms.GetStringSlice(),
	}

	return ret
}
