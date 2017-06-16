package ipc

import (
	"testing"
)

type EchoServer struct {

}

func (server *EchoServer) Handle(method, params string) *Response {
	body := "echo: " + params
	return &Response{"200", body}
}

func (server *EchoServer) Name() string {
	return "Echo Server"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	response1, _ := client1.Call("GET", "From Client1")
	response2, _ := client2.Call("GET", "From Client2")

	if (response1.Body != "echo: From Client1" || response1.Code != "200" ||
	response2.Body != "echo: From Client2" || response2.Code != "200") {
		t.Error("IpcClient Call failed.")
		t.Error("Expected: [ echo: From Client1 ], [ echo: From Client2 ]")
		t.Error("Actully:  [", response1.Body, "], [", response2.Body, "]")
	}

	client1.Close()
	client2.Close()
}