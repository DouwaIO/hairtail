# Hairtail


## data model

```
version: v1
kind: model
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

## pipeline

```
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
