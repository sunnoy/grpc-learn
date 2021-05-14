# 参考文章

https://eddycjy.com/posts/go/grpc/2018-09-24-stream-client-server/

# 示例代码仓库

https://github.com/EDDYCJY/go-grpc-example

# 项目初始化


- 基础命令安装
  
    官方文档 https://grpc.io/docs/languages/go/quickstart/

    ```bash
    # 安装protoc 代码生成
    brew install protobuf

    # 查看版本
    protoc --version

    # 安装golang插件
    export GO111MODULE=on
    go get google.golang.org/protobuf/cmd/protoc-gen-go \
             google.golang.org/grpc/cmd/protoc-gen-go-grpc

    # 添加环境变量
    export PATH="$PATH:$(go env GOPATH)/bin"

    ```

- 编写proto文件

  proto3语法参考

  [Protocol Buffers Version 3 Language Specification | Google Developers](https://developers.google.com/protocol-buffers/docs/reference/proto3-spec#top_level_definitions)

    ```bash
    syntax = "proto3";

    // 需要定义生成代码的的属于的包路径需要 / 斜线
    option go_package = "openvpn/";

    package openvpn;

    // The greeting service definition.
    service Greeter {
      // Sends a greeting
      rpc SayHello (HelloRequest) returns (HelloReply) {}
    }

    // The request message containing the user's name.
    message HelloRequest {
      string name = 1;
    }

    // The response message containing the greetings
    message HelloReply {
      string message = 1;
    }
    ```

- 生成示例代码

  [Go Generated Code | Protocol Buffers | Google Developers](https://developers.google.com/protocol-buffers/docs/reference/go-generated#invocation)

  每次修改完成proto后都要进行执行

    ```bash
    # source_relative 和proto文件同样的目录
    pwd
    # grpc-learn
    
    protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/hello.proto
    ```

- 依赖下载

    ```bash
    go mod tidy
    ```
  
# 持续输出命令行

核心代码

```go
cmd := exec.Command("ping", "127.0.0.1")

stdout, _ := cmd.StdoutPipe()

cmd.Start()

buf := bufio.NewReader(stdout)

line, _, _ := buf.ReadLine()
```


# grpc调试

```bash
go get github.com/fullstorydev/grpcui/...
# 在一个项目里面执行，需要 go module
go install github.com/fullstorydev/grpcui/cmd/grpcui
```

代码处理

```go
func main() {
	server := grpc.NewServer()

	// 。。。

	// 进行反射
	reflection.Register(server)
```

进行处理

```bash
grpcui -plaintext 127.0.0.1:1989

gRPC Web UI available at http://127.0.0.1:52191/
```