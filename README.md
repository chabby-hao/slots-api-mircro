# Slots API microservices

API 项目的微服务化拆分

## 框架

go-zero <https://go-zero.dev/cn/>

## 目录结构参考

### 项目目录

```
slots_api_micro // 项目跟目录
├── common // 通用库
│   ├── randx
│   └── stringx
├── go.mod
├── go.sum
└── service // 服务存放目录
    ├── order
    │   ├── api
    │   └── model
    │   └── rpc
    ├── pay
    │   ├── api
    │   └── model
    │   └── rpc
    └── user
        ├── api //  http访问服务，业务需求实现
        ├── cronjob // 定时任务，定时数据更新业务
        ├── model 
        ├── rmq // 消息处理系统：mq和dq，处理一些高并发和延时消息业务
        ├── rpc // rpc服务，给其他子系统提供基础数据访问
        └── script // 脚本，处理一些临时运营需求，临时数据修复
```

目录结构参考 <https://go-zero.dev/cn/service-design.html>

### api目录

```
.
├── etc
│   └── greet-api.yaml              // 配置文件
├── go.mod                          // mod文件
├── greet.api                       // api描述文件
├── greet.go                        // main函数入口
└── internal                        
    ├── config  
    │   └── config.go               // 配置声明type
    ├── handler                     // 路由及handler转发
    │   ├── greethandler.go
    │   └── routes.go
    ├── logic                       // 业务逻辑
    │   └── greetlogic.go
    ├── middleware                  // 中间件文件
    │   └── greetmiddleware.go
    ├── svc                         // logic所依赖的资源池
    │   └── servicecontext.go
    └── types                       // request、response的struct，根据api自动生成，不建议编辑
        └── types.go
```

目录结构参考 <https://go-zero.dev/cn/api-dir.html>

### rpc服务
```
.
├── etc             // yaml配置文件
│   └── greet.yaml
├── go.mod
├── greet           // pb.go文件夹①
│   └── greet.pb.go
├── greet.go        // main函数
├── greet.proto     // proto 文件
├── greetclient     // call logic ②
│   └── greet.go
└── internal        
    ├── config      // yaml配置对应的实体
    │   └── config.go
    ├── logic       // 业务代码
    │   └── pinglogic.go
    ├── server      // rpc server
    │   └── greetserver.go
    └── svc         // 依赖资源
        └── servicecontext.go
```

目录结构参考 <https://go-zero.dev/cn/rpc-dir.html>
