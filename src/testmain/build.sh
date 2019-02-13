rm -fr gowallet
export GOPATH=$GOPATH:`pwd`/../../
echo $GOPATH
# go build -o gowallet .
# ./gowallet -name test
go test -v testmain/gotest
go test -v -bench=. -run=^$ -cpu=1,2,4 -count=3 testmain/gotest #-run=^$表示不测试任何功能函数，