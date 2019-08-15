package task

import (
	"testing"
)

func TestSplitData(t *testing.T) {
	for _, unit := range []struct {
		data []byte
		ret1 []byte
    }{
        {[]byte("{\"name\":\"fabric_stock_in\",\"type\":\"add\",\"key\":\"F000323546\",\"time\":\"2019-04-23T03:10:33.775Z\",\"data\":{\"bill_no\":\"string\",\"details\":[{\"grade\":\"string\",\"location_no\":\"string\",\"line\":\"ASDF111\"},{\"fabric_no\":\"string\"}],\"bill_type_name\":\"采购入库\"}}"), []byte("[{\"data__bill_no\":\"string\",\"data__bill_type_name\":\"采购入库\",\"data__details__grade\":\"string\",\"data__details__line\":\"ASDF111\",\"data__details__location_no\":\"string\",\"key\":\"F000323546\",\"name\":\"fabric_stock_in\",\"time\":\"2019-04-23T03:10:33.775Z\",\"type\":\"add\"},{\"data__bill_no\":\"string\",\"data__bill_type_name\":\"采购入库\",\"data__details__fabric_no\":\"string\",\"key\":\"F000323546\",\"name\":\"fabric_stock_in\",\"time\":\"2019-04-23T03:10:33.775Z\",\"type\":\"add\"}]")},
        // {[]byte("asdd"), []byte("asddqwkkjhjkeqwe")},

    } {
        // 调用排列组合函数，与期望的结果比对，如果不一致输出错误
        if ret1 := SplitData(unit.data); string(ret1) != string(unit.ret1) {
            t.Errorf("ret1: [%v]", ret1)
        }
    }
}