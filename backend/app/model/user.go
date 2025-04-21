package model

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/xiao-en-5970/Goodminton/backend/app/global"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/bcrypts"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AccountStatus string



type User struct {
	ID                  int64         	`gorm:"primaryKey;autoIncrement;comment:用户编号"`
	Username            string         `gorm:"type:varchar(50);not null;uniqueIndex;comment:登录用户名（唯一）"`
	PasswordHash        string         `gorm:"type:varchar(255);not null;comment:加密后的密码"`
	Email               sql.NullString `gorm:"type:varchar(100);uniqueIndex;comment:邮箱（可选）"`
	Phone               sql.NullString `gorm:"type:varchar(20);uniqueIndex;comment:手机号（可选）"`
	AccountStatus       AccountStatus  `gorm:"type:ENUM('active', 'locked', 'disabled');default:'active';comment:账号状态"`
	LastLoginTime       sql.NullTime   `gorm:"comment:最后登录时间"`
	
	Nickname            string `gorm:"type:varchar(50);comment:可选，仅供查看"`
	AvatarPath          string         `gorm:"type:varchar(255);default:'default-avatar.png';comment:头像路径(相对路径)"`
	SelfEvaluatedLevel  sql.NullString `gorm:"type:varchar(20);comment:自评技术水平"`
	SystemScore         sql.NullInt32  `gorm:"comment:系统评估得分(0~100)"`
	PersonalityTags     datatypes.JSON `gorm:"comment:性格标签"`
	PlayStyleTags       datatypes.JSON `gorm:"comment:打球风格"`
	PreferredSkillLevel sql.NullString `gorm:"type:varchar(20);comment:希望对手技术水平"`
	PreferredTimeSlots  datatypes.JSON `gorm:"comment:时间偏好"`
	PreferredRegions    datatypes.JSON `gorm:"comment:常活动区域"`
	MaxCost             sql.NullInt32  `gorm:"comment:可接受的花销（单位：元）"`
	HistoricalPartners  datatypes.JSON `gorm:"comment:历史搭档ID列表"`
	RatingsGiven        datatypes.JSON `gorm:"comment:对别人的评价(用户ID:评分)"`
	CreateTime          time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime          time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
}

// TableName 设置表名
func (User) TableName() string {
	return "user"
}

// BeforeCreate 钩子 - 创建前设置默认值
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.AvatarPath == "" {
		u.AvatarPath = "default-avatar.png"
	}
	return nil
}

// CheckConstraints 自定义检查约束
func (u *User) CheckConstraints() error {
	if u.SystemScore.Valid && (u.SystemScore.Int32 < 0 || u.SystemScore.Int32 > 100) {
		return gorm.ErrInvalidValue
	}
	if u.MaxCost.Valid && u.MaxCost.Int32 < 0 {
		return gorm.ErrInvalidValue
	}
	return nil
}

// BeforeSave 钩子 - 保存前验证
func (u *User) BeforeSave(tx *gorm.DB) error {
	return u.CheckConstraints()
}

// BeforeUpdate 钩子 - 更新前验证
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	return u.CheckConstraints()
}


func FindUserByName(username string)(*User,error){
	user:=&User{}
	err:=global.Db.Model(user).Where("username=?",username).First(user).Error
	if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // 用户不存在
        }
        return nil, err // 其他数据库错误
    }
    return user, nil // 用户存在
}

func InsertUser(user *User)(err error){
	result := global.Db.Create(user) // 通过指针传递数据
	if result.Error != nil {
		// 处理错误
		global.Logger.Warnf("创建记录失败: %v", result.Error)
	}
	global.Logger.Infof("插入成功，ID: %d\n", user.ID)
	return nil
}

