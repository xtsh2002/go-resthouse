package main

import (
	"getCaptcha/handler"
	pb "getCaptcha/proto"

	"go-micro.dev/v5"
	// 匿名导入 Consul 注册中心，让 go-micro 自动识别并使用
	_ "go-micro.dev/v5/registry/consul" // 注意导入路径的变化
)

func main() {
	//初始化consul服务发现，（v5版本的正确写法）
	// consulReg, err := registry.NewRegistry(
	// 	registry.Addrs("127.0.0.1:8500"), // 这里填写你的consul地址
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// 1. 直接创建微服务实例，在选项中配置 Consul 地址和服务信息
	// 无需单独初始化 registry，通过 micro.Registry 选项自动关联 Consul
	// 创建服务
	service := micro.NewService(
		// 服务名称（必须唯一，用于服务发现）
		micro.Name("go.micro.srv.getCaptcha"),
		// micro.Registry(consulReg),
		// 配置 Consul 注册中心地址（关键：指定你的 Consul 服务地址）
		micro.RegistryAddr("127.0.0.1:8500"),
		micro.Address("192.168.1.35:18901"), //服务地址,防止随机生成端口冲突
	)

	// 2. 初始化服务（v5 版本必须调用，完成内部配置加载）
	service.Init()

	// 3. 注册业务处理器（将自定义 handler 与 proto 接口绑定）
	if err := pb.RegisterGetCaptchaHandler(service.Server(), handler.New()); err != nil {
		panic("注册处理器失败: " + err.Error())
	}

	// 4. 启动服务（监听请求并注册到 Consul）
	if err := service.Run(); err != nil {
		panic("服务启动失败: " + err.Error())
	}
}
