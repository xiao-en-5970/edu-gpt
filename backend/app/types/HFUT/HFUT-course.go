package types

type HFUTCoursesResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data HFUTCourses `json:"data"`
}

type HFUTCourses struct {
	List []Course `json:"list"`
	Page Page     `json:"page"`
}

type Course struct {
	CourseName string   `json:"courseName"`
	CourseCode string   `json:"courseCode"`
	CourseType string   `json:"courseType"`
	Credits    float64  `json:"credits"`
	ClassName  string   `json:"className"`
	ClassCode  string   `json:"classCode"`
	OpenDepart string   `json:"openDepart"`
	ExamMod    string   `json:"examMod"`
	Campus     string   `json:"campus"`
	Teachers   []string `json:"teachers"`
	Schedule   []string `json:"schedule"`
}

type Page struct {
	CurrentPage int `json:"currentPage"`
	RowsInPage  int `json:"rowsInPage"`
	RowsPerPage int `json:"rowsPerPage"`
	TotalRows   int `json:"totalRows"`
	TotalPages  int `json:"totalPages"`
}