func ConvertUserToUserInfo(dbUser *User) (*types.UserInfo, error) {
    // 处理JSON字段
    personalityTags, err := parseJSONToStringSlice(dbUser.PersonalityTags)
    if err != nil {
        return &types.UserInfo{}, fmt.Errorf("解析性格标签失败: %v", err)
    }

    playStyleTags, err := parseJSONToStringSlice(dbUser.PlayStyleTags)
    if err != nil {
        return &types.UserInfo{}, fmt.Errorf("解析打球风格失败: %v", err)
    }

    preferredTimeSlots, err := parseJSONToStringSlice(dbUser.PreferredTimeSlots)
    if err != nil {
        return &types.UserInfo{}, fmt.Errorf("解析时间偏好失败: %v", err)
    }

    preferredRegions, err := parseJSONToStringSlice(dbUser.PreferredRegions)
    if err != nil {
        return &types.UserInfo{}, fmt.Errorf("解析常活动区域失败: %v", err)
    }
    
    historicalPartners, err := parseJSONToStringSlice(dbUser.HistoricalPartners)
    if historicalPartners !=nil{
        if err != nil {
            return &types.UserInfo{}, fmt.Errorf("解析历史搭档失败: %v", err)
        }
    }
    

    ratingsGiven, err := parseJSONToStringSlice(dbUser.RatingsGiven)
    if ratingsGiven!=nil{
        if err != nil {
            return &types.UserInfo{}, fmt.Errorf("解析评价记录失败: %v", err)
        }
    }
    

    // 构建UserInfo
    userInfo := &types.UserInfo{
        ID:                 dbUser.ID,
        Username:           dbUser.Username,
        PasswordHash:       dbUser.PasswordHash,
        Email:              getStringFromNullString(dbUser.Email),
        Phone:              getStringFromNullString(dbUser.Phone),
        LastLoginTime:      formatNullTime(dbUser.LastLoginTime),
        Nickname:           dbUser.Nickname,
        AvatarPath:         dbUser.AvatarPath,
        SelfEvaluatedLevel: getStringFromNullString(dbUser.SelfEvaluatedLevel),
        SystemScore:        getInt32FromNullInt32(dbUser.SystemScore),
        PersonalityTags:    personalityTags,
        PlayStyleTags:      playStyleTags,
        PreferredSkillLevel: getStringFromNullString(dbUser.PreferredSkillLevel),
        PreferredTimeSlots: preferredTimeSlots,
        PreferredRegions:   preferredRegions,
        MaxCost:            getInt32FromNullInt32(dbUser.MaxCost),
        HistoricalPartners: historicalPartners,
        RatingsGiven:       ratingsGiven,
        CreateTime:         dbUser.CreateTime.Format("2006-01-02 15:04:05"),
        UpdateTime:         dbUser.UpdateTime.Format("2006-01-02 15:04:05"),
    }

    return userInfo, nil
}

// 辅助函数：从sql.NullString获取字符串
func getStringFromNullString(ns sql.NullString) string {
    if ns.Valid {
        return ns.String
    }
    return ""
}

// 辅助函数：从sql.NullInt32获取int32
func getInt32FromNullInt32(ni sql.NullInt32) int32 {
    if ni.Valid {
        return ni.Int32
    }
    return 0
}

// 辅助函数：格式化sql.NullTime
func formatNullTime(nt sql.NullTime) string {
    if nt.Valid {
        return nt.Time.Format("2006-01-02 15:04:05")
    }
    return ""
}

