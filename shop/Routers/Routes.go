package Routes

import (
	"github.com/gin-gonic/gin"
	"shop/Controllers"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/shop-api")
	{
		grp1.GET("product", Controllers.GetProducts)
		grp1.GET("product/:id", Controllers.GetProductByID)

		grp1.POST("retailer", Controllers.CreateProduct)
		grp1.PUT("retailer/:id", Controllers.UpdateProduct)
		grp1.DELETE("retailer/:id", Controllers.DeleteProduct)
		grp1.GET("retailer/order", Controllers.GetAllOrders)

		grp1.GET("customer", Controllers.GetCustomers)
		grp1.POST("customer/register", Controllers.RegisterCustomer)
		grp1.POST("customer/order", Controllers.PlaceOrder)
		grp1.GET("customer/:id", Controllers.CustomerOrders)
	}

	return r
}
