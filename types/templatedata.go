package types

type TableData struct {
	Fields     []string
	Data       map[int]TableRow
	TableName  string
	NextPageNo int
}
