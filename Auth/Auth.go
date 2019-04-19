package Auth

import (
	"database/sql"
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type AuthInfo struct {
	UserName string
	APIKey string

}

var apikey string;

func AuthenticatedUser(w http.ResponseWriter, r *http.Request) bool {


	var AuthInfo  AuthInfo
	_ = json.NewDecoder(r.Body).Decode(&AuthInfo)
	w.Header().Set("Content-Type", "application/json")
	var authenticated bool;
	viper.AddConfigPath("/etc/commservice/")
	viper.SetConfigName("credentials")
	viper.ReadInConfig()
	dbusername := viper.GetString("devdb.username")
	dbpass := viper.GetString("devdb.password")
	serverip := viper.GetString("devdb.dbhost")

	db, err := sql.Open("mysql", dbusername+":"+dbpass+"@tcp("+serverip+")"+"/")
	if err != nil {
		log.Fatal("Sorry there was a problem connecting to the database with user " + dbusername + "Please check /etc/commservice/credentials.yaml")

	}
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		defer tx.Rollback()
		stmt, err := tx.Query("SELECT apikey FROM accontrol.tbl_users WHERE username = %v", AuthInfo.UserName )
		if err != nil {
			log.Fatal("There was a problem looking you up.")
		}
		defer stmt.Close()


		for stmt.Next() {
			err := stmt.Scan(&apikey)
			if err != nil {
				log.Fatal(err)
			}
			if apikey != AuthInfo.APIKey  {
				log.Fatal("You are not authenticated")
				return authenticated == false
			}

			if apikey == AuthInfo.APIKey {
				log.Println("You are authenticated!")
				return authenticated == true
			}

		}

		return authenticated
	}
