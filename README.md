# wxcloudrun-golang
微信云托管 Go语言HTTP服务端示例

简介：了解在微信云托管上如何用GO语言创建简单的http服务。通过示例创建一张user表，并对其进行增删改查的操作，对应POST/DELETE/PUT/GET四种请求的实现。

版本：
Golang 1.17.1（修改版本，需要同步修改[Dockerfile](https://github.com/WeixinCloud/wxcloudrun-golang/blob/main/Dockerfile)中的基础镜像）

详细介绍：
1. 一键部署时将默认开通微信云托管中的MySQL，并自动将数据库基本信息传入了环境变量中，可直接使用。（数据库信息获取及配置详情见:[init.go](https://github.com/WeixinCloud/wxcloudrun-golang/blob/main/db/init.go)。）
2. [container.config.json](https://github.com/WeixinCloud/wxcloudrun-golang/blob/main/container.config.json)仅用于在微信云托管中创建流水线时配套使用。
   * 如果不使用流水线，而是用本项目的代码在微信云托管控制台手动「新建版本」，则container.config.json配置文件不生效。最终版本部署效果以「新建版本」窗口中手动填写的值为准。
   * 'dataBaseName'和‘executeSQLs’ 两个字段只有在服务第一次部署时生效，后续流水线触发的版本更新不会执行（避免重复初始化数据库）。
   
   
示例API列表：

1 根据ID查询用户

* URL路径：
  ```/user/:id```
  
* 请求示例：
```
curl -X GET  http://{ip}:{port}/user/8
```

* 响应示例：
```
{
	"code": 0,
	"errorMsg": "",
	"data": {
		"id": 17,
		"name": "1231231232131",
		"age": 10,
		"email": "m1779387qqwewqeqwe3123@163.com",
		"phone": "1779aqweqwea3873123@163.com",
		"description": "111",
		"create_time": "2021-11-05T10:08:55+08:00",
		"update_time": "2021-11-05T10:08:55+08:00"
	}
}
```


2 新增用户

* URL路径：
  ```/user```
  
* 请求示例：
```
curl http://{ip}:{port}/user \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{  
      "name":"1231231232131",
      "age":10,
      "email":"m1779387qqwewqeqwe3123@163.com",
      "phone":"1779aqweqwea3873123@163.com",
      "description":"111"
  }'
```

* 响应示例：
```
{
	"code": 0,
	"errorMsg": "",
	"data": {
		"id": 17,
		"name": "1231231232131",
		"age": 10,
		"email": "m1779387qqwewqeqwe3123@163.com",
		"phone": "1779aqweqwea3873123@163.com",
		"description": "111",
		"create_time": "2021-11-05T10:08:54.634282677+08:00",
		"update_time": "2021-11-05T10:08:54.634282677+08:00"
	}
}
```

3 根据ID修改用户

* URL路径：
  ```/user```
  
* 请求示例：
```
curl http://{ip}:{port}/user \
  -X PUT \
  -H 'Content-Type: application/json' \
  -d '{  
      "id":17,
      "name":"4585959595",
      "age":10,
      "email":"m1779387qqwewqeqwe3123@163.com",
      "phone":"1779aqweqwea3873123@163.com",
      "description":"222"
  }'
```

* 响应示例：
```
{
	"code": 0,
	"errorMsg": ""
}
```

4 根据ID删除用户

* URL路径：
  ```/user/:id```
  
* 请求示例：
```
curl http://{ip}:{port}/user/17 \
  -X DELETE \
  -H 'Content-Type: application/json' \
  -d '{   }'
```

* 响应示例：
```
{
	"code": 0,
	"errorMsg": ""
}
```

