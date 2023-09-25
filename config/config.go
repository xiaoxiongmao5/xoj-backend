package config

// 用于给用户分配accessKey,secretKey
const SALT = "xj"

// 用于给用户生成登录验证token（jwt）
const SecretKey = "your-secret-key"

// App配置数据
type AppConfiguration struct {
	Database struct {
		SavePath     string `json:"savePath"`
		MaxOpenConns int    `json:"maxOpenConns"`
	} `json:"database"`
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	Nacos struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"nacos"`
}

// App配置数据(可动态更新)
type AppConfigurationDynamic struct {
	IPWhiteList     []string `json:"ipWhiteList"`
	IPBlackList     []string `json:"ipBlackList"`
	IPAdminList     []string `json:"ipAdminList"`
	RateLimitConfig struct {
		RequestsPerSecond float64 `json:"requests_per_second"`
		BucketSize        int     `json:"bucket_size"`
	} `json:"rateLimitConfig"`
	Gatewayhost string `json:"gatewayhost"`
}

var (
	AppConfig        *AppConfiguration
	AppConfigDynamic *AppConfigurationDynamic
)
