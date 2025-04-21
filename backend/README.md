以下是三个接口的Markdown格式文档，包含完整的请求/响应示例和参数说明：

---

# 用户服务API文档


---

## 1. 用户登录

### 请求
`POST /api/v1/user/login`

```json
{
  "username": "2022210857",
  "password": "P@ssw0rd123"
}
```

### 参数说明
| 字段 | 类型 | 必填 | 验证规则 |
|------|------|------|----------|
| username | string | 是 | alphanum, min=4, max=20 |
| password | string | 是 | min=6 |

### 成功响应
`200 OK`
```json
{
    "code": 20001,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDUzMTUzNjMsInVzZXJuYW1lIjoiMjAyMjIxMDg1NyJ9.mQIEG1jl17mefFHF5KVA3E-0EbbX6_Ymv6GGLxWmZZs"
    },
    "msg": "登录成功"
}
```


### 错误代码
| 代码 | 说明 |
|------|------|
| 40101 | 用户名不存在 |
| 40102 | 密码错误 |

---


### 错误代码
根据提供的错误代码，以下是完整的Markdown接口文档补充错误代码部分：

# 错误代码规范

## 全局错误代码 (10xxx)

| 错误代码 | 状态码 | 说明 |
|---------|--------|------|
| 10001 | 200 OK | 操作成功 |
| 10002 | 500 Internal Server Error | 服务器内部错误 |
| 10003 | 400 Bad Request | 请求格式错误 |
| 10004 | 500 Internal Server Error | 未知错误 |
| 10005 | 502 Bad Gateway | 错误网关 |

## 用户相关错误 (20xxx)

| 错误代码 | 状态码 | 说明 |
|---------|--------|------|
| 20001 | 200 OK | 登录成功 |
| 20002 | 401 Unauthorized | 信息门户校验失败 |
| 20003 | 200 OK | 信息门户校验成功，请完成注册 |
| 20004 | 401 Unauthorized | 密码错误 |
| 20005 | 503 Service Unavailable | 信息门户服务不可用 |
| 20006 | 500 Internal Server Error | 信息门户未知错误 |
| 20007 | 201 Created | 注册成功 |
| 20008 | 404 Not Found | 用户不存在 |
| 20009 | 409 Conflict | 用户已存在 |
| 20010 | 500 Internal Server Error | 用户信息更新失败 |

## 鉴权错误 (30xxx)

| 错误代码 | 状态码 | 说明 |
|---------|--------|------|
| 30001 | 401 Unauthorized | 未授权访问 |
| 30002 | 401 Unauthorized | 无效Token |

## 使用示例

### 成功响应
```json
{
  "code": 10001,
  "message": "成功",
  "data": {
    "user_id": 123
  }
}
```

### 错误响应
```json
{
  "code": 20004,
  "message": "密码错误",
  "data": null
}
```
