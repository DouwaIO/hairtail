package main

import (
	"fmt"
	"github.com/DouwaIO/hairtail/src/pipeline"
)
func main() {

	test := `steps:
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
    `
	branches, err := pipeline.ParseString(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(branches)
}
