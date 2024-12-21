package domain

type Schema struct {
	Tables      map[string]Table
	Sequences   map[string]Sequence
	Functions   map[string]Function
	Triggers    map[string]Trigger
	Indexes     map[string]Index
	Constraints map[string]Constraint
}

type Table struct {
	Name        string
	Columns     map[string]Column
	Constraints []Constraint
	Indexes     []Index
}

type Column struct {
	Name       string
	Type       string
	Nullable   bool
	Default    *string
	References *Reference
}

type Constraint struct {
	Name       string
	Type       string
	Table      string
	Columns    []string
	References *Reference
}

type Reference struct {
	Table    string
	Columns  []string
	OnDelete string
	OnUpdate string
}

type Index struct {
	Name    string
	Table   string
	Columns []string
	Unique  bool
}

type Sequence struct {
	Name      string
	StartWith int64
	Increment int64
}

type Function struct {
	Name       string
	Arguments  []Argument
	ReturnType string
	Body       string
}

type Trigger struct {
	Name     string
	Table    string
	When     string
	Event    string
	Function string
}

type Argument struct {
	Name string
	Type string
}