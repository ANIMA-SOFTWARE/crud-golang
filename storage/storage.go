package storage

import "main/types"

type Storage interface {
	Get(int, types.Table) types.Table
	GetAll(types.Table) types.Table
	GetBySearch(string, types.Table) types.Table
	GetPage(int, types.Table) types.Table
	Delete(types.TableRow, types.Table) types.Table
	Append(types.TableRow, types.Table) types.Table
	Create(types.Table) types.Table
	GetTableType(string) types.Table
}
