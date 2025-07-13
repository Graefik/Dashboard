# Graefik Dashboard

Un dashboard personnalisé type Grafana pour visualiser les métriques de Traefik, développé en Vue.js + Go.

## 🎯 Objectif

Graefik est un dashboard de monitoring qui récupère et affiche les métriques de Traefik de manière simple et élégante. Alternative légère à Grafana, spécialement conçue pour Traefik.

## 🏗️ Architecture

- **Frontend** : Vue.js 3 + TypeScript + Vite
- **Backend** : Go + Gin
- **Conteneurisation** : Docker + Docker Compose

## 🚀 Démarrage rapide

### Développement

```bash
# Lancer l'environnement de développement avec hot-reload
docker-compose -f docker-compose-dev.yml up --build
```

- Frontend : http://localhost:5173
- Backend : http://localhost:8080

## 📊 Roadmap

### V1 - Dashboard basique (En cours)

- [x] Setup projet Vue.js + Go
- [x] Configuration Docker development
- [ ] Récupération des métriques depuis `/metrics` de Traefik
- [ ] Affichage des métriques en dur pas de stockage côté back, par service traefik
- [ ] Login simple (utilisateur unique configuré via docker-compose like .env)
- [ ] Interface basique de visualisation
- [ ] Docker-compose de production et publié sur Docker Hub

### V2 - Dashboard avancé

- [ ] Intégration InfluxDB pour historique des métriques
- [ ] Système de cache pour les métriques
- [ ] Interface de personnalisation complète :
  - [ ] Déplacement des graphiques dans l'interface (drag & drop)
  - [ ] Changement des couleurs et légendes
  - [ ] Sélection des données à afficher
  - [ ] Création de tableaux de bord personnalisés
- [ ] Gestion multi-utilisateurs
- [ ] Export des données
- [ ] Alertes et notifications email

## 🔧 Structure du projet

Le projet suit deux architectures distinctes adaptées à chaque partie :

**Frontend (Vue.js)** : Architecture **Feature-Based**

- Organisation par fonctionnalités métier (dashboard, login)
- Chaque feature contient ses propres composants, services et types
- Composants partagés dans un dossier `shared`

**Backend (Go)** : Architecture **Hexagonale**

- Séparation claire entre le code métier et les adaptateurs externes
- Domain isolé des dépendances techniques
- Ports et adaptateurs pour une architecture modulaire et testable

## 🛠️ Technologies

### Frontend

- Vue.js 3
- TypeScript
- Vite
- SCSS

### Backend

- Go 1.24.5
- Gin (framework web)
- Air (hot-reload en développement)

### DevOps

- Docker & Docker Compose
- Traefik (pour la récupération des metrics)
