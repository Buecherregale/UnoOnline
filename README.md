# UnoOnline

Ein vollständig funktionsfähiges Online-Uno-Spiel mit Echtzeit-Multiplayer-Funktionalität, entwickelt mit Go (Backend) und Nuxt.js (Frontend).

## 📋 Inhaltsverzeichnis

- [Projektübersicht](#projektübersicht)
- [Features](#features)
- [Technologie-Stack](#technologie-stack)
- [Architektur](#architektur)
- [Installation & Setup](#installation--setup)
- [API-Dokumentation](#api-dokumentation)
- [Frontend-Struktur](#frontend-struktur)
- [Backend-Struktur](#backend-struktur)
- [Entwicklung](#entwicklung)
- [Docker-Deployment](#docker-deployment)

## 🎯 Projektübersicht

UnoOnline ist ein webbasiertes Multiplayer-Uno-Kartenspiel, das es Spielern ermöglicht, Räume zu erstellen oder beizutreten und in Echtzeit gegeneinander zu spielen. Das Projekt verwendet eine moderne Microservice-Architektur mit separaten Frontend- und Backend-Services.

## ✨ Features

### Spielerfunktionen
- **Benutzerregistrierung**: Spieler können sich mit einem Namen registrieren und erhalten eine eindeutige UUID
- **Raum-Management**: Erstellen und Beitreten von Spielräumen
- **Echtzeit-Kommunikation**: WebSocket-basierte Kommunikation für Live-Spielupdates
- **Responsive Design**: Vollständig responsive Benutzeroberfläche mit Tailwind CSS

### Spielmechaniken
- **Raum-Erstellung**: Spieler können eigene Räume hosten
- **Raum-Beitritt**: Beitritt zu bestehenden Räumen über Raum-ID
- **Spieler-Verwaltung**: Automatisches Hinzufügen/Entfernen von Spielern
- **Verlassen-Schutz**: Warnung beim Verlassen einer aktiven Lobby

## 🛠 Technologie-Stack

### Backend
- **Go 1.23.1**: Hauptprogrammiersprache
- **Gorilla WebSocket**: WebSocket-Implementierung für Echtzeit-Kommunikation
- **UUID**: Eindeutige Identifikation von Spielern und Räumen
- **In-Memory Storage**: Temporäre Datenspeicherung

### Frontend
- **Nuxt.js 3.15.0**: Vue.js-Framework für SSR/SPA
- **Vue 3**: Reactive Frontend-Framework
- **TypeScript**: Typisierte JavaScript-Entwicklung
- **Tailwind CSS**: Utility-First CSS-Framework
- **Vue Router**: Client-seitiges Routing

### Infrastructure
- **Docker**: Containerisierung aller Services
- **Docker Compose**: Multi-Container-Orchestrierung
- **PostgreSQL 17.2**: Datenbank (vorbereitet für zukünftige Persistierung)

## 🏗 Architektur

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│                 │    │                 │    │                 │
│    Frontend     │◄──►│     Backend     │◄──►│   PostgreSQL    │
│   (Nuxt.js)     │    │      (Go)       │    │   (Database)    │
│   Port: 3000    │    │   Port: 8080    │    │   Port: 5432    │
│                 │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │
         │                       │
         └───────WebSocket───────┘
```

### Kommunikationsflow
1. **HTTP REST API**: Raum- und Spielerverwaltung
2. **WebSocket**: Echtzeit-Spielkommunikation
3. **Cookie-basierte Sessions**: Spieleridentifikation

## 🚀 Installation & Setup

### Voraussetzungen
- Docker & Docker Compose
- Git

### Schnellstart
```bash
# Repository klonen
git clone <repository-url>
cd UnoOnline

# Mit Docker Compose starten
docker-compose up --build

# Anwendung öffnen
# Frontend: http://localhost:3000
# Backend API: http://localhost:8080
```

### Entwicklungssetup

#### Backend (Go)
```bash
cd backend
go mod download
go run main.go
```

#### Frontend (Nuxt.js)
```bash
cd frontend/UnoFrontend
npm install
npm run dev
```

## 📡 API-Dokumentation

### REST Endpoints

#### Spieler-Endpoints
- **POST** `/player` - Neuen Spieler erstellen
  ```json
  {
    "name": "Spielername"
  }
  ```

#### Raum-Endpoints
- **POST** `/room` - Neuen Raum erstellen
- **GET** `/room/{id}` - Raum-Informationen abrufen
- **POST** `/room/{id}` - Spiel starten (nur Besitzer)
- **POST** `/room/{id}/players` - Raum beitreten
- **DELETE** `/room/{id}/players` - Raum verlassen

#### WebSocket
- **Endpoint**: `/ws`
- **Parameter**: `roomId`, `playerId`
- **Verwendung**: Echtzeit-Spielkommunikation

### Beispiel-Requests
```bash
# Spieler erstellen
curl -X POST http://localhost:8080/player \
  -H "Content-Type: application/json" \
  -d '{"name":"Klaus"}'

# Raum erstellen
curl -X POST http://localhost:8080/room \
  -H "Content-Type: application/json" \
  -d '{"id":"player-uuid"}'
```

## 🎨 Frontend-Struktur

### Seiten
- **`/login`**: Spielerregistrierung
- **`/hostOrJoin`**: Raum erstellen oder beitreten
- **`/lobby-[id]`**: Spielraum-Lobby

### Middleware
- **`checkUUID.global.ts`**: Globale Authentifizierungsprüfung
- **`checkJoin.ts`**: Lobby-Beitrittsvalidierung
- **`leaveIntersect.global.ts`**: Warnung beim Verlassen der Lobby

### Utilities
- **`models.ts`**: TypeScript-Typdefinitionen
- **`getIDFromCookie.ts`**: Spieler-ID aus Cookies extrahieren
- **`roomFetches.ts`**: API-Kommunikationshilfsfunktionen

## ⚙ Backend-Struktur

### Hauptkomponenten
- **`main.go`**: Server-Einstiegspunkt und Routing
- **`websocket.go`**: WebSocket-Handler

### API-Module
- **`controller/`**: HTTP-Request-Handler
- **`models/`**: Datenstrukturen (Player, Room, Card)
- **`data/`**: In-Memory-Datenspeicherung
- **`ws/`**: WebSocket-Server-Implementierung

### Datenmodelle
```go
type Player struct {
    Id   uuid.UUID `json:"id"`
    Name string    `json:"name"`
}

type Room struct {
    Id      uuid.UUID `json:"id"`
    Players []Player  `json:"players"`
    Owner   Player    `json:"owner"`
}
```

## 🔧 Entwicklung

### Frontend-Entwicklung
```bash
cd frontend/UnoFrontend
npm run dev          # Entwicklungsserver
npm run build        # Produktions-Build
npm run generate     # Statische Site-Generation
```

### Backend-Entwicklung
```bash
cd backend
go run main.go       # Server starten
go test ./...        # Tests ausführen
```

### Testdaten
Das Backend enthält vorgenerierte Testdaten:
- **Testspieler**: Klaus, Biggie Smalls, Wilhelm
- **Testraum**: `4d3e97bf-cc2e-4af0-9397-2a0e3b331c6f`

### API-Tests
HTTP-Testdateien befinden sich in `backend/api/controller/controller_tests/`:
- `create_player.http`
- `create_room.http`
- `join_room.http`
- `leave_room.http`
- `start.http`

## 🐳 Docker-Deployment

### Produktions-Deployment
```bash
# Alle Services starten
docker-compose up -d

# Logs anzeigen
docker-compose logs -f
```

### Umgebungsvariablen
- **API_BASE_URL**: Backend-URL für Frontend-Kommunikation
- **POSTGRES_PASSWORD**: Datenbank-Passwort

### Port-Mapping
- **Frontend**: 3000
- **Backend**: 8080
- **PostgreSQL**: 5432


