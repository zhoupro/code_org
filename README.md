# 代码分层

参考代码分层文档，理清代码分层的演进过程，整理出适合自己的代码分层架构。

框架有弹性，可适配简单项目和复杂项目。


## 对象定义


## 目录含义

- model/po 对应“存储层”的结构体，是对存储的一一映射。
- dal      数据读写层，统一处理数据库。
- conf     服务配置文件。
- service  基本业务逻辑。
- app      传统的service, 对基本业务逻辑编排。
- handler  入口层，参数校验。
- constant 常量定义
- message  mq异步消息入口
- cron     脚本服务入口
- task     包含异步任务的编排。
- convert  各种对象的转换
- bo       业务对象
- dto      传输对象
- vo       视图对象 
- util     工具
- rpc      外部业务服务
- infra    公司中台服务
- repo     聚合层，屏蔽rpc, infra细节
- script   脚本


##  开发流程
### 简单复杂度项目
#### 参与的对象
po->vo
dto->po
#### 参与的层

### 中等复杂度项目
#### 参与的对象
dto->bo->do->po

#### 参与的层

## 参考文章

- [1, 快速写出一个基于go的微服务](https://blog.csdn.net/qq_41080854/article/details/128804495)
- [2, DDD in golang](https://programmingpercy.tech/blog/how-to-domain-driven-design-ddd-golang/)
- [3, go-clean-architecture](https://github.com/luozhiyun993/go-clean-architecture)
- [4, 项目目录结构](https://lailin.xyz/post/go-training-week4-project-layout.html)
- [5, kratos 项目结构](https://go-kratos.dev/blog/go-project-layout/)
- [6, kratos 项目实践](https://juejin.cn/post/7097515746447589413)
- [7, 各种对象](https://zhuanlan.zhihu.com/p/102389552)
- [8, 阿里巴巴开发手册]()
- [9, 优秀代码如何分层](https://cloud.tencent.com/developer/article/1519017)
- [10, web 项目分层](https://chai2010.gitbooks.io/advanced-go-programming-book/content/ch5-web/ch5-07-layout-of-web-project.html)
- [11, 外观模式](https://zhuanlan.zhihu.com/p/612827175)
- [12, DDD go实战](https://juejin.cn/post/7123843139739058189)
- [13, 优秀代码是如何分层的]()

