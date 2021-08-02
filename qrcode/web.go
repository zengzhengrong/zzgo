package qrcode

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func WritePng(w http.ResponseWriter, content string, edges ...int) error {
	edgeLen := 300
	if len(edges) > 0 && edges[0] > 100 && edges[0] < 2000 {
		edgeLen = edges[0]
	}
	img, _ := qr.Encode(content, qr.L, qr.Unicode)
	img, _ = barcode.Scale(img, edgeLen, edgeLen)
	buff := bytes.NewBuffer(nil)
	png.Encode(buff, img)
	w.Header().Set("Content-Type", "image/png")
	_, err := w.Write(buff.Bytes())
	return err
}

func WriteJpg(w http.ResponseWriter, content string, edges ...int) error {
	edgeLen := 300
	if len(edges) > 0 && edges[0] > 100 && edges[0] < 2000 {
		edgeLen = edges[0]
	}
	img, _ := qr.Encode(content, qr.L, qr.Unicode)
	img, _ = barcode.Scale(img, edgeLen, edgeLen)
	buff := bytes.NewBuffer(nil)
	jpeg.Encode(buff, img, nil)
	w.Header().Set("Content-Type", "image/jpg")
	_, err := w.Write(buff.Bytes())
	return err
}
