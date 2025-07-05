package conf

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server  Server        `mapstructure:"server"`
	MySQL   MysqlConfig   `mapstructure:"mysql"`
	Redis   RedisConfig   `mapstructure:"redis"`
	HfutAPI HfutAPI       `mapstructure:"hfut-api"`
	Logging LoggingConfig `mapstructure:"logging"`
	Auth    Auth          `mapstructure:"auth"`
	Static  Static        `mapstructure:"static"`
	Cache   Cache         `mapstructure:"cache"` // 缓存配置
}
type Auth struct {
	MaxAge time.Duration `mapstructure:"max_age"`
}
type Server struct {
	Address string        `mapstructure:"address"`
	Port    int           `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type MysqlConfig struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Db              string        `mapstructure:"db"`
	Credentials     Credentials   `mapstructure:"credentials"`
	MaxIdleConns    int           `mapstructure:"maxidleconns"`
	SetMaxOpenConns int           `mapstructure:"maxopenconns"`
	ConnMaxLifetime time.Duration `mapstructure:"connmaxlifetime"`
}

type RedisConfig struct {
	Addr         string `mapstructure:"addr"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"poolsize"`
	CookieExpire int    `mapstructure:"cookie_expire"` //单位小时
	LikeExpire   int    `mapstructure:"like_expire"`
}

type HfutAPI struct {
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Retry int    `mapstructure:"retry"`
}
type Static struct {
	RootPath      string `mapstructure:"root_path"`
	AvatarsPath   string `mapstructure:"avatars_path"`
	BackImagePath string `mapstructure:"backimage_path"`
	PostPath      string `mapstructure:"post_path"`
	BookPath      string `mapstructure:"book_path"`
}
type LoggingConfig struct {
	StdOutLevel string `mapstructure:"stdout_level"` // 控制台日志级别
	FilePath    string `mapstructure:"file_path"`    // 日志文件路径
	FileLevel   string `mapstructure:"file_level"`   // 文件日志级别
	MaxSize     int    `mapstructure:"max_size"`     // 每个日志文件的大小，单位MB
	MaxBackups  int    `mapstructure:"max_backups"` // 保留的旧日志文件个数
	MaxAge      int    `mapstructure:"max_age"`    // 日志文件的最大保存天数
	Compress    bool   `mapstructure:"compress"`    // 是否压缩日志文件
}

type Credentials struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Cache struct {
	ReadExpiration  int `mapstructure:"read_expiration"`  // 读取缓存过期时间，单位秒
	WriteExpiration int `mapstructure:"write_expiration"` // 写入缓存过期时间，单位秒
}

func ConfInit(path string) (cfg *Config, err error) {
	viper.SetConfigFile(path)   // 指定配置文件路径
	viper.SetConfigType("yaml") // 显式设置配置类型
	cfg = &Config{}
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}
	os.MkdirAll(cfg.Static.RootPath, 0755)
	os.MkdirAll(cfg.Static.AvatarsPath, 0755)
	os.MkdirAll(cfg.Static.BackImagePath, 0755)
	os.MkdirAll(cfg.Static.PostPath, 0755)
	return cfg, nil
}
