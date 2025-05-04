package schema

import (
	"fmt"

	"github.com/draskenlabs/drasken-logger/logger"
)

type ScaleveraProvider struct {
	Require   string
	Name      string
	Variables ScaleveraGlobal
	log       *logger.Logger
}

func GetProvider(node *ScaleveraNode, log *logger.Logger) *ScaleveraProvider {
	p := ScaleveraProvider{
		log:       log,
		Variables: make(ScaleveraGlobal),
	}

	for index, dataNode := range node.Children {
		if index == 0 {
			p.Require = dataNode.Value
			continue
		} else if index == 1 {
			p.Name = dataNode.Value
			continue
		}

		if dataNode.IsAssignmentOperator() {
			globalData := ProcessGlobalBlockValue(dataNode, "")
			for key, value := range *globalData {
				p.Variables[key] = value
			}
		}
	}

	return &p
}

func (n *ScaleveraProvider) String() string {
	return fmt.Sprintf("ScaleveraProvider{Required: %s, Name: %s, Variables: %s}", n.Require, n.Name, n.Variables)
}
