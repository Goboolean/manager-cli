module github.com/Goboolean/manager-cli

go 1.19

require (
	github.com/Goboolean/shared v0.0.0-20230717060802-49f66c2bee25
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	github.com/notEpsilon/go-pair v0.0.0-20221220200415-e91ef28c6c0b
	github.com/spf13/cobra v1.7.0
	google.golang.org/grpc v1.56.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
)

replace (
	github.com/Goboolean/manager-cli/ => ./
	github.com/Goboolean/shared => /home/lsjtop10/projects/goboolean/shared
)
