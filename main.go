package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonToHtmlTable(c *gin.Context) {
	jsonData := `
        [{
            "id": 1,
            "name": "John Doe",
            "email": "john@example.com"
        },
        {
            "id": 2,
            "name": "Jane Smith",
            "email": "jane@example.com"
        }]
    `

	var data []map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing JSON data")
		return
	}

	c.HTML(http.StatusOK, "table.html", gin.H{
		"data": data,
	})
}
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("*.html")

	r.GET("/json-to-html", JsonToHtmlTable)

	r.Run(":8000")
}
