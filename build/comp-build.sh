deployname=web-crawler

echo "Building component $deployname"
rm main

export CGO_ENABLED=0
export GO111MODULE=on
go mod download

go env -w GOINSECURE=gorm.io,golang.org,go.etcd.io,gopkg.in
go get github.com/go-delve/delve/cmd/dlv
go build main.go

# Build the docker
docker image rm -f $deployname
docker build -t $deployname .