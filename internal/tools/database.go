package tools

import "github.com/jackc/pgx/v5"

var conn *pgx.Conn

func SetDatabaseConn(dbConn *pgx.Conn) {
	conn = dbConn
}

func GetDatabaseConn() *pgx.Conn {
	return conn
}