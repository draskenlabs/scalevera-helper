package helper

import (
	"fmt"

	"github.com/draskenlabs/drasken-logger/logger"
)

type ScaleveraTask struct {
	Name      string
	Variables ScaleveraGlobal
	Modules   []*ScaleveraModule
	log       *logger.Logger
}

func GetTask(node *ScaleveraNode, log *logger.Logger) *ScaleveraTask {
	p := ScaleveraTask{
		log:       log,
		Variables: make(ScaleveraGlobal),
	}

	for index, childNode := range node.Children {
		if index == 0 {
			p.Name = childNode.Value
		}

		if childNode.IsAssignmentOperator() {
			globalData := ProcessGlobalBlockValue(childNode, "")
			for key, value := range *globalData {
				p.Variables[key] = value
			}
		} else if childNode.IsModuleNode() {
			module := GetModule(childNode, log)
			p.Modules = append(p.Modules, module)
		}
	}

	return &p
}

func (p *ScaleveraTask) String() string {
	return fmt.Sprintf("ScaleveraTask { Name: %s, Variables: %s, Modules: %s }", p.Name, p.Variables, p.Modules)
}
