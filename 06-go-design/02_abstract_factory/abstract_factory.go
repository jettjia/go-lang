package abstract_factory

// OrderMainDAO 为订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详情纪录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory DAO 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO     //嵌套了子接口  创建订单-主要数据对象
	CreateOrderDetailDAO() OrderDetailDAO //嵌套了子接口  创建订单-详细数据对象
}
