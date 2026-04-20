package traefik

// Router est une règle de routing Traefik (http.routers.<name>).
type Router struct {
	Name        string
	Rule        string
	EntryPoints []string
	Service     string
	Middlewares []string
	Status      string // enabled | disabled | warning
	Using       []string
	Provider    string
	TLS         *RouterTLS
}

type RouterTLS struct {
	CertResolver string
	Domains      []CertDomain
}

type CertDomain struct {
	Main string
	SANs []string
}

// Service est un backend cible (http.services.<name>).
type Service struct {
	Name         string
	Type         string
	Provider     string
	Status       string
	UsedBy       []string
	ServerStatus map[string]string // URL → UP/DOWN
}
