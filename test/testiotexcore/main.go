package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

// Client is the blockchain API client.
type Client struct {
	api iotexapi.APIServiceClient
}

// New creates a new Client.
func New(serverAddr string) (*Client, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{
		api: iotexapi.NewAPIServiceClient(conn),
	}, nil
}

//// SendAction sends an action to blockchain.
//func (c *Client) SendAction(ctx context.Context, selp action.SealedEnvelope) error {
//	_, err := c.api.SendAction(ctx, &iotexapi.SendActionRequest{Action: selp.Proto()})
//	return err
//}
//
//// GetAccount returns a given account.
//func (c *Client) GetAccount(ctx context.Context, addr string) (*iotexapi.GetAccountResponse, error) {
//	return c.api.GetAccount(ctx, &iotexapi.GetAccountRequest{Address: addr})
//}
func (c *Client) readState(ctx context.Context) error {
	_, err := c.api.ReadState(ctx, &iotexapi.ReadStateRequest{
		ProtocolID: []byte("poll"),
		MethodName: []byte("getstorageat"),
		Arguments:  [][]byte{[]byte()}})
	return err
}
func main() {
	_, err := New("127.0.0.1:14014")
	if err != nil {
		fmt.Println(err)
		return
	}
}
