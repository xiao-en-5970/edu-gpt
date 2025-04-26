package types

type HFUTStudentInfoResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data HFUTStudentInfo `json:"data"`
}

type HFUTStudentInfo struct {
	StudentCode     string `json:"studentCode"`
	StudentId       string `json:"studentId"`
	UsernameEn      string `json:"usernameEn"`
	UsernameZh      string `json:"usernameZh"`
	Sex             string `json:"sex"`
	CultivateType   string `json:"cultivateType"`
	Department      string `json:"department"`
	Grade           string `json:"grade"`
	Level           string `json:"level"`
	StudentType     string `json:"studentType"`
	Major           string `json:"major"`
	Class           string `json:"class"`
	Campus          string `json:"campus"`
	Status          string `json:"status"`
	Length          string `json:"length"`
	EnrollmentDate  string `json:"enrollmentDate"`
	GraduateDate    string `json:"graduateDate"`
}