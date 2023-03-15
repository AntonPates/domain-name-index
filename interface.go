package domiannameindex

// Interface is an interface for domain name index.
type Interface interface {
	Insert(domainName string)
	Find(domainName string) (ok bool, fullPath string)
	Remove(domainName string)
	Print(prefix string)
}
