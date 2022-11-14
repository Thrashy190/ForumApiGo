# ForumApiGo

Backend api part for a fullstack project with Go and ReactJS

Frontend part [here](https://github.com/Thrashy190)

## Backend api part

The tools we use are:

- JWT for auth authorization
- bcrypt for password encryption
- GORM for database modeling and connection
- Gorilla mux for http router
- Docker container for lauch db
- PostgreSQL for database

### Docker config

```bash
docker run --name forum-go-app
    -e POSTGRES_PASSWORD=postgrespw
    -e POSTGRES_USER=postgres
    -e POSTGRES_DB=postgres
    -p 55002:5432
    -d postgres
```

### Run PGadmin server

```bash
docker run -p 5050:80
-e 'PGADMIN_DEFAULT_EMAIL=pgadmin4@pgadmin.org'
-e 'PGADMIN_DEFAULT_PASSWORD=admin'
-d
--name pgadmin4 dpage/pgadmin4
```
