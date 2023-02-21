package routes

import (
	"fmt"
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
		route: /User/Auth/Login
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
	
	c.BindJSON(&AuthUser)
	
	Result := database.AuthenticateUser(&AuthUser);

	if Result.Ok {
		c.JSON(OK, models.MakeServerResp(OK, AuthUser));
		return
	}

	c.JSON(OK, models.MakeServerRespFromResult(Result));
	
}


func SingUp(c *gin.Context) {
	
	/*
		route: /User/Auth/SignUp
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
	
	c.BindJSON(&New);
	
	fmt.Println("--Sign-up--");
	Result := database.SignUpAUser(&New);
	
	if Result.Ok {
		c.JSON(OK, models.MakeServerResp(OK, New));
		return
	}

	c.JSON(OK, models.MakeServerRespFromResult(Result));
}

func AddNewProductRoute(c *gin.Context) {

	/*
		route: /Product/Add
		
		Expects => {
			"img": string
			"owner_id": fk => int
			"label": string
			"price": int
			"vendor": string
			"jwt": string
		}
		
		Ok 	 	bool
		Desc 	string
		Code    int
	*/

	var P models.Product;
	c.BindJSON(&P);

	Result := database.AddNewProduct(&P);

	if Result.Ok {
		c.JSON(OK, models.MakeServerResp(OK, P));
		return
	}


	c.JSON(OK, models.MakeServerRespFromResult(Result));
}

func RemoveProductRoute(c *gin.Context) {
	NOTIMPLEMENTED(c, "/Product/Remove");
}

func BuyRoute(c *gin.Context) {
	NOTIMPLEMENTED(c, "/Product/Buy");
}
