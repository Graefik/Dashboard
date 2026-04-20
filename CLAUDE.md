# Graefik — Guide Claude

Instructions d'architecture et de conventions pour ce repo. Lis et respecte
avant toute modification.

## 1. Le projet

**Graefik** est un dashboard d'observabilité **on-premise** pour Traefik.
Objectif : alternative légère à Grafana, spécialisée, qui se branche sur
`/metrics` de Traefik et restitue l'état du reverse-proxy en temps réel
(RPS, latence p50/p95/p99, status codes, routers, services, middlewares,
certificats).

- Déployable par un seul `docker compose up`.
- Pas de dépendance cloud : tout tourne sur la machine du client.
- Vocation élitiste : UX pro, typo `Geist` + `Geist Mono`, accent unique
  teal, charts haute perf (uPlot).

## 2. Stack

| Couche             | Techno                                              |
| ------------------ | --------------------------------------------------- |
| Frontend           | Vue 3 + TypeScript + Vite + SCSS                    |
| UI primitives      | Reka UI (headless, Dialog / DropdownMenu / …)       |
| Charts temps réel  | uPlot (timeline, percentiles)                       |
| Charts catégoriels | ApexCharts (donut, bar horizontal, sparklines KPI)  |
| Formulaires        | vee-validate + zod                                  |
| Routing            | vue-router                                          |
| Backend            | Go 1.24 + Gin                                       |
| ORM                | Gorm + Gormigrate                                   |
| DB (V1)            | MySQL ; (V2) + InfluxDB pour l'historique métriques |
| Dev                | Docker Compose + Air (hot-reload Go) + Vite HMR     |

## 3. Layout du repo

```
Dashboard/
├── client/                  # Front Vue
├── server/                  # Back Go
├── docker-compose-dev.yml
├── README.md
└── CLAUDE.md                # ← tu es ici
```

## 4. Backend — Architecture hexagonale

### Principe

Le **domaine** est le cœur, isolé de tout framework. Tout ce qui est
technique (HTTP, SQL, Traefik API, crypto) est un **adaptateur** branché
sur un **port** (interface). Les dépendances pointent **vers le domaine**,
jamais l'inverse.

```
┌─────────────────────────────────────────────────────┐
│                   adapters (driving)                │
│  ┌───────────────┐                                  │
│  │  HTTP (Gin)   │  →  ports.inbound  →  usecase    │
│  └───────────────┘                         │        │
│                                            ↓        │
│                                        domain       │
│                                            ↑        │
│  ┌───────────────┐                         │        │
│  │ Gorm · Traefik│  ←  ports.outbound ─────┘        │
│  │ client        │                                  │
│  └───────────────┘                                  │
│                  adapters (driven)                  │
└─────────────────────────────────────────────────────┘
```

### Arborescence cible

