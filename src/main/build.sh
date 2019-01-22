rm -fr gowallet
export GOPATH=$GOPATH:/root/gowallet
echo $GOPATH
go build -o gowallet .
./gowallet