# Go Book API

Une API REST simple en Go pour gérer des livres, utilisant PostgreSQL sans ORM.  
Les opérations CRUD sont implémentées avec des requêtes SQL via `database/sql` et le driver `pgx`/`lib/pq`.  
Le routage est assuré par [Chi](https://github.com/go-chi/chi).

---

## Fonctionnalités

- Créer, lire, mettre à jour et supprimer des livres
- Endpoint de vérification de santé (health check)
- Connexion à PostgreSQL via variables d'environnement
- Requêtes SQL uniquement (pas d'ORM)
- Échanges au format JSON

---

## Stack technique

- **Langage :** Go 1.25+
- **Base de données :** PostgreSQL
- **Routage :** Chi v5
- **Accès DB :** database/sql + lib/pq
- **Migrations :** fichiers SQL

---

## Démarrage

### 1. Cloner le dépôt

```bash
git clone https://github.com/nassim-touissi/go-book-api.git
cd go-book-api
```

### 2. Démarrer PostgreSQL avec Docker

```bash
docker-compose up -d
```

•	Base : booksdb
•	Utilisateur : admin
•	Port : 5434

### 3. Appliquer les migrations

```bash
psql -h localhost -p 5434 -U admin -d booksdb -f migrations/001_create_books.up.sql
```

### 4. Définir les variables d’environnement

```bash
export DB_HOST=localhost
export DB_PORT=5434
export DB_USER=admin
export DB_PASSWORD=*****
export DB_NAME=booksdb
```

### 5. Lancer l’API

```bash
go run ./cmd/api/main.go
```

Edité par TOUISSI Nassim
M1 Dev&Data Student
H3 Hitema 