```
server/
├── cmd/
│   └── graefik/
│       └── main.go              # composition root : build les adapters,
│                                # injecte dans les use cases, démarre Gin
├── internal/
│   ├── domain/                  # entités + règles métier PURES
│   │   ├── user/
│   │   │   ├── user.go          # struct, value objects, invariants
│   │   │   └── errors.go
│   │   ├── metric/
│   │   │   ├── snapshot.go      # RPS, Latency, StatusCode…
│   │   │   └── aggregate.go
│   │   ├── router/
│   │   ├── service/
│   │   └── middleware/
│   │
│   ├── ports/                   # contrats (interfaces)
│   │   ├── inbound/             # use cases exposés (driving)
│   │   │   ├── auth.go
│   │   │   ├── overview.go
│   │   │   └── routers.go
│   │   └── outbound/            # dépendances requises (driven)
│   │       ├── user_repository.go
│   │       ├── traefik_client.go
│   │       └── clock.go
│   │
│   ├── app/                     # implémentations des use cases
│   │   ├── auth/
│   │   │   └── login.go         # implémente ports.inbound.Auth
│   │   ├── overview/
│   │   │   └── get_overview.go
│   │   └── routers/
│   │       └── list_routers.go
│   │
│   ├── adapters/
│   │   ├── http/                # ADAPTER INBOUND : Gin handlers
│   │   │   ├── router.go        # wiring des routes Gin
│   │   │   ├── auth_handler.go
│   │   │   ├── overview_handler.go
│   │   │   ├── middleware/
│   │   │   │   ├── auth.go      # JWT / session
│   │   │   │   └── logger.go
│   │   │   └── dto/             # objets de transport (requête/réponse)
│   │   │       └── overview.go
│   │   │
│   │   ├── persistence/
│   │   │   └── gorm/            # ADAPTER OUTBOUND : Gorm repos
│   │   │       ├── models/      # structs avec tags gorm (≠ domaine)
│   │   │       │   └── user.go
│   │   │       ├── migrations/  # gormigrate migrations
│   │   │       │   ├── 001_init.go
│   │   │       │   └── registry.go
│   │   │       ├── user_repo.go # implémente ports.outbound.UserRepository
│   │   │       └── db.go        # open, ping, shutdown
│   │   │
│   │   └── traefik/             # ADAPTER OUTBOUND : client API Traefik
│   │       ├── client.go        # implémente ports.outbound.TraefikClient
│   │       └── dto.go
│   │
│   ├── config/
│   │   └── config.go            # chargement env → struct typée
│   │
│   └── platform/                # utilitaires transverses (logger, errors…)
│       ├── logger/
│       └── errors/
│
├── go.mod
├── go.sum
├── .air.toml
└── Dockerfile.dev
```

### Règles dures (non négociables)

- `internal/domain/**` **n'importe RIEN** hors stdlib. Pas de `gorm`, pas
  de `gin`, pas de `context` provenant d'un framework, pas de tag ORM.
- `internal/ports/**` ne contient **que** des interfaces et des types
  d'entrée/sortie purs — aucune implémentation.
- `internal/app/**` dépend de `domain` + `ports.outbound`. Il n'importe
  **jamais** un adapter concret.
- `internal/adapters/**` dépendent de `ports` + `domain`. Ils
  **implémentent** les interfaces des ports.
- Le seul endroit où l'on câble tout ensemble est `cmd/graefik/main.go`
  (composition root). C'est là que `gorm.Open` est appelé, que les
  repositories sont instanciés, que les use cases reçoivent leurs deps,
  et que Gin est démarré.
- Un **modèle Gorm** (`adapters/persistence/gorm/models/*`) **n'est pas**
  une entité de domaine. Toujours mapper dans le repo.
- Migrations via **gormigrate** (`adapters/persistence/gorm/migrations/`).
  Pas de `AutoMigrate` en prod. Chaque migration a un ID daté.
- Erreurs métier dans `domain/*/errors.go` (ex: `ErrUserNotFound`).
  Les adapters **traduisent** en erreurs HTTP/SQL — jamais l'inverse.
- Un use case = une méthode publique par intention
  (`Login(ctx, email, password)`, `GetOverview(ctx)`).
- Tests :
  - `domain/**` : tests unitaires purs, zéro mock.
  - `app/**` : tests avec fakes des ports outbound (interfaces triviales
    à implémenter en mémoire).
  - `adapters/**` : tests d'intégration (Gorm contre MySQL jetable,
    handlers via `httptest`).

### Anti-patterns à refuser

- Handler Gin qui parle directement à Gorm.
- Struct de domaine avec `gorm:"primaryKey"`.
- Use case qui importe `gin.Context`.
- "Service" qui mélange parsing HTTP et règles métier.
- Retourner un modèle Gorm dans une réponse JSON.

## 5. Frontend — Architecture feature-based

### Principe

Une **feature** = une tranche verticale livrant une capacité métier
(dashboard, login, alerts…). Tout ce qui est propre à la feature vit
dedans. Ce qui est **transverse et générique** va dans `shared/`.

### Arborescence cible

