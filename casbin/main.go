package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {
	a, _ := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/hrm", true)
	e, _ := casbin.NewEnforcer("./model.conf", a)
	//sub := "alice"
	//obj := "data1"
	//act := "read"
	// 添加一些规则
	//_, err := e.AddPolicy("alice", "dat1", "write")
	//_, err = e.AddPolicy("alice", "data1", "write1")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//ok, err := e.Enforce(sub, obj, act)
	//if err != nil {
	//}
	//fmt.Println(ok)

	// 1.删除一些路由规则
	//success, _ := e.RemoveFilteredPolicy(0, "bo")
	//if success {
	//	fmt.Println("")
	//}
	// 2.获取路由
	list := e.GetFilteredPolicy(0, "alice")
	for _, v := range list {
		fmt.Println(v)
	}
}
