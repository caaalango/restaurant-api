package webctl

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/calango-productions/api/internal/adapters"
	"github.com/gin-gonic/gin"
)

type WebController struct {
	adapters *adapters.Adapters
}

func New(adapters *adapters.Adapters) WebController {
	return WebController{adapters: adapters}
}

func (w WebController) SetUpRoutes(r *gin.Engine) {
	r.GET("/menu", w.ShowMenu)
}

func ratingToStars(rating float64) string {
	fullStars := int(rating)
	if fullStars > 5 {
		fullStars = 5
	}
	stars := strings.Repeat("★", fullStars) + strings.Repeat("☆", 5-fullStars)
	return stars
}

func (w WebController) ShowMenu(c *gin.Context) {
	categories := []gin.H{
		{
			"Name": "Entradas",
			"Items": []gin.H{
				{
					"Name":               "Bruschetta",
					"Price":              12.50,
					"Description":        "Pão italiano torrado com tomate e manjericão",
					"ImageURL":           "https://storage.googleapis.com/restaurant-app/bruschetta.webp",
					"CommentsCount":      5,
					"RatingSum":          21,
					"RatingCount":        5,
					"RatingDistribution": []int{0, 1, 1, 2, 1}, // Exemplo
					"Comments": []map[string]interface{}{
						{"User": "João", "Text": "Delicioso!", "Rating": 5},
						{"User": "Maria", "Text": "Muito bom!", "Rating": 4},
						{"User": "Carlos", "Text": "Crocante e saboroso", "Rating": 4},
						{"User": "Ana", "Text": "Adorei", "Rating": 4},
						{"User": "Pedro", "Text": "Recomendo!", "Rating": 2},
					},
				},
				{
					"Name":               "Coxinha",
					"Price":              5.00,
					"Description":        "Coxinha de frango tradicional",
					"ImageURL":           "https://storage.googleapis.com/restaurant-app/coxinha.webp",
					"CommentsCount":      2,
					"RatingSum":          6,
					"RatingCount":        2,
					"RatingDistribution": []int{0, 0, 1, 0, 1},
					"Comments": []map[string]interface{}{
						{"User": "Paula", "Text": "Boa, mas podia ser maior", "Rating": 3},
						{"User": "Rafael", "Text": "Frango bem temperado", "Rating": 5},
					},
				},
			},
		},
	}

	for _, cat := range categories {
		items := cat["Items"].([]gin.H)
		for _, item := range items {
			ratingSum := float64(item["RatingSum"].(int))
			ratingCount := float64(item["RatingCount"].(int))
			var rating float64
			if ratingCount > 0 {
				rating = ratingSum / ratingCount
			} else {
				rating = 0
			}
			item["Stars"] = ratingToStars(rating)
			item["CommentsLabel"] = fmt.Sprintf("%d comentários", item["CommentsCount"])

			commentsJSON, _ := json.Marshal(item["Comments"])
			item["CommentsJSON"] = template.JS(string(commentsJSON))

			ratingDistJSON, _ := json.Marshal(item["RatingDistribution"])
			item["RatingDistJSON"] = template.JS(string(ratingDistJSON))
		}
	}

	data := gin.H{
		"title":      "Cardápio",
		"categories": categories,
	}

	c.HTML(http.StatusOK, "menu.html", data)
}
