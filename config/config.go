package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func readModuleConfig(cfg interface{}, path string, module string) error {
	environ := os.Getenv("ENV")
	if environ == "" {
		environ = "development"
	}

	getFormatFile := filePath(path)

	switch getFormatFile {
	case ".json":
		fname := path + "/" + module + "." + environ + ".json"
		jsonFile, err := ioutil.ReadFile(fname)
		if err != nil {
			return err
		}
		return json.Unmarshal(jsonFile, cfg)
	default:
		fname := path + "/" + module + "." + environ + ".yaml"
		yamlFile, err := ioutil.ReadFile(fname)
		if err != nil {
			return err
		}
		return yaml.Unmarshal(yamlFile, cfg)
	}

}

func filePath(root string) string {
	var file string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		file = filepath.Ext(info.Name())
		return nil
	})
	return file
}

var GlobalConfig *Config

func ReadModuleConfig(path string) interface{} {
	if GlobalConfig == nil {
		err := readModuleConfig(&GlobalConfig, path, "config")
		if err != nil {
			log.Fatalln("failed to read config for ", err)
		}
	}
	return GlobalConfig
}
