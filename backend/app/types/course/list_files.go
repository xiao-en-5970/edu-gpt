package types

type CourseFilesReq struct {
	CourseID uint `json:"course_id" form:"course_id"` // 课程ID
}

type CourseFilesResp struct {
	FileList []File `json:"file_list" form:"file_list"`
}

type File struct {
	FileName string `json:"file_name" form:"file_name"`
	FileSize int64  `json:"file_size" form:"file_size"`
	ModTime  string `json:"mod_time" form:"mod_time"` // 可以根据需要格式化时间
	FileType string `json:"file_type" form:"file_type"` // 文件类型
	FileURL  string `json:"file_url" form:"file_url"`   // 文件下载
}