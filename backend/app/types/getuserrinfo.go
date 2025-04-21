package types

type GetUserInfoReq struct {
    
}

type GetUserInfoResp struct {
	ID                  int64    `json:"id"`                   // 用户编号
	Username            string   `json:"username"`            // 登录用户名（唯一）
	PasswordHash        string   `json:"password_hash"`        // 加密后的密码
	Email               string   `json:"email,omitempty"`     // 邮箱（可选）
	Phone               string   `json:"phone,omitempty"`     // 手机号（可选）
	LastLoginTime       string   `json:"last_login_time"`     // 最后登录时间（字符串格式）
	Nickname            string   `json:"nickname,omitempty"`  // 可选，仅供查看
	AvatarPath          string   `json:"avatar_path"`         // 头像路径(相对路径)
	SelfEvaluatedLevel  string   `json:"self_evaluated_level"` // 自评技术水平
	SystemScore         int32    `json:"system_score"`        // 系统评估得分(0~100)
	PersonalityTags     []string `json:"personality_tags"`    // 性格标签
	PlayStyleTags       []string `json:"play_style_tags"`     // 打球风格
	PreferredSkillLevel string   `json:"preferred_skill_level"` // 希望对手技术水平
	PreferredTimeSlots  []string `json:"preferred_time_slots"` // 时间偏好
	PreferredRegions    []string `json:"preferred_regions"`   // 常活动区域
	MaxCost             int32    `json:"max_cost"`            // 可接受的花销（单位：元）
	HistoricalPartners  []string `json:"historical_partners"`  // 历史搭档ID列表
	RatingsGiven        []string `json:"ratings_given"`       // 对别人的评价(用户ID:评分)
	CreateTime          string   `json:"create_time"`         // 创建时间（字符串格式）
	UpdateTime          string   `json:"update_time"`         // 更新时间（字符串格式）
}