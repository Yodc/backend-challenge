package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type BeefSummary struct {
	Count int `json:"count"`
}

func countBeef(text string) map[interface{}]int {
	meats := strings.Fields(strings.ToLower(text))
	obj := make(map[interface{}]int)
	for _, meat := range meats {
		obj[meat]++
	}

	return obj
}

func replaceData(text string) string {
	new_text := strings.ReplaceAll(text, ",", "")
	new_text = strings.ReplaceAll(new_text, ".", "")

	return new_text
}

func beefSummaryHandler(c *gin.Context) {
	response, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch data"})
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read data"})
		return
	}

	beefCount := countBeef(replaceData(string(body)))

	stringMap := make(map[string]int)
	for k, v := range beefCount {
		strKey := fmt.Sprintf("%v", k)
		stringMap[strKey] = v
	}

	responseData := map[string]interface{}{
		"beef": stringMap,
	}

	c.JSON(http.StatusOK, responseData)
}

func main() {
	router := gin.Default()
	router.GET("/beef/summary", beefSummaryHandler)
	router.Run(":8080")
}
