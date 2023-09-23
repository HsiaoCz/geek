package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// fiber
// 号称最好用的go web框架

const (
	addr = "127.0.0.1:3001"
)

/*
// 自定义初始化配置信息
app := fiber.New(fiber.Config{
	// 开启多进程 默认关闭
    Prefork:       true,
    // 定义路由的大小写问题，开启之后/foo /Foo 将是不同的路由
    CaseSensitive: true,
	// 开启之后,/foo 和 /foo/将是不同的路由，否则路由时相同的
    StrictRouting: true,
	// 定义响应头中的Server标记
    ServerHeader:  "Fiber",
})
*/

/*
静态文件配置

	app.Static("/", "./public", fiber.Static{
	  Compress:      true, //是否开启压缩
	  ByteRange:     true, //是否启用字节范围请求。
	  Browse:        true, //是否启用目录浏览
	  Index:         "index.html" //默认的访问
	  CacheDuration: 10 * time.Second,//缓存时间
	  MaxAge:        3600,
	})
*/

/*
	    增加响应头信息
	    c.Accepts("png")

		//c.Accepts("json", "text")     // "json"
		//c.Accepts("application/json") // "application/json"

		//新增的响应头信息---下面的字段会出现再响应头
		c.Append("Link", "Test")
		c.Append("Link", "http://google.com", "http://localhost")

		//返回的APP路由堆栈
		return c.JSON(c.App().Stack())

		// 其他设置响应头的方法
		  c.Set("Content-Type", "text/plain")
	   // => "Content-type: text/plain"
	   // ...
	   c.Vary("Accept-Encoding", "Accept")
*/

// type H map[string]any
// 这个在fiber里面是fiber.Map

type Person struct {
	Name string `json:"name" form:"name"`
	Pass string `json:"pass" form:"pass"`
}

// Cookie 的设置和删除
type Cookie struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Path     string    `json:"path"`
	Domain   string    `json:"domain"`
	MaxAge   int       `json:"max_age"`
	Expires  time.Time `json:"expires"`
	Secure   bool      `json:"secure"`
	HTTPOnly bool      `json:"http_only"`
	SameSite string    `json:"same_site"`
}

// c.Download() 下载文件
// c.SendFile() 发送文件

/*
获取form表单文件
 file, _ := c.FormFile("5gmsg.conf")
      // Save file to root directory:
      fmt.Println("文件名称", file.Filename)
      return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
*/

/*
获取多文件
 //MultipartForm()。这将返回一个map[string][]string
      if form, err := c.MultipartForm(); err == nil {
         fmt.Println("输出表单信息,", form)
         files := form.File["5gmsg.conf"]
         fmt.Println("输出files信息,", files)
         for _, file := range files {
            fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
            if err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename)); err != nil {
               return c.SendString("上传失败！")
            }
         }
      }
      return c.SendString("上传成功！")
*/

// c.FormValue() 获取表单的值
// c.Get() 获取请求头信息
// c.IP(),c.IPS获取请求IP信息
// c.Is("html") 这个方法判断content-type的类型
//  c.XHR() // 判断是否XHRjQuery提交

// 中间件之间的变量传递
/*
//全局的中间件中我们的设置一个参数值信息---
   app.Use(func(c *fiber.Ctx) error {
      c.Locals("user", "小钟同学-变量存贮传递")
      return c.Next()
   })
   app.Get("/", func(c *fiber.Ctx) error {
      return c.JSON(fiber.Map{
         "name": c.Locals("user"),
         "age":  2000,
      })
   })

 */

// c.Redirect 重定向
// c.Params() 获取path参数
// c.Query() 和c.QueryParse() 路径上的参数匹配
// c.OriginUrl() 请求路径和参数
// c.SendStream() 响应字节流
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
	app.Get("/api/book", GetBook)

	// 路由组
	v1 := app.Group("/v1")
	v1.Get("/he", Hello)
	// 查看所有路由列表
	// app.Stack()

	// body解析
	app.Post("/user/register", UserRegister)
	// 设置cookie
	app.Get("/cookie", SetCookie)
	// 删除cookie
	app.Post("/cookie")
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

// body参数解析
func UserRegister(c *fiber.Ctx) error {
	p := new(Person)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "注册成功", "Data": p})
}

// 设置cookie
func SetCookie(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "john"
	cookie.Value = "doe"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// Set cookie
	c.Cookie(cookie)
	return c.SendString("设置cookie成功!")
}

// 删除cookie

func DeleteCookie(c *fiber.Ctx) error {
	//删除所有的ClearCookie
	c.ClearCookie()
	// 根据键值对删除
	c.ClearCookie("user")
	return c.SendString("删除cookie成功!")
}
