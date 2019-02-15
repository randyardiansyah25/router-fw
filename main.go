package main

import (
	"examps/router-fw/connection"
	"examps/router-fw/router"
	"github.com/joho/godotenv"
	"github.com/kpango/glg"
	"os"
)

func main() {
	_ = glg.Log("Starting....")
	router.Start()
}

func init(){
	_ = glg.Log("Loading configuration file....")
	err := godotenv.Load("config.env")
	if err != nil {
		_ = glg.Error("Error loading configuration file!")
		os.Exit(1)
	}

	log := glg.FileWriter("log/service.log", 0666)

	glg.Get().SetMode(glg.BOTH)
	if os.Getenv("DUMP_LOG") == "true" {
		glg.Get().AddLevelWriter(glg.LOG, log).
			AddLevelWriter(glg.ERR, log).
			AddLevelWriter(glg.WARN, log).
			AddLevelWriter(glg.DEBG, log).
			AddLevelWriter(glg.INFO, log)
	}

	err = connection.MysqlConnect(os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	if err != nil {
		_ = glg.Error("Connection database error : ", err.Error())
		os.Exit(1)
	}
}
