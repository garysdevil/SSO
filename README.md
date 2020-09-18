# OSS 单点登录(权限管理)平台

## 概览
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
## 与一个系统对接
### 单点登陆
1. 未登入过则跳转到http://10.200.79.81/static/login.html?redi
用户登入成功则返回跳转前的页面

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
