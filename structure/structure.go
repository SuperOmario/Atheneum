package structure

type Database struct {
	Name     string   `json : "name"`
	Entities []Entity `json : "entities"`
}

type Entity struct {
	Name        string      `json : "name"`
	Attributes  []Attribute `json : "attributes"`
	ForeignKeys []FK        `json : "foreignkeys"`
}

type Attribute struct {
	Name     string `json : "name"`
	Datatype string `json : "datatype"`
	Length   int    `json : "length"`
	Scale    int    `json : "scale"`
	Nullable bool   `json : "nullable"`
	PK       bool   `json : "PK"`
}

type FK struct {
	Name      string `json : "name"`
	Reference string `json : "reference"`
}
