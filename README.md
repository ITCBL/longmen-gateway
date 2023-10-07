# longmen-gateway 龙门关（网关）

## 基本信息
- 名称：龙门关（网关）
- 龙门关（Lóngmén Guān）："龙门"代表着进出通道，以及守护者的象征。
- 开发语言：go
- web框架：gin
- 数据库：mysql
- 缓存：redis
- 服务：网关代理服务 + 管理服务
- 

## 功能介绍

## 中间件

## 运行
- 执行`git pull` 拉取最新代码
- 执行`go mod tidy` 下载依赖包
- 修改配置文件 `config/config-dev.yaml` 配置文件
- 执行`go run main.go` 启动程序

## 部署

### 单机部署

### docker-compose部署

### k8s部署

## 开发计划
- 1、管理服务接口
  - 用户管理
    - [x] 注册
    - [x] 登录
    - [ ] 退出
    - [ ] 修改密码
    - [ ] 获取用户详细信息
    - [ ] 获取用户列表
  - 仪表盘/统计
    - ......
  - 服务管理
    - [ ] http/https
    - [ ] tcp
    - [ ] grpc
  - 租户管理
    - ......

- 2、代理服务接口
  - http/https
  - tcp
  - grpc


