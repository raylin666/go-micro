# MT 微服务框架 (基于 Kratos)

本框架是基于 `Kratos` 进行模块化设计的微服务框架，封装了常用的功能，使用简单，致力于进行快速的业务研发，同时增加了更多限制，约束项目组开发成员，规避混乱无序及自由随意的编码。<br />

提供了方便快捷的 `Makefile` 文件 (帮你快速的生成、构建、执行项目内容)。<br />

当你所需命令不存在时可添加到此文件中, 实现命令统一管理。这也大大的提高了开发者的开发效率, 让开发者更专注于业务代码。 <br />

### 目录介绍

| 目录 | 目录名称 | 目录描述 |
| --- | --- | --- |
| cmd | 项目启动 | 存放项目启动文件及依赖注入绑定 |
| config | 配置文件 | ProtoBuf 协议格式管理配置 |
| generate | 代码生成器 | 比如数据库查询器 |
| internal | 内部文件 | 存放项目业务开发文件 |
| pkg | 通用封装包 | 存放项目通用封装逻辑, 代码实现隔离项目内部业务 |
| static | 静态文件 | 比如图片、描述性文件、数据库SQL等 |
| bin | 运行文件 | |
| runtime | 临时/暂存 文件 | 比如日志文件 |

### 下载仓库

> git clone git@github.com:raylin666/go-mt-framework.git

### 初始化

> make init

### 下载依赖

> make generate

### 启动服务

> make run

访问服务 `curl 127.0.0.1:10010/heartbeat` , 返回 `200` 状态码则表示成功。
```shell
{
    "message": "PONE"
}
```

同时也支持采用 `Dockerfile` 和 `docker-compose` 启动哦 ！

### 编译执行文件 (需要有 .git 提交版本, 你也可以修改 `Makefile` 文件来取消这个限制)

> make build

编译成功后, 可以通过 `./bin/server` 命令运行服务。

<hr />

### 规范约束

> `service` 服务层 `对应/对接到 ProtoBuf 协议`, 职责上它只负责 `请求数据规划` 、`调用 biz 业务逻辑层` 及 `包装响应数据`, 不做其他任何逻辑处理。

> `biz` 逻辑层主要负责 `逻辑分工` 、`数据层调用` 及 `各数据块组装`, 返回给响应数据到服务层。

> `data` 数据层主要处理业务 `数据仓库` 的实例, `数据库逻辑处理`、`缓存逻辑处理`、`RPC 远程调用处理` 等相关操作。

> `pkg` 通用封装包内逻辑不允许调用 `internal` 内部包代码, 实现代码逻辑隔离, 也避免调用外部代码导致耦合和环境污染。

> 异常处理统一调用 `internal/constant/defined/errors.go` 内的变量, 该文件的所有变量都将对接到 `api` 层 `ProtoBuf 协议` 的 `error_reason.proto` 文件定义。

> 调用关系链: (服务层) `service` -> (业务逻辑层) `biz` -> (数据层) `data` , 逻辑代码只能下沉, 注意不要互调哦 ～

### 创建新模块

> 以 `account` 为例:
1. 复制 `api/v1/heartbeat.proto` 文件, 重命名为新模块名称, 编写API后执行 `make api` 命令生成API代码文件
2. 分别复制 `internal/service`, `internal/biz`, `internal/data` 的 `heartbeat.go` 文件处理层级关系链, `internal/service` 下的文件方法对应 `api/v1` 下的服务方法, 参数也是保持一致。
3. 处理完关系链后分别在 `internal/service/service.go`, `internal/biz/biz.go`, `internal/data/data.go` 文件的 `wire.NewSet` 方法绑定, 然后执行 `make wire` 命令更新依赖注入文件
4. 分别在 `internal/server/grpc.go` 和 `internal/server/http.go` 的参数添加服务, 然后分别调用类似 `v1.RegisterHeartbeatServer(srv, heartbeat)` 方法和调用 `v1.RegisterHeartbeatHTTPServer(srv, heartbeat)` 方法将服务注册即可～

### JWT 权限验证

