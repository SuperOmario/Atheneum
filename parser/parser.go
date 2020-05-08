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

	fmt.Println(byteValue)

	var database structure.Database

	json.Unmarshal(byteValue, &database)

	for i := 0; i < len(database.Entities); i++ {
		fmt.Println("Entity Name: " + database.Entities[i].Name)
		var Attributes = database.Entities[i].Attributes
		for j := 0; j < len(Attributes); j++ {
			fmt.Println("Attribute Name: " + Attributes[j].Name)
			fmt.Println("Attribute Datatype: " + Attributes[j].Datatype)
			fmt.Println("Attribute Length: ", Attributes[j].Length)
			fmt.Println("Nullable: ", Attributes[j].Nullable)
		}
	}

	return database

}
