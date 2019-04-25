package pipeline

import (
	"testing"
)

func TestTest(t *testing.T) {
	for _, unit := range []struct {
		data []byte
		ret1 []byte
		ret2 error
    }{
        {[]byte("asd"), []byte("asdqweqwe"), nil},
        {[]byte("asdd"), []byte("asddqwkkjhjkeqwe"), nil},

    } {
        // 调用排列组合函数，与期望的结果比对，如果不一致输出错误
        if ret1,ret2 := Test2(unit.data); string(ret1) != string(unit.ret1) {
            t.Errorf("ret1: [%v], ret2: [%v]", ret1, ret2)
        }
    }
}