中间件放在 `internal/middleware/auth/jwt.go` 文件, `NewAuthServer` 方法的 `Match` 调用用来进行<b>路由白名单过滤</b>, 可以用来指定路由是否需要经过权限验证, 代码示例:
```go
// NewAuthServer JWT Server 中间件
func NewAuthServer() func(handler middleware.Handler) middleware.Handler {
    return selector.Server(
        // JWT 权限验证
        JWTMiddlewareHandler(),
    ).Match(func(ctx context.Context, operation string) bool {
        // 路由白名单过滤 | 返回true表示需要处理权限验证, 返回false表示不需要处理权限验证
		r, err := regexp.Compile("/v1.Account/Login")
        if err != nil {
            // 自定义错误处理
            return true
        }
        return r.FindString(operation) != operation
    }).Build()
}
```

### 数据库模块

> 例如创建个 `account` 模型:
1. 编写模型文件, 在 `internal/repositories/dbrepo/model` 目录创建 `account.go` 文件, 内容如下:

```go
package model

type Account struct {
	UserName string `gorm:"column:username" json:"username"` // 用户名称
	Password string `gorm:"column:password" json:"password"` // 用户密码(加密串)
	Avatar   string `gorm:"column:avatar" json:"avatar"`     // 用户头像
	Status   int8   `gorm:"column:status" json:"status"`     // 用户状态 0:冻结 1:正常 2:暂停

	BaseModel
}
```

2. 如果需要制定 DIY 查询, 可在 `internal/repositories/dbrepo/method` 目录创建 `account.go` 文件, 内容如下:

```go
package method

import (
	"gorm.io/gen"
)

type Account interface {
	// where("`username`=@username")
	FindByUserName(username string) (gen.T, error)
}
```

3. 接下来就是添加到代码生成器中了，很简单。到 `generate/gormgen/db/default.go` 文件中添加 3 行代码即可:
```shell
// 代码 g.UseDB(dbInterface.Get().DB()) 后添加模型定义:
var accountModel = model.Account{}

// 代码 g.ApplyBasic 内添加注册模型:
g.ApplyBasic(
    accountModel,
)
  
// 代码 g.ApplyInterface 内添加注册DIY:
g.ApplyInterface(func(method method.Account) {}, accountModel)
```

4. 生成数据库查询器代码, 执行 `make gormgen` 命令, 成功后会在 `internal/repositories/dbrepo/query` 目录内生成对应的查询器文件。

5. 在 `internal/repositories/dbrepo/query.go` 文件中增加如下代码, 方便默认查询调用:
```go
// NewDefaultDbQuery 创建默认数据库查询
func NewDefaultDbQuery(dbRepo repositories.DbRepo) *query.Query {
return query.Use(dbRepo.DB(repositories.DB_CONNECTION_DEFAULT_NAME).Get().DB())
}
```

### 数据层处理

该层的设计目的是 <b>解藕数据与业务逻辑</b> 代码, 使层级更清晰, 业务逻辑层 `biz` 不需要引入 `repo` 来处理数据逻辑，通过接口方式类似 `HeartbeatRepo` 访问到数据层, 只需要专心处理业务逻辑即可。当业务复杂、多人协作开发、功能模块多的项目强烈建议采用数据层来降低后期维护成本。
具体的 `CURD` 操作文档请参考 `gorm`: <a href="https://gorm.io/zh_CN/">https://gorm.io/zh_CN/</a>


> 依照如上 `account` 模型为例:

在数据层添加查询一行记录代码, 打开 `internal/data/account.go` 文件, 示例内容如下：
```go
func (r *accountRepo) First(ctx context.Context, id int) (*model.Account, error) {
	var q = dbrepo.NewDefaultDbQuery(r.data.DbRepo)
	return q.Account.WithContext(ctx).Where(q.Account.ID.Eq(id)).First()
}
```

接下来只需要在 `internal/biz/account.go` 业务逻辑层定义接口, 然后调用该接口的方法即可, 定义接口方式:
```go
type AccountRepo interface {
	// 获取一条数据
    First(ctx context.Context, id int) (*model.Account, error)
}

// 调用方式
uc.repo.First(ctx, 1)
```
