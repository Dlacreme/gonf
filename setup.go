package gonf

import (
	"fmt"
	"path/filepath"
)

// If not set, default folder is "config"
var folder_pathname = "gonfig"

// If not set, default base config file is "base.gonf"
var base_config = "base.gonf"

// If not set, default file loaded is "dev.gonf"
var current_config = "dev.gonf"

func SetFolder(pathname string) {
	folder_pathname = pathname
}

func SetBase(filename string) {
	base_config = filename
}

func Use(envName string) {
	current_config = fmt.Sprintf("%s.gonf", envName)
}

func buildEnvFilename() string {
	return filepath.Join(folder_pathname, current_config)
}

func buildBaseFilename() string {
	return filepath.Join(folder_pathname, base_config)
}
