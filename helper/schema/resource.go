package schema

import (
	"fmt"

	"github.com/draskenlabs/drasken-logger/logger"
)

type ScaleveraResource struct {
	ProviderResource string
	Name             string
	Variables        ScaleveraGlobal
	log              *logger.Logger
}

func GetResource(node *ScaleveraNode, log *logger.Logger) *ScaleveraResource {
	p := ScaleveraResource{
		log:       log,
		Variables: make(ScaleveraGlobal),
	}

	for index, childNode := range node.Children {
		if index == 0 {
			p.ProviderResource = childNode.Value
		} else if index == 1 {
			p.Name = childNode.Value
		}

		if childNode.IsAssignmentOperator() {
			globalData := ProcessGlobalBlockValue(childNode, "")
			for key, value := range *globalData {
				p.Variables[key] = value
			}
		}
	}

	return &p
}

func (p *ScaleveraResource) String() string {
	return fmt.Sprintf("ScaleveraResource { ProviderResource: %s, Name: %s, Variables: %s }", p.ProviderResource, p.Name, p.Variables)
}
