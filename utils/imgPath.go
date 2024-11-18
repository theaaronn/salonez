package utils

import (
	"fmt"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetFirstImgPath(hallId int) string {
	imgPath := fmt.Sprintf("/hallImgs/%d-0.jpg", hallId)
	if fileExists(imgPath) {
		return imgPath
	}
	return "/assets/imgPlaceholder.jpg"
}

func GetAllImgPaths(hallId int) []string {
	imgPaths := make([]string, 5)
	for i := 1; i <= 5; i++ {
		path := fmt.Sprintf("/hallImgs/%d-%d", hallId, i)
		if fileExists(path) {
			imgPaths[i-1] = path
		} else {
			return imgPaths
		}
	}
	return imgPaths
}
