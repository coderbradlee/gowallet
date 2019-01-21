export GOPATH=$GOPATH:`pwd`/../../
echo $GOPATH
#go build -o coin-proxy .
go test -v -test.run TestNewMasterKey