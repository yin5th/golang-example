package blog

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.Run(":8088")
}
