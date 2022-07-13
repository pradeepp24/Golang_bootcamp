package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/Models"
)

func CreateProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, "product id not found")
		return
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func GetAllOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
