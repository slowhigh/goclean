# goclean

```bash
$ docker run -d --name goclean --env=POSTGRES_USER=goclean --env=POSTGRES_PASSWORD=goclean1! --env=POSTGRES_DB=goclean --env=TIMEZONE=Asia/Seoul -p 5432:5432 postgres
```

```bash
$ go run ./cmd/server

$ wire ./cmd/server

$ swag init -pd -g ./cmd/server/main.go -o ./docs/
```

#### swagger
- `URL` http://localhost:5000/v1/docs/index.html