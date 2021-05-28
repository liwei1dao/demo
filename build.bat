@echo "开始编译go_demo"
set GOOS=linux
set CGO_ENABLED=0
set GO111MODULE=auto
cd lego_demo
go build -o ../docker_demo/go_demo/console ./services/console/main.go
go build -o ../docker_demo/go_demo/gate ./services/gate/main.go
go build -o ../docker_demo/go_demo/live ./services/live/main.go

cd ../lgvue_demo
npm run build
REM pause