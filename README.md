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

# API介绍

使用Rest风格 和 基于Auth2.0的权限校验

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

## 
