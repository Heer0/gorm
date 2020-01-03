package mysql

import (
	"github.com/go-sql-driver/mysql"

	apmsql "github.com/Heer0/sql"
)

func init() {
	apmsql.Register("mysql", &mysql.MySQLDriver{}, apmsql.WithDSNParser(ParseDSN))
}
