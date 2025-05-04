package schema

import (
	"fmt"
	"os"

	"github.com/draskenlabs/drasken-logger/logger"
)

type ScaleveraRequirements struct {
	Requires []*ScaleveraRequire
	node     *ScaleveraNode
	log      *logger.Logger
}

func GetRequirements(node *ScaleveraNode, log *logger.Logger) *ScaleveraRequirements {
	r := ScaleveraRequirements{
		node: node,
		log:  log,
	}
	return &r
}

func (r *ScaleveraRequirements) GetRequires() {
	for _, requireNode := range r.node.Children {
		require := GetRequire(requireNode, r.log)
		if r.ContainsRequire(require.Name) {
			r.log.Error("`%s` require is already defined.", require.Name)
			os.Exit(1)
		}
		r.Requires = append(r.Requires, require)
	}
}

func (r *ScaleveraRequirements) ContainsRequire(name string) bool {
	hasRequire := false
	for _, require := range r.Requires {
		if require.Name == name {
			hasRequire = true
			break
		}
	}
	return hasRequire
}

func (r *ScaleveraRequirements) String() string {
	return fmt.Sprintf("ScaleveraRequirements(Requires: %s)", r.Requires)
}
