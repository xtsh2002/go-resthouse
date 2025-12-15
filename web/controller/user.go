package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"
	pb "restproject/web/proto" //给包起别名，防止包名冲突
	"restproject/web/utils"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"

	// "github.com/micro/go-micro/registry"
	"go-micro.dev/v5"
	"go-micro.dev/v5/registry" // 统一使用 v5 版本的 registry 接口

	// "go-micro.dev/v5/registry/consul"// 统一使用 v5 版本的 Consul 注册中心
	"github.com/micro/plugins/v5/registry/consul"
)

func GetSeession(c *gin.Context) {
	//初始化错误返回的 map
	resp := make(map[string]string)
	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.GetRecodeText(utils.RECODE_SESSIONERR)
	c.JSON(http.StatusOK, resp)

}

// 获取图片信息
func GetCaptcha(c *gin.Context) {
	//获取图片验证码
	// uuid := c.Param("uuid")

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"), // 你的 Consul 服务地址（默认端口 8500）
	)

	//指定consul服务发现
	// consulReg := consul.NewRegistry()
	consulService := micro.NewService(micro.Registry(consulReg))

	// 关键：初始化微服务客户端（v5 版本必须调用 Init，否则无法正常发现服务）
	consulService.Init()

	//初始化服务
	microClient := pb.NewGetCaptchaService("go.micro.srv.getCaptcha", consulService.Client())

	//调用远程函数
	resp, err := microClient.Call(context.TODO(), &pb.Request{})
	if err != nil {
		fmt.Println("未找到远程服务...")
		return
	}
	// 将得到的数据,反序列化,得到图片数据
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)

	// 将图片写出到 浏览器.
	png.Encode(c.Writer, img)

}
