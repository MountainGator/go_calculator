package main

import (
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func calc_mean(numbers []string) (float64, error) {
	sum := 0.00
	length := float64(len(numbers))
	for _, n := range numbers {
		i, err := strconv.ParseFloat(n, 64)
		if err == nil {
			sum += i
		} else {
			return 0.00, err
		}
	}
	var mean float64 = sum / length
	return mean, nil
}

func calc_median(numbers []string) (float64, error) {
	var median float64
	length := float64(len(numbers))
	var sorted = []int{}

	for _, n := range numbers {
		i, err := strconv.Atoi(n)
		if err != nil {
			return 0, err
		}
		sorted = append(sorted, i)
	}

	sort.Ints(sorted)
	if math.Mod(length, 2) == 0.000000 {

		x := length / 2
		x -= 1
		median = float64(sorted[int(x)])
		return median, nil
	} else {
		x := math.Ceil(length / 2)
		x -= 1
		y := x - 1
		median = (float64(sorted[int(x)]) + float64(sorted[int(y)])) / 2
		return median, nil
	}

}

func main() {
	r := gin.Default()

	r.GET("/mean/:nums", func(c *gin.Context) {
		numbers := strings.Split(c.Param("nums"), ",")

		mean, err := calc_mean(numbers)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "not a number"})
		} else {
			c.IndentedJSON(http.StatusAccepted, gin.H{"operation": "mean", "value": mean})
		}

	})

	r.GET("/median/:nums", func(c *gin.Context) {
		numbers := strings.Split(c.Param("nums"), ",")

		median, err := calc_median(numbers)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "not a number"})
		} else {
			c.IndentedJSON(http.StatusAccepted, gin.H{"operation": "median", "value": median})
		}

	})

	r.Run()
}
