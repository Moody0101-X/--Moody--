package routes

import (
	"github.com/gin-gonic/gin"
	"moody.com/api/database"
	"moody.com/api/models"
)

func GetUserProductByUUIDRoute(c *gin.Context) {
	NOTIMPLEMENTED(c, "/User/Products/{ UUID }");
}

func GetProductByIdRoute(c *gin.Context) {
	NOTIMPLEMENTED(c, "/Product/{ ID }");
}

func GetAllProductsRoute(c *gin.Context) {	
	err, products := database.GetAllProducts();	
	
	if err != nil {
		c.JSON(OK, models.MakeServerResp(204, "No content!"));
		return
	}

	c.JSON(OK, models.MakeServerResp(OK, products));
}