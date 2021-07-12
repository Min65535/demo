package dao

import (
	"demo/dfm-test/pkg/model"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestNewOrderDao(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockOrderDao(ctrl)

	m.EXPECT().QueryByNo(gomock.Eq("1")).Return(model.DemoOrder{}, nil)// errors.New("not exist"))
	if v,err := m.QueryByNo("1"); err != nil {
		fmt.Println(v.UserName,err)
		t.Fatal("expected -1, but got", err)
	}

}
