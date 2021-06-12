package config

import "fmt"

// HTTPType http 配置信息
type HTTPType struct {
	Port             string `yaml:"Port"`
	MaxContentLength int    `yaml:"MaxContentLength"`
	ShutdownTimeout  int    `yaml:"ShutdownTimeout"`
	MaxLoggerLength  int    `yaml:"MaxLoggerLength"`
}

// CORSType 跨域设置
type CORSType struct {
	Enable           bool     `yaml:"Enable"`
	AllowOrigins     []string `yaml:"AllowOrigins"`
	AllowMethods     []string `yaml:"AllowMethods"`
	AllowHeaders     []string `yaml:"AllowHeaders"`
	AllowCredentials bool     `yaml:"AllowCredentials"`
	MaxAge           int      `yaml:"MaxAge"`
}

// GORMType gorm 配置信息
type GORMType struct {
	Debug             bool `yaml:"Debug"`
	MaxLifetime       int  `yaml:"MaxLifetime"`
	MaxOpenConns      int  `yaml:"MaxOpenConns"`
	MaxIdleConns      int  `yaml:"MaxIdleConns"`
	EnableAutoMigrate bool `yaml:"EnableAutoMigrate"`
}

// DBType 数据库配置定义
type DBType struct {
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	User       string `yaml:"User"`
	Password   string `yaml:"Password"`
	DBName     string `yaml:"DBName"`
	Parameters string `yaml:"Parameters"`
}

// DSN 得到数据库连接
func (d *DBType) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DBName,
		d.Parameters,
	)
}

// LogType 日志配置类型定义
type LogType struct {
	Level  int8   `yaml:"Level"`
	Output string `yaml:"Output"`
}

// LogFileHookType 文件归档钩子配置
type LogFileHookType struct {
	Filename   string `yaml:"Filename"`
	MaxSize    int    `yaml:"Maxsize"`
	MaxBackups int    `yaml:"MaxBackups"`
	MaxAge     int    `yaml:"Maxage"`
	Compress   bool   `yaml:"Compress"`
}

// JWT 数据库配置定义
type JwtType struct {
	Secret  string `yaml:"Secret"`
	Iss     string `yaml:"Iss"`
	ExpTime int    `yaml:"ExpTime"`
}

//Redis 配置定义
type RedisType struct {
	Enable string `yaml:"Enable"`
	Port   int    `yaml:"Port"`
}

type OssType struct {
	UpUrl         string `yaml:"UpUrl"`
	UpUrlInternal string `yaml:"UpUrlInternal"`
	DownUrl       string `yaml:"DownUrl"`
	AK            string `yaml:"AccessKey"`
	SK            string `yaml:"SecretKey"`
	Bucket        string `yaml:"Bucket"`
}

// CType 配置文件类型定义
type CType struct {
	Mode        string          `yaml:"Mode"`
	HTTP        HTTPType        `yaml:"HTTP"`
	CORS        CORSType        `yaml:"CORS"`
	GORM        GORMType        `yaml:"GORM"`
	DB          DBType          `yaml:"DB"`
	Log         LogType         `yaml:"Log"`
	LogFileHook LogFileHookType `yaml:"LogFileHook"`
	JWT         JwtType         `yaml:"JWT"`
	REDIS       RedisType       `yaml:"Redis"`
	Oss         OssType         `yaml:"OSS"`
}
