package memory

import (
	"math/rand"

	"github.com/itohio/itohio.github.io/pkg/config"
	"github.com/itohio/itohio.github.io/pkg/ports"
)

type Adapter struct {
	cfg       *config.Config
	greetings []string
}

var _ ports.DbPort = &Adapter{}

func New(cfg *config.Config) *Adapter {
	return &Adapter{
		cfg:       cfg,
		greetings: []string{"Hello,", "Hi,", "Ahoy,", "Aloha,"},
	}
}

func (a *Adapter) GetRandomGreeting() string {
	return a.greetings[rand.Intn(len(a.greetings))]
}

func (a *Adapter) GetGreetings() []string {
	return a.greetings
}
