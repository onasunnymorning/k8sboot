# gin-template
template repo with boilerplate gin and containerization


# Notes

## Dir structure
```
/my-go-app
├── .gitignore
├── README.md
├── Makefile
├── docker-compose.yml
├── go.mod
├── go.sum
├── cmd
│   ├── component1
│   │   ├── main.go
│   │   └── Dockerfile
│   └── component2
│       ├── main.go
│       └── Dockerfile
├── internal
│   ├── component1
│   │   ├── handler.go
│   │   └── service.go
│   └── component2
│       ├── handler.go
│       └── service.go
├── pkg
│   └── common
│       ├── utils.go
│       └── config.go
├── scripts
│   ├── build.sh
│   └── deploy.sh
└── config
    ├── component1
    │   └── config.yaml
    └── component2
        └── config.yaml
```