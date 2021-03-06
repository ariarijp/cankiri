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

func getFieldNames() []string {
	return []string{
		"Host",
		"HostName",
		"User",
		"Port",
		"IdentityFile",
		"ProxyCommand",
		"SSHCmd",
	}
}

func toMap(h *sshconfig.SSHHost) map[string]interface{} {
	keys := getFieldNames()
	values := toArray(h)

	return map[string]interface{}{
		keys[0]: values[0],
		keys[1]: values[1],
		keys[2]: values[2],
		keys[3]: values[3],
		keys[4]: values[4],
		keys[5]: values[5],
		keys[6]: values[6],
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

func renderJSON(hosts []*sshconfig.SSHHost) {
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
}

func renderJSONL(hosts []*sshconfig.SSHHost) {
	for _, h := range hosts {
		fmt.Println(toJSON(h))
	}
}

func renderDSV(hosts []*sshconfig.SSHHost, sep *string) {
	fmt.Println(strings.Join(getFieldNames(), *sep))

	for _, h := range hosts {
		fmt.Println(toString(h, *sep))
	}
}

func main() {
	optFormat := flag.String("format", "tsv", "Format")
	optSep := flag.String("sep", "\t", "Separator")
	flag.Parse()

	hosts, err := sshconfig.ParseSSHConfig(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	switch *optFormat {
	case "json":
		renderJSON(hosts)
	case "jsonl":
		renderJSONL(hosts)
	default:
		renderDSV(hosts, optSep)
	}
}
