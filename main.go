package main

import (
	"bytes"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	// buf := &bytes.Buffer{}
	// buf.WriteString("DB_HOST=localhost")
	// buf.WriteString("\nDB_PORT=3306")
	// buf.WriteString("\nDB_USER=root")
	// buf.WriteString("\nDB_PASSWORD=123456")
	// buf.WriteString("\nDB_NAME=test")
	// myEnv, err := godotenv.Parse(buf)
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// fmt.Println("name: ", myEnv["DB_USER"])
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// fmt.Println("Hello, World!")
	// r.Run()

	buf := &bytes.Buffer{}
	buf.WriteString("DB_HOST=localhost")
	buf.WriteString("\nDB_PORT=3306")
	buf.WriteString("\nDB_USER=root")
	buf.WriteString("\nDB_PASSWORD=123456")
	buf.WriteString("\nDB_NAME=test")
	env, err := godotenv.Parse(buf)
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Write(env, "./.env")
	if err != nil {
		log.Fatal(err)
	}

}
