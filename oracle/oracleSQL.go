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

	for i := 0; i < len(database.Entities); i++ {
		l, err := file.WriteString("CREATE TABLE " + database.Entities[i].Name + " (\n")
		if err != nil {
			fmt.Println(err)
			file.Close()
		}
		fmt.Println(l, "Bytes written successfully")
		var Attributes = database.Entities[i].Attributes
		for j := 0; j < len(Attributes); j++ {
			file.WriteString(`"` + Attributes[j].Name + `" `)
			if isNumber(Attributes[j]) {
				file.WriteString("NUMBER")
			} else {
				file.WriteString(Attributes[j].Datatype)
			}
			if !isDate(Attributes[j]) {
				file.WriteString("(" + strconv.Itoa(Attributes[j].Length))
				if isNumber(Attributes[j]) {
					file.WriteString("," + strconv.Itoa(Attributes[j].Scale) + ")")
				} else if !isDate(Attributes[j]) {
					file.WriteString(")")
				}
				if Attributes[j].Nullable == false {
					file.WriteString(" NOT NULL ENABLE")
				}
			}
			if j == len(Attributes)-1 {
				comp := getPrimaryKey(database.Entities[i])
				if len(comp) == 1 {
					file.WriteString(",\nPRIMARY KEY (" + comp[0].Name + ")")
				} else if len(comp) > 1 {
					file.WriteString(",\nPRIMARY KEY (")
					for k := 0; k < len(comp); k++ {
						file.WriteString(comp[k].Name)
						if k == len(comp)-1 {
							file.WriteString(", ")
						}
					}
					file.WriteString(")")
				}
				if hasForeignKeys(database.Entities[i]) {
					for k := 0; k < len(database.Entities[i].ForeignKeys); k++ {
						file.WriteString(",\nFOREIGN KEY (" + database.Entities[i].ForeignKeys[k].Name + ") REFERENCES ")
						file.WriteString(database.Entities[i].ForeignKeys[k].Reference + "(" + database.Entities[i].ForeignKeys[k].Name + ")")
					}
				}
				file.WriteString("\n)\n\n")
			} else {
				file.WriteString("," + "\n")
			}
		}
	}
	file.Close()
}

//Returns the Primary Key Attribute (will default to the first attribute in the entity if no Primary Key is found)
func getPrimaryKey(entity structure.Entity) []structure.Attribute {
	var comp []structure.Attribute
	for i := 0; i < len(entity.Attributes); i++ {
		if entity.Attributes[i].PK == true {
			comp = append(comp, entity.Attributes[i])
		}
	}
	return comp
}

//Determines if the entity contains foreign key attributes
func hasForeignKeys(entity structure.Entity) bool {
	if len(entity.ForeignKeys) > 0 {
		return true
	}
	return false
}

//checks if the data type is supported by Oracle
func isSupportedDatatType(attribute structure.Attribute) bool {
	if isBinaryType(attribute) || isCharType(attribute) || isDate(attribute) || isNumber(attribute) {
		return true
	}
	return false
}

//determines if the datatype is a SUPPORTED character data type
func isCharType(attribute structure.Attribute) bool {
	switch attribute.Datatype {
	case "CHAR":
		return true
	case "VARCHAR":
		return true
	case "NCHAR":
		return true
	case "NVARCHAR":
		return true
	case "CLOB":
		return true
	case "NCLOP":
		return true
	case "LONG":
		return true
	default:
		return false
	}
}

//determines if datatype is numeric
func isNumber(attribute structure.Attribute) bool {
	switch attribute.Datatype {
	case "BIT":
		return true
	case "TINYINT":
		return true
	case "SMALLINT":
		return true
	case "INT":
		return true
	case "BIGINT":
		return true
	case "DECIMAL":
		return true
	case "NUMERIC":
		return true
	case "FLOAT":
		return true
	case "REAL":
		return true
	case "NUMBER":
		return true
	default:
		return false
	}
}

//determines if datatype is Date (oracle does not support other Date and Time data types)
func isDate(attribute structure.Attribute) bool {
	switch attribute.Datatype {
	case "DATE":
		return true
	default:
		return false
	}
}

//determines if datatype is a SUPPORTED binary data type
func isBinaryType(attribute structure.Attribute) bool {
	switch attribute.Datatype {
	case "BLOB":
		return true
	case "BFILE":
		return true
	case "RAW":
		return true
	case "LONG RAW":
		return true
	default:
		return false
	}
}
