# Work-6
This is a microservice refactoring project that uses the Kitex framework and RPC communication. In this project, we will learn and deploy knowledge related to microservices. In addition, this project also needs to be optimized and error resolved based on the previous project.


## 完成情况

完整项目目录: /Tree.md

## 准备工作

运行kitex_gen.sh脚本文件 [ ./kitex_gen.sh] 


#### 进行Makefile指令的初始化操作 

- make -B idl
#### 用于生成kitex框架下的idl文件对应的各种服务

- make -B go 

#### 完成运行环境的搭建与将.go源文件编译为可执行的二进制文件
- make api
- make user
- make video
- make publish
- make comment
- make relation
- make favorite

#### 以上的几种指令用于启动各个不同的服务 



`docker`
```

```
接口文档

[https://apifox.com/apidoc/shared-16d8fec9-45b1-4607-b6ff-6957b9a4e215]

业务文档

[https://a0fgr0eyrxd.feishu.cn/docx/PWVndH9sUofOy0x4iXpcAXmVnce?openbrd=1&doc_app_id=501&blockId=doxcn7jJLeQWQTxiUXmsQpRsuzf&blockType=whiteboard&blockToken=HQxRw2qCwh4I94bN1azckkZ5ndd#doxcn7jJLeQWQTxiUXmsQpRsuzf]


接口文档

[https://a0fgr0eyrxd.feishu.cn/docx/Vz86drRYHoCNTixo62ocWxwbnHh?openbrd=1&doc_app_id=501&blockId=doxcnUukRabp79vH85sYqJiSYcb&blockType=whiteboard&blockToken=Wikbw1R1PhOytsbvggPcDOLln7g#doxcnUukRabp79vH85sYqJiSYcb]

项目文档

[https://a0fgr0eyrxd.feishu.cn/docx/MNt9dTVXboCi6fxMEoOcw6xankd?from=from_copylink]


`业务架构`

```
详细见Tree.md文件
```
## 日志
`3.14`
```
完成了将原来的Gin框架并且使用自动迁移模式进行数据表的创建 改为使用Hertz框架进行重构，
并且在此过程中引入了Hertz认证的JWT进行Token鉴权验证
```

`3.18`
```
学着使用自己构建SQL脚本去构建数据表，并且在对数据进行修改的过程中引入了触发器的概念，
对点赞关注等操作进行一个自动关联，同时也是有了外键关联，建立起数据表之间的联系
```

`3.20`
```
遇到了websocket的无法连接问题
----------------------------
下午解决了 在项目中看注释
----------------------------
完成了websocket两人聊天的demo
```
----------------------------
`3.23`
```
引入了RabbitMQ，完成对于离线状态下消息的通道缓存，并且引入了Redis完成了基础的增删改查，减缓了数据库的查询压力
```
`3.25号`
```
项目结构不规范 没有建立三层架构 于是进行修改
```
`3.26号`
```
更改为三层架构 将Bonus完成 完成了所有的接口 剩下的就是对其进行优化
```
`3.27号`
```
完成了Github Action的学习 与Pull Request的使用 同时尝试部署golangci-lint
```
`3.28号`
```
完成了Redis接口的飞书文档编写
完成所有接口的编写
完成了golangci-lint本地静态部署
如果想将输出的日志定向到本地 可以使用golangci-lint run >[文件路径] 命令

完成了接口文档的编写
```

`4.9号`
```
完成了细节优化 尝试优化Sql 与Redis
```

`4.11号`
```
完成了dockerfile地构建编写
准备学习微服务
```
`4.19号`
```
项目的答辩 指出了错误和给出了建议 开始思考如何优化  完善
```
`4.20号`
```
开始尝试理解并发的本质 以及哪些场景下适合使用 以及学习benchmark测试
```
`4.24号`
```
开始学习ElasticSearch 了解到除了mysql外的另一种非关系型数据库 尝试将其部署到项目中
```
`4.25号`
```
完成Es的基本增删改查操作 同时将其运用到service中 查看效果
```

`5.5号`
```
使用Kitex完成了项目的重构 同时开始尝试了解其服务注册等内容
```
`5.6号`
```
根据上次答辩遗留的error分层处理进行了项目的完善
```
`

### 需要优化的地方：

`Redis的引入`
```
完成
```
`RabiitMq 消息队列的引入`
```
完成
```

## 待完成项
~~1 .评论接口需要实现对评论进行评论~~

~~2 .点赞接口需要处理对评论的点赞~~

~~3 .社交模块：完成基于 websocket 的聊天功能，考虑到聊天的实时性，请使用 Redis + MySQL 方式实现~~

~~4 .实现对视频的排行~~
## Bonus:



## 下一阶段

Kitex

k8s 

