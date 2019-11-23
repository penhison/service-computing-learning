# REST API for blog v0.1

## 前言
本次设计的api用于开发环境，获取的不是网页，而是json格式的数据。
本次演示的博客api网址假设为https://api.blog.penhison.com/

## api总览
|描述|方法|地址|
|-|-|-|
|获取用户信息|GET|/[username]|
|获取文章信息|GET|/article/[id]|
|用户认证|GET|/|
|修改文章|POST|/article/[id]|
|删除文章|DELETE|/article/[id]|
|创建文章|POST|/article/new|
|创建评论|POST|/article/[id]/comment|

## 架构
所有API访问都是通过HTTPS进行的。主地址https://api.blog.penhison.com/。所有数据都以JSON的形式发送和接收。

```
curl -i https://api.github.com/users/octocat/orgs
HTTP/1.1 200 OK
Server: nginx
Date: Fri, 12 Oct 2012 23:33:14 GMT
Content-Type: application/json; charset=utf-8
Connection: keep-alive
Status: 200 OK
ETag: "a00049ba79152d03380c34652f2cb612"
X-Blog-Media-Type: v1
X-RateLimit-Limit: 5000
X-RateLimit-Remaining: 4987
X-RateLimit-Reset: 1350085394
Content-Length: 5
Cache-Control: max-age=0, private, must-revalidate
X-Content-Type-Options: nosniff
```
包含空白字段null而不是将其省略。

所有时间戳以ISO 8601格式返回：
```
YYYY-MM-DDTHH:MM:SSZ
```

## 认证
```
curl -u“用户名” https://api.blog.penhison.com
```
使用无效的凭据进行身份验证将返回`401 Unauthorized`：
```
curl -i https://api.blog.penhison.com -u foo:bar
HTTP/1.1 401 Unauthorized
{
  "message": "Bad credentials",
  "documentation_url": "api.blog.penhison.com/v1"
}
```


在短时间内检测到多个具有无效凭据的请求后，API会临时拒绝该用户的所有身份验证尝试（包括具有有效凭据的请求）`403 Forbidden`：
```
curl -i https://api.blog.penhison.com -u valid_username:valid_password
HTTP/1.1 403 Forbidden
{
  "message": "Maximum number of login attempts exceeded. Please try again later.",
  "documentation_url": "https://api.blog.penhison.com/v1"
}
```

## 获取用户信息
已知用户名，获取详细信息
```
GET https://api.blog.penhison.com/penhison
```
当用户不存在或限制权限时回复`404 NOT FOUND`
否则回复用户的基本信息，包括用户名，发布的所有自己可访问的文章
```
{
  "username":"penhison"
  "article_count": 10,
  "visits": 12345,
  "articles":[
    {
      "title":"title",
      "href":"https://api.blog.penhison.com/article/1234",
      "abstract":"abstract",
      "words": 1629,
      "visits": 508
    },
	...
   ]
}

```

## 获取文章信息
已知文章id，获取详细信息
```
GET https://api.blog.penhison.com/article/233
```
当文章不存在或限制权限时回复`404 NOT FOUND`
否则回复文章的基本信息，包括发布者用户名，单词数，发布时间，文章详细信息，评论信息
```
{
  "title":"title"
  "words": 5678,
  "visits": 12345,
  "date": "2019-11-23"
  "detail": "detail"
  "comments": [
	{
		"username": "username",
		"date": "xxxx-xx-xx",
		"detail": "detail"
	},...
   ]
}

```

## 修改文章
已知文章id，修改文章内容
```
PUT https://api.blog.penhison.com/article/233 -d '{"title": "title", "detail": "detail"}
```
当文章不存在时回复`404 NOT FOUND`,
当用户未登录时或用户没有修改权限时回复`403 FORBIDDEN`,
否则修改文章信息，并返回修改后的文章内容
```
{
  “status”: "ok"
  "title":"title"
  "words": 5678,
  "visits": 12345,
  "date": "2019-11-23"
  "detail": "detail"
  "comments": [
	{
		"username": "username",
		"date": "xxxx-xx-xx",
		"detail": "detail"
	},...
   ]
}

```

## 删除文章
已知文章id，删除文章
```
DElETE https://api.blog.penhison.com/article/233
```
当文章不存在时回复`404 NOT FOUND`,
当用户未登录时或用户没有修改权限时回复`403 FORBIDDEN`,
否则删除文章,返回删除是否成功
```
{
  “status”: "ok"
}

```

## 创建文章
发布文章
```
POST https://api.blog.penhison.com/article/new -d '{"title": "title", "detail": "detail"}
```
当用户未登录时或用户没有修改权限时回复`403 FORBIDDEN`,
否则新建文章，并返回修改后的文章信息
```
{
  “status”: "ok"
  “id”: 12366
  "title":"title"
  "words": 5678,
  "visits": “0”,
  "date": "2019-11-23"
  "detail": "detail"
  "comments": [
	{
		"username": "username",
		"date": "xxxx-xx-xx",
		"detail": "detail"
	},...
   ]
}

```

## 创建评论
已知文章id，创建评论
```
POST https://api.blog.penhison.com/article/233/comment -d '{"username": "username","detail": "detail"}
```
当文章不存在时回复`404 NOT FOUND`,
当用户未登录时或用户没有修改权限时回复`403 FORBIDDEN`,
否则创建评论，并返回评论信息
```
{
  “status”: "ok"
  "username": "username",
  "date": "xxxx-xx-xx",
  "detail": "detail"
}

```