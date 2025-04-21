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
  "password": "QQQQQAQWERqwer2."
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
    "code": 10000,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDUzMzc5NDIsInVzZXJuYW1lIjoiMjAyMjIxMDg1NyJ9.mE604OfYgRU5bdqn_S99aAL7rW9oJiYmGHE032CzjCQ"
    },
    "msg": "成功"
}
```

### 错误响应
#### 信息门户校登录失败
```json
{
    "code": 40001,
    "data": "null",
    "msg": "信息门户校登录失败"
}
```

#### 信息门户内部问题，请重试
```json
{
    "code": 40002,
    "data": "null",
    "msg": "信息门户内部问题，请重试"
}
```



## 2. 获取用户信息

### 请求
`POST /api/v1/user/auth/get_userinfo`

**认证方式**: Bearer Token  
**请求头**:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDUzMzczMTYsInVzZXJuYW1lIjoiMjAyMjIxMDg1NyJ9.6p6iPuRqZr2tDI7Q9bbUS3Xa_Z0Tk47sIao0DY4VXQE
```

### 请求参数
无请求体

### 成功响应
`200 OK`
```json
{
    "code": 10000,
    "data": {
        "studentCode": "163348",
        "studentId": "2022210857",
        "usernameEn": "Li Tie",
        "usernameZh": "李铁",
        "sex": "男",
        "cultivateType": "主修",
        "department": "管理学院",
        "grade": "2022",
        "level": "本科",
        "studentType": "一般本科生",
        "major": "大数据管理与应用",
        "class": "大数据22-1班",
        "campus": "屯溪路校区",
        "status": "正常",
        "length": "4.0",
        "enrollmentDate": "2022-09-01",
        "graduateDate": "2026-07-01",
        "id": 1,
        "created_at": "2025-04-21T23:55:16.676+08:00",
        "updated_at": "2025-04-21T23:55:16.676+08:00",
        "username": "2022210857",
        "account_status": "active",
        "nickname": "李铁",
        "avatar_path": "default-avatar.png"
    },
    "msg": "成功"
}
```

### 响应字段说明

#### 基础信息
| 字段 | 类型 | 说明 |
|------|------|------|
| code | int | 响应码 (10000表示成功) |
| msg | string | 响应消息 |
| data | object | 用户数据对象 |

#### 用户数据(data对象)
| 字段 | 类型 | 说明 |
|------|------|------|
| studentCode | string | 学生编码 |
| studentId | string | 学号 |
| usernameEn | string | 英文姓名 |
| usernameZh | string | 中文姓名 |
| sex | string | 性别 |
| cultivateType | string | 培养类型 |
| department | string | 学院 |
| grade | string | 年级 |
| level | string | 学历层次 |
| studentType | string | 学生类型 |
| major | string | 专业 |
| class | string | 班级 |
| campus | string | 校区 |
| status | string | 学籍状态 |
| length | string | 学制 |
| enrollmentDate | string | 入学日期 |
| graduateDate | string | 预计毕业日期 |
|------|------|------|
| id | int | 用户ID |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |
| username | string | 用户名 |
| account_status | string | 账号状态 (active/locked/disabled) |
| nickname | string | 昵称 |
| avatar_path | string | 头像路径 |

### 错误响应

#### 未授权访问
```json
{
    "code": 30001,
    "msg": "未授权访问"
}
```

#### 无效Token
```json
{
    "code": 30002,
    "msg": "无效Token"
}
```

#### 用户不存在
```json
{
    "code": 20002,
    "msg": "用户不存在"
}
```
#### 信息门户未登录
```json
{
    "code": 40004,
    "data": "null",
    "msg": "信息门户未登录"
}
```


## 3. 更改用户信息

### 请求
`POST /api/v1/user/auth/update_userinfo`

