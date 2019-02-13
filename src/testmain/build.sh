rm -fr gowallet
export GOPATH=$GOPATH:`pwd`/../../
echo $GOPATH
# go build -o gowallet .
# ./gowallet -name test
# go test -v -cpu=1,2,4 -count=1 -benchmem -coverprofile=xxxx.out -covermode=count testmain/gotest
# go test -v -bench=. -run=^$ -cpu=1,2,4 -count=1 -benchmem -coverprofile=xxxx.out -covermode=count testmain/gotest #-run=^$表示不测试任何功能函数，
# go tool cover -html=xxxx.out

go test -v -test.run TestCC ./gotest/demo53_test.go ./gotest/demo53.go