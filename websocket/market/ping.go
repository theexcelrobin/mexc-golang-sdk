package mexcwsmarket

import (
	mexcws "github.com/theexcelrobin/mexc-golang-sdk/websocket"
)

func (s *Service) Ping() error {
	req := &mexcws.WsReq{
		Method: "PING",
	}

	return s.client.Send(req)
}
