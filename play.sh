export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

go build -o ./build/apiServer ./api &&
go build -o ./build/dataServer ./data &&
docker-compose up -d --scale data=5