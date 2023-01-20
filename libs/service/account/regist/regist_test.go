package regist_service

import (
	"fmt"
	"testing"
)

// 测试创建用户
func TestCreateUser(t *testing.T) {
	usr, e := CreateUser("testUser1", "test_password", "test_salt")
	if e != nil {
		e.Panic()
	}

	fmt.Println(usr)
}

func TestRuleUsername(t *testing.T) {
	// 可用校验
	name := "0lakjd"
	name2 := "_-_"
	// 错误名称校验
	name3 := "1,z"
	// 长度溢出校验
	name4 := "123123123123123123123123123"
	type TestStruct struct {
		name   string
		target bool
	}
	testList := []TestStruct{
		{name: name, target: true},
		{name: name2, target: true},
		{name: name3, target: false},
		{name: name4, target: false},
	}
	for _, test := range testList {
		if ruleUsername(test.name) != test.target {
			t.Errorf("error test: %s", test.name)
		}
	}
}
