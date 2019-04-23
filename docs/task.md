# task

- 数据获取
  - mq
  - url
    - 主动
    - 被动
- 基本操作
  - format格式化
      - 日期
      - 金额
  - calculate计算
      - 对两个字段进行计算，一般是数量计算
  - even打平
      - 主从表转换成单表
  - aggregate聚合
      - 根据groupby条件进行聚合
      - 聚合类型：count, avg, sum
      - 支持定义输出格式
  - filter过滤(支持数据分离，filter后会得到满足条件和不满足条件的两个数据集）
  - join连接(左、右、全、内)
  - select选择(选择需要的字段)
      - 过滤掉不需要的字段
- 特殊操作
  - dispatch分摊
      - 根据特定条件，将数量按目标顺序进行从头到为分摊
  - accumulate累加
      - 根据特定条件，将数量更新到目标记录中
- 输出数据
  - output
  - url


## even

数据平整器，用于将主从数据进行平整处理

1. 输入

一个JSON列表，支持列表的每一行进行平整处理

1. 输出

- 成功：状态0，输出一个平整后的JSON列表
- 失败：状态1，并带有错误信息

1. 配置

```yaml
type: event
settings:
  dept: 5
  char: __
  include: ['data.sales', 'data.order_no']
  exclude: ['data.customers']
```

- dept: 解析深度，默认5，对JSON的数据层级别进行平整解析的最大深度
- char: 字段分离符号，默认`__`，用于分割解析后上下级key名称，如上例中：`data__sales`
- exclude: 排除需要进行解析的key
- include：包含要解析的key

### 例子

- 例1

配置
```yaml
type: event
settings:
  dept: 5
  char: __
  include: ['data.sales', 'data.order_no']
```

输入
```json
{"datetime": "2019-04-03 10:22:20",
"data": {
  "order_no": "001",
  "sales": [{
    "name": "张三"
  }, {
    "name": "李四"
  }],
  "customers": [{
    "name": "安踏"
  }, {
    "name": "李宁"
  }]
}
}
```

输出
```json
[{"data__order_no": "001", "data__sales__name": "张三"}
,{"data__order_no": "001", "data__sales__name": "李四"}
]
```

- 例2
配置
```yaml
type: event
settings:
  dept: 5
  char: __
```

输入
```json
{"datetime": "2019-04-03 10:22:20",
"data": {
  "order_no": "001",
  "sales": [{
    "name": "张三"
  }, {
    "name": "李四"
  }],
  "customers": [{
    "name": "安踏"
  }, {
    "name": "李宁"
  }]
}
}
```

输出
```json
[{"datatime": "2019-04-03 10:22:20", "data__order_no": "001", "data__sales__name": "张三", "data__customers_name": "安踏"}
,{"datatime": "2019-04-03 10:22:20", "data__order_no": "001", "data__sales__name": "张三", "data__customers_name": "李宁"}
,{"datatime": "2019-04-03 10:22:20", "data__order_no": "001", "data__sales__name": "李四", "data__customers_name": "安踏"}
,{"datatime": "2019-04-03 10:22:20", "data__order_no": "001", "data__sales__name": "李四", "data__customers_name": "李宁"}
]
```
