// Provide a quick & easy way to load & use your configuration
package gonf

import (
	"fmt"
	"path/filepath"
)

// If not set, default folder is "config"
var folder_pathname = "config"

// If not set, default root config file is "config.gonf"
var root_config = "config.gonf"

// If not set, default file loaded is "dev.gonf"
var current_config = "dev.gonf"

// Cached loaded key/value
var cache map[string]string = make(map[string]string)

func SetFolder(pathname string) {
	folder_pathname = pathname
}

func Use(envName string) {
	current_config = fmt.Sprintf("%s.gonf", envName)
}

func LoadMany(keys []string) {
	for _, k := range keys {
		Get(k)
	}
}

func Load(key string) {
	Get(key)
}

func Get(key string) string {
	// load from cache if available
	if v, ok := cache[key]; ok {
		return v
	}
	// load from env config file if available
	tree, err := fileToTree(buildFilenameForEnv())
	if err != nil {
		fmt.Printf("ERROR ! > %s\n", err)
		return ""
	}
	printTree(tree)
	// if rawValue == "" {
	// // backup on root config
	// rawValue = fileToTree(buildRootFilename(), key)
	// }
	// parse & consume value
	//	value, err := interpretValue(rawValue)
	//if err != nil {
	//// the value is invalid
	//return ""
	//}
	//// all good. Save in cache and return
	//pushCache(key, value)
	return ""
}

func interpretValue(rawValue string) (string, error) {
	return rawValue, nil
}

func pushCache(key string, value string) {
	cache[key] = value
}

func buildFilenameForEnv() string {
	return filepath.Join(folder_pathname, current_config)
}

func buildRootFilename() string {
	return filepath.Join(folder_pathname, root_config)
}
