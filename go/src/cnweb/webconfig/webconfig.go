/*
Package for web app configuration
*/
package webconfig

import (
	"bufio"
	"cnweb/applog"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	configVars map[string]string
	domain *string
)

func init() {
	localhost := "localhost"
	domain = &localhost
	site_domain := os.Getenv("SITEDOMAIN")
	if site_domain != "" {
		domain = &site_domain
	}
	configVars = readConfig()
}

func DBConfig() string {
	dbhost := "mariadb"
	host := os.Getenv("DBHOST")
	if host != "" {
		dbhost = host
	}
	dbport := "3306"
	port := os.Getenv("DBPORT")
	if port != "" {
		dbport = port
	}
	dbuser := "app_user"
	user := os.Getenv("DBUSER")
	if user != "" {
		dbuser = user
	}
	dbpass := os.Getenv("DBPASSWORD")
	dbname := "corpus_index"
	d := os.Getenv("DATABASE")
	if d != "" {
		dbname = d
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost,
		dbport, dbname)
}

// Gets all configuration variables
func GetAll() map[string]string {
	return configVars
}

// Get environment variable for sending email from
func GetFromEmail() string {
	fromEmail := os.Getenv("FROM_EMAIL")
	if fromEmail == "" {
		fromEmail = GetVar("FromEmail")
	}
	return fromEmail
}

// Get the domain name of the site
func GetPasswordResetURL() string {
	passwordResetURL := os.Getenv("PASSWORD_RESET_URL")
	if passwordResetURL == "" {
		passwordResetURL = GetVar("PasswordResetURL")
	}
	return passwordResetURL
}

// Get environment variable for serving port
func GetPort() int {
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}
	port, err := strconv.Atoi(portString)
	if err != nil {
		port = 8080
	}
	return port
}

// Get the domain name of the site
func GetSiteDomain() string {
	return *domain
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
	cnwebHome := os.Getenv("CNWEB_HOME")
	if cnwebHome == "" {
		applog.Error("config.readConfig: CNWEB_HOME is not defined")
		cnwebHome = "."
	}
	fileName := fmt.Sprintf("%s/webconfig.yaml", cnwebHome)
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
