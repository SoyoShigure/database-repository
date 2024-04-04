package option

type SQLSelectOption struct{
	Table string
	Columns *[]string
	WherePhrase SQLWhereOption
	OrderBy *SQLOrderByOption
	Limit int
	Offset *int
}

func NewSelectOption(table string) *SQLSelectOption{
	return &SQLSelectOption{
		Table: table,
		Columns: nil,
		WherePhrase: nil,
		OrderBy: nil,
		Limit: 0,
		Offset: nil,
	}
}