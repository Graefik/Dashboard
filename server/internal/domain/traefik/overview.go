package traefik

// Overview résume l'état de l'instance Traefik à un instant T.
// Pur : pas de tag json, pas d'ORM. L'adapter traduit le JSON Traefik vers ça.
type Overview struct {
	HTTP     Section
	TCP      Section
	UDP      Section
	Features Features
	Version  string
}

type Section struct {
	Routers     Counts
	Services    Counts
	Middlewares Counts
}

type Counts struct {
	Total    int
	Warnings int
	Errors   int
}

type Features struct {
	Tracing    string
	Metrics    string
	AccessLog  bool
	Hub        bool
}
