## Map的底层原理

### 1、map的核心结构

map的核心是两个结构体:hmap和bmap

创建map---->hmap{
    //字典的键值对个数
    //当使用len()获取map中的容量的时候，返回的就是这个值
    count
    //创建桶的个数为2的B次方
    //B用来决定创建多少个bmap
    B
    //当前map中桶的数组
    buckets
    //哈希因子，用于对key生成hash值
    //用于确定key应该放在哪儿
    hash0
    .....
}

其中buckets代表的桶数组，用于存储键值对
每个数组的值为bmap{
    //8个元素的数组，存储字典中每个key的哈希值的高8位
    tophash
    //8个元素的数组，存储字典key
    keys
    //8个元素的数组，存储字典的value
    values
    //指针，当前桶存不下的时候创建的溢出桶
    overflow
}

bmap真正存储键值对的单元，hmap统领这些bmap

### 2、map进行初始化

```go
// 初始化一个可容纳10个元素的map
info=map(map[string]string,10)
```

- 第一步：创建一个hmap结构体对象
- 第二步：生成一个哈希因子hash0并赋值刀hmap对象中(用于后续为key创建hash值)
- 第三步：根据hint=10，并根据算法规则来创建B,当前B应该为1

这个计算规则:
hint        B
0-8         0
9-13        1
14-26       2
.....

- 第四步：根据B去创建桶(bmap),并存放在buckets数组中，当前bmap的数量为2
  - 当B<4时，根据B创建的桶的个数规则为:2^B(标准桶)
  - 当B>4时，根据B创建的桶的个数的规则为:2^B+2^(B-4)（标准桶+溢出桶）

每个bmap中可以春初8个键值对，当不够存储时需要使用溢出桶，并将当前bmap中的overflow字段指向溢出桶的位置

### 3、写入数据

```go
info["name"]="lisi"
```

在map中写入数据时，内部的执行流程为:
- 第一步：结合哈希因子和键`name`生成hash值0100011010101101
- 第二步：获取哈希值的后B位，并根据后B位的值来确定此键值对存放到哪个桶中

这里其实是将哈希值和桶掩码进行与运算，将最终得到的值的后B位，假设为0，那么使用这个值作为桶数组的索引
将键值对存储到这个桶中

存储在这个桶的数据包括tophash,key,value分别写入到桶中的三个数组中

将数据写入到map后，还需要将count+1

### 4、读取数据

```go
value:=info["name"]
```

- 第一步：结合哈希因子和键name生成哈希值
- 第二步：获取哈希值的后B位，并根据后B位的值决定将此键值对存放到哪个桶中
- 第三步：确定桶后，再根据key的hash值计算出tophash（高8位），根据tophash和key去桶中查找数据

如果当前桶没找到，则根据overflow桶中找，都没找到就表明值不存在


### 5、扩容

在向map中添加数据时，当达到某个条件时，则会引发字典扩容

扩容条件：
- map中数据总个数/桶个数>6.5，引发翻倍扩容
- 使用了太多溢出桶(溢出桶使用的太多会导致map处理速度降低)
  - B<=15，已使用的溢出桶个数>=2^B时，引发等量扩容。
  - B>15，已使用的溢出桶的个数>=2^15时，引发等量扩容

当发生扩容的时候，hmap里的buckets会指向新桶，而oldbuckets会指向旧桶
这时候新桶里面还没有数据，数据还在旧桶里面

### 6、迁移

将旧桶中的数据迁移到新桶

如果是翻倍扩容，那么迁移就是将旧桶中的数据分流到新的两个桶中，并且桶位置编号：同编号位置和翻倍后对应编号位置
比如原来为0号桶的数据分流到新的零号和翻倍扩容的另一个零号位置

怎么实现的呢？

迁移时会遍历某个旧桶中所有的key(包括溢出桶)，并且根据key重新生成哈希值，根据哈希值的低B位来决定此键值对分流到哪个新的桶中

如果时等量扩容(溢出桶太多引发引发的扩容)，那么数据迁移的机制就比较简单，就是将旧桶(含溢出桶)中的值迁移到新桶中。
这种扩容和迁移的意义在于：当溢出桶较多而每个桶中的数据又不多时，可以通过等量扩容和迁移让数据更紧凑，从而减少溢出桶