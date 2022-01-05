package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// Config struct for describe configuration of the app.
type Config struct {
	Server *serverConfig
}

type serverConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

var (
	once     sync.Once // create sync.Once primitive
	instance *Config   // create nil Config struct
)

// NewConfig function to prepare config variables from .env file and return config.
func NewConfig() *Config {
	// Configuring config one time.
	once.Do(func() {
		// Server host (should be string):
		host := os.Getenv("SERVER_HOST")
		// Server port (should be int):
		port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
		if err != nil {
			panic("wrong server port (check your .env)")
		}
		// Server read timeout (should be int):
		readTimeout, err := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
		if err != nil {
			panic("wrong server read timeout (check your .env)")
		}
		// Server write timeout (should be int):
		writeTimeout, err := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
		if err != nil {
			panic("wrong server write timeout (check your .env)")
		}
		// Server idle timeout (should be int):
		idleTimeout, err := strconv.Atoi(os.Getenv("SERVER_IDLE_TIMEOUT"))
		if err != nil {
			panic("wrong server idle timeout (check your .env)")
		}

		// Set all variables to the config instance.
		instance = &Config{
			Server: &serverConfig{
				Addr:         fmt.Sprintf("%s:%d", host, port),
				ReadTimeout:  time.Duration(readTimeout) * time.Second,
				WriteTimeout: time.Duration(writeTimeout) * time.Second,
				IdleTimeout:  time.Duration(idleTimeout) * time.Second,
			},
		}
	})

	// Return configured config instance.
	return instance
}
