package ahp

import (
	"encoding/json"
)

type ahpTree struct {
	Parent            *ahpLeaf `json:"-"`
	Leafs             map[string]*ahpLeaf
	IsComputed        bool
	ResultTableFormat []*ahpLeaf `json:"-"`
	Comparisons       map[string]ahpTreeComparison
}

type ahpLeaf struct {
	Id                string
	Parent            *ahpLeaf   `json:"-"`
	Children          []*ahpLeaf `json:"-"`
	LocalPonderation  float64
	GlobalPonderation float64
}

type ahpInputTreeLeaf struct {
	Id       string `json:"id"`
	ParentId string `json:"parent_id"`
}

type ahpInputTree struct {
	LeafArray   []ahpInputTreeLeaf           `json:"leaf_array"`
	Comparisons map[string]ahpTreeComparison `json:"comparisons"`
}

type ahpTreeComparison struct {
	ComparisonOrder  []string    `json:"order"`
	ComparisonMatrix [][]float64 `json:"comparison_matrix"`
}

func GetTree(jsonStr string, compute bool, saveTree bool) (ahpTree, error) {
	input, err := ParseFrom(jsonStr)
	if err != nil {
		return ahpTree{}, err
	}
	tree, err := ConvertInputToAhpTree(input)
	if err != nil {
		return ahpTree{}, err
	}
	if compute {
		err = ComputeTree(&tree)
		if err != nil {
			return ahpTree{}, err
		}
	}

	return tree, nil
}

func ParseFrom(jsonStr string) (ahpInputTree, error) {
	inputTree := ahpInputTree{}
	err := json.Unmarshal([]byte(jsonStr), &inputTree)
	return inputTree, err
}

func ComputeLeaf(leaf *ahpLeaf, tree *ahpTree) error {
	if len(leaf.Children) == 0 {
		tree.ResultTableFormat = append(tree.ResultTableFormat, leaf)
		return nil
	}
	localComp := tree.Comparisons[leaf.Id]
	eigen := Eigen(localComp.ComparisonMatrix)
	for index, child := range leaf.Children {
		child.LocalPonderation = eigen[index]
		child.GlobalPonderation = child.LocalPonderation * leaf.GlobalPonderation
	}
	for _, chil := range leaf.Children {
		ComputeLeaf(chil, tree)
	}
	return nil
}

func ComputeTree(tree *ahpTree) error {
	currentLeaf := tree.Parent
	err := ComputeLeaf(currentLeaf, tree)
	tree.IsComputed = true
	return err
}

func Eigen(matrix [][]float64) []float64 {
	//Sum cols
	matrixSize := len(matrix)
	eigen := make([]float64, matrixSize)
	colSums := make([]float64, matrixSize)
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			colSums[j] += matrix[i][j]
		}
	}
	tempMatrix := make([][]float64, 0)
	for i := 0; i < matrixSize; i++ {
		tempRow := make([]float64, matrixSize)
		for j := 0; j < len(matrix[i]); j++ {
			tempRow = append(tempRow, matrix[i][j]/colSums[j])
		}
		tempMatrix = append(tempMatrix, tempRow)
	}
	for i := 0; i < matrixSize; i++ {
		localSum := float64(0)
		for _, j := range tempMatrix[i] {
			localSum += j
		}
		eigen[i] = localSum / float64(matrixSize)
	}
	return eigen
}

func ConvertInputToAhpTree(inputTree ahpInputTree) (ahpTree, error) {
	//Parent Leaf is always first
	//Is computed is always false when creating a tree
	//COmparison
	//resultTableFormat
	tree := ahpTree{}
	tree.Leafs = make(map[string]*ahpLeaf)
	if len(inputTree.LeafArray) == 0 {
		return tree, nil
	}
	parentLeaf := Leaf(inputTree.LeafArray[0])
	parentLeaf.GlobalPonderation = 1
	parentLeaf.LocalPonderation = 1
	tree.Parent = &parentLeaf
	tree.Leafs[parentLeaf.Id] = &parentLeaf
	for i := 1; i < len(inputTree.LeafArray); i++ {
		currentInput := inputTree.LeafArray[i]
		currentLeaf := Leaf(currentInput)
		if val, ok := tree.Leafs[currentInput.ParentId]; ok {
			currentLeaf.Parent = val
			val.Children = append(val.Children, &currentLeaf)
		}
		tree.Leafs[currentInput.Id] = &currentLeaf
	}
	FillComparison(inputTree, &tree)
	return tree, nil
}

func FillComparison(inputTree ahpInputTree, tree *ahpTree) {
	tree.Comparisons = inputTree.Comparisons
}

func Leaf(inputTreeLeaf ahpInputTreeLeaf) ahpLeaf {
	return ahpLeaf{
		Id:                inputTreeLeaf.Id,
		LocalPonderation:  0,
		GlobalPonderation: 0,
	}
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
