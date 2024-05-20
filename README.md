# goclean

```bash
$ docker run -d --name goclean --env=POSTGRES_USER=goclean --env=POSTGRES_PASSWORD=goclean1! --env=POSTGRES_DB=goclean --env=TIMEZONE=Asia/Seoul -p 5432:5432 postgres
```

```bash
$ go run ./cmd/server

$ go install github.com/google/wire/cmd/wire@latest
$ wire ./cmd/server

$ swag init -pd -g ./cmd/server/main.go -o ./docs/

$ go install golang.org/x/tools/cmd/goimports@latest
$ goimports -w .

├── .gitignore
├── .vscode
│  └── settings.json
├── cmd
│  └── server
│    ├── main.go
│    ├── wire_gen.go
│    └── wire.go
├── docs
│  ├── docs.go
│  ├── swagger.json
│  └── swagger.yaml
├── go.mod
├── go.sum
├── infra
│  ├── config
│  │  ├── config.go
│  │  └── config.json
│  ├── database
│  │  ├── database.go
│  │  └── repository
│  │    └── memo_repo.go
│  ├── infra.go
│  └── router
│    ├── handler
│    │  └── memo_handler.go
│    ├── middleware
│    │  └── middleware.go
│    └── router.go
├── internal
│  ├── controller
│  │  ├── controller.go
│  │  └── rest
│  │    ├── dto
│  │    │  ├── memo_dto
│  │    │  │  ├── memo_create_dto.go
│  │    │  │  ├── memo_delete_dto.go
│  │    │  │  ├── memo_find_all_dto.go
│  │    │  │  ├── memo_find_one_dto.go
│  │    │  │  └── memo_update_dto.go
│  │    │  └── pagination_dto.go
│  │    └── memo_controller.go
│  ├── entity
│  │  └── memo.go
│  └── usecase
│    ├── memo
│    │  └── memo_ucase.go
│    └── usecase.go
├── LICENSE
└── README.md
```

#### swagger
- `URL` http://localhost:5000/v1/docs/index.html
