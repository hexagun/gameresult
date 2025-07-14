package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func initLogging() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
}

func setConfig() {
	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/config") // path to look for the config file in
	viper.AddConfigPath(".")       // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	log.Info().Msg("configs:")
	log.Debug().Msg(fmt.Sprintf("%s%s", "environment:", viper.GetString("environment")))
	log.Debug().Msg(fmt.Sprintf("%s%d", "gameresult.port:", viper.GetInt("gameresult.port")))
	// log.Debug().Msg(fmt.Sprintf("%s%s", "db.host:", viper.GetString("db.host")))
	// log.Debug().Msg(fmt.Sprintf("%s%s", "db.user:", viper.GetString("db.user")))
	// log.Debug().Msg(fmt.Sprintf("%s%s", "db.password:", viper.GetString("db.password")))
	// log.Debug().Msg(fmt.Sprintf("%s%s", "db.name:", viper.GetString("db.name")))
}

func main() {
	initLogging()
	setConfig()
	gameResultPort := viper.GetInt("gameresult.port")
	// // host := viper.GetString("db.host")
	// // user := viper.GetString("db.user")
	// // password := viper.GetString("db.password")
	// // name := viper.GetString("db.name")
	// // dbPort := viper.GetString("db.port")

	// log.Debug().Msg(fmt.Sprintf("%s", "Pre db init"))
	// database.InitDB(dbPort, host, user, password, name)
	// log.Debug().Msg(fmt.Sprintf("%s", "Post db init"))
	r := gin.Default()
	// //r.POST("/users", handlers.CreateUser)

	r.POST("/gameresult", func(c *gin.Context) {
		log.Debug().Msg(fmt.Sprintf("%s", "gameresult"))
	})

	port := fmt.Sprintf("%s%d", ":", gameResultPort)
	log.Debug().Msg(fmt.Sprintf("%s%s", "Handling gameresults on port ", port))
	r.Run(port)
}
