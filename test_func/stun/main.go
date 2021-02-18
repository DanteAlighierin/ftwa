package test

import (
	"fmt"
	"testing"

	"github.com/ccding/go-stun/stun"
)

func test(t *testing.T) {
	client := stun.NewClient()
	client.SetServerAddr("stun1.l.google.com:19302")
	// client.SetServerAddr(stun.DefaultServerAddr)
	client.SetVerbose(true)  //for debug
	client.SetVVerbose(true) //for debug
	nat, host, err := client.Discover()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("NAT Type:", nat)
	if host != nil {
		fmt.Println("External IP Family:", host.Family())
		fmt.Println("External IP:", host.IP())
		fmt.Println("External Port:", host.Port())
		return
	}
	return
}
