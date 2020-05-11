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
					Scale:    0,
					Nullable: false,
				},
				{
					Name:     "FName",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "LName",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "DOB",
					Datatype: "DATE",
					Length:   0,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "Address",
					Datatype: "VARCHAR2",
					Length:   30,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "City",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "County",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "Email",
					Datatype: "VARCHAR2",
					Length:   30,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "Mobile",
					Datatype: "VARCHAR2",
					Length:   20,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "Notes",
					Datatype: "VARCHAR2",
					Length:   50,
					Scale:    0,
					Nullable: true,
				},
			},
		},
		{
			Name: "Physio",
			Attributes: []structure.Attribute{
				{
					Name:     "PHYID",
					Datatype: "NUMBER",
					Length:   2,
					Scale:    0,
					Nullable: false,
				},
				{
					Name:     "FName",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "LName",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "HireDate",
					Datatype: "DATE",
					Length:   0,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "Qual",
					Datatype: "VARCHAR2",
					Length:   30,
					Scale:    0,
					Nullable: true,
				},
				{
					Name:     "Salary",
					Datatype: "NUMBER",
					Length:   9,
					Scale:    2,
					Nullable: true,
				},
				{
					Name:     "PhyRole",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
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
