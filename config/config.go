package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const config_file_name = "config.json"

var Config configFile

func init() {
	file, err := os.Open(config_file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&Config)
	if err != nil {
		fmt.Println(err)
		return
	}

}

type configFile struct {
	Dest    destination `json:"Destination"`
	Runtime runtime     `json:"Runtime"`
}

type destination struct {
	IP   string `json:"ip"`
	Port uint   `json:"port"`
}

type runtime struct {
	PacketsToSend int `json:"packetsToSend"`
}
