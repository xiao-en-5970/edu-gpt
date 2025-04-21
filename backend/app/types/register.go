package types


type RegisterReq struct {
    Username             string   `json:"username"`               // 用户名/学号
    Password             string   `json:"password"`               // 密码
    Nickname             string   `json:"nickname"`               // 昵称
    AvatarPath           string   `json:"avatar_path"`            // 头像路径
    SelfEvaluatedLevel   string   `json:"self_evaluated_level"`   // 自评水平
    SystemScore          int      `json:"system_score"`           // 系统评分
    PersonalityTags      []string `json:"personality_tags"`       // 性格标签
    PlayStyleTags        []string `json:"play_style_tags"`        // 打球风格标签
    PreferredSkillLevel  string   `json:"preferred_skill_level"`  // 期望对手水平
    PreferredTimeSlots   []string `json:"preferred_time_slots"`   // 偏好时间段
    PreferredRegions     []string `json:"preferred_regions"`      // 偏好区域
    MaxCost              int      `json:"max_cost"`               // 最高消费限制
}
type RegisterResp struct{

}