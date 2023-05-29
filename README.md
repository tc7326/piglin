# piglin-猪灵图床

猪灵图床 一款图床API 一键部署 可以嵌入你的网页、APP、等。

# 所用框架

## web框架：[Gin](https://gin-gonic.com/)

```go
go get -u github.com/gin-gonic/gin
```

## 数据库：[GORM](https://gorm.io/)

```go
go get -u gorm.io/gorm
```

## 缓存：[Go Redis](https://redis.uptrace.dev/zh/)

```go
go get github.com/redis/go-redis/v9
```

# MVC

## Controller 层

主要做鉴权、基本数据校验 和 构造业务层需要的数据 比如校验用户昵称长度等

## Service 层

具体的业务实现 比如 新增用户这样一个业务的实现逻辑

## DAO 层

使用MySQL持久化数据 最基本的增删改查 比如向数据库的用户表插入一条用户数据

# API介绍

使用Rest风格 和 基于OAuth2.0的权限校验

## 响应基类

```json
{
  "code": 0,
  "err": "err msg",
  "result": {
    "key1": "value1",
    "key2": "value2",
  }
}
```

## 用户API

### 注册

### 登录

### 更新token

