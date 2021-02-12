package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:   "test_hotel",
		User:     "testuser",
		Password: "12345",
	}
	if conn.ConnectionURL() != "testuser:12345@/test_hotel" {
		t.Error("Unexpected connection string")
	}
}
