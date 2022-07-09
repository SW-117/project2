package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/SW-117/project2/library"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
)

func GetUserDb(dbname, user, pass, host string) *sql.DB {
	db, err := manager.New(dbname, user, pass, host).Set(manager.SetCharset("utf8"),
		manager.SetAllowCleartextPasswords(true),
		manager.SetInterpolateParams(true),
		manager.SetTimeout(1*time.Second),
		manager.SetReadTimeout(1*time.Second)).Port(3306).Open(true)
	if err != nil {
		log.Default().Println("init database error", err)
		return nil
	}
	return db
}

func RegisteruUser(userData *library.UserData, ctx context.Context) (int64, error) {
	db := GetUserDb("test", "root", "", "127.0.0.1")
	if db == nil {
		return 0, errors.New("database init error")
	}
	now := time.Now().Unix()
	mdata := map[string]interface{}{
		"name":        userData.Name,
		"sex":         userData.Sex,
		"pass":        userData.PassWard,
		"create_time": now,
		"update_time": now,
	}
	var data []map[string]interface{}
	data = append(data, mdata)
	cond, vals, err := builder.BuildInsert("user", data)
	if err != nil {
		fmt.Println("insert user error : ", err)
		return 0, errors.New("insert user error:" + err.Error())
	}
	res, err := db.Exec(cond, vals...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
