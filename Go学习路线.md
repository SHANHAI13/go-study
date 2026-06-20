# Go 从零到可面试 — 完整学习路线

> 开始日期：2025年6月20日
> 目标：2025年9月秋招拿到 Go 后端实习 offer
> 学历：二本，大三
> 可用时间：暑假 7-8 月，每天 6-8 小时

---

## 第一周：基础语法

### Day 1（今天 — 环境 + 变量和类型）

- [ ] 安装 Go：https://go.dev/dl/
- [ ] 验证安装：`go version`
- [ ] 新建文件夹 `gostudy/`，写第一个 `main.go`
- [ ] 变量声明、基本类型（int, string, bool, float64）
- [ ] `fmt.Println`、`fmt.Printf`
- [ ] 写 3 个小程序：
  - 打印乘法口诀表
  - 输入两个数字，输出和差积商
  - 判断一个数是奇数还是偶数
- [ ] 注册 GitHub，创建仓库，代码推上去

### Day 2

- [ ] if-else、switch-case
- [ ] for 循环（Go 只有 for，没有 while）
- [ ] 数组（array）和切片（slice）——搞清楚区别
- [ ] 写 3 个小程序：
  - 冒泡排序
  - 判断一个字符串是不是回文
  - 生成前 20 个斐波那契数列

### Day 3

- [ ] map 的增删改查
- [ ] range 遍历
- [ ] 函数定义、多返回值
- [ ] 值传递 vs 引用传递
- [ ] 写 2 个小程序：
  - 统计一篇文章里每个单词出现的次数
  - 写一个函数，接收一个整数切片，返回最大值和最小值

### Day 4

- [ ] struct 定义和使用
- [ ] 方法（func (接收者) 方法名）
- [ ] 接口（interface）——Go 最核心的抽象
- [ ] 写 1 个程序：学生管理系统
  - struct Student {name, score}
  - 实现新增、删除、按分数排序、求平均分
  - 用接口定义 StudentStore，用 map 实现

### Day 5

- [ ] 指针 `*` 和 `&`
- [ ] new 和 make 的区别
- [ ] error 类型、fmt.Errorf
- [ ] panic 和 recover（了解即可，别滥用）
- [ ] defer 的用法和执行顺序
- [ ] 改造学生管理系统：加错误处理（找不到学生返回 error）

### Day 6 — 并发

- [ ] go 关键字启动 goroutine
- [ ] channel 的创建、发送、接收
- [ ] 无缓冲 channel vs 有缓冲 channel
- [ ] select 多路复用
- [ ] sync.WaitGroup
- [ ] 写 2 个程序：
  - 3 个 goroutine 并发打印数字（用 channel 同步）
  - 生产者-消费者模型（1 个生产者，3 个消费者，用 channel 通信）

### Day 7 — 复习 + 刷题

- [ ] 重写这周写过的所有代码（不看答案）
- [ ] 注册 LeetCode 账号
- [ ] 刷 5 道简单题（用 Go 写）：
  - Two Sum（两数之和）
  - Reverse String（反转字符串）
  - Valid Parentheses（有效括号）
  - Linked List Cycle（环形链表）
  - Binary Tree Inorder Traversal（中序遍历）
- [ ] 题解和代码推到 GitHub

---

## 第一周检查清单

到第七天晚上，必须能**不看任何资料**写出：

- [ ] 启动一个 HTTP 服务器（net/http 包，10 行代码）
- [ ] 开 3 个 goroutine，用 channel 收发数据
- [ ] 定义 interface 并实现它
- [ ] 用 slice 和 map 完成增删改查

做不到就加练。做到了进入第二周。

---

## 完整八周计划

| 周 | 主题 | 关键产出 |
|---|------|----------|
| 第 1 周 | 基础语法 | 能写完整的小程序 |
| 第 2 周 | 标准库：net/http、encoding/json、io、os、time | 写一个 REST API 服务 |
| 第 3-4 周 | 数据库：MySQL + Redis + GORM | API 服务接数据库 + 缓存 |
| 第 5-6 周 | 项目一：手写迷你 HTTP 框架 | 路由、中间件、上下文 |
| 第 7 周 | 项目二：分布式缓存 | LRU、一致性哈希、HTTP 通信 |
| 第 8 周 | 刷题冲刺 + 简历打磨 | 90 道题、两个项目上简历 |

---

## 两个核心项目

### 项目一：迷你 HTTP 框架（第 5-6 周）

```
□ 路由注册（GET/POST/PUT/DELETE）
□ 动态路由参数 /user/:id
□ 分组路由 /api/v1/*
□ 中间件（日志、恢复、计时）
□ 上下文 Context 封装
□ JSON 绑定和响应

→ 面试最爱问的项目，体现对 net/http 和接口的理解
```

### 项目二：分布式缓存（第 7 周）

```
□ LRU 缓存淘汰
□ 并发安全（sync.Mutex）
□ HTTP 节点间通信
□ 一致性哈希选择节点
□ 单机 → 分布式渐进

→ Go 面试经典项目，字节几乎必问
```

---

## 每天的铁律

1. 每天至少写 200 行代码（不算复制粘贴）
2. 所有代码推到 GitHub，提交记录要绿
3. 遇到 bug 先自己调 30 分钟，调不出来再搜
4. 晚上写下今天学了什么（一句话就行）
5. 周末不学新东西，复习本周的

---

## 投递策略（8月中旬开始）

```
重点投：二线互联网 + 传统行业 IT + 创业公司
别投：字节/阿里/腾讯（9月机筛大概率挂，浪费时间；寒假可以试）
海投量：100+ 份简历

范围：Java 岗 80%，Go 岗 20%（有 Go 项目可以投 Java）
```

---

## 学习资源

- Go 官方文档：https://go.dev/doc/
- Go by Example：https://gobyexample.com/
- Go 语言圣经（中文）：https://gopl-zh.github.io/
- LeetCode：https://leetcode.cn/

---

> 从 `fmt.Println("hello, 世界")` 开始。
> 今天是 2025年6月20日，距离 9 月还有两个半月。
> 够用，但每一天都得用上。
