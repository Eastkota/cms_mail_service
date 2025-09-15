package config

import (
	"log"
	"os"
	"strconv"
)


func PostgresUser() 	string { return os.Getenv("POSTGRES_USER") }
func PostgresPassword() string { return os.Getenv("POSTGRES_PASSWORD") }
func PostgresHost() 	string {return os.Getenv("POSTGRES_HOST") }
func PostgresPort() 	string { return os.Getenv("POSTGRES_PORT") }
func PostgresDB() 		string { return os.Getenv("POSTGRES_DB") }
func MailUsername() 	string {return os.Getenv("MailUsername")}
func MailPassword() 	string {return os.Getenv("MailPassword")} 
func MailHost()			string {return os.Getenv("MailHost")}
func MailEncryption() 	string {return os.Getenv("MailEncryption")}
func MailFromAddress()	string {return os.Getenv("MailFromAddress")}
func MailFromNAme()		string {return os.Getenv("MailFromName")}
  
func MailPort() int {
	portStr := os.Getenv("MailPort")

	if portStr == "" {
		log.Fatal("MailPort environment variable is not set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid MailPort value: %s. Must be an integer.", portStr)
	}

	return port
}
