package api

import (
	"firebase.google.com/go/auth"
	"github.com/aiiro/youtube-manager-go/middlewares"
	"github.com/aiiro/youtube-manager-go/models"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type ToggleFavoriteVideoResponse struct {
	VideoId    string `json:"video_id"`
	IsFavorite bool   `json:"is_favorite"`
}

func ToggleFavoriteVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		videoId := c.Param("id")
		token := c.Get("auth").(*auth.Token)
		user := models.User{}
		if dbs.DB.Table("users").
			Where(models.User{UID: token.UID}).First(&user).RecordNotFound() {

			user = models.User{UID: token.UID}
			dbs.DB.Create(&user)
		}

		favorite := models.Favorite{}
		isFavorite := false
		if dbs.DB.Table("favorites").
			Where(models.Favorite{UserId: user.ID, VideoId: videoId}).First(&favorite).RecordNotFound() {

			favorite = models.Favorite{UserId: user.ID, VideoId: videoId}
			dbs.DB.Create(&favorite)
			isFavorite = true
		} else {
			dbs.DB.Delete(&favorite)
		}

		res := ToggleFavoriteVideoResponse{
			VideoId:    videoId,
			IsFavorite: isFavorite,
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
