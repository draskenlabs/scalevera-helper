package helper

import (
	"fmt"

	"github.com/draskenlabs/drasken-logger/logger"
)

type ScaleveraStep struct {
	Name      string
	Variables ScaleveraGlobal
	Tasks     []*ScaleveraTask
	log       *logger.Logger
}

func GetStep(node *ScaleveraNode, log *logger.Logger) *ScaleveraStep {
	p := ScaleveraStep{
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
		} else if childNode.IsTaskNode() {
			step := GetTask(childNode, log)
			p.Tasks = append(p.Tasks, step)
		}
	}

	return &p
}

func (p *ScaleveraStep) String() string {
	return fmt.Sprintf("ScaleveraStep { Name: %s, Variables: %s, Tasks: %s }", p.Name, p.Variables, p.Tasks)
}
