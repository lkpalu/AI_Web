package controller

import (
	"ai_web/internal/model"
	"ai_web/internal/user"
	"ai_web/internal/util"
	"context"
	"database/sql"
	//"github.com/gin-gonic/gin"
	gin "ai_web/Gwe/Gwe"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tmc/langchaingo/llms"
	"log"
	"net/http"
	"time"
)

type Chat struct{}

func (this *Chat) Dochat(c *gin.Context) {
	var body model.Chat
	err := c.ShouldBindJSON(&body)
	//fmt.Println(body.Text)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	open, err := sql.Open("sqlite3", "C:\\Users\\lkpalu\\GolandProjects\\ai_web\\internal\\SQL\\message.db")
	if err != nil {
		log.Println("打开数据库失败")
	}
	defer open.Close()
	_, err = open.Exec("INSERT INTO messages(name,time,text) VALUES (?,?,?)", user.U.Name, time.Now(), body.Text)
	prompt := util.CreatePrompt()

	data := map[string]any{
		"text": body.Text,
	}

	messages, err := prompt.FormatMessages(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	//fmt.Println(messages)
	contents := []llms.MessageContent{
		llms.TextParts(messages[0].GetType(), messages[0].GetContent()),
		llms.TextParts(messages[1].GetType(), messages[1].GetContent()),
	}
	llm := util.CteateModel(c, "qwen")
	resp, err := llm.GenerateContent(context.Background(), contents)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": resp.Choices[0].Content,
	})
	//fmt.Println(resp.Choices[0].Content)
	_, err = open.Exec("INSERT INTO messages(name,time,text) VALUES (?,?,?)", "AI", time.Now(), resp.Choices[0].Content)
}