**认证方式**: Bearer Token  
**请求头**:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDUzMzczMTYsInVzZXJuYW1lIjoiMjAyMjIxMDg1NyJ9.6p6iPuRqZr2tDI7Q9bbUS3Xa_Z0Tk47sIao0DY4VXQE
```

### 请求参数
```json
{
  "id": "1",
  "nickname": "李铁2333"
}
```

### 参数说明
| 字段 | 类型 | 必填 | 验证规则 |
|------|------|------|----------|
| id | int | 是 | 用户ID |
| username | string |否 | 用户名 |
| account_status | string |否 | 账号状态 (active/locked/disabled) |
| nickname | string |否 | 昵称 |
| avatar_path | string |否 | 头像路径 |

### 成功响应
`200 OK`
```json
{
    "code": 10000,
    "data": {
        "studentCode": "163348",
        "studentId": "2022210857",
        "usernameEn": "Li Tie",
        "usernameZh": "李铁",
        "sex": "男",
        "cultivateType": "主修",
        "department": "管理学院",
        "grade": "2022",
        "level": "本科",
        "studentType": "一般本科生",
        "major": "大数据管理与应用",
        "class": "大数据22-1班",
        "campus": "屯溪路校区",
        "status": "正常",
        "length": "4.0",
        "enrollmentDate": "2022-09-01",
        "graduateDate": "2026-07-01",
        "id": 1,
        "created_at": "2025-04-21T23:55:16.676+08:00",
        "updated_at": "2025-04-21T23:55:16.676+08:00",
        "username": "2022210857",
        "account_status": "active",
        "nickname": "李铁2333",
        "avatar_path": "default-avatar.png"
    },
    "msg": "成功"
}
```
#### 未授权访问
```json
{
    "code": 30001,
    "msg": "未授权访问"
}
```

#### 无效Token
```json
{
    "code": 30002,
    "msg": "无效Token"
}
```

#### 用户不存在
```json
{
    "code": 20002,
    "msg": "用户不存在"
}
```
#### 信息门户未登录
```json
{
    "code": 40004,
    "data": "null",
    "msg": "信息门户未登录"
}
```
#### 请求格式错误
```json
{
    "code": 10002,
    "msg": "请求格式错误"
}
```

### 响应字段说明

#### 基础信息
| 字段 | 类型 | 说明 |
|------|------|------|
| code | int | 响应码 (10000表示成功) |
| msg | string | 响应消息 |
| data | object | 用户数据对象 |

#### 用户数据(data对象)
| 字段 | 类型 | 说明 |
|------|------|------|
| studentCode | string | 学生编码 |
| studentId | string | 学号 |
| usernameEn | string | 英文姓名 |
| usernameZh | string | 中文姓名 |
| sex | string | 性别 |
| cultivateType | string | 培养类型 |
| department | string | 学院 |
| grade | string | 年级 |
| level | string | 学历层次 |
| studentType | string | 学生类型 |
| major | string | 专业 |
| class | string | 班级 |
| campus | string | 校区 |
| status | string | 学籍状态 |
| length | string | 学制 |
| enrollmentDate | string | 入学日期 |
| graduateDate | string | 预计毕业日期 |
|------|------|------|
| id | int | 用户ID |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |
| username | string | 用户名 |
| account_status | string | 账号状态 (active/locked/disabled) |
| nickname | string | 昵称 |
| avatar_path | string | 头像路径 |

### 错误响应

#### 未授权访问
```json
{
    "code": 30001,
    "msg": "未授权访问"
}
```

#### 无效Token
```json
{
    "code": 30002,
    "msg": "无效Token"
}
```

#### 用户不存在
```json
{
    "code": 20002,
    "msg": "用户不存在"
}
```
#### 信息门户未登录
```json
{
    "code": 40004,
    "data": "null",
    "msg": "信息门户未登录"
}
```


# 错误代码


## 全局错误代码 (10xxx)

| 错误代码 | 状态码 | 说明 |
|---------|--------|------|
| 10000 | 200 OK | 操作成功 |
| 10001 | 500 Internal Server Error | 服务器内部错误 |
| 10002 | 400 Bad Request | 请求格式错误 |
| 10003 | 500 Internal Server Error | 未知错误 |
| 10004 | 502 Bad Gateway | 错误网关 |

## 用户相关错误 (20xxx)

| 错误代码 | 状态码 | 说明 |
|---------|--------|------|
| 20001 | 401 Unauthorized | 密码错误 |
| 20002 | 404 Not Found | 用户不存在 |
| 20003 | 409 Conflict | 用户已存在 |
| 20004 | 500 Internal Server Error | 用户信息更新失败 |

## 鉴权错误 (30xxx)

| 错误代码 | 状态码 | 说明 |
|---------|--------|------|
| 30001 | 401 Unauthorized | 未授权访问 |
| 30002 | 401 Unauthorized | 无效Token |

## 信息门户(HFUT)相关错误 (40xxx)

| 错误代码 | 状态码 | 说明 |
|---------|--------|------|
| 40001 | 401 Unauthorized | 信息门户登录失败 |
| 40002 | 503 Service Unavailable | 信息门户内部问题，请重试 |
| 40003 | 500 Internal Server Error | 信息门户未知错误 |
| 40004 | 401 Unauthorized | 信息门户未登录 |



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
