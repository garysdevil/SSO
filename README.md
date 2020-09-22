# OSS 单点登录(权限管理)平台

## 概览
0. go版本
go1.15.1

1. 依赖的服务
mysql
redis

2. 生成文档
swag init --output ./src/docs

3. 访问文档的url
/swagger/index.html

4. 运行程序
go run .

5. 编译为可执行命令
go build -o sso

5. 自动生成表格
./sso initdb

6. 启动服务端
./sso

## 首次部署运行sso系统
1. 部署mysql数据库服务，redis服务
2. 创建数据库sso
3. 修改配置文件
cp settings.yaml settings.dev.yaml
vi settings.dev.yaml
4. go build -o sso
5. ./sso initdb
6. ./sso

## 与一个系统对接
### 单点登陆
1. 登入
  - 业务系统设置未登入过则跳转到 http://sso.wxblockchain.com/v1/sso/static/login.html?redirectURL=登入成功跳转页面的URI
  - sso系统，用户登入成功则 在 wxblockchain.com 域名下设置cookie保存token 和 重定向到 redirectURL?token=string

2. 登出
/v1/sso/logout
输入
{
  "token": "string"
}
登出成功输出
{
  "code": 0,
  "message": "OK",
  "data": {
    "flag": true
  }
}

3. 验证token是否有效
/v1/sso/check
输入
{
  "token": "string"
}
token有效则返回 
{
  "code": 0,
  "message": "OK",
  "data": {
    "flag": true
  }
}
