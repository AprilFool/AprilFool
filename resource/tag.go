package resource

import (
	"fmt"
	"github.com/AprilFool/AprilFool/resource/db"
	"io"
	"log"
	"os"
)

type Tag struct {
	Id   int
	Name string
}

type TagResource struct{}

func (r TagResource) Name() string {
	return "tag"
}

func (r TagResource) Logger() *log.Logger {
	return log.New(io.Writer(os.Stderr), r.Name(), log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

func (tag TagResource) Get(vars map[string][]string) (interface{}, error) {
	var id int
	var name string
	conn := db.GetConnection("")

	if len(vars["id"]) == 1 {
		conn.QueryRow(fmt.Sprintf("SELECT * FROM Tag WHERE id=%v", vars["id"][0])).Scan(&id, &name)
		return Tag{Id: id, Name: name}, nil
	} else {
		rows, _ := conn.Query("SELECT * FROM Tag")
		data := make([]Tag, 0)
		defer rows.Close()
		for rows.Next() {
			_ = rows.Scan(&id, &name)
			data = append(data, Tag{Id: id, Name: name})
		}
		return data, nil
	}
}
