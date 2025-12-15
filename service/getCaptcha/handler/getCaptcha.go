package handler

import (
	"context"
	"fmt"

	// log "go-micro.dev/v5/logger"
	"encoding/json"
	pb "getCaptcha/proto"
	"image/color" // 导入颜色包

	"github.com/afocus/captcha"
)

type GetCaptcha struct{}

// Return a new handler
func New() *GetCaptcha {
	return &GetCaptcha{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetCaptcha) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {

	// 生成图片验证码

	// 初始化对象
	cap := captcha.New()

	// 设置字体
	cap.SetFont("./conf/comic.ttf")

	// 设置验证码大小
	cap.SetSize(128, 64)

	// 设置干扰强度
	cap.SetDisturbance(captcha.NORMAL)

	// 设置前景色
	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})

	// 设置背景色
	cap.SetBkgColor(color.RGBA{100, 0, 255, 255}, color.RGBA{255, 0, 127, 255}, color.RGBA{255, 255, 10, 255})

	// 生成字体
	img, str := cap.Create(4, captcha.NUM)
	fmt.Println(str)
	// 存储图片验证码到 redis 中
	// err := model.SaveImgCode(str, req.Uuid)
	// if err != nil {
	// 	return err
	// }

	// 将 生成成的图片 序列化.
	imgBuf, _ := json.Marshal(img)

	// 将 imgBuf 使用 参数 rsp 传出
	rsp.Img = imgBuf
	return nil
}
