package pq

import (
	pq "github.com/lib/pq"

	apmsql "github.com/Heer0/sql"
)

func init() {
	apmsql.Register("postgres", &pq.Driver{}, apmsql.WithDSNParser(ParseDSN))
}
