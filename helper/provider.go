package helper

import (
	"fmt"
	"os"

	"github.com/draskenlabs/drasken-logger/logger"
)

type ScaleveraProvider struct {
	Name string
	Data ScaleveraProviderParam
	log  *logger.Logger
}

type ScaleveraProviderParam struct {
	Source  string
	Version string
}

func GetProvider(node *ScaleveraNode, log *logger.Logger) *ScaleveraProvider {
	p := ScaleveraProvider{
		log: log,
	}

	for index, dataNode := range node.Children {
		if index == 0 {
			p.Name = dataNode.Value
			continue
		}

		switch dataNode.Children[0].Value {
		case "source":
			p.Data.Source = dataNode.Children[1].Value
		case "version":
			p.Data.Version = dataNode.Children[1].Value
		default:
			p.log.Error("Invalid provider param `%s`", dataNode.Children[0].Value)
			os.Exit(1)
		}
	}

	return &p
}

func (n *ScaleveraProvider) String() string {
	return fmt.Sprintf("ScaleveraProvider{Name: %s, Data: %s}", n.Name, n.Data)
}

func (n *ScaleveraProviderParam) String() string {
	return fmt.Sprintf("ScaleveraProviderParam{Source: %s, Version: %s}", n.Source, n.Version)
}
