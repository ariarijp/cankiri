package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
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

func toArray(h *sshconfig.SSHHost) []string {
	return []string{
		h.Host[0],
		h.HostName,
		h.User,
		strconv.Itoa(h.Port),
		h.IdentityFile,
		h.ProxyCommand,
		toSSHCmdString(h),
	}
}

func toMap(h *sshconfig.SSHHost) map[string]interface{} {
	arr := toArray(h)

	return map[string]interface{}{
		"Host":         arr[0],
		"HostName":     arr[1],
		"User":         arr[2],
		"Port":         arr[3],
		"IdentityFile": arr[4],
		"ProxyCommand": arr[5],
		"SSHCmd":       arr[6],
	}
}

func toString(h *sshconfig.SSHHost, sep string) string {
	return strings.Join(toArray(h), sep)
}

func toJSON(h *sshconfig.SSHHost) string {
	b, err := json.Marshal(toMap(h))
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func main() {
	optFormat := flag.String("format", "tsv", "Format")
	optSep := flag.String("sep", "\t", "Separator")
	flag.Parse()

	hosts, err := sshconfig.ParseSSHConfig(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	if *optFormat == "json" {
		fmt.Println("[")

		for i, h := range hosts {
			fmt.Print("  ", toJSON(h))

			if i == len(hosts)-1 {
				fmt.Println("")
			} else {
				fmt.Println(",")
			}
		}

		fmt.Println("]")
	} else if *optFormat == "jsonl" {
		for _, h := range hosts {
			fmt.Println(toJSON(h))
		}
	} else {
		fmt.Println(strings.Join([]string{
			"Host",
			"Host Name",
			"User",
			"Port",
			"Identity File",
			"Proxy Command",
			"SSH Command",
		}, *optSep))

		for _, h := range hosts {
			fmt.Println(toString(h, *optSep))
		}
	}
}
