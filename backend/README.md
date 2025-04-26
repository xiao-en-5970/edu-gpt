[TOC]

---

# 公共前缀

`/api/v1`



# 鉴权方式

## 1.认证方式: 

Bearer Token  

## 2.请求头:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDUzMzczMTYsInVzZXJuYW1lIjoiMjAyMjIxMDg1NyJ9.6p6iPuRqZr2tDI7Q9bbUS3Xa_Z0Tk47sIao0DY4VXQE
```

## 3.Url示例:

 `/api/v1/user/auth/`

所有带auth的url都需要token鉴权



# 用户服务API文档

## 0. user组前缀

`/api/v1/user`



## 1. 用户登录

### 请求
`POST /api/v1/user/login`

### 请求参数

```json
{
  "username": "2022210857",
  "password": "QQQQQAQWERqwer2."
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|----------|
| username | string | 是 | 学号 |
| password | string | 是 | 密码 |

### 成功响应
`200 OK`

```json
{
    "code": 10000,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDUzMzc5NDIsInVzZXJuYW1lIjoiMjAyMjIxMDg1NyJ9.mE604OfYgRU5bdqn_S99aAL7rW9oJiYmGHE032CzjCQ",
        "id": 1
    },
    "msg": "成功"
}
```

| 类型   | 字段  | 说明          |
| ------ | ----- | ------------- |
| string | token | token用于鉴权 |
| int    | id    | 用户号        |



## 2. 获取用户信息【auth】

### 请求
`POST /api/v1/user/auth/get_userinfo`

### 请求参数
```json
无
```

### 成功响应
`200 OK`
```json
{
    "code": 10000,
    "data": {
        "id": 1,
        "username_zh": "李铁",
        "sex": "男",
        "cultivate_type": "主修",
        "department": "管理学院",
        "grade": "2022",
        "level": "本科",
        "major": "大数据管理与应用",
        "class": "大数据22-1班",
        "campus": "屯溪路校区",
        "enrollment_date": "2022-09-01",
        "graduate_date": "2026-07-01",
        "created_at": "2025-04-25T07:07:01.907Z",
        "username": "2022210857",
        "account_status": "active",
        "nickname": "李铁668879",
        "avatar_path": "127.0.0.1:8080/api/v1/user/auth/imageurl/1",
        "signature": "这人啥也没说",
        "tags": []
    },
    "msg": "成功"
}
```

### 响应字段说明

| 字段 | 类型 | 说明 |
|------|------|------|
| id | int | 用户ID |
| username_zh | string | 中文姓名 |
| sex             | string   | 性别                              |
| cultivate_type  | string   | 培养类型                          |
| department      | string   | 学院                              |
| grade           | string   | 年级                              |
| level           | string   | 学历层次                          |
| student_type    | string   | 学生类型                          |
| major           | string   | 专业                              |
| class           | string   | 班级                              |
| campus          | string   | 校区                              |
| status          | string   | 学籍状态                          |
| length          | string   | 学制                              |
| enrollment_date | string   | 入学日期                          |
| graduate_date   | string   | 预计毕业日期                      |
| ------          | ------   | ------                            |
| created_at      | string   | 创建时间                          |
| username        | string   | 用户名                            |
| account_status  | string   | 账号状态 (active/locked/disabled) |
| nickname        | string   | 昵称                              |
| avatar_path     | string   | 头像路径                          |
| signature       | string   | 个性签名                          |
| tags            | []string | tag                               |
|                 |          |                                   |
|                 |          |                                   |



## 3. 更改用户信息【auth】

### 请求
`POST /api/v1/user/auth/update_userinfo`

### 请求参数
```json
{
  "id": "1",
  "nickname": "李铁2333"
}
```

### 参数说明
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|----------|
| avatar_path    | string | 否   | 头像路径                          |
| username | string |否 | 用户名 |
| account_status | string |否 | 账号状态 (active/locked/disabled) |
| nickname | string |否 | 昵称 |

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
### 响应字段说明
| 字段 | 类型 | 说明 |
|------|------|------|
| student_code    | string | 学生编码                          |
| studentId       | string | 学号                              |
| username_en     | string | 英文姓名                          |
| username_zh     | string | 中文姓名                          |
| sex             | string | 性别                              |
| cultivate_type  | string | 培养类型                          |
| department      | string | 学院                              |
| grade           | string | 年级                              |
| level           | string | 学历层次                          |
| student_type    | string | 学生类型                          |
| major           | string | 专业                              |
| class           | string | 班级                              |
| campus          | string | 校区                              |
| status          | string | 学籍状态                          |
| length          | string | 学制                              |
| enrollment_date | string | 入学日期                          |
| graduate_date   | string | 预计毕业日期                      |
| ------          | ------ | ------                            |
| id              | int    | 用户ID                            |
| created_at      | string | 创建时间                          |
| updated_at      | string | 更新时间                          |
| username        | string | 用户名                            |
| account_status  | string | 账号状态 (active/locked/disabled) |
| nickname        | string | 昵称                              |



## 4. 用户信息【auth】

### 请求

`POST /api/v1/user/auth/:id`

### 请求参数

```json
无
```

### 成功响应

`200 OK`

```json
{
    "code": 10000,
    "data": {
        "id": 1,
        "create_at": "2025-04-26T14:45:47.751Z",
        "department": "数学学院",
        "nickname": "小奀66666666",
        "avatar_path": "remote.xiaoen.xyz/api/v1/user/auth/imageurl/1",
        "sex": "男",
        "grade": "2022",
        "campus": "翡翠湖校区",
        "signature": "这人啥也没说",
        "tags": "[]"
    },
    "msg": "成功"
}
```

### 响应字段说明

| 字段        | 类型     | 说明         |
| ----------- | -------- | ------------ |
| id          | int      | 用户号       |
| create_at   | string   | 创建用户时间 |
| department  | string   | 学院         |
| nickname    | string   | 昵称         |
| avatar_path | string   | 图片url      |
| sex         | string   | 性别         |
| grade       | string   | 年级         |
| campus      | string   | 校区         |
| signature   | string   | 签名         |
| tags        | []string | 标签         |



## 5. 上传用户头像【auth】

### 请求

`POST /api/v1/user/auth/upload_image`

### 请求参数

```
form-data

key:avatar

value:图片文件
```

### 成功响应

`200 OK`

```json
{
    "code": 10000,
    "data": {
        "url": "127.0.0.1:8080/api/v1/user/imageurl/1"
    },
    "msg": "成功"
}
```

### 响应字段说明

| 字段 | 类型   | 说明    |
| ---- | ------ | ------- |
| url  | string | 图片url |



## 6. 获取用户头像【auth】

### 请求

`POST /api/v1/user/auth/imageurl/:id`

### 请求参数

```
无
```

### 成功响应

`200 OK`

```
图片内容
```

# 帖子服务API文档

## 0. post组前缀

`/api/v1/post`



## 1. 上传帖子【auth】

### 请求

`POST /api/v1/post/auth/create`

### 请求参数

```json
{
    "title":"震惊！他居然这样做！",
    "content":"其实啥也没干"
}
```

### 成功响应

`200 OK`

```json
{
    "code": 10000,
    "data": {
        "id": 4
    },
    "msg": "成功"
}
```

### 响应字段说明

| 字段 | 类型 | 说明   |
| ---- | ---- | ------ |
| id   | int  | 帖子id |



## 2. 编辑帖子【auth】

### 请求

`POST /api/v1/post/auth/edit`

### 请求参数

```json
{
    "id": 1,
    "title": "震惊",
    "content": "没什么好震惊的"
}
```

### 成功响应

`200 OK`

```json
{
    "code": 10000,
    "data": {
        "id": 1,
        "poster_id": 1,
        "title": "震惊",
        "content": "没什么好震惊的",
        "view_count": 0,
        "like_count": 0,
        "collect_count": 0,
        "comment_count": 0,
        "create_at": "2025-04-25T14:59:36Z",
        "update_at": "2025-04-26T12:51:16Z",
        "poster_nickname": "李铁668879",
        "poster_grade": "2022",
        "poster_campus": "屯溪路校区",
        "poster_department": "管理学院"
    },
    "msg": "成功"
}
```

### 响应字段说明

| 字段              | 类型   | 说明       |
| ----------------- | ------ | ---------- |
| id                | int    | 帖子id     |
| poster_id         | int    | 发帖人id   |
| title             | string | 帖子标题   |
| content           | string | 帖子内容   |
| view_count        | int    | 浏览量     |
| like_count        | int    | 点赞量     |
| collect_count     | int    | 收藏量     |
| comment_count     | int    | 评论量     |
| create_at         | string | 创建时间   |
| update_at         | string | 修改时间   |
| poster_nickname   | string | 发帖人昵称 |
| poster_grade      | string | 发帖人年级 |
| poster_campus     | string | 发帖人校区 |
| poster_department | string | 发帖人学院 |

## 5. 获取帖子内容【auth】

### 请求

`POST /api/v1/post/auth/:id`

### 请求参数

```
无
```

### 成功响应

`200 OK`

```json
{
    "code": 10000,
    "data": {
        "id": 1,
        "poster_id": 1,
        "title": "震惊",
        "content": "没什么好震惊的",
        "view_count": 0,
        "like_count": 0,
        "collect_count": 0,
        "comment_count": 0,
        "create_at": "2025-04-25T14:59:36Z",
        "update_at": "2025-04-26T12:51:16Z",
        "poster_nickname": "李华668879",
        "poster_grade": "2021",
        "poster_campus": "翡翠湖校区",
        "poster_department": "计算机学院"
    },
    "msg": "成功"
}
```

### 响应字段说明

| 字段              | 类型   | 说明       |
| ----------------- | ------ | ---------- |
| id                | int    | 帖子id     |
| poster_id         | int    | 发帖人id   |
| title             | string | 帖子标题   |
| content           | string | 帖子内容   |
| view_count        | int    | 浏览量     |
| like_count        | int    | 点赞量     |
| collect_count     | int    | 收藏量     |
| comment_count     | int    | 评论量     |
| create_at         | string | 创建时间   |
| update_at         | string | 修改时间   |
| poster_nickname   | string | 发帖人昵称 |
| poster_grade      | string | 发帖人年级 |
| poster_campus     | string | 发帖人校区 |
| poster_department | string | 发帖人学院 |

## 5. 待定

### 请求

`POST /api/v1/user/auth/upload_image`

### 请求参数

```
form-data

key:avatar

value:图片文件
```

### 成功响应

`200 OK`

```json
{
    "code": 10000,
    "data": {
        "url": "127.0.0.1:8080/api/v1/user/imageurl/1"
    },
    "msg": "成功"
}
```

### 响应字段说明

| 字段 | 类型   | 说明    |
| ---- | ------ | ------- |
| url  | string | 图片url |

## 5. 待定

### 请求

`POST /api/v1/user/auth/upload_image`

### 请求参数

```
form-data

key:avatar

value:图片文件
```

### 成功响应

`200 OK`

```json
{
    "code": 10000,
    "data": {
        "url": "127.0.0.1:8080/api/v1/user/imageurl/1"
    },
    "msg": "成功"
}
```

### 响应字段说明

| 字段 | 类型   | 说明    |
| ---- | ------ | ------- |
| url  | string | 图片url |

## 5. 待定

### 请求

`POST /api/v1/user/auth/upload_image`

### 请求参数

```
form-data

key:avatar

value:图片文件
```

### 成功响应

`200 OK`

```json
{
    "code": 10000,
    "data": {
        "url": "127.0.0.1:8080/api/v1/user/imageurl/1"
    },
    "msg": "成功"
}
```

### 响应字段说明

| 字段 | 类型   | 说明    |
| ---- | ------ | ------- |
| url  | string | 图片url |

## 

# 错误代码


## 全局错误代码 (10xxx)

| 错误代码 | 说明 |
|---------|------|
| 10000 | 操作成功 |
| 10001 | 服务器内部错误 |
| 10002 | 请求格式错误 |
| 10003 | 未知错误 |
| 10004 | 错误网关 |

## 用户相关错误 (20xxx)

| 错误代码 | 说明 |
|---------|------|
| 20001 | 密码错误 |
| 20002 | 用户不存在 |
| 20003 | 用户已存在 |
| 20004 | 用户信息更新失败 |

## 鉴权错误 (30xxx)

| 错误代码 | 说明 |
|---------|------|
| 30001 | 未授权访问 |
| 30002 | 无效Token |

## 信息门户(HFUT)相关错误 (40xxx)

| 错误代码 | 说明 |
|---------|------|
| 40001 | 信息门户登录失败 |
| 40002 | 信息门户内部问题，请重试 |
| 40003 | 信息门户未知错误 |
| 40004 | 信息门户未登录 |

## 图片相关错误 (50xxx)

| 错误代码 | 说明         |
| :------- | ------------ |
| 50001    | 图片格式错误 |

## 帖子相关错误 (60xxx)

| 错误代码 | 说明       |
| -------- | ---------- |
| 60001    | 帖子不存在 |

## 

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

------

