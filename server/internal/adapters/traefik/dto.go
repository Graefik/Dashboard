package traefikclient

// DTO : miroir exact du JSON Traefik. Jamais exposé hors du package.

type overviewDTO struct {
	HTTP     sectionDTO  `json:"http"`
	TCP      sectionDTO  `json:"tcp"`
	UDP      sectionDTO  `json:"udp"`
	Features featuresDTO `json:"features"`
}

type sectionDTO struct {
	Routers     countsDTO `json:"routers"`
	Services    countsDTO `json:"services"`
	Middlewares countsDTO `json:"middlewares"`
}

type countsDTO struct {
	Total    int `json:"total"`
	Warnings int `json:"warnings"`
	Errors   int `json:"errors"`
}

type featuresDTO struct {
	Tracing   string `json:"tracing"`
	Metrics   string `json:"metrics"`
	AccessLog bool   `json:"accessLog"`
	Hub       bool   `json:"hub"`
}

type versionDTO struct {
	Version   string `json:"Version"`
	Codename  string `json:"Codename"`
	StartDate string `json:"startDate"`
}

type routerDTO struct {
	Name        string      `json:"name"`
	Rule        string      `json:"rule"`
	EntryPoints []string    `json:"entryPoints"`
	Service     string      `json:"service"`
	Middlewares []string    `json:"middlewares"`
	Status      string      `json:"status"`
	Using       []string    `json:"using"`
	Provider    string      `json:"provider"`
	TLS         *routerTLSDTO `json:"tls,omitempty"`
}

type routerTLSDTO struct {
	CertResolver string          `json:"certResolver"`
	Domains      []certDomainDTO `json:"domains"`
}

type certDomainDTO struct {
	Main string   `json:"main"`
	SANs []string `json:"sans"`
}

type serviceDTO struct {
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	Provider     string            `json:"provider"`
	Status       string            `json:"status"`
	UsedBy       []string          `json:"usedBy"`
	ServerStatus map[string]string `json:"serverStatus"`
}
