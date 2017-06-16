package ipc

import (
	"encoding/json"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string) (res *Response, err error) {
	req := &Response{method, params}
	
	b, err := json.Marshal(req)
	
	if err != nil {
		return
	}

	client.conn <- string(b)
	str := <- client.conn // wait for returning
	
	var response Response
	err =  json.Unmarshal([]byte(str), &response)
	res = &response

	return
}

func (client * IpcClient) Close() {
	client.conn <- "CLOSE"
}