package parser

import (
	"atheneum/structure"
	"fmt"
	"reflect"
	"testing"
)

func TestOpen(t *testing.T) {
	database := Open("../JSON/basic.json")

	expectedDatabase := structure.Database{Name: "Clinic", Entities: []structure.Entity{
		{
			Name: "Client",
			Attributes: []structure.Attribute{
				{
					Name:     "CID",
					Datatype: "NUMBER",
					Length:   4,
					Replace:  "0",
					Nullable: false,
				},
				{
					Name:     "FName",
					Datatype: "VARCHAR2",
					Length:   15,
					Replace:  "",
					Nullable: true,
				},
				{
					Name:     "LName",
					Datatype: "VARCHAR2",
					Length:   15,
					Replace:  "",
					Nullable: true,
				},
				{
					Name:     "DOB",
					Datatype: "DATE",
					Length:   0,
					Replace:  "",
					Nullable: true,
				},
			},
		},
	},
	}

	if !reflect.DeepEqual(database, expectedDatabase) {
		fmt.Println(expectedDatabase)
		t.Errorf("Database is not equal to expectedDatabase")
	}
}
