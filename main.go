package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {

	directorys := []string{"Imagens", "Videos", "Compactados", "Instaladores", "Documentos"}

	for _, directory := range directorys {
		CreateDirectory(directory)
	}

	file_extension := map[string]string{
		".png":  directorys[0],
		".jpeg": directorys[0],
		".jpg":  directorys[0],
		".mp4":  directorys[1],
		".wmv":  directorys[1],
		".avi":  directorys[1],
		".rar":  directorys[2],
		".exe":  directorys[3],
		".xlsx": directorys[4],
		".pdf":  directorys[4],
	}

	filepath.WalkDir(".", func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			log.Panic(err)
		}

		if file.IsDir() && file.Name() != "." {
			return filepath.SkipDir
		}

		if !file.IsDir() && path != "Organizer.exe" {

			ex := filepath.Ext(path)

			if file_extension[ex] != "" {
				MoveFile(path, "./"+file_extension[ex])
			} else {
				fmt.Printf("The followig file has no directory to be moved: %s\n", path)
			}
		}

		return nil
	})
}

func MoveFile(sourcePath, destPath string) error {

	inputFile, err := os.ReadFile(sourcePath)
	if err != nil {
		log.Panic(err)
	}

	err = os.Chdir(destPath)
	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile(sourcePath, inputFile, 0666)
	if err != nil {
		log.Panic(err)
	}

	os.Chdir("..")
	err = os.Remove(sourcePath)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%s has been moved to %s\n", sourcePath, destPath)

	return nil
}

func CreateDirectory(directory string) {

	err := os.Mkdir(directory, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	} else {
		fmt.Printf("Directory has been created: %s\n", directory)
	}

}
