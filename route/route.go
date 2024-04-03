package route

import (
	"log"
	"os"
	"strings"
)

func Init() {
	hiepWedding()
	khanhWedding()

}

func getImageFileNamesInFolder(folderPath string) []string {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string
	for _, file := range files {
		name := file.Name()
		if strings.Contains(name, ".jpg") || strings.Contains(name, ".png") || strings.Contains(name, ".jpeg") {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}
