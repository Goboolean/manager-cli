module github.com/Goboolean/manager-cli

go 1.19

require (
	github.com/Goboolean/shared v0.0.0-20230717060802-49f66c2bee25
	github.com/joho/godotenv v1.5.1
	github.com/notEpsilon/go-pair v0.0.0-20221220200415-e91ef28c6c0b
	github.com/spf13/cobra v1.7.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace (
	github.com/Goboolean/manager-cli/internal/adaptor/transaction => ./internal/adaptor/transaction
	github.com/Goboolean/shared => /home/lsjtop10/projects/goboolean/shared
)
