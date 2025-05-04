package schema

type ScaleveraGlobal map[string]interface{}

func GetGlobal(node *ScaleveraNode, prefixParam string) *ScaleveraGlobal {
	globalData := make(ScaleveraGlobal)

	for _, dataNode := range node.Children {
		parsedData := ProcessGlobalBlockValue(dataNode, prefixParam)
		for key, value := range *parsedData {
			globalData[key] = value
		}
	}

	return &globalData
}

func ProcessGlobalBlockValue(node *ScaleveraNode, prefixParam string) *ScaleveraGlobal {
	globalData := make(ScaleveraGlobal)

	if node.IsAssignmentOperator() {
		key := node.Children[0].Value
		if prefixParam != "" {
			key = prefixParam + "." + key
		}

		nestedData := GetGlobal(node.Children[0], key)
		for nestedKey, nestedValue := range *nestedData {
			globalData[nestedKey] = nestedValue
		}

		if len(node.Children) > 1 {
			valueNode := node.Children[1]
			if valueNode.IsAssignmentOperator() {
				// Recursively process nested assignment
				nestedData := GetGlobal(valueNode, key)
				for nestedKey, nestedValue := range *nestedData {
					globalData[nestedKey] = nestedValue
				}
			} else {
				// Direct key-value assignment
				globalData[key] = valueNode.Value
			}
		}
	}

	return &globalData
}
