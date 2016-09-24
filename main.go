package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mikkeloscar/sshconfig"
)

func toSSHCmdString(h *sshconfig.SSHHost) string {
	host := h.HostName
	if host == "" {
		host = h.Host[0]
	}

	return fmt.Sprintf("ssh -p %d %s@%s", h.Port, h.User, host)
}

func toString(h *sshconfig.SSHHost) string {
	return strings.Join([]string{
		h.Host[0],
		h.HostName,
		h.User,
		strconv.Itoa(h.Port),
		h.IdentityFile,
		h.ProxyCommand,
		toSSHCmdString(h),
	}, "\t")
}

func main() {
	hosts, err := sshconfig.ParseSSHConfig(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Join([]string{
		"Host",
		"Host Name",
		"User",
		"Port",
		"Identity File",
		"Proxy Command",
		"SSH Command",
	}, "\t"))

	for _, h := range hosts {
		fmt.Println(toString(h))
	}
}
