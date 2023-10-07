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

### 5、ES的文档操作

**文档创建操作**

1. 单个文档添加

```go
func CreateDoc(){
    user := models.UserModel{
    ID:        12,
    UserName:  "lisi",
    Age:       23,
    NickName:  "夜空中最亮的lisi",
    CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
    Title:     "今天天气很不错",
  }
  indexResponse, err := global.ESClient.Index().Index(user.Index()).BodyJson(user).Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("%#v\n", indexResponse)
}
```

2. 批量添加文档

```go
func DocCreateBatch() {

  list := []models.UserModel{
    {
      ID:        12,
      UserName:  "fengfeng",
      NickName:  "夜空中最亮的枫枫",
      CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
    },
    {
      ID:        13,
      UserName:  "lisa",
      NickName:  "夜空中最亮的丽萨",
      CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
    },
  }

  bulk := global.ESClient.Bulk().Index(models.User{}.Index()).Refresh("true")
  for _, model := range list {
    req := elastic.NewBulkCreateRequest().Doc(model)
    bulk.Add(req)
  }
  res, err := bulk.Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(res.Succeeded())
}
```

mapping里面没有的字段，再创建的时候会动态的添加
text类型的数据不能排序

**删除文档**

删除文档有两种：

1. 根据id删除文档

```go

// 这里文档如果不存在会报错

func DocDelete() {

  // 这里Refresh 这个参数意思是是否立即删除索引的意思
  // 如果设置为false 会过一段时间才会删除
  deleteResponse, err := global.ESClient.Delete().
    Index(models.UserModel{}.Index()).Id("tmcqfYkBWS69Op6Q4Z0t").Refresh("true").Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(deleteResponse)
}
```

2. 根据id批量删除

```go
func DocDeleteBatch() {
  idList := []string{
    "tGcofYkBWS69Op6QHJ2g",
    "tWcpfYkBWS69Op6Q050w",
  }
  bulk := global.ESClient.Bulk().Index(models.UserModel{}.Index()).Refresh("true")
  for _, s := range idList {
    req := elastic.NewBulkDeleteRequest().Id(s)
    bulk.Add(req)
  }
  res, err := bulk.Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(res.Succeeded())  // 实际删除的文档切片
}

// 如果文档不存在，不会有错误， res.Succeeded() 为空

```

**查询文档**

1. 列表查询

```go
func DocFind() {

  limit := 2
  page := 4
  from := (page - 1) * limit
  
  // es的查询条件
  // From和Size分别为分页查询的条件
  // 这里代表查询全部数据  
  query := elastic.NewBoolQuery()
  res, err := global.ESClient.Search(models.User{}.Index()).Query(query).From(from).Size(limit).Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  count := res.Hits.TotalHits.Value  // 总数
  fmt.Println(count)
  for _, hit := range res.Hits.Hits {
    fmt.Println(string(hit.Source))
  }
}
```

2. 精确匹配查询

```go
// 给query设置条件
query := elastic.NewTermQuery("user_name", "fengfeng")
```

3. 模糊匹配

模糊匹配主要查询text,也可以查询keyword
模糊匹配keyword需要搜索完整的

```go
query := elastic.NewMatchQuery("nick_name", "夜空中最亮的枫枫")
```

4. 嵌套字段的查询

```json
"title": {
    "type": "text",
    "fields": {
        "keyword": {
            "type": "keyword",
            "ignore_above": 256
        }
    }
},
```

因为title是text类型，只能模糊匹配，但是需要精确匹配的时候，也能通过title.keyword的形式进行精确匹配

```go
query := elastic.NewTermQuery("title.keyword", "这是我的枫枫") // 精确匹配
//query := elastic.NewMatchQuery("title", "这是我的枫枫")  // 模糊匹配
```

**更新文档**

```go
func DocUpdate() {
  res, err := global.ESClient.Update().Index(models.UserModel{}.Index()).
    Id("vmdnfYkBWS69Op6QEp2Y").Doc(map[string]any{
    "user_name": "你好枫枫",
  }).Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("%#v\n", res)
}
```