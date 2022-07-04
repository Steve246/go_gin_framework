package main

//beda form data dan json

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	routerEngine := gin.Default()
	//pake extension thunder client atau pake postman buat masukin localhost di get

	// routerEngine := gin.Default()
	// routerEngine.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Healthy Check")
	// })

	// routerEngine.GET("/greeting", greetingDuqa)
	//http://localhost:8080/greeting

	// routerEngine.GET("/greeting/:name", greeting)

	// routerEngine.POST("/login", login)

	/*
	* Grouping: auth
	* www.enigmacamp.com/auth/login
	* www.enigmacamp.com/auth/register
	* www.enigmacamp.com/auth/logout

	* grouping: master
	* www.enigmacamp.com/master/users
	* www.enigmacamp.com/master/students
	* www.enigmacamp.com/api/auth/login

	 */

	routerEngine.GET("/", func(c *gin.Context) {
		c.String(200, "Healthy Check")
	})

	rgAuth := routerEngine.Group("/api/auth")
	rgAuth.POST("/login", login)
	//http://localhost:8080/api/auth/login

	rgMaster := routerEngine.Group("/api/master")
	rgMaster.GET("/greeting/:name", greeting)
	//http://localhost:8080/api/master/greeting/:name

	err := routerEngine.Run() //default, kalau kosong pake port 8080
	//di run, bisa masukin custom port --> misal ": 888"
	if err != nil {
		panic(err)
	}
}

// func greeting(c *gin.Context) {
// 	c.String(200, "Healthy Ok")
// }

/*
*Model Binding & Validation
* -> ShouldBind, ShouldBindJSON -> using struct tag binding required
 */

func login(c *gin.Context) {

	//http://localhost:8080/login?username=SteveJo&password=1234

	var uc UserCredential

	//ShouldBind dan ShouldBind Json

	//MustBind

	if err := c.BindJSON(&uc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": "ini_token",
		})
	}

	// username := c.PostForm("username")
	// c.PostForm("password")
	// c.String(http.StatusOK, "hello %s", username)

	//Cara akses di postman
	// Ubah jadi post
	// masukin http://localhost:8080/login
	// body --> form data
	//masukin username and password
	// yang nampak cma username
}

func greeting(c *gin.Context) {

	//kapan pake querry atau param

	// 	Any required or mandatory attributes should be added as path param
	// Any optional attributes should be added as query param
	// params used for filtering data are usually used as query param

	// name := c.Param("name") //with parameter name at path /greeting

	//querry param string in path -> ?key=value

	name := c.Param("name")

	kec := c.Query("kecamatan")
	kel := c.Query("kelurahan")
	c.String(http.StatusOK, "Helo....%s saat ini kamu ada di kec %s kel %s", name, kec, kel)

	//http://localhost:8080/greeting/Steven?kecamatan=Kalideres&kelurahan=Rawa Buata

}
