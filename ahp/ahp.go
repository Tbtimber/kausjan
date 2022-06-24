package ahp

import "encoding/json"

func testA() string {
	return "hello from AHP"
}

type ahpNode struct {
	nodeType string
	id       string
	children []*ahpNode
}

type ahpInputTreeLeaf struct {
	Id       string `json:"id"`
	ParentId string `json:"parent_id"`
}

type ahpInputTree struct {
	LeafArray   []ahpInputTreeLeaf       `json:"leaf_array"`
	Comparisons []ahpInputTreeComparison `json:"comparisons"`
}

type ahpInputTreeComparison struct {
	ParentId         string      `json:"parent_id"`
	ComparisonMatrix [][]float64 `json:"comparison_matrix"`
}

func ParseFrom(jsonStr string) (ahpInputTree, error) {
	inputTree := ahpInputTree{}
	err := json.Unmarshal([]byte(jsonStr), &inputTree)
	return inputTree, err
}

/*
{
    "inputtree": {
        "leaf_array" : [
            {
                "id": "testA",
                "parent_id": ""
            },
			{
                "id": "testB",
                "parent_id": "testA"
            }
        ],
        "comparisons": [
            {
                "parent_id": "str",
                "comparisons_matrix": [[
                    "float"
                ]]
            }
        ]
    },
    "gotree": {
        "parentNode" : "leaf",
        "leafs" : {
            "id1": {
                "id": "id-node",
                "local_ponderation": "float",
                "global_ponderation": "float",
                "parent": "leaf",
                "children": ["leaf"]
            }
        },
        "isComputed": true,
        "resultTableFormat": ["leaf"]
    }
}

*/
