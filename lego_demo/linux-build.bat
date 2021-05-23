set GOOS=linux
set CGO_ENABLED=0
del bin/backstage,bin/gate,bin/live
go build -o bin/backstage services/backstage/main.go
go build -o bin/gate services/gate/main.go
go build -o bin/live services/live/main.go
REM pause