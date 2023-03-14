package nats

import (
	"fmt"

	"github.com/iamthe1whoknocks/hezzl_test_task/config"
	"github.com/nats-io/nats.go"
)

// Nats struct
type Nats struct {
	Conn *nats.Conn
}

// Constructor
func New(n *config.Nats) (*Nats, error) {
	nc, err := nats.Connect(fmt.Sprintf("http://%s:%s", n.Host, n.Port))
	if err != nil {
		return nil, fmt.Errorf("nats - New - Connect: %w", err)
	}
	return &Nats{
		Conn: nc,
	}, nil

}
