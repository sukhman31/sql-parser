// Support for parsing statements of the following kind
//
// create table TABLENAME(COLUMNS DATATYPES)
// insert into TABLENAME(COLUMNS) values(VALUES)
// select * from TABLENAME
// update TABLENAME set COLUMNS = VALUES where CONDITIONS
// delete from TABLENAME where CONDITIONS

package parser

type QueryType int

const (
	UNKNOWN_TYPE QueryType = iota
	CREATE
	INSERT
	SELECT
	UPDATE
	DELETE
)

var QueryTypeString = []string{
	"UNKNOWN_TYPE",
	"CREATE",
	"INSERT",
	"SELECT",
	"UPDATE",
	"DELETE",
}

type Operator int

const (
	UNKNOWN_OPERATOR Operator = iota
	EQ                          // =
	NE                          // !=
	GT                          // >
	LT                          // <
	GTE                         // >=
	LTE                         // <=
)

var OperatorAsString = []string{
	"UNKNOWN_CONDITION",
	"EQ",
	"NE",
	"GT",
	"LT",
	"GTE",
	"LTE",
}

type Condition struct {
	lhs      string
	rhs      string
	operator Operator
}

type ParsedQuery struct {
	queryType    QueryType
	columns      map[string]string
	columnValues map[string]string
	allColumns   bool
	conditions   []Condition
	tableName    string
}