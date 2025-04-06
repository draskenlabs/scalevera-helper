package helper

import (
	"fmt"
	"os"

	"github.com/draskenlabs/drasken-logger/logger"
)

type ScaleveraRequire struct {
	node      *ScaleveraNode
	log       *logger.Logger
	Providers []*ScaleveraProvider
}

func GetRequire(node *ScaleveraNode, log *logger.Logger) *ScaleveraRequire {
	r := ScaleveraRequire{
		node: node,
		log:  log,
	}
	return &r
}

func (r *ScaleveraRequire) GetProviders() {
	for _, providerNode := range r.node.Children {
		provider := GetProvider(providerNode, r.log)
		if r.ContainsProvider(provider.Name) {
			r.log.Error("`%s` provider is already defined.", provider.Name)
			os.Exit(1)
		}
		r.Providers = append(r.Providers, provider)
	}
}

func (r *ScaleveraRequire) ContainsProvider(name string) bool {
	hasProvider := false
	for _, provider := range r.Providers {
		if provider.Name == name {
			hasProvider = true
			break
		}
	}
	return hasProvider
}

func (r *ScaleveraRequire) String() string {
	return fmt.Sprintf("ScaleveraRequire(Providers: %s)", r.Providers)
}