// 辅助函数：解析JSON字段到[]string
func parseJSONToStringSlice(jsonData datatypes.JSON) ([]string, error) {
    if jsonData == nil || len(jsonData) == 0 {
        return []string{}, nil
    }

    var result []string
    err := json.Unmarshal(jsonData, &result)
    if err != nil {
        return nil, err
    }
    return result, nil
}
func ConvertUserInfoToUser(userInfo *types.UserInfo) (*User, error) {
    // 处理JSON字段序列化
    personalityTags, err := json.Marshal(userInfo.PersonalityTags)
    if err != nil {
        return &User{}, fmt.Errorf("序列化性格标签失败: %v", err)
    }

    playStyleTags, err := json.Marshal(userInfo.PlayStyleTags)
    if err != nil {
        return &User{}, fmt.Errorf("序列化打球风格失败: %v", err)
    }

    preferredTimeSlots, err := json.Marshal(userInfo.PreferredTimeSlots)
    if err != nil {
        return &User{}, fmt.Errorf("序列化时间偏好失败: %v", err)
    }

    preferredRegions, err := json.Marshal(userInfo.PreferredRegions)
    if err != nil {
        return &User{}, fmt.Errorf("序列化常活动区域失败: %v", err)
    }

    historicalPartners, err := json.Marshal(userInfo.HistoricalPartners)
    if err != nil {
        return &User{}, fmt.Errorf("序列化历史搭档失败: %v", err)
    }

    ratingsGiven, err := json.Marshal(userInfo.RatingsGiven)
    if err != nil {
        return &User{}, fmt.Errorf("序列化评价记录失败: %v", err)
    }

    // 处理时间字段（字符串 -> time.Time）
    createTime, err := time.Parse("2006-01-02 15:04:05", userInfo.CreateTime)
    if err != nil {
        return &User{}, fmt.Errorf("解析创建时间失败: %v", err)
    }

    updateTime, err := time.Parse("2006-01-02 15:04:05", userInfo.UpdateTime)
    if err != nil {
        return &User{}, fmt.Errorf("解析更新时间失败: %v", err)
    }

    // 处理最后登录时间（可能为空字符串）
    var lastLoginTime sql.NullTime
    if userInfo.LastLoginTime != "" {
        t, err := time.Parse("2006-01-02 15:04:05", userInfo.LastLoginTime)
        if err != nil {
            return &User{}, fmt.Errorf("解析最后登录时间失败: %v", err)
        }
        lastLoginTime = sql.NullTime{Time: t, Valid: true}
    }

    // 构建User结构体
    user := &User{
        ID:                  userInfo.ID,
        Username:           userInfo.Username,
        PasswordHash:       userInfo.PasswordHash,
        Email:              toNullString(userInfo.Email),
        Phone:              toNullString(userInfo.Phone),
        AccountStatus:      "active", // 默认值，可根据业务调整
        LastLoginTime:      lastLoginTime,
        Nickname:           userInfo.Nickname,
        AvatarPath:         userInfo.AvatarPath,
        SelfEvaluatedLevel: toNullString(userInfo.SelfEvaluatedLevel),
        SystemScore:        toNullInt32(userInfo.SystemScore),
        PersonalityTags:    personalityTags,
        PlayStyleTags:      playStyleTags,
        PreferredSkillLevel: toNullString(userInfo.PreferredSkillLevel),
        PreferredTimeSlots: preferredTimeSlots,
        PreferredRegions:   preferredRegions,
        MaxCost:            toNullInt32(userInfo.MaxCost),
        HistoricalPartners: historicalPartners,
        RatingsGiven:       ratingsGiven,
        CreateTime:         createTime,
        UpdateTime:         updateTime,
    }

    return user, nil
}

// 辅助函数：string -> sql.NullString
func toNullString(s string) sql.NullString {
    return sql.NullString{String: s, Valid: s != ""}
}

// 辅助函数：int32 -> sql.NullInt32
func toNullInt32(i int32) sql.NullInt32 {
    return sql.NullInt32{Int32: i, Valid: i != 0}
}

func RegisterReqToUser(req *types.RegisterReq) (*User, error) {
	// 1. 密码加密（bcrypt加密）
	hashedPassword, err := bcrypts.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("密码加密失败: %v", err)
	}

	// 2. 序列化JSON字段
	personalityTags, err := json.Marshal(req.PersonalityTags)
	if err != nil {
		return nil, fmt.Errorf("性格标签序列化失败: %v", err)
	}

	playStyleTags, err := json.Marshal(req.PlayStyleTags)
	if err != nil {
		return nil, fmt.Errorf("打球风格序列化失败: %v", err)
	}

	preferredTimeSlots, err := json.Marshal(req.PreferredTimeSlots)
	if err != nil {
		return nil, fmt.Errorf("时间偏好序列化失败: %v", err)
	}

	preferredRegions, err := json.Marshal(req.PreferredRegions)
	if err != nil {
		return nil, fmt.Errorf("偏好区域序列化失败: %v", err)
	}

	// 3. 构建User对象
	user := &User{
		Username:            req.Username,
		PasswordHash:        string(hashedPassword),
		Nickname:            req.Nickname,
		AvatarPath:          req.AvatarPath,
		SelfEvaluatedLevel:  toNullString(req.SelfEvaluatedLevel),
		SystemScore:         toNullInt32(int32(req.SystemScore)),
		PersonalityTags:     personalityTags,
		PlayStyleTags:       playStyleTags,
		PreferredSkillLevel: toNullString(req.PreferredSkillLevel),
		PreferredTimeSlots:  preferredTimeSlots,
		PreferredRegions:    preferredRegions,
		MaxCost:             toNullInt32(int32(req.MaxCost)),
		HistoricalPartners:  datatypes.JSON("[]"), // 初始化为空数组
		RatingsGiven:        datatypes.JSON("{}"), // 初始化为空对象
	}

	return user, nil
}


func UpdateUser(newuser * User)(error){
    return global.Db.Model(&User{ID: newuser.ID}).Updates(*newuser).Error
}
