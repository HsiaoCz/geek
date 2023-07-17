# geek

## 1、SSO(单点登录)和Oauth2.0

### 1.1 单机登录的基本流程

使用session cookie的方式实现登录鉴权

```go
// 使用cookie
// 当我们拿到用户注册信息后，生成一个唯一ID
// 将用户名放到cookie里面
// 使用一个map让id作为key,name作为值
&http.Cookie{

}

```