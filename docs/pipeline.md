# Hairtail

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