```
client/src/
├── assets/                       # images, logos (icon.png, logo.png)
├── scss/
│   ├── _colors.scss              # DESIGN TOKENS (source unique de vérité)
│   └── layout.scss               # reset, typo globale, background,
│                                 # overrides Apex/uPlot globaux
├── router/
│   └── index.ts                  # routes (lazy imports des views)
│
├── features/
│   ├── dashboard/
│   │   ├── Dashboard.view.vue    # page principale
│   │   ├── components/           # composants UTILISÉS uniquement ici
│   │   │   ├── KpiRow.vue
│   │   │   ├── ServicesTable.vue
│   │   │   └── …
│   │   ├── service/              # couche d'appels API de la feature
│   │   │   └── overview.service.ts
│   │   ├── schema/               # zod schemas propres à la feature
│   │   ├── types/                # types domaine frontend
│   │   └── utils/
│   ├── login/
│   │   ├── Login.view.vue
│   │   ├── components/
│   │   │   └── login.form.vue
│   │   ├── service/
│   │   │   └── auth.service.ts
│   │   ├── schema/
│   │   │   └── login.schema.ts
│   │   ├── types/
│   │   └── utils/
│   └── <nouvelle-feature>/       # même moule exact
│
├── shared/
│   ├── ui/                       # PRIMITIVES génériques stylées
│   │   ├── Button.vue
│   │   ├── Input.vue
│   │   ├── Panel.vue             # conteneur cardable avec header/footer
│   │   ├── KpiCard.vue
│   │   ├── StatusPill.vue
│   │   └── UPlot.vue             # wrapper uPlot
│   ├── components/               # éléments de layout réutilisés
│   │   ├── Header.vue
│   │   ├── Sidebar.vue
│   │   ├── Footer.vue
│   │   ├── CommandPalette.vue
│   │   └── Wrapper.vue
│   ├── layouts/
│   │   └── MainLayout.vue        # shell : sidebar + topbar + statusbar
│   ├── api/                      # client HTTP partagé, interceptors
│   │   └── http.ts
│   ├── types/                    # types transversaux
│   └── utils/
│
├── App.vue
├── main.ts
└── vite-env.d.ts
```

### Règles dures

- **Isolation des features** : `features/a` ne peut **pas** importer
  depuis `features/b`. Si un besoin est partagé, il remonte dans
  `shared/`.
- **UI primitives** (Button, Input, Panel…) → **toujours** dans
  `shared/ui/`. Ne pas dupliquer un composant générique dans une feature.
- **Services** : chaque feature a son `service/` pour les appels API
  spécifiques. Le client HTTP de base (axios/fetch wrapper) vit dans
  `shared/api/`.
- **Schemas** : validation formulaire via zod, dans `<feature>/schema/`,
  consommés par vee-validate.
- **Types** : miroirs TS des DTO back dans `<feature>/types/` si propres
  à la feature, sinon `shared/types/`.
- **Alias `@/`** déjà configuré (vite.config.ts) → toujours utiliser
  `@/features/…` / `@/shared/…`, jamais de `../../../`.
- **Lazy loading** : chaque route charge sa view via dynamic import
  (`() => import("@/features/dashboard/Dashboard.view.vue")`).

### Règles de design (appliquées partout)

- **Tokens uniquement**. Jamais un hex en dur dans un composant. Tout
  passe par `scss/_colors.scss` (`$accent`, `$bg-surface`, `$text-muted`,
  `$radius-md`, `$ease-out`, …).
- **Typographie** : `Geist` (UI) + `Geist Mono` (nombres, code, labels
  techniques). Tous les nombres affichés en `font-mono` +
  `font-variant-numeric: tabular-nums`.
- **Accent unique** : `$accent` (teal `#00c2b2`). Jamais d'autre
  couleur UI chrome. Les statuts sémantiques (`success`/`warning`/
  `danger`/`info`) ne sont **pas** des accents — ils codent un état.
- **BEM en SCSS** : `.panel`, `.panel__header`, `.panel--loose`. Pas de
  descendant selectors profonds.
