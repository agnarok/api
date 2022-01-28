package main

import (
	"os"
  "time"
	"github.com/gin-gonic/gin"
)

func writeAnything(c *gin.Context) {
    f, _ :=  os.OpenFile("/data/log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660);
    f.WriteString("Logged at: "+ time.Now().String()+"\n")
    f.Close()
    c.JSON(200, gin.H{
      "message":"logged!",
    })
}


func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

  r.GET("/log", writeAnything)

  r.GET("/bye", func (c *gin.Context) {
    panic("bye!")
  })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
