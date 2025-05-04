package schema

import (
	"fmt"

	"github.com/draskenlabs/drasken-logger/logger"
)

type ScaleveraModule struct {
	Provider  string
	Name      string
	Variables ScaleveraGlobal
	Resources []*ScaleveraResource
	log       *logger.Logger
}

func GetModule(node *ScaleveraNode, log *logger.Logger) *ScaleveraModule {
	p := ScaleveraModule{
		log:       log,
		Variables: make(ScaleveraGlobal),
	}

	for index, childNode := range node.Children {
		if index == 0 {
			p.Provider = childNode.Value
		} else if index == 1 {
			p.Name = childNode.Value
		}

		if childNode.IsAssignmentOperator() {
			globalData := ProcessGlobalBlockValue(childNode, "")
			for key, value := range *globalData {
				p.Variables[key] = value
			}
		} else if childNode.IsResourceNode() {
			resource := GetResource(childNode, log)
			p.Resources = append(p.Resources, resource)
		}
	}

	return &p
}

func (p *ScaleveraModule) String() string {
	return fmt.Sprintf("ScaleveraModule { Provider: %s, Name: %s, Variables: %s, Resources: %s }", p.Provider, p.Name, p.Variables, p.Resources)
}
