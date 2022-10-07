package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Todo: create an interface here instead of a concrete type. This is tightly coupled with the invoice-service.env format.
// Datatype to store environment variables.
type Config struct {
	ApiVersion  string // API version.
	ServiceIp   string // Service IP address.
	ServicePort string // Service port.

	RepoHost     string // DB IP address.
	RepoPort     string // DB port.
	RepoTimeout  int    // DB request timeout in seconds.
	RepoDriver   string // DB driver name - eg: postgres.
	RepoName     string // Database name.
	RepoUsername string // DB username.
	RepoPassword string // DB password.
	SslMode      string // SSL status can be enabled or disabled.

	TestDb    string // DB Env flag for choosing DB server.
	TestDbCon string // DB connection for test env.
}

// Returns Config object
func NewConfig() *Config {
	return &Config{}
}

// 2 default seconds timeout
const DEFAULT_TIMEOUT int = 2

// Reads environment variables from a .env file
func GetEnvVariables(fileName string) (*Config, error) {

	err := godotenv.Load(fileName)
	if err != nil {
		fmt.Println("Error loading occurred. Err:", err)
		return nil, err
	}
	env := NewConfig()

	env.ApiVersion = os.Getenv("API_VERSION")
	env.ServiceIp = os.Getenv("SERVICE_IP")
	env.ServicePort = os.Getenv("SERVICE_PORT")

	env.RepoHost = os.Getenv("DB_HOST")
	env.RepoPort = os.Getenv("DB_PORT")
	env.RepoUsername = os.Getenv("DB_USER")
	env.RepoPassword = os.Getenv("DB_PASS")
	env.RepoName = os.Getenv("DB_NAME")
	env.SslMode = os.Getenv("SSL")
	env.RepoDriver = os.Getenv("DB_DRIVER")

	env.RepoTimeout, err = strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		fmt.Printf("Error converting string %s to int. Setting timeout to %d", os.Getenv("DB_TIMEOUT"), DEFAULT_TIMEOUT)
		env.RepoTimeout = DEFAULT_TIMEOUT
	}

	env.TestDb = os.Getenv("TEST_DB")
	env.TestDbCon = os.Getenv("TEST_DB_CON")

	return env, nil
}

// Returns obj with default config for all env variables.
func NewDefaultConfig() *Config {
	return &Config{
		ApiVersion:   "v1",
		ServiceIp:    "127.0.0.1",
		ServicePort:  "3000",
		RepoHost:     "db.bit.io",
		RepoPort:     "5432",
		RepoTimeout:  DEFAULT_TIMEOUT,
		RepoDriver:   "postgres",
		RepoName:     "GajanSoorian/invoice_db",
		RepoUsername: "GajanSoorian",
		RepoPassword: "v2_3udqt_WawXZtQTZ3rqyvCuKJ7SnZ9",
		SslMode:      "require",
		TestDb:       "enabled",
		TestDbCon:    "",
	}
}
