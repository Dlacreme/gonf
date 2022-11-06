package gonf

import (
	"fmt"
	"path/filepath"
)

// If not set, default folder is "config"
var folder_pathname = "gonfig"

// If not set, default root config file is "config.gonf"
var root_config = "base.gonf"

// If not set, default file loaded is "dev.gonf"
var current_config = "dev.gonf"

func SetFolder(pathname string) {
	folder_pathname = pathname
}

func SetRoot(filename string) {
	root_config = filename
}

func Use(envName string) {
	current_config = fmt.Sprintf("%s.gonf", envName)
}

func buildFilenameForEnv() string {
	return filepath.Join(folder_pathname, current_config)
}

func buildRootFilename() string {
	return filepath.Join(folder_pathname, root_config)
}
