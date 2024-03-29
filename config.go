package hcpairing

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	Config ConfigSet = NewConfigSet()
)

type ConfigSet interface {
	GetDebugMode() bool
	GetPostgresHost() string
	GetPostgresPort() string
	GetPostgresUser() string
	GetPostgresPassword() string
	GetPostgresDBName() string
	GetRedisHost() string
	GetRedisPort() string
	GetRedisPassword() string
	GetGoogleMapAPIKey() string
	GetHCPairingDNSName() string
}

type configSet struct {
	debugMode        bool
	postgresHost     string
	postgresPort     string
	postgresUser     string
	postgresPassword string
	postgresDBName   string
	redisHost        string
	redisPort        string
	redisPassword    string
	googleMapAPIKey  string
	hcpairingDNSName string
}

func NewConfigSet() ConfigSet {
	err := godotenv.Load()
	if err != nil {
		Logger.Warn(
			"error loading .env file",
		)
	}
	instance := &configSet{
		debugMode:        os.Getenv("DEBUG_MODE") == "True",
		postgresHost:     os.Getenv("POSTGRES_HOST"),
		postgresPort:     os.Getenv("POSTGRES_PORT"),
		postgresUser:     os.Getenv("POSTGRES_USER"),
		postgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		postgresDBName:   os.Getenv("POSTGRES_DB"),
		redisHost:        os.Getenv("REDIS_HOST"),
		redisPort:        os.Getenv("REDIS_PORT"),
		redisPassword:    os.Getenv("REDIS_PASSWORD"),
		googleMapAPIKey:  os.Getenv("GOOGLE_MAP_API_KEY"),
		hcpairingDNSName: os.Getenv("HCPAIRING_DNS_NAME"),
	}
	return instance
}

func (c *configSet) GetDebugMode() bool {
	return c.debugMode
}

func (c *configSet) GetPostgresHost() string {
	return c.postgresHost
}

func (c *configSet) GetPostgresPort() string {
	return c.postgresPort
}

func (c *configSet) GetPostgresUser() string {
	return c.postgresUser
}

func (c *configSet) GetPostgresPassword() string {
	return c.postgresPassword
}

func (c *configSet) GetPostgresDBName() string {
	return c.postgresDBName
}

func (c *configSet) GetGoogleMapAPIKey() string {
	return c.googleMapAPIKey
}

func (c *configSet) GetRedisHost() string {
	return c.redisHost
}

func (c *configSet) GetRedisPort() string {
	return c.redisPort
}

func (c *configSet) GetRedisPassword() string {
	return c.redisPassword
}

func (c *configSet) GetHCPairingDNSName() string {
	return c.hcpairingDNSName
}
