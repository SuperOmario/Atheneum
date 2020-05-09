package oracle

import (
	"atheneum/parser"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenCreateStatement(t *testing.T) {
	db := parser.Open("../JSON/basic.json")
	GenCreateStatements(db)

	testFile, testerr := os.Open("../output/ClinicCreateStatements.txt")
	if testerr != nil {
		fmt.Println(testerr, "testErr")
	}
	testCreateStatements, testReadErr := ioutil.ReadAll(testFile)
	if testReadErr != nil {
		fmt.Println(testReadErr, "testReadErr")
		testFile.Close()
	}

	expectedFile, expectederr := os.Open("../output/output.txt")
	if expectederr != nil {
		fmt.Println(expectederr, "expectedErr")
	}
	expectedCreateStatements, expectedReadErr := ioutil.ReadAll(expectedFile)
	if expectedReadErr != nil {
		fmt.Println(expectedReadErr, "expectedReadErr")
		testFile.Close()
	}

	if len(testCreateStatements) != len(expectedCreateStatements) {
		fmt.Println(testCreateStatements, expectedCreateStatements)
		t.Fatalf("Different Lengths")
	}
	for i, v := range testCreateStatements {
		if v != expectedCreateStatements[i] {
			t.Fatalf("Different Characters")
		}
	}
}
