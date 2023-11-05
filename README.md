# longmen-gateway 龙门关（网关）

## 基本信息
- 名称：龙门关（网关）
- 龙门关（Lóngmén Guān）："龙门"代表着进出通道，以及守护者的象征。
- 开发语言：go (go协程的并发性能好)
- web框架：gin
- 数据库：mysql
- 缓存：redis
- 限流：自研 or sentinel-go
- 链路追踪：自研 or jaeger
- 日志：zap + es

## 功能介绍

主要分为两大模块，网关管理服务和网关代理服务。网关代理服务的信息，来源于网关管理系统。

### 网关管理服务

### 网关代理服务

## 中间件

### 基础/通用

#### 跨域

#### 权限

#### 日志

### 代理专用

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
- 1、管理端接口
  - 用户管理
    - [x] 注册
    - [x] 登录
    - [x] 图形验证
    - [ ] 短信验证
    - [ ] 退出
    - [ ] 修改密码
    - [ ] 获取用户详细信息
    - [ ] 获取用户列表
  - 仪表盘/统计
    - [ ] qps
    - [ ] qpd
    - [ ] ......
  - 服务管理
    - [ ] http/https
    - [ ] tcp
    - [ ] grpc
  - 租户管理
    - [ ] ......

- 2、代理端接口
  - [ ] http/https
  - [ ] tcp
  - [ ] grpc

- 3、管理端前端


