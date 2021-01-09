### Go BBS

> This is a backend project, you can see the frontend project in [Golang-Vue-BBS](https://github.com/Cyberist-Edgar/Golang-Vue-BBS)

To run this project, you should go to the config directory to configure first, the config file is as follows:


```yaml
# database configuration
database: 
  mysql:
    database: "bbs" # the database name you use
    username: "Edgar" # username 
    password: "Edgar"  # password

# log configuration
log:
  path: "log" # log file directory
  filename: "bbs.log"  # log filename

# file directory to store file like avatar, carousel and upload file
filepath:
  base: "file"  # the base file directory
  detail:
    avatar: "avatar"  # avatar directory like file/avatar
    carousel: "carousel"  # carousel directory
    download: "download"  # upload file directory

# baidu api key configuration
baidu:
  key: "xxx" # baidu api key 

# email configuration, when there is error, send a email to report
email:
  server: "xxx"  # email server address, should be form of "host:port", such as smtp.126.com:25
  username: "xxx"  # email username
  password: "xxx"  # email password, usually not you login password, but the agent password like pop3 service
  target: # email address send report to, you can have as many as you like
    - "xxx@qq.com"
```
You must configure your database, baidu key and email, others have their default value
> To get a baidu api key, you should visit [HERE](http://lbsyun.baidu.com/apiconsole/key#/home), but I am sorry that it seems this website has no English version, then just use the Google Translate