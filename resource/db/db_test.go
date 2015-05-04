package db

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	var err error

	conn := GetConnection("")

	if conn == nil {
		t.Fail()
	} else {
		_, err = conn.Exec("DROP TABLE IF EXISTS APRILFOOL_DB_TEST")
		if err != nil {
			t.Error(err)
		}
		_, err = conn.Exec(`CREATE TABLE IF NOT EXISTS APRILFOOL_DB_TEST (
								id    SERIAL PRIMARY KEY,
							    dt    TIMESTAMP WITHOUT TIME ZONE,
                                vc    VARCHAR(50))`)
		if err != nil {
			t.Error(err)
		}

		id := []int{1, 2}
		dt := make([]time.Time, 2, 2)
		dt[0] = time.Now().UTC()
		vc := []string{"1st description", "2nd description"}

		_, err = conn.Exec("INSERT INTO APRILFOOL_DB_TEST VALUES ($1, $2, $3)", id[0], dt[0], vc[0])
		if err != nil {
			t.Error(err)
		}

		var id2 int
		var dt2 time.Time
		var vc2 string

		err = conn.QueryRow("SELECT * FROM APRILFOOL_DB_TEST WHERE dt=$1", dt[0]).Scan(&id2, &dt2, &vc2)
		if err != nil {
			t.Error(err)
		} else {
			if err != nil {
				t.Error(err)
			}
			if dt[0].Format(time.RFC3339) != dt2.Format(time.RFC3339) || vc[0] != vc2 {
				t.Error(errors.New(fmt.Sprintf("saved values are not equal %v %v %v %v", dt[0].Format(time.RFC3339), dt2.Format(time.RFC3339), vc[0], vc2)))
			}
		}

		dt[1] = time.Now().UTC()
		_, err = conn.Exec("INSERT INTO APRILFOOL_DB_TEST VALUES ($1, $2, $3)", id[1], dt[1], vc[1])
		if err != nil {
			t.Error(err)
		}

		rows, err := conn.Query("SELECT * FROM APRILFOOL_DB_TEST")
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&id2, &dt2, &vc2)
			switch id2 {
			case id[0]:
				if dt[0].Format(time.RFC3339) != dt2.Format(time.RFC3339) || vc[0] != vc2 {
					t.Error("saved multi values are not equal %v %v %v %v", dt[0], dt2, vc[0], vc2)
				}
			case id[1]:
				if dt[1].Format(time.RFC3339) != dt2.Format(time.RFC3339) || vc[1] != vc2 {
					t.Error("saved multi values are not equal %v %v %v %v", dt[0], dt2, vc[0], vc2)
				}
			}
		}

		_, err = conn.Exec("DROP TABLE APRILFOOL_DB_TEST")
		if err != nil {
			t.Error(err)
		}
	}
}
