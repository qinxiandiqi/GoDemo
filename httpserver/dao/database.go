package dao
import "database/sql"

var Db sql.DB

func init() {
	Db = sql.Open("mydql", "user@password@/dbname")
}
