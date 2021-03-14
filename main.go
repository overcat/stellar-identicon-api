package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	identicon "github.com/overcat/stellar-identicon-go"
	"image/png"
	"strconv"
)

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"_link": gin.H{
			"avatar": gin.H{
				"href": "/avatar/{account_id}/{?width,height}",
			},
		},
		"stellar_identicon_api_source_code": "https://github.com/overcat/stellar-identicon-api",
		"stellar_identicon_go_source_code":  "https://github.com/overcat/stellar-identicon-go",
		"sep_0033_file":                     "https://github.com/stellar/stellar-protocol/blob/master/ecosystem/sep-0033.md",
	})
}

func NoRouterHandler(c *gin.Context) {
	c.JSON(404, gin.H{
		"code":    "PAGE_NOT_FOUND",
		"message": "Page not found",
	})
}

func AvatarHandler(c *gin.Context) {
	accountId := c.Param("account_id")
	widthQuery := c.DefaultQuery("width", "")
	heightQuery := c.DefaultQuery("height", "")

	width := identicon.Width
	if widthQuery != "" {
		width_, err := strconv.Atoi(widthQuery)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    "INVALID_WIDTH",
				"message": "invalid width",
			})
			return
		}
		width = width_
	}

	height := identicon.Height
	if heightQuery != "" {
		height_, err := strconv.Atoi(heightQuery)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    "INVALID_HEIGHT",
				"message": "invalid height",
			})
			return
		}
		height = height_
	}
	img, err2 := identicon.Generate(accountId, width, height)
	if err2 != nil {
		c.JSON(400, gin.H{
			"code":    "INVALID_STELLAR_ID",
			"message": "invalid stellar id",
		})
		return
	}
	c.Header("Content-Type", "image/png")
	c.Header("Cache-Control", "public, max-age=86400")
	png.Encode(c.Writer, img)
}

func main() {
	r := gin.Default()
	r.GET("/", IndexHandler)
	r.GET("/avatar/:account_id", AvatarHandler)
	r.NoRoute(NoRouterHandler)
	r.Use(cors.Default())
	r.Run()
}
