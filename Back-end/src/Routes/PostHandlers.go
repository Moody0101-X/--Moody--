package routes

import (
	// "net/http"
	"github.com/gin-gonic/gin"
	"moody.com/api/models"
	"moody.com/api/database"
)



/*
Priority Queue.

	[*] SignUp
	[*] Login
	
	[] Get/remove/Add products
	[] Buy product.
	[] Get User/Users.
	[] Get User product.

*/

func Login(c *gin.Context) {
	/*
		Expects => {
			"email": string
			"pwd": string
		} >> or << {
			"jwt ": string
		}
			Ok 	 	bool
			Desc 	string
			Code    int
	*/
	var AuthUser models.User;
	
	c.bindJSON(&AuthUser)
	
	Result := database.AuthenticateUser(&AuthUser);

	if Result.Ok {
		c.JSON(Result.Code, AuthUser);
		return
	}

	c.JSON(Result.Code, Result.Desc);
}


func SingUp(c *gin.Context) {
	/*
		Expects => {
			"email": string
			"pwd": string
			"name": string
			"phone_number": string
		}
			Ok 	 	bool
			Desc 	string
			Code    int
	*/

	var New models.User;

	c.bindJSON(&New);
	
	Result := database.SignUpAUser(&New);
	
	if Result.Ok {
		c.JSON(Result.Code, New);
		return
	}

	c.JSON(Result.Code, Result.Desc);
}

func AddNewProductRoute(c *gin.Context) {
	NOTIMPLEMENTED(c, "/Product/Add");
}

func RemoveProductRoute(c *gin.Context) {
	NOTIMPLEMENTED(c, "/Product/Remove");
}

func BuyRoute(c *gin.Context) {
	NOTIMPLEMENTED(c, "/Product/Buy");
}
