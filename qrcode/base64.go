package qrcode

// 生成登录二维码图片, 方便在网页上显示

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GetJPGBase64(content string, edges ...int) string {
	edgeLen := 300
	if len(edges) > 0 && edges[0] > 100 && edges[0] < 2000 {
		edgeLen = edges[0]
	}
	img, _ := qr.Encode(content, qr.L, qr.Unicode)
	img, _ = barcode.Scale(img, edgeLen, edgeLen)

	emptyBuff := bytes.NewBuffer(nil) // 开辟一个新的空buff缓冲区
	jpeg.Encode(emptyBuff, img, nil)
	dist := make([]byte, 50000)                        // 开辟存储空间
	base64.StdEncoding.Encode(dist, emptyBuff.Bytes()) // buff转成base64
	return "data:image/png;base64," + string(dist)     // 输出图片base64(type = []byte)
}

func GetPNGBase64(content string, edges ...int) string {
	edgeLen := 300
	if len(edges) > 0 && edges[0] > 100 && edges[0] < 2000 {
		edgeLen = edges[0]
	}
	img, _ := qr.Encode(content, qr.L, qr.Unicode)
	img, _ = barcode.Scale(img, edgeLen, edgeLen)

	emptyBuff := bytes.NewBuffer(nil) // 开辟一个新的空buff缓冲区
	png.Encode(emptyBuff, img)
	dist := make([]byte, 50000)                        // 开辟存储空间
	base64.StdEncoding.Encode(dist, emptyBuff.Bytes()) // buff转成base64
	return string(dist)                                // 输出图片base64(type = []byte)
}
