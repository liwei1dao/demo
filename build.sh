cd lego_demo
GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../docker_demo/go_demo/console ./services/console/main.go
GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../docker_demo/go_demo/gate ./services/gate/main.go
GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../docker_demo/go_demo/live ./services/live/main.go
