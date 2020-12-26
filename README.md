### 用Golang实现BBS

> 此为后端项目，前端Vue项目见 [Golang-Vue-BBS](https://github.com/Cyberist-Edgar/Golang-Vue-BBS)


首先需要在config目录下进行相应的配置，其中`database`，`baidu key`，以及`邮箱`必须进行配置
```yaml
# 数据库配置
database: 
  mysql:
    database: "bbs" # 数据库名称
    username: "Edgar" # 用户名
    password: "Edgar"  # 密码

# 日志配置
log:
  path: "log" # 日志目录
  filename: "bbs.log"  # 日志名称

# 文件目录配置
filepath:
  base: "file"  # 文件夹
  detail:
    avatar: "avatar"  # 头像目录
    carousel: "carousel"  # 轮播图目录
    download: "download"  # 上传之后的文件存放目录

# 百度api key
baidu:
  key: "xxx" # 百度访问key

# 邮件配置，发生错误时自动发送邮件
email:
  server: "xxx"  # 需要端口号，比如smtp.126.com:25
  username: "xxx"  # 邮箱代理用户名
  password: "xxx"  # 邮箱密码
  target: # 需要发送到的邮箱地址
    - "xxx@qq.com"
```