- **Pas d'emoji** dans le code ni les textes utilisateur. SVG inline ou
  rien.
- **Pas de glow décoratif** (box-shadow teal énorme). Les seules lueurs
  autorisées sont **fonctionnelles** (dot "live" qui pulse, état actif).
- **Pas de gradient text** sur gros titres. Hiérarchie par poids +
  couleur, pas par dégradé arc-en-ciel.

### Quel lib de chart utiliser ?

- **uPlot** (via `shared/ui/UPlot.vue`) dès qu'il s'agit de **séries
  temporelles** (>1 point par seconde, streaming, >500 points). C'est
  ce que fait Grafana.
- **ApexCharts** pour le reste : donut, bar horizontal, radial,
  sparklines courtes (<30 points) — Apex fait ça mieux out-of-the-box
  et le coût perf est négligeable.
- **Jamais Chart.js, Recharts, Highcharts**.

### Quel composant Reka utiliser ?

- Dialog / Popover / DropdownMenu / Combobox / Tooltip / Toast →
  **Reka UI**, stylé via nos tokens.
- Bannis : PrimeVue, Naive UI, Element Plus, Vuetify. Ils imposent
  leurs thèmes.

## 6. Conventions générales

### Git

- Branches : `feat/<slug>`, `fix/<slug>`, `chore/<slug>`, `refacto/…`.
- Commits en français, format `type : sujet court` (cohérent avec
  l'historique actuel : `feat : …`, `chore : …`).
- PR : une feature, une PR. Description explique le **pourquoi**, pas le
  quoi (le diff suffit).

### Secrets & config

- Tout passe par **variables d'environnement** lues au boot
  (`server/internal/config/config.go`).
- Jamais de secret commit. `.env` est ignoré.
- Pour le front : préfixe `VITE_` uniquement (les autres ne sont pas
  exposées).

### Tests

- Back : `go test ./...` doit être vert avant merge.
- Front : pas de stack de tests en V1 ; à introduire en V2 (Vitest +
  Testing Library). En attendant, **tester manuellement** via
  `pnpm dev` sur chaque changement UI.

### Documentation

- Pas de commentaires qui paraphrasent le code. Un commentaire
  n'existe que pour expliquer le **pourquoi** non évident (contrainte,
  workaround, invariant caché).
- Pas de docstring multi-paragraphes.
- Documentation produit dans `README.md`. Documentation agent dans
  `CLAUDE.md`. C'est tout.

## Principes de code - Obligatoire

Respecter les principes suivants dans tout le code :

- **DRY** (Don't Repeat Yourself) : pas de duplication, extraire les logiques communes
- **KISS** (Keep It Simple, Stupid) : solutions simples, pas de sur-ingenierie
- **SOLID** : responsabilite unique, ouvert/ferme, substitution, segregation d'interfaces, inversion de dependances
- **Pas de magic strings/numbers** : utiliser des constantes nommees (`const MAX_RETRIES = 3`, pas `3` en dur dans le code)

## 7. Commandes utiles

```bash
# Dev full stack
docker compose -f docker-compose-dev.yml up --build

# Front seul
cd client && pnpm install && pnpm dev

# Front : formatter
cd client && pnpm format

# Back seul (avec hot-reload)
cd server && air

# Back : build
cd server && go build ./cmd/graefik

# Back : tests
cd server && go test ./...
```

Frontend : `http://localhost:5173` · Backend : `http://localhost:8080`
· Adminer : `http://localhost:8081`.

## 8. Ce qu'il ne faut **jamais** faire

- Mélanger les couches hexagonales côté back.
- Importer d'une feature à une autre côté front.
- Introduire une lib "à la mode" (Tailwind, Chakra, Element Plus…)
  sans accord explicite.
- Utiliser un hex couleur en dur dans un composant.
- Committer un mot de passe, une clé, un token.
- Supprimer ou contourner les migrations gormigrate pour "aller plus
  vite".
- Ajouter une dépendance front sans vérifier qu'une alternative déjà
  installée ne fait pas le job (check `client/package.json`).
