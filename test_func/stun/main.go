package main

// go get "github.com/ccding/go-stun/stun"
func NewExternalIP() (string, error) {
	client := stun.NewClient()
	client.SetServerAddr(stun.DefaultServerAddr)
	client.SetVerbose(true)  //for debug
	client.SetVVerbose(true) //for debug
	nat, host, err := client.Discover()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("NAT Type:", nat)
	if host != nil {
		fmt.Println("External IP Family:", host.Family())
		fmt.Println("External IP:", host.IP())
		fmt.Println("External Port:", host.Port())
		return host.IP(), nil
	}
	return "", errors.New("Some error")
}
