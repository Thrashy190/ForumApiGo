# ForumApiGo

Backend api part for a fullstack project with Go and ReactJS

Frontend part [here] (https://github.com/Thrashy190)

## Backend api part

The tools we use are:

    - JWT for auth authorization
    - GORM for database modeling and connection
    - Gorilla mux for http router
    - Docker container for lauch db
    - PostgreSQL for database

### Docker config

    docker run --name forum-go-app -e POSTGRES_PASSWORD=forumgoapp -e POSTGRES_USER=forum-go -e POSTGRES_DB=forum-go -p 5433:5433 -d postgres
