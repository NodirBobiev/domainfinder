package whois

type Whois interface {
	Exists(domain string) (bool, error)
}
