package dto

type OverviewResponse struct {
	HTTP     SectionDTO  `json:"http"`
	TCP      SectionDTO  `json:"tcp"`
	UDP      SectionDTO  `json:"udp"`
	Features FeaturesDTO `json:"features"`
	Version  string      `json:"version"`
}

type SectionDTO struct {
	Routers     CountsDTO `json:"routers"`
	Services    CountsDTO `json:"services"`
	Middlewares CountsDTO `json:"middlewares"`
}

type CountsDTO struct {
	Total    int `json:"total"`
	Warnings int `json:"warnings"`
	Errors   int `json:"errors"`
}

type FeaturesDTO struct {
	Tracing   string `json:"tracing"`
	Metrics   string `json:"metrics"`
	AccessLog bool   `json:"accessLog"`
	Hub       bool   `json:"hub"`
}

type RouterDTO struct {
	Name        string         `json:"name"`
	Rule        string         `json:"rule"`
	EntryPoints []string       `json:"entryPoints"`
	Service     string         `json:"service"`
	Middlewares []string       `json:"middlewares"`
	Status      string         `json:"status"`
	Using       []string       `json:"using"`
	Provider    string         `json:"provider"`
	TLS         *RouterTLSDTO  `json:"tls,omitempty"`
}

type RouterTLSDTO struct {
	CertResolver string          `json:"certResolver"`
	Domains      []CertDomainDTO `json:"domains,omitempty"`
}

type CertDomainDTO struct {
	Main string   `json:"main"`
	SANs []string `json:"sans,omitempty"`
}

type ServiceDTO struct {
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	Provider     string            `json:"provider"`
	Status       string            `json:"status"`
	UsedBy       []string          `json:"usedBy,omitempty"`
	ServerStatus map[string]string `json:"serverStatus,omitempty"`
}
