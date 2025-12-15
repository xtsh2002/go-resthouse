package main

import (
	"getCaptcha/handler"
	pb "getCaptcha/proto"

	// 新增：用于设置环境变量
	"go-micro.dev/v5"

	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5/registry"
)

func main() {
	// 1. 显式初始化 Consul 注册中心（指定 Consul 地址，兼容所有 v5 版本）
	// 这里直接使用 consul.NewRegistry()，并通过 consul.Addrs() 配置地址
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"), // 你的 Consul 服务地址（默认端口 8500）
	)
	// if err != nil {
	// 	panic(err)
	// }

	// 关键步骤：通过环境变量指定 Consul 地址（go-micro 所有版本都支持）
	// 环境变量 "MICRO_REGISTRY" 用于指定注册中心类型（这里填 consul）
	// 环境变量 "MICRO_REGISTRY_ADDRESS" 用于指定 Consul 服务地址
	// os.Setenv("MICRO_REGISTRY", "consul")
	// os.Setenv("MICRO_REGISTRY_ADDRESS", "127.0.0.1:8500")

	// 1. 创建微服务实例：直接在选项中配置 Consul 地址，无需单独初始化 registry
	// 核心：通过 registry.Option 类型的 "registry.address" 配置 Consul 地址
	service := micro.NewService(
		// 服务名称（必须唯一，用于服务发现）
		micro.Name("go.micro.srv.getCaptcha"), // 服务唯一名称（服务发现用）
		micro.Registry(consulReg),             // 关键：绑定 Consul 注册中心
		// 配置 Consul 注册中心地址（关键：指定你的 Consul 服务地址）
		// micro.RegistryAddr("127.0.0.1:8500"),
		micro.Address("192.168.1.35:18901"), //服务地址,防止随机生成端口冲突
		// 关键：通过 micro.Registry 选项配置 Consul 地址（兼容所有 v5 版本）
		// micro.Registry(registry.NewRegistry(
		// 	// 通过 registry.Addrs() 配置 Consul 服务地址（兼容所有 v5 版本）
		// 	registry.Addrs("127.0.0.1:8500"), // 这里填写你的 Consul 地址和端口
		// 	// 显式指定使用 consul 作为注册中心驱动
		// 	// registry.Plugin("consul"),
		// )),
		// 指定服务版本（可选，用于区分不同版本的微服务）
		// micro.Version("v1"),

	)

	//  2. 必须调用 Init() 完成服务内部初始化（v5 版本强制要求，加载配置/注册中心等）
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
