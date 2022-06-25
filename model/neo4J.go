package model

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jConfiguration struct {
	Url      string
	Username string
	Password string
	Database string
}

func (nc *Neo4jConfiguration) NewDriver() (neo4j.Driver, error) {
	return neo4j.NewDriver(nc.Url, neo4j.BasicAuth(nc.Username, nc.Password, ""))
}

func ParseConfiguration() *Neo4jConfiguration {
	database := LookupEnvOrGetDefault("NEO4J_DATABASE", "kausjan")
	if !strings.HasPrefix(LookupEnvOrGetDefault("NEO4J_VERSION", "4"), "4") {
		database = ""
	}
	return &Neo4jConfiguration{
		Url:      LookupEnvOrGetDefault("NEO4J_URI", "neo4j://localhost"),
		Username: LookupEnvOrGetDefault("NEO4J_USER", "neo4j"),
		Password: LookupEnvOrGetDefault("NEO4J_PASSWORD", "kausjan"),
		Database: database,
	}
}

func LookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}

func UnsafeClose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}
