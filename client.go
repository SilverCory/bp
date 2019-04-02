package bp

import (
	"errors"
	"git.cory.red/coryory/bp/idk"
	httpclient "git.cory.red/coryory/bp/internal/clients/http"
	protoclient "git.cory.red/coryory/bp/internal/clients/protobuf"
	"os"
)

// Client an interface for clients.
type Client interface{
	GetA() idk.A
	GetBByID(int) idk.B
	SetA(idk.A) error
	SetBID(idk.B, int) error
}

// NewClient will create a new environment depending on the environment.
func NewClient() (Client, error) {
	url, ok := os.LookupEnv("BP_SERVICE")
	if ok {
		return httpclient.NewClient(url)
	}

	proto, ok := os.LookupEnv("BP_PROTO")
	if ok {
		return protoclient.NewClient(url)
	}

	return nil, errors.New("IDK")
}