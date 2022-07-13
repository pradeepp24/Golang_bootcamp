package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"retail_shop/Models"
)

//grp1.POST("customer/register", Controllers.RegisterCustomer)
//grp1.POST("customer/order", Controllers.PlaceOrder)
//grp1.GET("customer/:id", Controllers.CustomerOrders)

func GetCustomers(c *gin.Context) {
	var customer []Models.Customer
	err := Models.GetAllCustomers(&customer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

func PlaceOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	id := order.CustomerId
	fmt.Println(id)
	err := Models.PlaceOrder(&order, id)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(err)
		c.JSON(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, order)
	}
}
func RegisterCustomer(c *gin.Context) {
	var customer Models.Customer
	c.BindJSON(&customer)
	err := Models.RegisterCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
func CustomerOrders(c *gin.Context) {
	id := c.Params.ByName("id")
	var order []Models.Order
	err := Models.GetCustomerOrders(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
