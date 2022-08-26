package main

//importing packages
import(
    "net/http"
    "github.com/gin-gonic/gin"
)
//it represents data of fruits
type fruit struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Price float32  `json:"price"`
    Quantity  string `json:"quantity"`
}

//record of fruits
var fruits = []fruit {
    {ID:"1", Name:"Apple", Price: 10, Quantity:"1 dozens"},
    {ID:"2", Name:"Banana", Price: 10, Quantity:"2 dozens"},
    {ID:"3", Name:"Mango", Price: 11.5, Quantity:"3 dozens"},
    {ID:"4", Name:"Guava", Price: 6.5 , Quantity:"4 dozens"},
}

func main() {
    router := gin.Default()
    router.GET("/fruits", getFruits)
    router.GET("/fruits/:id", getFruitByID)
    router.POST("/fruits", postFruits)
    router.DELETE("/fruits/:id", deleteFruitsByID)
    router.PUT("/fruits/:id", putFruitsByID)
    router.Run("localhost:8080")
}

//getFruits responds with the list of all fruits as JSON
func getFruits(c *gin.Context){
    c.IndentedJSON(http.StatusOK, fruits)
}

//adds a fruit from JSON recieved in the request body.
// postFruits adds an fruit from JSON received in the request body.
func postFruits(c *gin.Context) {
    var newFruit fruit

    // Call BindJSON to bind the received JSON to
    // newFruit.
    if err := c.BindJSON(&newFruit); err != nil {
        return
    }

    // Add the new fruit to the slice.
    fruits = append(fruits, newFruit)
    c.IndentedJSON(http.StatusCreated, newFruit)
}
//getFruitByID gets the fruit whose ID value matches the iD
func getFruitByID(c *gin.Context)  {
    id := c.Param("id")
    //looping though list of fruits
    for _, a := range fruits {
        if a.ID == id{
            c.IndentedJSON(http.StatusOK,a)
            return
        }
    } 
    c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Fruit not found"})   
}

func deleteFruitsByID(c *gin.Context)  {
    id := c.Param("id")
    for index, a := range fruits {
        if a.ID==id{
            fruits = append(fruits[:index], fruits[index+1:]...)
            break
        }
    }
    c.IndentedJSON(http.StatusOK,fruits)
}

func putFruitsByID(c *gin.Context)  {
    id := c.Param("id")
    for index, v := range fruits {
        if v.ID == id{
            fruits = append(fruits[:index], fruits[index+1:]...)
            var addFruit fruit
             if err := c.BindJSON(&addFruit); err != nil {
                return
            }
            fruits=append(fruits, addFruit)
            c.IndentedJSON(http.StatusCreated, fruits)
            return
        }
    }
}


