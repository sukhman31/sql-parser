// Support for parsing statements of the following kind
//
// create table TABLENAME(COLUMNS DATATYPES)
// insert into TABLENAME(COLUMNS) values(VALUES)
// select * from TABLENAME
// update TABLENAME set COLUMNS = VALUES where CONDITIONS
// delete from TABLENAME where CONDITIONS

package parser

type queryType int

const (
	UNKNOWN queryType = iota
	CREATE
	INSERT 
	SELECT
	UPDATE
	DELETE
)

var queryTypeString = []string{
	"UNKNOWN",
	"CREATE",
	"INSERT",
	"SELECT",
	"UPDATE",
	"DELETE",
}

type ParsedQuery struct {
	queryType queryType
	columns map[string]string
	columnValues map[string]string
	allColumns bool
	conditions []string
	tableName string
}