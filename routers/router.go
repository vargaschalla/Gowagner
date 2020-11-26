package routers

import (
	"net/http"
	"strings"

	"github.com/vargaschalla/Gowagner/models"

	"github.com/gin-gonic/gin"
	"github.com/vargaschalla/Gowagner/apis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {

	conn, err := connectDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
		//return
	}

	conn.AutoMigrate(
		&models.Persona{},
		&models.User{},
	)
	r := gin.Default()

	//config := cors.DefaultConfig() https://github.com/rs/cors
	//config.AllowOrigins = []string{"http://localhost", "http://localhost:8086"}

	r.Use(CORSMiddleware())

	r.Use(dbMiddleware(*conn))

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", apis.ItemsIndex)
		v1.GET("/persons", apis.PersonsIndex)
		v1.POST("/persons", apis.PersonsCreate)
		//v1.GET("/persons/:id", apis.PersonsGet)
		v1.PUT("/persons/:id", apis.PersonsUpdate)
		v1.DELETE("/persons/:id", apis.PersonsDelete)

		v1.GET("/users", apis.UsersIndex)
		v1.POST("/users", apis.UsersCreate)
		v1.GET("/users/:id", apis.UsersGet)
		v1.PUT("/users/:id", apis.UsersUpdate)
		v1.DELETE("/users/:id", apis.UsersDelete)
		v1.POST("/login", apis.UsersLogin)
		v1.POST("/logout", apis.UsersLogout)
	}

	return r
}

func connectDBmysql() (c *gorm.DB, err error) {

	dsn := "root:aracelybriguit@tcp(localhost:3306)/academico?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	return conn, err
}

func connectDB() (c *gorm.DB, err error) {
	////dsn := "docker:docker@tcp(mysql-db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "user=artuqahvnfvpxi password=4350a71d518b779f4f83fb21efaecf2f1046ab9512a9123438f58ba6c4615b28 host=ec2-54-157-88-70.compute-1.amazonaws.com dbname=d161cscs5jpnth port=5432 sslmode=require TimeZone=Asia/Shanghai"
	//dsn := "user=postgres password=postgres2 dbname=users_test host=localhost port=5435 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return conn, err
}

func dbMiddleware(conn gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//c.Header("Access-Control-Allow-Origin", "http://localhost, http://localhost:8086,")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func authMiddleWare() gin.HandlerFunc { //ExtractToken
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := models.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated (IsTokenValid)."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
