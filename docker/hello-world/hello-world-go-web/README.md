GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o hello-web main.go


docker build -t hello-world-go-web .