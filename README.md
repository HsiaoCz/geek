# geek

## 1、SSO(单点登录)和 Oauth2.0

### 1.1 单机登录的基本流程

使用 session cookie 的方式实现登录鉴权

```go
// 使用cookie
// 当我们拿到用户注册信息后，生成一个唯一ID
// 将用户名放到cookie里面
// 使用一个map让id作为key,name作为值
&http.Cookie{

}

```

### 1.2 什么是单点登录

单点登录就是在一个多系统共存的环境下，用户在一次登录后，就不用在其他系统
中登录，也就是用户的一次登录能得到其他所有系统的信任

就比如在支付宝登录了，在淘宝，天猫也可以登录

单点登录问题出在哪儿？难点在哪儿？

### 1.3 Oauth2

OAuth2.0 的四个方面：

- 资源拥有方 user （我拥有信息）

- 资源方 微信（帮我们保管、存储信息）

- 资源请求方 client （简书，简书想拿到资源）

- 授权方 oauth2 （独立的授权服务）

一个例子，比如简书，支持第三方登录，比如微信
会跳出来一个二维码，通过扫码，会提示你的信息正在被使用
可以同意也可以拒绝

在这里，client(简书) 和各方的关系：

- client(简书) -----> 资源拥有方(我，即用户)

client 申请授权，用户给予一个授权

- client(简书) ------> 授权方 OAuth2

client 申请授权,授权方颁发 access token

- client(简书) ------> 资源方 (微信)

client 发送 access token 申请资源
资源方验证 access token，同时颁发给 client 资源

**OAuth2 四种授权流程**

- 授权码(authorization_code)
  用户登录授权，先拿 code
  用 code 换 token

- 隐藏式(implicit)
  用户登录授权，不拿 code,直接拿 token

- 密码式(password)
  用户提前给客户端用户名和密码
  验证客户端，直接用用户名和密码，拿 token

- 客户端凭证(client_credentials)
  验证客户端，直接拿 token

go 官方oauth2库

```go
go get -u github.com/go-oauth2/oauth2
```

**sso**

分离出一个服务，去维护一个统一的会话，大家去共享

恰好，OAuth2的authorization_code流程能够帮我们实现SSO

**llaoj/oauth2**

