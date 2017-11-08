package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type XmlFileName struct {
	Name string
}

var Filename []string //Необходим для получения названия имени файла

func Unzip(archive, target string) (error error, Filename []string) {

	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err, Filename
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err, Filename
	}
		for _, file := range reader.File {
			path := filepath.Join(target, file.Name)
			if file.FileInfo().IsDir() {
				os.MkdirAll(path, file.Mode())
				continue
			}

			fileReader, err := file.Open()
			if err != nil {

				if fileReader != nil {
					fileReader.Close()
				}

				return err, Filename
			}

			targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				fileReader.Close()

				if targetFile != nil {
					targetFile.Close()
				}

				return err, Filename
			}

			if _, err := io.Copy(targetFile, fileReader); err != nil {
				fileReader.Close()
				targetFile.Close()

				return err, Filename
			}

			Filename = append(Filename, file.FileInfo().Name())

			fileReader.Close()
			targetFile.Close()
		}


	return nil, Filename
}



