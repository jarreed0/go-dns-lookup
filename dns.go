package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/miekg/dns"
)

func generate(t uint16) *dns.Msg {
	mTemp := new(dns.Msg)
	mTemp.SetQuestion(dns.Fqdn(os.Args[1]), t)
	mTemp.RecursionDesired = true
	return mTemp
}

func createR(m *dns.Msg, c *dns.Client, config dns.ClientConfig) *dns.Msg {
	r, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
	if r == nil {
		log.Fatalf("*** error: %s\n", err.Error())
	}
	return r
}

func printInfo(r *dns.Msg) {
	for _, a := range r.Answer {
		fmt.Printf("%v\n", a)
		//words := string(%v)
	}
}

func query(s string, ui uint16, c *dns.Client, config dns.ClientConfig) {
	//print(s)
	r := createR(generate(ui), c, config)
	printInfo(r)
}

func main() {
	recs := map[string]uint16{
		"a":       dns.TypeA,
		"aaaa":    dns.TypeAAAA,
		"cname":   dns.TypeCNAME,
		"mx":      dns.TypeMX,
		"ns":      dns.TypeNS,
		"ptr":     dns.TypePTR,
		"soa":     dns.TypeSOA,
		"srv":     dns.TypeSRV,
		"txt":     dns.TypeTXT,
		"dnskey":  dns.TypeDNSKEY,
		"ds":      dns.TypeDS,
		"nsec":    dns.TypeNSEC,
		"nsec3":   dns.TypeNSEC3,
		"rrsig":   dns.TypeRRSIG,
		"afsdb":   dns.TypeAFSDB,
		"atma":    dns.TypeATMA,
		"caa":     dns.TypeCAA,
		"cert":    dns.TypeCERT,
		"dhcid":   dns.TypeDHCID,
		"dname":   dns.TypeDNAME,
		"hinfo":   dns.TypeHINFO,
		"isdn":    dns.TypeISDN,
		"loc":     dns.TypeLOC,
		"mb":      dns.TypeMB,
		"mg":      dns.TypeMG,
		"minfo":   dns.TypeMINFO,
		"mr":      dns.TypeMR,
		"naptr":   dns.TypeNAPTR,
		"nsapptr": dns.TypeNSAPPTR,
		"rp":      dns.TypeRP,
		"rt":      dns.TypeRT,
		"tlsa":    dns.TypeTLSA,
		"x25":     dns.TypeX25,
	}

	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	c := new(dns.Client)
	recType := "a"
	if len(os.Args) >= 3 {
		res := strings.ToLower(os.Args[2])
		if res[0:1] == "@" {
			config.Servers[0] = res[1:]
			if len(os.Args) >= 4 {
				recType = strings.ToLower(os.Args[3])
			}
		} else {
			recType = res
		}
	}
	query(recType, recs[recType], c, *config)
}
