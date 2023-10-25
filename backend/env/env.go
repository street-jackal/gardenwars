package env

import "os"

var (
	dbURL  string
	dbName string

	userJWTSecret string
)

func Init() {
	if dbURL = os.Getenv("DB_URL"); dbURL == "" {
		dbURL = "mongodb://localhost:27017"
	}

	if dbName = os.Getenv("DB_NAME"); dbName == "" {
		dbName = "gardenwars"
	}

	if userJWTSecret = os.Getenv("USER_JWT_SECRET"); userJWTSecret == "" {
		userJWTSecret = "secret"
	}

}

func GetDbURL() string {
	return dbURL
}

func GetDbName() string {
	return dbName
}

func GetUserJWTSecret() string {
	return userJWTSecret
}
