package db

import (
	"testing"
)

func TestConnection(t *testing.T) {
	var err error

	conn := GetConnection()

	if conn == nil {
		t.Fail()
	} else {
		_, err = conn.Exec("CREATE TABLE IF NOT EXISTS APRILFOOL_DB_TEST (ID INT)")
		if err != nil {
			t.Fail()
		}
		_, err = conn.Exec("DROP TABLE APRILFOOL_DB_TEST")
		if err != nil {
			t.Fail()
		}
	}
}
