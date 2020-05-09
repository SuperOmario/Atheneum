package structure

type Database struct {
	Name     string   `json : "name"`
	Entities []Entity `json : "entities"`
}

type Entity struct {
	Name       string      `json : "name"`
	Attributes []Attribute `json : "attributes"`
}

type Attribute struct {
	Name     string `json : "name"`
	Datatype string `json : "datatype"`
	Length   int    `json : "length"`
	Replace  string `json : "replace`
	Nullable bool   `json : "nullable"`
}
