package parser

import (
	"fmt"
)

func parseSql(sql string) ParsedQuery {
	fmt.Println(sql)
	return ParsedQuery{}
}