package resource

import (
	"database/sql"
	"errors"
	"github.com/AprilFool/AprilFool/resource/db"
	bc "golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"os"
)

type Signin struct {
	Session  string
	Name     string
	Photo    string
	Status   int
	password string
	sugar    string
}

type SigninResource struct{}

func (r SigninResource) Name() string {
	return "signin"
}

func (r SigninResource) Logger() *log.Logger {
	return log.New(io.Writer(os.Stderr), r.Name(), log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

func (r SigninResource) Get(vars map[string][]string) (interface{}, error) {

	conn := db.GetConnection("")

	if len(vars["user"]) == 1 && len(vars["password"]) == 1 {
		user, password := vars["user"][0], vars["password"][0]

		var si Signin

		err := conn.QueryRow("SELECT name,photo,status,password,sugar FROM \"User\" WHERE (mail=$1 OR name=$1)", user).Scan(&si.Name, &si.Photo, &si.Status, &si.password, &si.sugar)
		switch {
		case err == sql.ErrNoRows:
			r.Logger().Println("no valid account %s\n", user)
		case err != nil:
			r.Logger().Println(err)
		default:
			hp, err := bc.GenerateFromPassword([]byte(password+si.sugar), 10)
			r.Logger().Printf("%s %s", si.password, string(hp))

			err = bc.CompareHashAndPassword([]byte(si.password), []byte(vars["password"][0]+si.sugar))
			if err == nil {
				si.Session = ""
				si.password = ""
				si.sugar = ""
				return si, nil
			}
		}
	}
	return nil, errors.New("signin fails")
}
