package index

import (
	"strings"

	"github.com/miekg/dns"
)

func normalize(domainName string) string {
	domainName = strings.ToLower(domainName)
	if !dns.IsFqdn(domainName) {
		domainName = dns.Fqdn(domainName)
	}
	return domainName
}
