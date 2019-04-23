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

## pipeline

```
version: v1
kind: pipeline
steps:
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
services:
  - name: input_data_from_mq
    desc: 从mq中监听数据的获取
    type: mqGet
    settings:
      protocol: amop
      host:
      user:
      topic:
      ackPolicy: immediate
```

### services

- `services`和`steps`的差别是，`steps`需要触发才会执行，而`services`在`pipeline`生效后会一直启动着
- `services`一般使用前置监听任务，比如`mq`和`http post`用于等待数据源信息

## Transcation（事务）

后续涉及事务会考虑使用`Saga`或`Process Manager`概念，详见: [](https://cloud.tencent.com/developer/article/1160083)和[](https://blog.csdn.net/ethanwhite/article/details/53766018)
