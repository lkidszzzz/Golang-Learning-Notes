//go.mod 依赖管理
package main

import "go.uber.org/zap"

//go mod init	初始化当前文件夹，创建 go.mod 文件
//go mod why	解释为什么需要依赖
//go mod edit	编辑 go.mod 文件
//go mod vendor	将依赖复制到 vendor 目录下
//go mod verify	校验依赖
//go mod graph	打印模块依赖图
//go mod tidy	增加缺少的包，删除无用的包
//go mod download	下载依赖包到本地（默认为 GOPATH/pkg/mod 目录）

//受影响的命令：
//go get
//go list
//go clean —modcache
//go build

func main() {
	logger, _ := zap.NewProduction()
	logger.Warn("warning test")
}
