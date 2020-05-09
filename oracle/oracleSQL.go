package oracle

import (
	"atheneum/structure"
	"fmt"
	"os"
	"strconv"
)

func GenCreateStatements(database structure.Database) {
	file, err := os.Create("../output/" + database.Name + "CreateStatements.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully Created File: " + file.Name())
	}

	l, err := file.WriteString("CREATE DATABASE " + database.Name + "\n")
	if err != nil {
		fmt.Println(err)
		file.Close()
	}

	fmt.Println(l, "Bytes written successfully")

	for i := 0; i < len(database.Entities); i++ {
		l, err := file.WriteString("\nCREATE TABLE " + database.Entities[i].Name + " (\n")
		if err != nil {
			fmt.Println(err)
			file.Close()
		}
		fmt.Println(l, "Bytes written successfully")
		var Attributes = database.Entities[i].Attributes
		for j := 0; j < len(Attributes); j++ {
			file.WriteString(`"` + Attributes[j].Name + `" `)
			file.WriteString(Attributes[j].Datatype)
			if Attributes[j].Datatype != "DATE" {
				file.WriteString("(" + strconv.Itoa(Attributes[j].Length))
				if Attributes[j].Replace != "" {
					file.WriteString("," + Attributes[j].Replace + ")")
				} else if Attributes[j].Datatype != "DATE" {
					file.WriteString(")")
				}
				if Attributes[j].Nullable == false {
					file.WriteString(" NOT NULL ENABLE")
				}
			}
			if j == len(Attributes)-1 {
				fmt.Println(j)
				file.WriteString("\n)\n\n")
			} else {
				fmt.Println("comma new line")
				file.WriteString("," + "\n")
			}
		}
	}
}
