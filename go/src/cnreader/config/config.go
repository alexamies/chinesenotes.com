/*
Package for command line tool configuration
*/
package config

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"strconv"
)

const conversionsFile = "data/corpus/html-conversion.csv"

var projectHome, dictionaryDir string
var configVars map[string]string

// A type that holds the source and destination files for HTML conversion
type HTMLConversion struct {
	SrcFile, DestFile, Template string
	GlossChinese bool
	Title string
}

func init() {
	projectHome = os.Getenv("CNREADER_HOME")
	log.Printf("config.init: projectHome: '%s'\n", projectHome)
	if projectHome == "" {
		projectHome = "../../.."
	}
	configVars = readConfig()
}

// Subdomains to avoid whne loading the dictionary, default: empty
func AvoidSubDomains() map[string]bool {
  avoidSub := make(map[string]bool)
	if val, ok := configVars["AvoidSubDomains"]; ok {
		values := strings.Split(",", val)
		for _, value := range values {
			avoidSub[value] = true
		}
	}
	return avoidSub
}

// Returns the directory where the corpus metadata is stored
func CorpusDataDir() string {
	return projectHome + "/data/corpus"
}

// Returns the directory where the raw corpus text files are read from
func CorpusDir() string {
	return projectHome + "/corpus"
}

// The name of the directory containing the dictionary files
func DictionaryDir() string {
	val, ok := configVars["DictionaryDir"]
	if ok {
		return projectHome + "/" + val
	}
	return projectHome + "/data"
}

// Gets the list of source and destination files for HTML conversion
func GetHTMLConversions() []HTMLConversion {
	log.Printf("GetHTMLConversions: projectHome: '%s'\n", projectHome)
	conversionsFile := projectHome + "/" + conversionsFile
	convFile, err := os.Open(conversionsFile)
	if err != nil {
		log.Fatal("config.GetHTMLConversions fatal error: ", err)
	}
	defer convFile.Close()
	reader := csv.NewReader(convFile)
	reader.FieldsPerRecord = -1
	reader.Comma = rune('\t')
	reader.Comment = rune('#')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	conversions := make([]HTMLConversion, 0)
	for _, row := range rawCSVdata {
		if len(row) == 3 {
			conversions = append(conversions, HTMLConversion{row[0], row[1],
				row[2], true, ""})
		} else if len(row) == 4  {
			glossChinese, err := strconv.ParseBool(row[3])
			if err != nil {
				conversions = append(conversions, HTMLConversion{row[0], row[1],
				row[2], true, ""})
			} else {
				conversions = append(conversions, HTMLConversion{row[0], row[1],
					row[2], glossChinese, ""})
			}
		} else if len(row) == 5  {
			glossChinese, err := strconv.ParseBool(row[3])
			if err != nil {
				conversions = append(conversions, HTMLConversion{row[0], row[1],
				row[2], true, row[4]})
			} else {
				conversions = append(conversions, HTMLConversion{row[0], row[1],
					row[2], glossChinese, row[4]})
			}
		}
	}
	return conversions
}

// Gets a configuration variable value
func GetVar(key string) string {
	val, ok := configVars[key]
	if !ok {
		log.Printf("config.GetVar: could not find key: '%s'\n", key)
		val = ""
	}
	return val
}

// The name of the directory containing the dictionary files
func IndexDir() string {
	return projectHome + "/index"
}

// The name of the text files with lexical units (word senses)
func LUFileNames() []string {
	fileNames := []string{DictionaryDir() + "/words.txt"}
	val, ok := configVars["LUFiles"]
	if ok {
		tokens := strings.Split(val, ",")
		fileNames = []string{}
		for _, token := range tokens {
			fileNames = append(fileNames, DictionaryDir() + "/" + token)
		}
	}
	return fileNames
}

// Gets the public home directory, relative to the cnreader command line tool
func ProjectHome() string {
	return projectHome
}

// Reads the configuration file with project variables
func readConfig() map[string]string {
	vars := make(map[string]string)
	fileName := projectHome + "/config.yaml"
	configFile, err := os.Open(fileName)
	if err != nil {
		projectHome = "../../../.."
		log.Printf("config.readConfig: setting projectHome to: '%s'\n",
			projectHome)
		fileName = projectHome + "/config.yaml"
		configFile, err = os.Open(fileName)
		if err != nil {
			log.Fatal("config.init fatal error: config.yaml not found")
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
			log.Fatal("config.readConfig: error reading config file", err)
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

// Sets the public home directory, relative to the cnreader command line tool
func SetProjectHome(home string) {
	projectHome = home
}

// Gets the name of the directory where the HTML templates are stored
func TemplateDir() string {
	templateHome := os.Getenv("TEMPLATE_HOME")
	//log.Printf("config.TemplateDir: templateHome: '%s'\n", templateHome)
	if templateHome == "" {
		templateHome = "html/templates"
	}
	return ProjectHome() + "/" + templateHome
}

// Gets the Web directory, as used for serving HTML files
func WebDir() string {
	webDir := os.Getenv("WEB_DIR")
	//log.Printf("config.TemplateDir: webDir: '%s'\n", webDir)
	if webDir == "" {
		webDir = "web"
	}
	return projectHome + "/" + webDir
}

// Gets the Web directory, as used for serving HTML files
func VocabFormat() string {
	val, ok := configVars["VocabFormat"]
	if ok {
		return val
	}
	return "<a title=\"%s | %s\" class=\"%s\" href=\"/words/%d.html\">%s</a>"
}