package router

func (s Server) RegisterRouter() {
	s.eng.GET("order/:order_no", s.hld.QueryOrder())
	s.eng.POST("order", s.hld.AddOrder())
	s.eng.DELETE("order/:order_no", s.hld.DeleteOrder())
	s.eng.PUT("order", s.hld.UpdateOrder())
	s.eng.GET("order/list", s.hld.QueryOrderList())
	s.eng.POST("upload", s.hld.Upload())
	s.eng.GET("download", s.hld.Download())
	s.eng.GET("save_excel", s.hld.SaveExcel())
}
