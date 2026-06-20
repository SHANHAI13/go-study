# 校园招聘平台 - Campus Recruit

Go + Gin 后端项目，专为秋招简历准备。

## 技术栈

- **语言**：Go 1.21+
- **Web 框架**：Gin
- **ORM**：GORM
- **数据库**：MySQL 8.0
- **缓存**：Redis
- **认证**：JWT (golang-jwt)

## 功能模块

- [ ] 用户注册/登录（JWT 鉴权）
- [ ] 职位发布与管理（企业端）
- [ ] 职位列表查询（分页 + 条件筛选 + Redis 缓存）
- [ ] 简历上传与管理（学生端）
- [ ] 职位投递
- [ ] 职位收藏
- [ ] 统一响应格式 & 全局错误处理
- [ ] Swagger 接口文档

## 项目结构

```
campus-recruit/
├── config/          # 配置
├── controller/      # 控制器（处理 HTTP 请求）
├── service/         # 业务逻辑
├── dao/             # 数据访问层（数据库操作）
├── model/           # 数据模型
├── middleware/      # 中间件（JWT、CORS 等）
├── router/          # 路由
├── utils/           # 工具函数
└── main.go          # 入口
```

## 快速开始

### 前置条件

- Go 1.21+
- MySQL 8.0+
- Redis 7.0+

### 运行

```bash
# 安装依赖
go mod tidy

# 创建数据库
mysql -u root -p -e "CREATE DATABASE IF NOT EXISTS campus_recruit"

# 启动服务
go run main.go
```

### 测试接口

```bash
# 健康检查
curl http://localhost:8080/ping

# 注册
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456","role":"student"}'

# 登录
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'

# 获取用户信息（需要替换 YOUR_TOKEN）
curl http://localhost:8080/api/v1/user/info \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## 开发计划

参见 `Go后端学习求职计划.md`
