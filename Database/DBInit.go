package Database

import (
	"database/sql"
	"log"
	"github.com/spf13/viper"
	_ "github.com/go-sql-driver/mysql"
)

// Specific for the auth db
func DatabaseInitAuth() (*sql.DB)  {
	viper.AddConfigPath("/etc/commservice/")
	viper.SetConfigName("comconfig")
	viper.ReadInConfig()
	dbusername := viper.GetString("authdb.username")
	dbpass := viper.GetString("authdb.password")
	serverip := viper.GetString("authdb.dbhost")

	db, err := sql.Open("mysql", dbusername + ":" + dbpass +  "@tcp(" + serverip + ")" + "/")
	if err != nil {
		log.Fatal("Sorry there was a problem connecting to the database with user " + dbusername + " host " + serverip +  " pass " + dbpass + " Please check /etc/commservice/credentials.yaml")
		log.Fatal(err)

	}
	return db

}

git