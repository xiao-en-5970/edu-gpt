package conf
import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)



type Config struct {
	Server   Server   		`mapstructure:"server"`
	MySQL    MysqlConfig 	`mapstructure:"mysql"`
	Redis    RedisConfig    `mapstructure:"redis"`
	HfutAPI  HfutAPI      	`mapstructure:"hfut-api"`
	Logging  LoggingConfig  `mapstructure:"logging"`
	Auth 	 Auth 			`mapstructure:"auth"`
}
type Auth struct{ 
	MaxAge time.Duration `mapstructure:"max_age"`
}
type Server struct {
	Port    int           `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type MysqlConfig struct {
	Host        string     	`mapstructure:"host"`
	Port        int        	`mapstructure:"port"`
	Db 			string 		`mapstructure:"db"`
	Credentials Credentials `mapstructure:"credentials"`
	MaxIdleConns int 		`mapstructure:"maxidleconns"`
	SetMaxOpenConns int 	`mapstructure:"maxopenconns"`
	ConnMaxLifetime time.Duration `mapstructure:"connmaxlifetime"`
}

type RedisConfig struct {
	Addr string `mapstructure:"addr"`
	DB   int    `mapstructure:"db"`
	PoolSize int `mapstructure:"poolsize"`
}

type HfutAPI struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Retry int 	`mapstructure:"retry"`
}

type LoggingConfig struct {
	Level string `mapstructure:"level"`
}

type Credentials struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func ConfInit(path string) (cfg *Config,err error){
	viper.SetConfigFile(path)  // 指定配置文件路径
	viper.SetConfigType("yaml") // 显式设置配置类型
	cfg = &Config{}
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	return cfg, nil
}