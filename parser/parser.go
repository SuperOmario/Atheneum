package parser

import (
	"atheneum/structure"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Open(inputFile string) structure.Database {
	jsonFile, err := os.Open(inputFile)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully opened JSON file")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var database structure.Database

	json.Unmarshal(byteValue, &database)

	fmt.Println(database)

	return database
}
