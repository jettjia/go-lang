## 一、概念与功能

Gaea 是一款数据库中间件，准确的说是 MySQL 中间件。它实现了 MySQL 协议，将自己伪装成一个 MySQL 服务器，应用程序通过 MySQL 客户端访问 Gaea，向 Gaea 发起 SQL 请求，Gaea 将 SQL 请求转发到后端 MySQL 执行，得到响应结果后再返回给客户端。



![img](https://static001.infoq.cn/resource/image/00/e4/00b0b58afdb308765962d0212fe146e4.png)



使用中间件可以集中管理用户和数据库配置信息，当用户和数据库实例规模增长时，可以有效减轻 DBA 的运维负担。



Gaea 抽象出 namespace, user, slice 这几个概念。namespace 对应于一个业务，是 Gaea 划分资源的基本单位。user 基本等同于 MySQL 的用户，user 通过 username 和 password 连接到 Gaea。username 和 password 可以唯一确定一个 namespace，一个 namespace 可以有多个用户。用户有权限的区别，存在只读用户，读写用户。



slice 对应于 MySQL 实例资源，一个 slice 必须包含一个主库，可以包含 0-n 个从库，可实现读写分离。namespace 中可以包含多个 slice，通过制定路由规则可实现分库分表的功能。图中绿色部分就是 Gaea 中的逻辑概念。



![img](https://static001.infoq.cn/resource/image/de/d6/de915f51899a193f382fafaf2477e2d6.png)



那么 Gaea 是如何管理以上这些配置信息的呢？这里要提到 Gaea 系统的 3 个组成部分，Gaea 系统由 Proxy，CC 和 Web 组成。



![img](https://static001.infoq.cn/resource/image/31/71/314e5184e2022da36e2921e3cddc4071.png)



上面几张图中展示的 Gaea 中间件准确来说叫作 Gaea Proxy，负责 MySQL 流量接入。Gaea CC 是中控服务，通过 Proxy 提供的管理接口与 Proxy 交互，主要用于配置管理和下发。



Gaea Web 提供了一个可视化的管理控制台，用于 DBA 管理配置信息和开发工程师查看配置信息。Gaea Web 通过操作 Gaea CC 来控制配置下发到 Gaea Proxy。



![img](https://static001.infoq.cn/resource/image/b4/51/b400c4303d4f5c872f068133f571a651.png)



目前我们在生产环境中将配置数据存放在 etcd 中，Gaea CC 和 Gaea Proxy 通过 etcd 进行配置数据交互。



Gaea 的主要功能：



- 非分片表支持大部分 SQL；
- 分片表支持 MySQL 路由，Kingshard 路由;
- 聚合函数支持常用的 max, min, sum, count, group by, order by；
- 支持多个分片表 join（需满足条件），分片表和全局表的 join。



## 二、快速使用

### 1、编译安装

使用 Gaea 比较简单，源码编译需要依赖 go 1.12。Gaea 使用 go module 管理依赖，克隆代码仓库后，进入 Gaea 目录执行 make 即可下载依赖并编译生成 gaea 的二进制文件。



### 2、编译配置

启动 Gaea Proxy 需要两份配置，一份是 Proxy 本身配置，如监听端口，日志路径，namespace 配置方式，等等。另一份是 namespace 配置，主要包含用户信息、DB 列表、实例配置信息等。



在生产环境，我们采用 etcd 存储 namespace 配置，使用一个可视化管理控制台进行配置查看和变更。为方便单机使用和测试，Gaea 也支持基于文件的 namespace 配置，这样就不需要依赖 etcd 了。



### 3、启动 Proxy

配置好这两份配置以后，执行 bin/gaea -config=etc/gaea.ini，进程启动后查看默认监听端口 13306 如果正常监听，则说明启动成功。



### 4、发送请求

使用 MySQL 客户端用 namespace 中的用户名和密码，即可像连接到 MySQL 一样连接到 Gaea。连接到 Gaea 后可以执行常用 SQL，对于非分片表来说，大部分 SQL 都是支持的。



分片表目前仅支持 DML 和一些常用管理命令，对分片表执行 DDL 会在 Gaea 层面报错，请求不会发往后端 MySQL 实例。



### 5、查看监控

Gaea 提供了较为完善的监控指标，方便开发和 DBA 根据监控指标排查问题。目前内部使用 prometheus 作为监控后端，用 grafana 作数据展示。



![img](https://static001.infoq.cn/resource/image/d1/23/d1beeeb13acdf73675518ff381f00723.png)



监控项主要包含了两个方面，一个是业务层面也就是 SQL 请求，主要包括 QPS、请求耗时、会话数、慢 SQL、错误 SQL 等指标。



另一个是机器层面，主要包括 CPU 使用率、内存使用率、协程数、GC 耗时等指标。业务监控指标采用 namespace 和 proxy 两个维度，方便从不同的维度进行统计，从而获得更为直观的监控数据。



![img](https://static001.infoq.cn/resource/image/e2/45/e2d0b615c49b48d568ee94010fa6b245.png)



除了 grafana 监控面板之外，Gaea 还提供了 SQL 反查功能。通过 grafana 面板上获取到的 SQL 范式的 MD5，用户可以查看自己 namespace 下面的慢 SQL 和错误 SQL 的具体的 SQL 范式，方便用户定位问题。



## 三、实现细节

### 1、整体架构

Gaea 的整体架构比较简洁，从上到下包括协议解析、会话管理、SQL 解析、路由调度、SQL 执行这几个模块。



![img](https://static001.infoq.cn/resource/image/59/56/5993fd0a6e321ab6faac9a2a957c9c56.png)



用户建立会话后，通过 MySQL 协议与 Gaea 交互，并且该会话接受 Gaea 的管理。发起的 SQL 请求经协议解析后得到具体的 SQL，经过 SQL 解析器解析成语法树，交由路由调度层做进一步处理，并得到最终发往后端执行的 SQL，也就是 Gaea 的执行计划。



SQL 执行模块根据执行计划，从后端 MySQL 连接池中拿到对应实例的连接，执行 SQL 并获取执行结果。如果是分片表的 SQL，还需要对多个执行结果做聚合。最后将执行结果按照 MySQL 协议返回给客户端。



### 2、网络协议

Gaea 支持 MySQL 协议，包含 MySQL 文本协议和 MySQL 二进制 prepare 协议。文本协议的实现主要参考了另外一款开源数据库中间件 Kingshard，二进制 prepare 协议则是按照 MySQL 的官方文档实现。



![img](https://static001.infoq.cn/resource/image/92/cc/9285a98385a46d3ed75796addec31fcc.png)



二进制 prepare 协议由前端会话保存 prepare 的 stmt，执行时首先由客户端发起 COM_STMT_PREPARE 请求，Gaea 处理请求成功后会返回给客户端 StmtId，随后客户端再发起 COM_STMT_EXECUTE 请求，Gaea 会用参数替换 SQL 中的占位符，然后以文本协议向后端 MySQL 发起请求并得到结果集，再将结果集转换成二进制协议后写回客户端。



### 3、SQL 执行

Gaea 使用 TiDB 的 SQL 解析器解析 SQL。也有一些其他开源项目在使用该解析器，如小米云平台开源的 SQL 优化工具 soar。



![img](https://static001.infoq.cn/resource/image/65/09/65015d5b4930e4865e03985b82eebd09.png)



当时我们调研了 Vitess, Kingshard, TiDB 这几款中间件的解析器，发现 Vitess 和 Kingshard 的解析器处理复杂 SQL 会遇到问题，解析得到的语法树抽象以及提供的语法树接口在实现分库分表的路由计算和 SQL 改写时会比较复杂。



TiDB 的解析器 SQL 兼容性较好，能够满足我们的使用需求，解析得到的语法树接口提供了一个 Restore()方法，方便对分片表 SQL 进行改写。



计算路由需要使用 SQL 解析后得到的语法树 (AST)。我们借助 AST 提供了 Visitor 机制，根据 SQL 中存在的表名判断是否包含分片表。如果只要包含一个分片表，就会走到分片表逻辑，计算路由，改写 SQL。



而如果不包含任何一个分片表，则将该 SQL 直接发往默认 slice 执行。计算路由时，只需要找到对应的 AST Node，改写 SQL 时只需要把对应的 Node 替换成一个装饰器 Node，这样只需要对原 AST 的根节点调用 Restore()，即可得到改写后的 SQL。



![img](https://static001.infoq.cn/resource/image/c8/41/c8d161d5cb001a6a5743d59f605c1e41.png)



### 4、连接管理

Gaea 的客户端会话支持超时清理，默认清理 1 小时内无任何请求的会话。超时清理采用时间轮实现。



后端与 MySQL 实例的连接采用连接池实现，连接池通过空闲连接数、最大连接数、空闲时间、超时时间等几个参数进行配置。



连接池机制可以保证后端 MySQL 的连接数可控，即使前端有大量的长连接，也不会对后端 MySQL 造成不良影响。



Gaea 连接池支持从库负载均衡和容错，当某台从库宕机，获取连接失败时，Gaea 会自动找到下一台从库获取连接，当所有从库均失败时，Gaea 会从主库获取连接。



这里有一点需要注意，一些 session 级别的系统变量设置需要保证前后端连接一致。



Gaea 处理 session 变量是拦截前端的 SQL 请求，将这些值保存到 session 中，当发起 SQL 查询请求，获取后端连接执行时，拿到后端连接后先将 session 中的这些变量拼接成一个 SET SQL 执行一次，再执行 SQL 查询。



目前 Gaea 仅支持有限的几个 session 变量，后面根据需要可能还会考虑增加。



### 5、配置热加载

对数据库中间件来说，配置热加载可以说是一个强需求。面对多租户场景，大量的 namespace 配置需要在线修改，如果每次修改都需要重启才能生效，会给 DBA 带来非常大的运维负担。Gaea 在设计时考虑到以上问题，实现了 namespace 级别的配置热加载。



配置热加载过程分为 3 个阶段，对 Gaea Proxy 来说分为两个阶段。



![img](https://static001.infoq.cn/resource/image/64/ed/6449a35a1ccd5209763d7e5eb19d45ed.png)



第 1 阶段，修改 etcd 中的 namespace 配置。第 2 阶段，向集群中的所有 Gaea Proxy 发起 prepare 请求，触发 Gaea Proxy 读取 etcd 中新的 namespace 配置，并初始化 namespace 的资源。如果所有 prepare 请求均成功, 则可以进入第 3 阶段 commit。



![img](https://static001.infoq.cn/resource/image/ea/b6/ea722b4c22994ab2fc7701f2735f34b6.png)



Gaea Proxy 以无锁方式切换到新的 namespace，并在延迟 1 分钟后关闭旧的 namespace。这主要是考虑到有一些长耗时的 SQL 请求还未完成，延迟 1 分钟尽量保证这些请求能够完成。配置热加载的过程对前端会话来说是无感知的，最大限度降低配置修改对业务的影响。



## 四、未来规划

Gaea 从去年下半年开始立项，开发，内部上线，到今年 5 月份开源，相比其他老牌的数据库中间件来说还比较年轻，功能上还有待完善，一些细节还需要打磨。



我们初步计划在未来支持事务追踪功能，分布式事务功能，进一步扩宽 Gaea 的应用场景。内部使用的 Gaea Web 也会在适配后尽快开源，方便大家通过可视化界面管理和操作 Gaea。



另外一个重要组件 Gaea Agent 目前还在内部开发中，主要用于 MySQL 实例管理，分片表自动扩容，缩容等操作，相信在不久的将来也会开源。



细节方面，目前有些 SQL 还不能支持，未来也会不断改进和完善，这里也希望大家在使用 Gaea 遇到问题时可以反馈给我们，我们会尽力解决。



性能也是另一个需要不断改进的方面，在内部使用 sysbench 初步测试结果来看，目前 Gaea Proxy 的性能在点查询场景 QPS 比 MyCAT 高 20%左右，虽然我们已经在协议层，SQL 执行层做了一部分优化，但一定还有性能提升的空间。



Gaea 项目现已开源，并且我们内部使用的版本与开源版本完全一致。因此该项目会长期维护，欢迎大家试用，目前已经有社区的伙伴为 Gaea 贡献代码，期待您的参与。



Github 地址：https://github.com/XiaoMi/Gaea