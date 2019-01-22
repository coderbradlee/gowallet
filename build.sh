rm -fr gowallet
export GOPATH=$GOPATH:`pwd`
echo $GOPATH
go build -o gowallet .
./gowallet