# xoj-backend（在线判题系统-后端）

## 项目概述

一个编程题目评测系统。能够根据管理员预设的题目用例对用户提交的代码进行执行和评测。

## 项目本地启动

1. 修改/conf 下的配置
    * appconfig.json：修改 `database` 和 `redis` 的连接地址
    * dubbogo.yaml：修改 `nacos` 的连接地址
2. 启动项目
    ```cmd
    go mod tidy
    go run main.go
    ```

## 运行项目中的单元测试

```bash
go test -v ./
go clean -testcache //清除测试缓存
```

## 关于 RPC 远程调用

该项目内的部分业务使用了dubbo-go 框架的rpc远程调用模式。

* 该项目角色是提供方（Provide）

* 配置文件位置：/conf/dubbogo.yaml

* 具体业务为为：
     * 获得题目信息 `Question.GetById`
     * 更新题目通过数+1 `Question.Add1AcceptedNum`
     * 获取提交题目信息 `QuestionSubmit.GetById`
     * 更新提交题目信息 `QuestionSubmit.UpdateById`

### 相关命令

1. 运行注册中心nacos：[见文档](https://blog.csdn.net/trinityleo5/article/details/132622712?spm=1001.2014.3001.5502)

2. 在项目根目录下运行下面命令，生成 rpc 相关go文件，然后共享 ./rpc_api 文件夹 给远程调用方的项目使用
    ```bash
    protoc --go_out=. --go-triple_out=. ./api.proto
    ```

## 其他补充

* 使用 `swag` 生成接口文档命令
    ```bash
    swag fmt
    swag init 
    ```
