package utils

import (
    "encoding/json"
  	"fmt"
  	"os"
)

type Config struct {

    AvKey string `json:"avKey"`
    WtdKey string `json:"wtdKey"`
}

func LoadConfiguration(file string) Config {
    var config Config
    configFile, err := os.Open(file)
   
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config
}