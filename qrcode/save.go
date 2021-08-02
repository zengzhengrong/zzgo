package qrcode

// 生成登录二维码图片

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func SaveJpegFile(filePath, content string, edges ...int) error {
	edgeLen := 300
	if len(edges) > 0 && edges[0] > 100 && edges[0] < 2000 {
		edgeLen = edges[0]
	}
	img, _ := qr.Encode(content, qr.L, qr.Unicode)
	img, _ = barcode.Scale(img, edgeLen, edgeLen)

	return writeFile(filePath, img, "jpg")
}

func SavePngFile(filePath, content string, edges ...int) error {
	edgeLen := 300
	if len(edges) > 0 && edges[0] > 100 && edges[0] < 2000 {
		edgeLen = edges[0]
	}
	img, _ := qr.Encode(content, qr.L, qr.Unicode)
	img, _ = barcode.Scale(img, edgeLen, edgeLen)

	return writeFile(filePath, img, "png")
}

func writeFile(filePath string, img image.Image, format string) error {
	if err := createDir(filePath); err != nil {
		return err
	}
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}
	defer file.Close()
	switch strings.ToLower(format) {
	case "png":
		err = png.Encode(file, img)
	case "jpg":
		err = jpeg.Encode(file, img, nil)
	default:
		return errors.New("format not accept")
	}
	if err != nil {
		return err
	}
	return nil
}

func createDir(filePath string) error {
	var err error
	dirPath := filepath.Dir(filePath)
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		if !os.IsExist(err) {
			err = os.MkdirAll(dirPath, 0777)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if dirInfo.IsDir() {
			return nil
		}
		return errors.New("directory is a file")
	}
	return nil
}
