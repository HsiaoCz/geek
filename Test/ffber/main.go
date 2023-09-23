package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// fiber
// 号称最好用的go web框架

const (
	addr = "127.0.0.1:3001"
)

/*
app := fiber.New(fiber.Config{
    Prefork:       true,
    CaseSensitive: true,
    StrictRouting: true,
    ServerHeader:  "Fiber",
})
*/

func main() {
	app := fiber.New()
	// 路由 基础路由
	app.Get("/user/hello", Hello)
	// 路由，动态路由
	app.Get("/helloworld/:username")
	// 静态文件
	app.Static("/prefix", "./")
	// 中间件
	// 全局中间件
	app.Use(PrintMethod)
	// 匹配某个路由或以某个前缀的所有路由
	app.Use("/user/hello", PrintMethod)
	app.Use("/user", PrintMethod)

	// 匹配参数
	app.Get("/api/book")
	app.Listen(addr)
}

// 基础路由
func Hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(map[string]any{"Message": "hello,My man"})
}

// 动态路由
func HelloWorld(c *fiber.Ctx) error {
	username := c.Params("username")
	return c.Status(fiber.StatusOK).JSON(map[string]any{"Message": "hello:" + username})
}

// 全局中间件
func PrintMethod(c *fiber.Ctx) error {
	method := c.Request().Header.Method()
	fmt.Println("this handler method is :", method)
	return c.Next()
}

// 路径上的查询参数
func GetBook(c *fiber.Ctx) error {
	book_name := c.Query("book_name")
	id := c.Query("id")
	return c.Status(fiber.StatusOK).JSON(map[string]any{"message": "获取数据成功", "data": book_name + id})
}
