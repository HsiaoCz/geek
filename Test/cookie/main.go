package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// go 使用cookie和session

type User struct {
	Identity int64  `json:"identity"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// gin 框架中set cookie的方法
/*
func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	if path == "" {
		path = "/"
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: c.sameSite,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
}
*/
/*
- name和Value键值对，设置Cookie的名称及相对应的值，都必须是字符串类型，会对具体的值进行编码。
- MaxAge 失效的时间，单位秒，默认为 -1。具体含义：

整数，则该Cookie在maxAge秒后失效。
负数，该Cookie为临时Cookie，关闭浏览器后立即失效。
0，表示删除该Cookie。


- Path 指定Cookie在哪个路径（路由）下生效。默认是”/“，意思是当前Domain下的所有路径。
如果设置为/index，则只有/index下的路由可以获取到该Cookie。比如：/index/title。
- Domain 指定Cookie所属域名，默认是当前域名。注意：”localhost“ 和 ”127.0.0.1“是完全不同的，你在”localhost“下设置的Cookie，不会被”127.0.0.1“访问到。
- Secure 该Cookie是否仅被使用安全协议传输。当Secure值为true时，Cookie在Http中是无效，在Https中才有效。引申：Http和Https的区别。
- HttpOnly 默认为False。 如果设置为True，则无法通过JS获取Cookie的值，但仍然可以通过Application手动修改Cookie。一定程度上能防止XSS攻击。引申，什么是XSS攻击。
- SameSite 可以防止第三方站点使用Cookie。
*/

// 写cookie
func writeCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 3)
	cookie := http.Cookie{Name: "username", Value: "geek", Expires: expiration}
	http.SetCookie(w, &cookie)
	w.Write([]byte("set cookie sccess"))
}

// 读cookie
func readCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	json.NewEncoder(w).Encode(cookie)
}

func main() {
	http.HandleFunc("/writecookie", writeCookie)
	http.HandleFunc("/readcookie", readCookie)

	log.Fatal(http.ListenAndServe("127.0.0.1:9911", nil))
}
