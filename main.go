package main

import (
	"os"
	"fmt"
	"sync"
	"net/http"
	"mymodule/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadDotEnvVariables()
}

var store = make(map[string]string)  //For key-value datastore
var queue = make(map[string][]int32) //For queue implementation
var lock = sync.RWMutex{}            //For concurrent operations


//Already defined structure for inputs for validation
type setFormat struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type qpushFormat struct {
	Key   string  `json:"key" binding:"required"`
	Value []int32 `json:"value" binding:"required"`
}

func SET(c *gin.Context) {
	//Checking if the values passed to set follow the format or not.
	//By referencing to the already created setFormat
	var req setFormat
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	lock.Lock()
	store[req.Key] = req.Value
	lock.Unlock()

	c.Status(200)
}

func GET(c *gin.Context) {

	//Geting key passed in the parameter using 'Param()'
	key := c.Param("key")

	lock.RLock()
	//Two values are returned 1st is value and 2nd is bool whether present or not.
	value, ok := store[key]
	lock.RUnlock()

	//Checking if the key is present or not.
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "key not found"})
		return
	}

	//If present printing the value in desired format.
	c.JSON(http.StatusOK, gin.H{"value": value})
}

func QPUSH(c *gin.Context) {

	var req qpushFormat
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	key := req.Key
	values := req.Value

	lock.Lock()
	defer lock.Unlock()

	// _, ok := queue[key]

	// if ok {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "queue already present"})
	// 	return
	// }

	// if !ok {
	// 	q = []int32{}
	// }

	// for _, item := range values {
	// 	q = append(q, item)
	// }

	queue[key] = append(queue[key], values...)

	// queue["Key"] = q

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func QPOP(c *gin.Context) {
	list := c.Param("list")

	lock.Lock()
	defer lock.Unlock()

	q,ok := queue[list]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "queue not found"})
		return
	}

	if len(q) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "queue is empty"})
		return
	}

	value := q[len(q)-1]
	queue[list] = q[:len(q)-1]

	c.JSON(http.StatusOK, gin.H{"value": value})

}

func main() {

	r := gin.Default()

	r.POST("/set", SET)
	r.GET("/get/:key", GET)
	r.POST("/qpush", QPUSH)
	r.GET("/qpop/:list", QPOP)

	Port := os.Getenv("PORT")

	fmt.Println("\nSupported Endpoints are: \nPOST\t\t/set\nGET\t\t/get/:key\nPOST\t\t/qpush\nGET\t\t/qpop/:list")
	fmt.Println()
	fmt.Println("Listening and serving HTTP on port:", Port,".....")

	err := r.Run("localhost:" + Port)

	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}


}
