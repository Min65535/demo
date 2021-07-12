package service

import (
	"demo/dfm-test/inter/order/dao"
	"demo/dfm-test/pkg/model"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceTestSuite struct {
	ctl   *gomock.Controller
	dao   *dao.MockOrderDao
	svc   OrderSvc
	clean func()
	suite.Suite
}

func (s *ServiceTestSuite) TearDownTest() {
	s.clean = func() {
		defer s.ctl.Finish()
	}
}

func (s *ServiceTestSuite) SetupTest() {
	s.ctl = gomock.NewController(s.T())
	s.dao = dao.NewMockOrderDao(s.ctl)
	s.svc = NewOrderService(s.dao)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) Test_Create() {
	var req = &model.DemoOrder{
		OrderNo:  "sss",
		UserName: "s",
		Amount:   11,
		Status:   "sds",
		FileUrl:  "sadasd",
	}
	s.dao.EXPECT().Create(req).AnyTimes().Return(nil)
	err := s.svc.Create(req)
	if err != nil {
		s.T().Log("err:", err.Error())
	}
	assert.NoError(s.T(), err)
	var req1 = &model.DemoOrder{
		OrderNo:  "ssssss",
		UserName: "ssss",
		Amount:   11,
		Status:   "sds",
		FileUrl:  "sadasd",
	}
	s.dao.EXPECT().Create(req1).AnyTimes().Return(errors.New("bad error"))
	err = s.svc.Create(req1)
	if err != nil {
		s.T().Log("err:", err.Error())
	}
	assert.EqualError(s.T(), err, "bad error")
}
