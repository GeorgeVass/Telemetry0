package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:   "test_hotel",
		User:     "root",
		Password: "19751975",
	}
	if conn.ConnectionURL() != "root:19751975@/test_hotel" {
		t.Error("Unexpected connection string")
	}
}
