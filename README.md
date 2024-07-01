## Graphql-service

### Options

- Create posts
- Find posts and pagination
- Disable comments for posts
- Create comments
- Comment text limitation 2000 characters

## Tools

- Go 1.21
- GraphQL
- PostgreSQL 16
- Docker compose

## Getting Started

1. Setup configuration .env 

```env
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_HOST=postgres
DB_PORT=5432
IN_MEMORY_MODE=false

HTTP_PORT=8080
```

2. Run
```bash
docker compose up --build 
```

