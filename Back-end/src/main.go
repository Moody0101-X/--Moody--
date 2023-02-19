package main

import (   	
    "fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"moody.com/api/database"
	"moody.com/api/routes"
	"log"
)

/*
	
	This App needs this functionality.
		[] Ability to auth an admin. if an admin is authenticated they can perform various functions such as
			An admin is {
				"id",
				"name",
			}

			-> add products.
			-> remove products.
			
			A product is {
				""
			}
*/

func RequestCancelRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			
			if err := recover(); err != nil {
				fmt.Println("The Request was cancelled because an unexpected error interupted.")
				fmt.Println("err:\n")
				log.Fatal(err);
				
				c.Request.Context().Done()
			}

		}()
		
		c.Next()
	}	
}

func run() {
	// GET THE PORT :)
	var PORT string = ":8888"
	// var CDN string = "http://" + networking.GetCurrentMachineIp() + ":8500"
	// var APP_LINK string = "http://" + networking.GetCurrentMachineIp() + PORT
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())	
	router.Use(gin.Logger(), RequestCancelRecover())
	// HTML/JS/CSS/IMG loaders

	// router.Static("/static", "./public/static")
	// router.Static("/img", "./public/img")
	// router.Static("/Global.css", "./public/Global.css")
	// router.LoadHTMLGlob("public/*.html")
	// router.Static("/", "./public")
	
	router.GET("/", routes.Index)
    router.NoRoute(routes.Index)
    
	// running the server.
	fmt.Println("[!] Serving in port -> ", PORT)
	// fmt.Println("[!] Using cdn ->", CDN)
	// fmt.Println("[!] Go to this link to review the app: ", APP_LINK)
	fmt.Println()
	router.Run(PORT)
}

func main() {
    
    fmt.Println()
	fmt.Println()

    // var env string = models.GetEnv("ENV")
    
    // fmt.Println("[!] Environement ->", env)

    var DB string = "./Moody.db"
    
    err, path := database.InitializeDb(DB);

    if err != nil {
        fmt.Println("Error opening the database! ", err.Error())
        return
    }

    fmt.Println("[!] connected to database ->", path)
    run()
}

