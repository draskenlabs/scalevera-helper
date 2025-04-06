package helper

import (
	"fmt"

	"github.com/draskenlabs/drasken-logger/logger"
)

type ScaleveraPlan struct {
	Name      string
	Variables ScaleveraGlobal
	Steps     []*ScaleveraStep
	log       *logger.Logger
}

func GetPlan(node *ScaleveraNode, log *logger.Logger) *ScaleveraPlan {
	p := ScaleveraPlan{
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
		} else if childNode.IsStepNode() {
			step := GetStep(childNode, log)
			p.Steps = append(p.Steps, step)
		}
	}

	return &p
}

func (p *ScaleveraPlan) String() string {
	return fmt.Sprintf("ScaleveraPlan { Name: %s, Variables: %s, Steps: %s }", p.Name, p.Variables, p.Steps)
}
