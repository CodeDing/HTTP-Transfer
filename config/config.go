package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var path string
var g_config *File

func init() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed to get current directory!")
		os.Exit(-1)
	}
	mode := os.Getenv("GGXT_MODE")
	if mode == "" {
		mode = "development"
	}
	path = filepath.Join(dir, fmt.Sprintf("data/config_%s.ini", mode))
}

func Init() {
	var err error
	g_config, err = NewFile(path)
	if err != nil {
		panic(err)
	}
}

func stringDefault(section, key, def string) string {
	if g_config == nil {
		Init()
	}
	if g_config.HasSection(section) {
		return g_config.GetString(section,key, def)
	}
	return def
}

func String(section, key string) string {
	return  stringDefault(section,key, "")
}

func StringDefault(section, key, def string) string {
	return stringDefault(section,key, def)
}

func boolDefault(section, key string, def bool) bool {
	if g_config == nil {
		Init()
	}
	if g_config.HasSection(section) {
		return g_config.GetBool(section,key,def)
	}
	return false
}

func Bool(section, key string) bool {
	return boolDefault(section,key,false)
}

func BoolDefault(section, key string, def bool) bool {
	return boolDefault(section, key,def)
}

func intDefault(section, key string, def int) int {
	if g_config == nil {
		Init()
	}
	if g_config.HasSection(section) {
		return g_config.GetInt(section, key, def)
	}
	return 0
}

func Int(section, key string) int {
	return intDefault(section, key, 0)
}

func IntDefault(section, key string, def int) int {
	return intDefault(section,key, def)
}
