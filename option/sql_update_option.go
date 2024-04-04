package option

type SQLUpdateOption struct{
	Table string
	Columns *[]string
	Values *[]interface{}
	WherePhrase SQLWhereOption
}

func NewUpdateOption(table string) *SQLUpdateOption{
	return &SQLUpdateOption{
		Table: table,
		Columns: nil,
		Values: nil,
		WherePhrase: nil,
	}
}