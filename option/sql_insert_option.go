package option

type SQLInsertOption struct{
	Table string
	Columns *[]string
	Values *[]interface{}
	WherePhrase SQLWhereOption
}

func NewInsertOption(table string) *SQLInsertOption{
	return &SQLInsertOption{
		Table: table,
		Columns: nil,
		Values: nil,
		WherePhrase: nil,
	}
}