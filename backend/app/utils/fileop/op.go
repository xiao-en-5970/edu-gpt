package fileop

import (
	"os"
	"path/filepath"
)

// FileInfo 是一个结构体，用于存储文件的信息
type FileInfo struct {
	Name    string
	Size    int64
	ModTime string // 可以根据需要格式化时间
	FileType string // 文件类型（如：txt, jpg, pdf等）
}

// GetAllFilesInfo 返回指定路径下的所有文件信息
func GetAllFilesInfo(dirPath string) ([]FileInfo, error) {
	var filesInfo []FileInfo

	// 使用 filepath.Walk 遍历目录
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 如果有错误，返回错误
		}

		// 只处理文件，不处理目录
		if !info.IsDir() {
			// 创建 FileInfo 结构体并添加到切片中
			fileInfo := FileInfo{
				Name:    info.Name(),
				Size:    info.Size(),
				ModTime: info.ModTime().Format("2006-01-02 15:04:05"),
				FileType: filepath.Ext(info.Name()), // 获取文件扩展名作为文件类型
			}
			filesInfo = append(filesInfo, fileInfo)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return filesInfo, nil
}
