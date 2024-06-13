package client

import (
  "context"
  "fmt"

  rpchttp "github.com/cometbft/cometbft/rpc/client/http"
  coretypes "github.com/cometbft/cometbft/rpc/core/types"
  "github.com/mingi3442/go-grpc/log"
)

type WsClient struct {
  RpcClient *rpchttp.HTTP
}

func Connect(url string) (*WsClient, error) {

  rpcWsClient, err := rpchttp.New(url, "/websocket")
  if err != nil {
    return nil, err
  }

  if err := rpcWsClient.Start(); err != nil {
    msg := fmt.Sprintf("Failed to start RPC client: %v", err)
    log.Log(log.ERROR, msg)
    return nil, err
  }
  log.Log(log.INFO, "RPC client started")
  return &WsClient{
    RpcClient: rpcWsClient,
  }, nil
}

func (c *WsClient) DisConnect() error {
  if c.RpcClient != nil {
    return c.RpcClient.Stop()
  }
  log.Log(log.INFO, "RPC client stopped")
  return nil
}

func (c *WsClient) Subscribe(ctx context.Context, subscriber, query string) (<-chan coretypes.ResultEvent, error) {
  events, err := c.RpcClient.Subscribe(ctx, subscriber, query)
  if err != nil {
    return nil, err
  }

  msg := fmt.Sprintf("Subscribed to events with query: %s", query)
  log.Log(log.INFO, msg)
  return events, nil
}
