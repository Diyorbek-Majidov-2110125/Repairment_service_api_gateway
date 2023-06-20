package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	ServiceName string
	ServiceHost string
	ServicePort string

	Environment string // debug, test, release
	Version     string

	UserServiceHost string
	UserServicePort string

	OrderServiceHost string
	OrderServicePort string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	PostgresMaxConnections int32

	RmqUsername string
	RmqPassword string
	RmqUriHOST  string
	RmqUriPORT  int
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "order_user_api_gateway"))
	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.ServicePort = cast.ToString(getOrReturnDefaultValue("SERVICE_PORT", ":9090"))

	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.UserServiceHost = cast.ToString(getOrReturnDefaultValue("USER_SERVICE_HOST", "localhost"))
	config.UserServicePort = cast.ToString(getOrReturnDefaultValue("USER_SERVICE_PORT", ":8080"))

	config.OrderServiceHost = cast.ToString(getOrReturnDefaultValue("ORDER_SERVICE_HOST", "localhost"))
	config.OrderServicePort = cast.ToString(getOrReturnDefaultValue("ORDER_SERVICE_PORT", ":8081"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "diyorbek"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "qwer"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", config.ServiceName))

	config.RmqUsername = cast.ToString(getOrReturnDefaultValue("RMQ_USERNAME", "rmq_user"))
	config.RmqPassword = cast.ToString(getOrReturnDefaultValue("RMQ_PASSWORD", "hello"))
	config.RmqUriHOST = cast.ToString(getOrReturnDefaultValue("RMQ_URI_HOST", "82.148.6.63"))
	config.RmqUriPORT = cast.ToInt(getOrReturnDefaultValue("RMQ_URI_PORT", 5672))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
