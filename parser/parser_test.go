package parser

import (
	"atheneum/structure"
	"reflect"
	"testing"
)

func TestOpen(t *testing.T) {
	database := Open("../JSON/basic.json")

	expectedDatabase := structure.Database{Name: "Clinic", Entities: []structure.Entity{
		{
			Name: "Customer",
			Attributes: []structure.Attribute{
				{
					Name:     "CID",
					Datatype: "int",
					Length:   4,
					Nullable: false,
				},
				{
					Name:     "FName",
					Datatype: "varchar2",
					Length:   15,
					Nullable: true,
				},
				{
					Name:     "SName",
					Datatype: "varchar2",
					Length:   15,
					Nullable: true,
				},
				{
					Name:     "DOB",
					Datatype: "date",
					Length:   0,
					Nullable: true,
				},
			},
		},
	},
	}

	if !reflect.DeepEqual(database, expectedDatabase) {
		t.Errorf("Database is not equal to expectedDatabase")
	}
}
