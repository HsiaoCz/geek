## ElasticSearch

### 1、ES 安装

使用 docker 进行安装

```docker
docker pull elasticsearch:7.12.8
```

创建 docker 容器挂载的目录:

```bash
# linux的命令
mkdir -p /opt/es/config & mkdir -p /opt/es/data & mkdir -p /opt/es/plugins

chmod 777 /opt/es/data
```

创建配置文件

```bash
echo "http.host: 0.0.0.0" > /opt/es/config/elasticsearch.yml
```

创建容器

```bash
# linux
docker run --name es -p 9200:9200  -p 9300:9300 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms84m -Xmx512m" -v /opt/es/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v /opt/es/data:/usr/share/elasticsearch/data -v /opt/es/plugins:/usr/share/elasticsearch/plugins -d elasticsearch:7.12.0

```

这时，访问:9200 就可以了

### 2、连接 ElasticSearch

这里使用

```go
go get github.com/olivere/elastic/v7
```

```go
package core

import (
  "es_study/global"
  "fmt"
  "github.com/olivere/elastic/v7"
)

func EsConnect() {

  client, err := elastic.NewClient(
    // 连接es的地址
    elastic.SetURL("http://127.0.0.1:9200"),
    // 这里的es是使用docker搭建的，需要关闭一个检测才能连接成功
    elastic.SetSniff(false),
    // 设置用户名和密码
    elastic.SetBasicAuth("", ""),
  )
  if err != nil {
    fmt.Println(err)
    return
  }
   // 利用全局变量的方法使用client
  global.ESClient = client
}
```

### 3、ES 的认证

给 es 设置用户名和密码

可以使用这个教程:

`https://blog.csdn.net/qq_38669698/article/details/130529829`

### 4、ES 的索引操作

mapping 索引的常见类型

```json
{
  "mappings": {
    "properties": {
      "title": {
        "type": "text" // 查询的时候是分词匹配
      },
      "key": {
        "type": "keyword" // 完整匹配
      },
      "user_id": {
        "type": "integer"
      },
      "created_at": {
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
```

**索引的创建**

```go
// 这里先写哥model

package model

type User struct{
    ID uint `json:"id"`
    Username string `json:"user_name"`
    NickName string `json:"nick_name"`
    CreatedAt time.Time `json:"created_at"`
}

func (u User)Mapping()string{
    return `{
  "mappings": {
    "properties": {
      "nick_name": {
        "type": "text" // 查询的时候是分词匹配
      },
      "user_name": {
        "type": "keyword" // 完整匹配
      },
      "id": {
        "type": "integer"
      },
      "created_at": {
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}`
}
```

```go
package es

func CreateIndex(){
    createIndex, err := global.ESClient.
    CreateIndex("user_index").
    BodyString(model.User{}.Mapping()).Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(createIndex)
  fmt.Println("索引创建成功")
}
```

索引已经存在会报错

所以在索引创建的时候需要先判断一下是索引是否存在
```go
func ExistsIndex(index string) bool {
  exists, _ := global.ESClient.IndexExists(index).Do(context.Background())
  return exists
}
```

索引创建完成访问127.0.0.1:9200/user_idenx/_mapping可以查看创建的索引

**删除索引**

```go
func DeleteIndex(index string) {
  _, err := global.ESClient.
    DeleteIndex(index).Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(index, "索引删除成功")
}
```

在创建索引的时候，可以先判断索引是否存在，如果存在，先删除再创建

```go
func CreateIndex(){
    index:="user_index"
    if ExistsIndex(index){
        // 索引存在，先删除
        DeleteIndex(index)
    }
    createIndex, err := global.ESClient.
    CreateIndex(index).
    BodyString(model.User{}.Mapping()).Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(createIndex)
  fmt.Println("索引创建成功")
}
```

