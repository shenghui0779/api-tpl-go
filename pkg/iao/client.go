package iao

import (
	"os"
	"tplgo/pkg/iao/client"
)

type Client struct {
	User client.Client
}

var API *Client

func InitClient() {
	API = &Client{
		User: client.NewDefaultClient(os.Getenv("SERVICE_USER")),
	}
}
