module github.com/lzxm160/gowallet/test/testiotexcore

go 1.12

require (
	github.com/ethereum/go-ethereum v1.8.27
	github.com/iotexproject/iotex-core v0.8.6 // indirect
	github.com/iotexproject/iotex-proto v0.2.1-0.20190814190638-f74c55ffedf5
	google.golang.org/grpc v1.21.0
)

replace github.com/ethereum/go-ethereum => github.com/lzxm160/go-ethereum v10.1.5+incompatible
