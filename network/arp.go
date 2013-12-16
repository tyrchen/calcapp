package network

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

const (
	IP_PREFIX = "192.168.0."
	IP_MAX    = 30
	ARP_REGEX = `.*?\s+\((?P<ip>[\d\.]+)\)\s+at\s+(?P<mac>[a-zA-Z\d:]+)`
)

func SendPing() (seen map[string]bool) {
	seen = make(map[string]bool, IP_MAX)

	for i := 1; i <= IP_MAX; i++ {
		ip := fmt.Sprintf("%s%d", IP_PREFIX, i)
		if exec.Command("ping", "-c1", "-W1", ip).Run() == nil {
			seen[ip] = true
		}
	}
	return
}

func ParseArpTable() (captures []map[string]string) {
	captures = make([]map[string]string, 0)
	data, err := exec.Command("arp", "-a").Output()
	if err != nil {
		fmt.Printf("Error running arp command: %v\n", err)
		os.Exit(1)
	}

	regex := regexp.MustCompile(ARP_REGEX)
	names := regex.SubexpNames()
	matches := regex.FindAllStringSubmatch(string(data), -1)
	for _, match := range matches {
		cmap := make(map[string]string)
		for pos, val := range match {
			name := names[pos]
			if name != "" {
				cmap[name] = val
			}
		}
		captures = append(captures, cmap)
	}
	return
}
