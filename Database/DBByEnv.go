package Database


import (
	"log"
)

func DatabaseByEnvironment(Environment string) (string){
	db := DatabaseInitAll("/etc/dm/","GenService", "GenService.username", "GenService.password", "GenService.dbhost")
	var db_ip string;
	//var databaseip string;
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	stmt, err := tx.Query("SELECT db_ip FROM clarity_tools.tbl_database_info  WHERE db_env = ?", Environment)
	if err != nil {
		log.Fatal("There was a problem up the db.")
	}
	defer stmt.Close()



	for stmt.Next() {
		err := stmt.Scan(&db_ip)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db_ip

}

// Returns ip address of all databases
func DatabaseAllIP() []string {
	var databaseip  []string
	db := DatabaseInitAll("/etc/dm/","GenService", "GenService.username", "GenService.password", "GenService.dbhost")
	var db_ip string;
	//var databaseip string;
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	stmt, err := tx.Query("SELECT db_ip FROM clarity_tools.tbl_database_info")
	if err != nil {
		log.Fatal("There was a problem up the db.")
	}
	defer stmt.Close()



	for stmt.Next() {
		err := stmt.Scan(&db_ip)
		if err != nil {
			log.Fatal(err)
		}
		databaseip = append(databaseip, db_ip)
	}

	return databaseip

}
