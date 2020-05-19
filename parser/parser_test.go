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
					PK:       true,
				},
				{
					Name:     "FName",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "LName",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "DOB",
					Datatype: "DATE",
					Length:   0,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "Address",
					Datatype: "VARCHAR2",
					Length:   30,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "City",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "County",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "Email",
					Datatype: "VARCHAR2",
					Length:   30,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "Mobile",
					Datatype: "VARCHAR2",
					Length:   20,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "Notes",
					Datatype: "VARCHAR2",
					Length:   50,
					Scale:    0,
					Nullable: true,
					PK:       false,
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
					PK:       true,
				},
				{
					Name:     "FName",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "LName",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "HireDate",
					Datatype: "DATE",
					Length:   0,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "Qual",
					Datatype: "VARCHAR2",
					Length:   30,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "Salary",
					Datatype: "NUMBER",
					Length:   9,
					Scale:    2,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "PhyRole",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
			},
		},
		{
			Name: "TestTreatment",
			Attributes: []structure.Attribute{
				{
					Name:     "TID",
					Datatype: "NUMBER",
					Length:   2,
					Scale:    0,
					Nullable: false,
					PK:       true,
				},
				{
					Name:     "TName",
					Datatype: "VARCHAR2",
					Length:   20,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "TDesc",
					Datatype: "VARCHAR2",
					Length:   35,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "Fee",
					Datatype: "NUMBER",
					Length:   9,
					Scale:    2,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "Duration",
					Datatype: "NUMBER",
					Length:   2,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
			},
		},
		{
			Name: "TestAppointment",
			Attributes: []structure.Attribute{
				{
					Name:     "APPID",
					Datatype: "NUMBER",
					Length:   5,
					Scale:    0,
					Nullable: false,
					PK:       true,
				},
				{
					Name:     "CID",
					Datatype: "NUMBER",
					Length:   4,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "PHYID",
					Datatype: "NUMBER",
					Length:   2,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "TID",
					Datatype: "NUMBER",
					Length:   2,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "AppDate",
					Datatype: "DATE",
					Length:   0,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "Paid",
					Datatype: "VARCHAR2",
					Length:   1,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "PayMethod",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "AppTime",
					Datatype: "VARCHAR2",
					Length:   15,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "Discount",
					Datatype: "NUMBER",
					Length:   2,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
				{
					Name:     "Notes",
					Datatype: "VARCHAR2",
					Length:   50,
					Scale:    0,
					Nullable: true,
					PK:       false,
				},
			},
			ForeignKeys: []structure.FK{
				{
					Name:      "CID",
					Reference: "Client",
				},
				{
					Name:      "PHYID",
					Reference: "Physio",
				},
				{
					Name:      "TID",
					Reference: "TestTreatment",
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
