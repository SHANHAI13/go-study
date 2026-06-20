package config

// 项目配置（实际项目从 yaml 文件读取，这里先用常量简单处理）
const (
	ServerPort = ":8080"

	// MySQL
	MySQLDSN = "root:123456@tcp(127.0.0.1:3306)/campus_recruit?charset=utf8mb4&parseTime=True&loc=Local"

	// Redis
	RedisAddr     = "127.0.0.1:6379"
	RedisPassword = ""
	RedisDB       = 0

	// JWT
	JWTSecret = "campus-recruit-secret-key"
	JWTExpire = 24 // 小时
)
