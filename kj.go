package krakenjwt

import (
	"net/http"
	"strconv"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type KJ struct {
	//router *gin.Engine
	server     *http.Server
	listenAddr string
}

func (k *KJ) getRandomName(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	gofakeit.Seed(int64(id))
	c.JSON(http.StatusOK, gin.H{"name": gofakeit.Name()})
}

func (k *KJ) getRandomEmail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	gofakeit.Seed(int64(id))
	c.JSON(http.StatusOK, gin.H{"name": gofakeit.Email()})
}

func (k *KJ) getRandomPhone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	gofakeit.Seed(int64(id))
	c.JSON(http.StatusOK, gin.H{"name": gofakeit.Phone()})
}

func (k *KJ) getRandom(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	gofakeit.Seed(int64(id))

	field := c.Param("field")

	var v interface{}

	switch field {
	case "name":
		v = gofakeit.Name()
	case "email":
		v = gofakeit.Email()
	case "phone":
		v = gofakeit.Phone()
	case "person":
		v = gofakeit.Person()
	case "car":
		v = gofakeit.Car()
	case "job":
		v = gofakeit.Job()
	case "contact":
		v = gofakeit.Contact()
	case "currency":
		v = gofakeit.Currency()
	case "color":
		v = gofakeit.Color()
	case "url":
		v = gofakeit.URL()
	case "domain":
		v = gofakeit.DomainName()
	case "ipv4":
		v = gofakeit.IPv4Address()
	case "ipv6":
		v = gofakeit.IPv6Address()
	case "useragent":
		v = gofakeit.UserAgent()
	default:
		v = gofakeit.Number(0, 999999999999999)
	}

	c.JSON(http.StatusOK, gin.H{field: v})
}

func (k *KJ) Run() {
	log.Tracef("starting HTTP API server at %s...", k.listenAddr)
	if err := k.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start the HTTP API server at %s: %s", k.listenAddr, err)
	}
	log.Tracef("serving done!?!?!")
}

func New(listen string) (*KJ, error) {

	k := KJ{
		listenAddr: listen,
	}

	router := gin.Default()

	authMiddleware, err := AuthMiddleware()
	if err != nil {
		return nil, err
	}

	crs := cors.DefaultConfig()
	crs.AddAllowHeaders("Authorization")
	crs.AllowAllOrigins = true
	router.Use(cors.New(crs))

	router.POST("/login", authMiddleware.LoginHandler)

	rnd := router.Group("/random")
	rnd.GET("/refresh-token", authMiddleware.RefreshHandler)
	//rnd.Use(authMiddleware.MiddlewareFunc())

	rnd.GET("/:field/:id", k.getRandom)

	k.server = &http.Server{
		Addr:    k.listenAddr,
		Handler: router,
	}

	return &k, nil
}
