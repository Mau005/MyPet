package configuration

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"gopkg.in/yaml.v3"
)

var Store *sessions.CookieStore
var Security string
var Config *Configuration

type DataBase struct {
	Host       string `yaml:"Host"`
	Port       uint   `yaml:"Port"`
	User       string `yaml:"User"`
	Password   string `yaml:"Password"`
	NameDB     string `yaml:"NameDB"`
	SqlitePath string `yaml:"SqlitePath"`
	Engine     string `yaml:"Engine"`
}

type Server struct {
	Ip             string `yaml:"Ip"`
	Port           uint   `yaml:"Port"`
	Debug          bool   `yaml:"Debug"`
	LengthSecurity uint   `yaml:"LengthSecurity"`
}

type Configuration struct {
	DataBase DataBase `yaml:"DataBase"`
	Server   Server   `yaml:"Server"`
}

func LoadConfiguration(fileName string) error {
	config := Configuration{}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}

	Config = &config //asigno la configuracion
	Security = "secret"
	if !config.Server.Debug {
		Security, _ = GenerateSecretKey(int(Config.Server.LengthSecurity))
	}

	if Config.Server.Debug {
		log.Println("Mode Debug")
	} else {
		log.Println("Mode Normal")
	}

	Store = sessions.NewCookieStore([]byte(Security))
	return nil

}

func GenerateSecretKey(lenSecurity int) (string, error) {
	key := make([]byte, lenSecurity)
	_, err := rand.Read(key)
	if err != nil {
		return "secret", err
	}

	encodedKey := base64.StdEncoding.EncodeToString(key)

	return encodedKey, nil
}
