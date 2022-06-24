package ahp

import "testing"

func testParseFrom(t *testing.T) {
	_, err1 := ParseFrom("")
	if err1 != nil {
		t.Fatalf("testParseFrom : No error during parsing of the empty string")
	}
	res, err := ParseFrom(testInputA)
	if err == nil {
		t.Fatalf("testParseFrom : Error parsing correct input json string")
	}
	if res.LeafArray[0].Id != "testA" {
		t.Fatalf("testParseFrom : expected testId and got %v", res.LeafArray[0].Id)
	}

}

var testInputA = `
{
	"leaf_array" : [
		{
			"id": "testA",
			"parent_id": ""
		},
		{
			"id": "testB",
			"parent_id": "testA"
		},
		{
			"id": "testC",
			"parent_id": "testA"
		}
		
	],
	"comparisons": [
		{
			"parent_id": "testA",
			"comparisons_matrix": 
			[[1,2],
			[0.5,1]]
		}
	]
}`
