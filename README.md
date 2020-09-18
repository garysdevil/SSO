# 1
0. go版本
go1.15.1

1. 生成文档
swag init --output ./src/docs
访问文档的url
/swagger/index.html

2. 运行程序
go run .

4. 编译为可执行命令
go build -o sso

5. 依赖的服务
mysql
redis
# 与一个系统对接
## 单点登陆
1. 登入
/sso/login
输入 账户密码
{
  "password": "string",
  "username": "string"
}
输出
{
  "code": 0,
  "message": "OK",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDAzMjY1NDcsImlhdCI6MTYwMDMyMjk0NywidXNlcm5hbWUiOiJ4aWVzaGlnYW5nIiwicm9sZWlkIjpudWxsfQ.jhyIwARMok8H5Gl9yx7vEY_u33HuRHTXJFfNNagvhdI"
  }
}
2. 登出
/sso/logout
输入
{
  "token": "string"
}
输出
{
  "code": 0,
  "message": "OK",
  "data": {
    "flag": true
  }
}

3. 验证token是否有效，有效则返回true
/sso/check
输入
{
  "token": "string"
}
输出
{
  "code": 0,
  "message": "OK",
  "data": {
    "flag": true
  }
}
