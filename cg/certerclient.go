package cg

import (
	"dumb-lobby/ipc"
	"encoding/json"
	"errors"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil {
		return err
	}

	res, err := client.Call("addPlayer", string(b))
	if err == nil && res.Code == "200" {
		return nil
	}

	return err
}

func (client *CenterClient) RemovePlayer(name string) error {
	ret, _ := client.Call("removeplayer", name)
	if ret.Code == "200" {
		return nil
	}
	return errors.New(ret.Code)
}

func (client *CenterClient) ListPlayer(params string) (ps []*Player, err error) {
	res, _ := client.Call("listplayer", params)
	if res.Code != "200" {
		err = errors.New(res.Code)
		return
	}

	err = json.Unmarshal([]byte(res.Body), &ps)
	return
}

func (client *CenterClient) Broadcast(message string) error {
	m := &Message{Content: message}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	res, _ := client.Call("broadcast", string(b))
	if res.Code == "200" {
		return nil
	}
	return errors.New(res.Code)
}
