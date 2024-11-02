package user

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

var U User

func Login(c *gin.Context) {
	open, err := sql.Open("sqlite3", "../SQL/message.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer open.Close()
	err = c.ShouldBindJSON(&U)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	rows, err := open.Query("SELECT * FROM User")
	if err != nil {
		log.Println("查询失败", err)
	}
	for rows.Next() {
		var name string
		var account string
		var password string
		err = rows.Scan(&name, &account, &password)
		if err != nil {
			log.Println(err)
		}
		if U.Account == account && U.Password == password {
			c.JSON(http.StatusOK, "登录成功")
		}

	}
	_, err = open.Exec("INSERT INTO User(name,account,password) VALUES (?,?,?)", U.Name, U.Account, U.Password)
}
