package Auth

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

type Authenticated struct {
	Authenticated string
}




func AuthenticatedUser(akey, username string) (bool) {

	var apikey string
	var authenticated bool;
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
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		defer tx.Rollback()
		stmt, err := tx.Query("SELECT apikey FROM accontrol.tbl_users WHERE username = ?", username )
		if err != nil {
			log.Fatal("There was a problem looking you up.")
		}
		defer stmt.Close()



		for stmt.Next() {
			err := stmt.Scan(&apikey)
			if err != nil {
				log.Fatal(err)
			}
			if apikey != akey {
				log.Println("You are not authenticated")
				 authenticated := false

				return authenticated
			}

			if apikey == akey {
				log.Println("You are authenticated!")
				authenticated := true
				return authenticated
			}

		}

		return authenticated

	}
