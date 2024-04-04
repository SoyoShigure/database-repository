package option

type SQLDeleteOption struct{
	Table string
	WherePhrase SQLWhereOption
}

func NewDeleteOption(table string) *SQLDeleteOption{
	return &SQLDeleteOption{
		Table: table,
		WherePhrase: nil,
	}
}

