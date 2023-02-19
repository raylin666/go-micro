package utils

import (
	"database/sql"
	"reflect"
)

// ScanRowsMap 遍历结果集
func ScanRowsMap(rows *sql.Rows) []map[string]string {
	res := make([]map[string]string, 0)               // 定义结果 map
	colTypes, _ := rows.ColumnTypes()                 // 列信息
	var rowParam = make([]interface{}, len(colTypes)) // 传入到 rows.Scan 的参数 数组
	var rowValue = make([]interface{}, len(colTypes)) // 接收数据一行列的数组
	for i, colType := range colTypes {
		rowValue[i] = reflect.New(colType.ScanType())           // 跟据数据库参数类型，创建默认值 和类型
		rowParam[i] = reflect.ValueOf(&rowValue[i]).Interface() // 跟据接收的数据的类型反射出值的地址

	}

	// 遍历数据
	for rows.Next() {
		rows.Scan(rowParam...) // 赋值到 rowValue 中
		record := make(map[string]string)
		for i, colType := range colTypes {
			if rowValue[i] == nil {
				record[colType.Name()] = ""
			} else {
				record[colType.Name()] = ByteStr(rowValue[i].([]byte))
			}
		}
		res = append(res, record)
	}

	return res
}
