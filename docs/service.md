# service

- [x] mq: 监听mq
    - amop
    - ServiceBus(Azure)
- [x] webhook: 通过`api`触发，所有`pipeline`都有效，不需要配置，无法关闭
- [ ] file: 通过读取文件获取数据(文本、Excel)，是否监视
- [ ] table: 通过读取表获取数据(数据库中的表)，是否监视
- [ ] cron: 定时任务

## mq

通过监听`mq`获取`pipeline`触发数据

1. 输入
    - 无
1. 输出
    - 获取到的`mq`数据
1. 参数
    - protocol: 支持MQ协议
        - amop: 默认
    - host: 服务器地址
    - user: 账号
    - password: 密码
    - topic: 订阅主题
    - ackPolicy: 确认方式
        - immediate: 默认

## webhook

使用`hairtail`提供的接口，并指定`pipeline`的名称就可以调用`POST`接口提交数据并触发相应的`pipeline`的执行

所有`pipeline`有效，无法关闭

## file

## table

## cron

定时任务，支持定义`pipeline`在特定的时间内触发，最小支持到`秒`级别

1. 参数
    - year
        - *
        - 1000 ~ 9999
    - month
        - *
        - 1 ~ 12
    - day
        - *
        - 1 ~ 31
    - hour
        - *
        - 0 ~ 23
    - minute
        - *
        - 1 ~ 59
    - second
        - *
        - 1 ~ 59
