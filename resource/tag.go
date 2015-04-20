package resource

import (
	"fmt"
	"github.com/AprilFool/AprilFool/resource/db"
)

type Tag struct {
	Id int
	Name string
}

type TagResource struct{}

func (tag TagResource) Name() string {
	return "Tag"
}

func (tag TagResource) Get(vars map[string][]string) (interface{}, error) {
	var id int
	var name string
	conn := db.GetConnection()	

	if len(vars["id"]) == 1 {
		conn.QueryRow(fmt.Sprintf("SELECT * FROM Tag WHERE id=%v", vars["id"][0])).Scan(&id, &name)
		return Tag{Id:id, Name:name}, nil
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
