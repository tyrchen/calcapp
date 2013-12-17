package network

import (
	"bufio"
	"encoding/json"
	"fmt"
	osutil "github.com/tyrchen/goutil/osutil"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
)

type KVPair map[string]string
type IpMacTable []KVPair

const (
	IP_PREFIX = "192.168.0."
	IP_MAX    = 30
	ARP_REGEX = `.*?\s+\((?P<ip>[\d\.]+)\)\s+at\s+(?P<mac>[a-zA-Z\d:]+)`
	DATA_PATH = "data/mac/arp.dat"
)

func sendPing() (seen map[string]bool) {
	seen = make(map[string]bool, IP_MAX)

	for i := 1; i <= IP_MAX; i++ {
		ip := fmt.Sprintf("%s%d", IP_PREFIX, i)
		if exec.Command("ping", "-c1", "-W1", ip).Run() == nil {
			seen[ip] = true
		}
	}
	return
}

func parseArpTable() (captures IpMacTable) {
	captures = make(IpMacTable, 0)
	data, err := exec.Command("arp", "-a").Output()
	if err != nil {
		fmt.Printf("Error running arp command: %v\n", err)
		os.Exit(1)
	}

	regex := regexp.MustCompile(ARP_REGEX)
	names := regex.SubexpNames()
	matches := regex.FindAllStringSubmatch(string(data), -1)
	for _, match := range matches {
		cmap := make(KVPair)
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

func saveArpTable(filename string) {
	ips := parseArpTable()
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	s := bufio.NewWriter(f)
	b, _ := json.Marshal(ips)
	s.Write(b)
	s.Flush()
}

func loadArpTable(filename string) (ret IpMacTable) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(bytes, &ret)
	return
}

func GetIpMac() IpMacTable {
	filename := DATA_PATH
	if !osutil.FileExists(filename) {
		saveArpTable(filename)
	}

	return loadArpTable(filename)
}
