/*
Package for web app configuration
*/
package webconfig

import (
	"bufio"
	"cnweb/applog"
	"io"
	"os"
	"strings"
)

var configVars map[string]string

func init() {
	configVars = readConfig()
}

// Gets all configuration variables
func GetAll() map[string]string {
	return configVars
}

// Gets a configuration variable value
func GetVar(key string) string {
	val, ok := configVars[key]
	if !ok {
		applog.Error("config.GetVar: could not find key: ", key)
		val = ""
	}
	return val
}

// Reads the configuration file with project variables
func readConfig() map[string]string {
	vars := make(map[string]string)
	fileName := "webconfig.yaml"
	configFile, err := os.Open(fileName)
	if err != nil {
		if err != nil {
			applog.Fatal("config.init fatal error: ", err)
		}
	}
	defer configFile.Close()
	reader := bufio.NewReader(configFile)
	eof := false
	for !eof {
		var line string
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			applog.Fatal("config.readConfig: error reading config file", err)
		}
		// Ignore comments
		if strings.HasPrefix(line, "#") {
			continue
		}
		i := strings.Index(line, ":")
		if i > 0 {
			varName := line[:i]
			val := strings.TrimSpace(line[i+1:])
			vars[varName] = val
		}
	}
	return vars
}
