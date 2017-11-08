package utils

import (
	"os"
	"fmt"
	"encoding/json"
	"path/filepath"
	"path"
)
type Config struct {
	Database struct {
			 Host string `json:"host"`
			 Username string `json:"username"`
			 Password string `json:"password"`
			 Db string `json:"db"`
			 Port string `json:"port"`
		 } `json:"database"`
	Upload_dir struct {
			 Patch string `json:"patch"`
	} `json:"upload_dir"`
}

func LoadConfiguration(file string) Config {
	var config Config

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))


	configFile, err := os.Open(path.Join(dir, "config.json"))
	if err != nil { fmt.Println(err.Error()) }
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}