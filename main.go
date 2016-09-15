package main

import (
	"fmt"
	"os"

	"github.com/mikkeloscar/sshconfig"
)

func main() {
	hosts, err := sshconfig.ParseSSHConfig(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Host\tHost Name\tUser\tPort\tIdentity File\tSSH Command")
	for _, h := range hosts {
		host := h.HostName
		if host == "" {
			host = h.Host[0]
		}

		sshCmd := fmt.Sprintf("ssh -p %d %s@%s", h.Port, h.User, host)
		fmt.Printf("%s\t%s\t%s\t%d\t%s\t%s\n", h.Host[0], h.HostName, h.User, h.Port, h.IdentityFile, sshCmd)
	}
}
