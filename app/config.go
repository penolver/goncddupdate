// load JSON config

package app

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "os"
)

type Configuration struct {
  CurrentIP   string `json:"CurrentIP"`
  Namecheap_Host string `json:"Namecheap_Host"`
  Namecheap_Domain string `json:"Namecheap_Domain"`
  Namecheap_Pass string `json:"Namecheap_Pass"`
}

// Load JSON Config file
func LoadConfig(path string) Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return config
} // LoadConfig

// Save JSON config file
func SaveConfig(v interface{}, path string) {
    fo, err := os.Create(path)
    if err != nil {
      log.Fatal("Config save file write error: ",err)
    }
    defer fo.Close()
    e := json.NewEncoder(fo)
    if err := e.Encode(v); err != nil {
      log.Fatal("Config save file encode error: ",err)
    }
} // SaveConfig
