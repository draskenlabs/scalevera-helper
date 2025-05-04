package schema

import (
	"fmt"
	"os"

	"github.com/draskenlabs/drasken-logger/logger"
)

type ScaleveraRequire struct {
	Name string
	Data ScaleveraRequireParam
	log  *logger.Logger
}

type ScaleveraRequireParam struct {
	Source  string
	Version string
}

func GetRequire(node *ScaleveraNode, log *logger.Logger) *ScaleveraRequire {
	p := ScaleveraRequire{
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

func (n *ScaleveraRequire) String() string {
	return fmt.Sprintf("ScaleveraRequire{Name: %s, Data: %s}", n.Name, n.Data)
}

func (n *ScaleveraRequireParam) String() string {
	return fmt.Sprintf("ScaleveraRequireParam{Source: %s, Version: %s}", n.Source, n.Version)
}
