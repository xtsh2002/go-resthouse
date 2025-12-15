## 三、web端对接微服务实现

### 1、拷贝密码本。 将 service 下的 proto/  拷贝 web/   下、

在web去调用微服务，实现验证码功能：在controller/user.go中写方法：GetCaptcha

如果web崩溃，微服务端服务仍然存在

客户端与微服务端要进行交互，交互的时候就要用到proto密码本，

微服务和web服务可能不在同一主机上，所以在使用的时候，要把微服务端的proto目录，拷贝到web端

![](H:\GO语言学习\image\ScreenShot_2025-12-14_191054_212.png)





### 2、在控制器的GetCaptcha中导入包，起别名

```go
pb "restproject/web/proto" //给包起别名，防止包名冲突
```

### 3、指定consul 服务发现：

```go
//指定consul服务发现
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"), // 你的 Consul 服务地址（默认端口 8500）
	)

	consulService := micro.NewService(micro.Registry(consulReg))

// 关键：初始化微服务客户端（v5 版本必须调用 Init，否则无法正常发现服务）
	consulService.Init()
```

### 4、初始化客户端

```go
//初始化服务
	microClient := pb.NewGetCaptchaService("go.micro.srv.getCaptcha", consulService.Client())

```

### 5、调用远程函数

```go
//调用远程函数
	resp, err := microClient.Call(context.TODO(), &pb.Request{})
	if err != nil {
		fmt.Println("未找到远程服务...")
		return
	}
```

### 6、将得到的数据,反序列化,得到图片数据

```go
// 将得到的数据,反序列化,得到图片数据
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)
```



### 7、将图片写出到 浏览器.

```go
png.Encode(c.Writer, img)
```



### 8、测试：

1. 启动 consul  ，  consul agent -dev
2. 启动 service/getCaptcha  下的  main.go
3. 启动 web/ 下的  main.go
4. 浏览器中 192.168.IP: port/home    点击注册 查看图片验证码！



### 9、上传代码到github



```bash
# 初始化仓库
git init

# 添加文件
git add .

# 提交文件
git commit -m "初始提交"

# 关联远程仓库
git remote add origin https://github.com/xtsh2002/go-resthouse.git

# 上传代码
git push -u origin master

```

以后提交,**本地代码修改后再次提交的步骤如下：**

```bash
# 添加修改的文件
git add .

# 提交修改
git commit -m "修复了bug"

# 推送修改
git push origin master

```





