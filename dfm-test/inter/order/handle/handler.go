package handle

import (
	"github.com/min65535/demo/dfm-test/inter/order/service"
	"github.com/min65535/demo/dfm-test/pkg/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go/build"
	"net/http"
)

type Handler struct {
	order service.OrderSvc
}

func NewHandler(order service.OrderSvc) *Handler {
	return &Handler{
		order: order,
	}
}

func (s *Handler) QueryOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		no := c.Param("order_no")
		if len(no) == 0 {
			s.httpErr(c, http.StatusBadRequest, "Not found no.")
			return
		}
		o, err := s.order.QueryByNo(no)
		if err != nil {
			s.httpErr(c, http.StatusInternalServerError, err.Error())
			return
		}
		s.httpSuccess(c, &Order{
			// ID:       o.ID,
			OrderNo:  o.OrderNo,
			UserName: o.UserName,
			Amount:   o.Amount,
			Status:   o.Status,
			FileUrl:  o.FileUrl,
		})
	}
}

func (s *Handler) QueryOrderList() gin.HandlerFunc {
	return func(c *gin.Context) {
		orders, err := s.order.QueryList(0, 100)
		if err != nil {
			s.httpErr(c, http.StatusInternalServerError, err.Error())
			return
		}
		s.httpSuccess(c, orders)
	}
}

func (s *Handler) AddOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddOrderReq
		if err := c.BindJSON(&req); err != nil {
			s.httpErr(c, http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println(req)

		if err := req.IsValid(); err != nil {
			s.httpErr(c, http.StatusBadRequest, err.Error())
			return
		}

		o := model.DemoOrder{
			OrderNo:  req.OrderNo,
			UserName: req.UserName,
			Amount:   req.Amount,
			Status:   req.Status,
			FileUrl:  req.FileUrl,
		}

		if err := s.order.Create(&o); err != nil {
			s.httpErr(c, http.StatusInternalServerError, err.Error())
			return
		}

		s.httpSuccess(c, &Order{
			// ID:       o.ID,
			OrderNo:  o.OrderNo,
			UserName: o.UserName,
			Amount:   o.Amount,
			Status:   o.Status,
			FileUrl:  o.FileUrl,
		})
	}
}

func (s *Handler) UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Order

		if err := c.BindJSON(&req); err != nil {
			s.httpErr(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := req.IsValid(); err != nil {
			s.httpErr(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := s.order.UpdateByNo(req.OrderNo, map[string]interface{}{
			"order_no":  req.OrderNo,
			"user_name": req.UserName,
			"amount":    req.Amount,
			"status":    req.Status,
			"file_url":  req.FileUrl,
		}); err != nil {
			s.httpErr(c, http.StatusInternalServerError, err.Error())
			return
		}

		s.httpSuccess(c, "success")
	}
}

func (s *Handler) DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		no := c.Param("order_no")
		if len(no) == 0 {
			s.httpErr(c, http.StatusBadRequest, "Not found no.")
			return
		}

		if _, err := s.order.DeleteByNo(no); err != nil {
			s.httpErr(c, http.StatusInternalServerError, err.Error())
			return
		}
		s.httpSuccess(c, "success")
	}
}

func (s *Handler) Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		no := c.PostForm("order_no")
		if len(no) == 0 {
			s.httpErr(c, http.StatusBadRequest, "Not found no.")
			return
		}
		mt, err := c.FormFile("file")
		if err != nil {
			s.httpErr(c, http.StatusBadRequest, err.Error())
			return
		}
		path := fmt.Sprintf("%s/%s", build.Default.GOPATH, mt.Filename)
		fmt.Println(path)
		if err = c.SaveUploadedFile(mt, path); err != nil {
			s.httpErr(c, http.StatusInternalServerError, err.Error())
			return
		}
		if err = s.order.UpdateByNo(no, map[string]interface{}{
			"file_url": path,
		}); err != nil {
			s.httpErr(c, http.StatusInternalServerError, err.Error())
			return
		}
		s.httpSuccess(c, path)
	}
}

// 下载文件
func (s *Handler) Download() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename, err := s.order.OrderDownload()
		if err != nil {
			fmt.Println("err:", err.Error())
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		c.File(filename)
	}
}

// 下载文件
func (s *Handler) SaveExcel() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename, err := s.order.OrderDownload()
		if err != nil {
			fmt.Println("err:", err.Error())
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		c.File(filename)
	}
}

func (s *Handler) httpErr(c *gin.Context, code int, err string) {
	c.JSON(code, NewHttpErrResp(err))
}

func (s *Handler) httpSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewHttpSuccessResp(data))
}
