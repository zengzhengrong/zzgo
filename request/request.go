package request

import (
	"mime/multipart"
	"net/http"

	"github.com/spf13/cast"
)

const defaultMultipartMemory = 32 << 20 // 32 MB
var MaxUploadThreads int = 20

// GetFileOrFiles is 获取单个文件或者多个文件
// 多个文件以file_{0}开头读取form-data，最多20个 在map中 键名以文件名为键名
// 单文件以file读取form-data
func GetFileOrFiles(req *http.Request) (bool, map[string]*multipart.FileHeader, error) {
	var name string
	var IsSingle bool
	files := make(map[string]*multipart.FileHeader, 20)
	for i := 0; i < MaxUploadThreads; i++ {
		name = "file_" + cast.ToString(i)
		fs, err := fromfile(req, name)
		if err != nil && i == 0 {
			IsSingle = true
			break

		}
		if err != nil {
			// Not found break
			break
		}
		files[fs.Filename] = fs
	}
	if IsSingle && len(files) == 0 {
		sf, err := fromfile(req, "file")
		if err != nil {
			return IsSingle, files, err
		}
		files[sf.Filename] = sf
		return IsSingle, files, nil
	}
	return IsSingle, files, nil
}

func fromfile(req *http.Request, name string) (*multipart.FileHeader, error) {
	if req.MultipartForm == nil {
		if err := req.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := req.FormFile(name)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}
