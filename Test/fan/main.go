package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:anything", HandleCEOFan)
	log.Fatal(r.Run("127.0.0.1:3001"))
}

func HandleCEOFan(c *gin.Context) {
	name := c.Param("anything")
	if name == "xiaofanzong" {
		name = "小樊总"
		c.Header("Content-Type", "text/html;charset=utf-8")
		c.String(http.StatusOK, fmt.Sprintf(
			`<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>CEO</title>
		</head>
		<body>
		</body>
		<script>
			let str="%s加油!"
			alert(str)
		</script>
		</html>`, name))
	} else {
		c.Header("Content-Type", "text/html;charset=utf-8")
		c.String(http.StatusOK, fmt.Sprintf(
			`<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>CEO</title>
		</head>
		<body>
		</body>
		<script>
			let str="%s爬"
			alert(str)
		</script>
		</html>`, name))
	}
}
