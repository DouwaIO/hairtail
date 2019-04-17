# Hairtail


## schema

```
version: v1
kind: schema
name: storage
columns:
  - name: id
    desc: a identity data
    type: int
    default: 0
  - name: name
    desc: the name of people
    type: int
    default: ''
  - name: sex
    desc: '0: male, 1: female'
    type: int
    default: 0
```

字段说明：

version: `API的版本号`
columns:
- name: `名称，必须且不能重复`
- desc: `描述，可不填`
- type: `类型，必填，只能是下面几种`
  - int
  - date
  - datetime
  - string
  - float
  - text
- default: `默认值，选填，用于获取不到数据的时候提供默认值`

## task

```
version: v1
kind: task
```

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
  - 打平
      - 主从表转换成单表
  - aggregate聚合
      - 根据groupby条件进行聚合
      - 聚合类型：count, avg, sum
      - 支持定义输出格式
  - filter过滤
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

## pipeline

```
version: v1
kind: pipeline
steps:
  - name: input_data_from_mq
    desc: 从mq中监听数据的获取
    type: mqGet
    settings:
      protocol: amop
      host:
      user:
      topic:
      ackPolicy: immediate

  - name: checkmqdatafrombeforestep
    desc: 验证上个步骤获取的数据是否符合定义的数据模型
    type: dataCheck
    settings:
      model: storage

  - name: outputToOtherModel
    desc: 验证上个步骤获取的数据是否符合定义的数据模型
    type: outputDataModel
    settings:
      model: storage1
```
