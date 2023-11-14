package main

import (
	"github.com/gin-gonic/gin"
	"golang-rest/album"
	"golang-rest/config"
	"strconv"
)

const HOST = "SERVER.HOST"
const PORT = "SERVER.PORT"
const COLON = ":"

func main() {
	router := gin.Default()
	album.RegisterAlbumEndpoints(router)

	config.LoadConfig()
	var serverAddress = config.GetStringConfig(HOST) + COLON + strconv.Itoa(config.GetIntConfig(PORT))

	router.Run(serverAddress)
}
