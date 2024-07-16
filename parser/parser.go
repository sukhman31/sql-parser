package parser

func parseSql(sql string) ParsedQuery {
	parsedQuery := ParsedQuery{}
	parse(sql, &parsedQuery)
	return parsedQuery
}

func parse(sql string, pq *ParsedQuery) {
}