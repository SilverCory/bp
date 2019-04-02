package protobuf

import (
	"errors"
	"git.cory.red/coryory/bp/idk"
)

type Client struct {
	url string
}

func NewClient(url string) (*Client, error) {
	if url == "" {
		return nil, errors.New("INVALID URL")
	}

	return &Client{
		url: url,
	}, nil
}

func (c *Client) GetA() idk.A {
	panic("AAAA")
}
func (c *Client) GetBByID(int) idk.B {
	panic("AAAA")
}
func (c *Client) SetA(idk.A) error {
	panic("AAAA")
}
func (c *Client) SetBID(idk.B, int) error {
	panic("AAAA")
}