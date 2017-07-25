package main

import (
	"bookshelf/controllers"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

// main ...
func main() {
	router := gin.Default()

	// 1. http://localhost:8080/ へアクセスすると「Hello world」と表示する。
	/*router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world")
	})

	//2. http://localhost:8080/hoge へアクセスすると、「fuga」と表示する。
	router.GET("/hoge", func(c *gin.Context) {
		c.String(200, "fuga")
	})*/

	router.GET("/:id", func(c *gin.Context) {
		// Pramを処理する
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(400, err)
			return
		}

		if id <= 0 {
			c.JSON(400, gin.H{"status": "id should be bigger than 0"})
			return
		}

		// データを処理する
		ctrl := controllers.NewUser()
		result := ctrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{})
			return
		}

		c.JSON(200, result)
	})

	router.Run(":8080")
}
