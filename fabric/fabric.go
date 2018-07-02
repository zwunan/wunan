package fabric

import(
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)


func Fabric(username string, password string) string {	
	db,err := sql.Open("sqlite3","./UserDB.db")
	checkErr(err)
	var userName string
	userName = ""
	rows,err :=db.Query("SELECT USERNAME FROM USER WHERE USERNAME=? AND PASSWORD=?",username,password)
	for rows.Next() {
		err = rows.Scan(&userName)
		checkErr(err)
		fmt.Println(userName)
	}

	rows.Close()
	db.Close()
	return userName
}

func Register(username string, password string) string {
	db,err:=sql.Open("sqlite3","./UserDB.db")
	checkErr(err)
	var userName string
	rows,err :=db.Query("SELECT USERNAME FROM USER WHERE USERNAME=? AND PASSWORD=?",username,password)
	for rows.Next() {
		err = rows.Scan(&userName)
		checkErr(err)
		fmt.Println(userName)
	}
	if (userName !=""){
		return "00"
	}
	stmt, err := db.Prepare("insert into USER(USERNAME,PASSWORD) values(?,?)")
	checkErr(err)
	_, err = stmt.Exec(username, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(username)
	db.Close()
	return "11"
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}