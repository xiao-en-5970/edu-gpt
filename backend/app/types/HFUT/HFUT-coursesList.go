package types

type CourseSchedule struct {
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
	ID        int    `json:"id"`
	Room      string `json:"room"`
	Weekday   int    `json:"weekday"`
	WeekIndex int    `json:"weekIndex"`
}

type CourseTable struct {
	ID           int              `json:"id"`
	Code         string           `json:"code"`
	AdminClasses string           `json:"adminClasses"`
	Name         string           `json:"name"`
	Type         string           `json:"type"`
	Teachers     []string         `json:"teachers"`
	StudentCount int              `json:"studentCount"`
	Weeks        string           `json:"weeks"`
	Credits      float64          `json:"credits"`
	ExamMode     string           `json:"examMode"`
	Schedule     []CourseSchedule `json:"schedule"`
}

type CourseListResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []CourseTable `json:"data"`
}
