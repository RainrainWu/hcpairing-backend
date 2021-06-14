package hcpairing

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
	GetRouter() *gin.Engine
	Start()
}

type server struct {
	router *gin.Engine
}

func NewServer() Server {
	instance := server{
		router: gin.Default(),
	}
	instance.router.GET("v1/tags", TagsGetHandler)
	instance.router.GET("v1/records", RecordsGetHandler)
	instance.router.POST("v1/records", RecordsPostHandler)
	return &instance
}

func TagsGetHandler(c *gin.Context) {
	prefix := c.DefaultQuery("prefix", "")
	c.JSON(
		http.StatusOK,
		gin.H{
			"tags": SearchTags(prefix),
		},
	)
}

type recordPayload struct {
	Zipcode string   `json:"zipcode"`
	Tags    []string `json:"tags"`
}

func RecordsGetHandler(c *gin.Context) {
	zipcode := c.DefaultQuery("zipcode", "")
	c.JSON(
		http.StatusOK,
		gin.H{
			"results": DBConn.GetRecordsByZipcode(zipcode),
		},
	)
}

func RecordsPostHandler(c *gin.Context) {
	payload := &recordPayload{}
	err := c.BindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid paylaod",
		})
		return
	}
	DBConn.AppendRecord(payload.Zipcode, payload.Tags)
	c.JSON(http.StatusOK, gin.H{
		"specialties": DirectConversion(payload.Tags, -1),
	})
}

func (s *server) GetRouter() *gin.Engine {
	return s.router
}

func (s *server) Start() {
	s.router.Run()
}
