# ECHO 学习Demo

## 简介

此项目主要用于学习使用传说中比其他Go web框架性能高10倍的[ECHO](https://github.com/labstack/echo)

> 项目使用 glide 构建

## 项目中使用到包

日志： github.com/donnie4w/go-logger/logger

单元测试：github.com/smartystreets/goconvey

数据库查询：github.com/jmoiron/sqlx

mysql驱动：github.com/go-sql-driver/mysql

配置文件解析：github.com/BurntSushi/toml

## 使用说明

- 执行 doc目录下的sql文件，创建数据库

- 使用glide 构建项目，如果有包无法下载请自行想办法解决

- 启动项目访问 http://localhost:8080/user.html 

  - 单元测试

    - 在service目录执行 `go test` 命令
    - 使用`goconvey`命令启动goconvey测试工具，查看测试结果。

## 开发体验

​	之前开发项目有用过beego，Echo相比beego的开发体验差距还是蛮大的。beego已经是一个非常完善的web框架了。Echo的好处就是其他组件可以自由选择比如日志、数据库操作等

​	Golang 的 web框架性能对比

​	 http://blog.csdn.net/qq849635649/article/details/56017559

​	http://colobu.com/2017/04/07/go-webframework-benchmark-2017-Spring/

​	http://colobu.com/2016/04/06/the-fastest-golang-web-framework/

选择 github.com/jmoiron/sqlx	作为数据库操作包的原因是个人对ORM没什么好感，ORM开发的时候时挺爽，但是如果接手维护一个老项目是用ORM写的就不爽了。想到一个长颈鹿和兔子的笑话

```
长颈鹿说：夏天喝水，凉水流过我长长的脖子，很凉快
兔问：你吐过么？
```


**有疑问可以给我留言，或者发邮件给我（很长时间才会看一次邮件）。代码里注释很完善提问前请先看一下**

**如果有大神发现有错误的地方，烦请赐教**

