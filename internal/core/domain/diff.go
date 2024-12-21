package domain

type SchemaDiff struct {
	AddedTables      []TableDiff
	RemovedTables    []string
	ModifiedTables   []TableDiff
	AddedSequences   []Sequence
	RemovedSequences []string
	AddedFunctions   []Function
	RemovedFunctions []string
	ModifiedFunctions []FunctionDiff
}

type TableDiff struct {
	Name              string
	AddedColumns      []Column
	RemovedColumns    []string
	ModifiedColumns   []ColumnDiff
	AddedConstraints  []Constraint
	RemovedConstraints []string
	AddedIndexes      []Index
	RemovedIndexes    []string
}

type ColumnDiff struct {
	Name          string
	OldType       string
	NewType       string
	OldNullable   bool
	NewNullable   bool
	OldDefault    *string
	NewDefault    *string
	OldReferences *Reference
	NewReferences *Reference
}

type FunctionDiff struct {
	Name         string
	OldFunction  Function
	NewFunction  Function
}