package models

import (
	"database/sql"
	_ "github.com/nakagami/firebirdsql"
	"../db"
	"../utils"
	"fmt"
)

const programname  = "vaccinations"

type ModePatientFindOfArena struct {
	Fio string `json:"Fio"`
	Pasport string `json:"Pasport"`
	DateRogd string `json:"DateRogd"`
	Pol string `json:"Pol"`
}

type CheckUser struct {
	Count string
}


func ModelsAuth(name, pass string) []*CheckUser {
	rows, err := db.Select(`SELECT count(*) FROM users WHERE login = $1 and pass = $2 and access_program = $3`,
						name, pass, programname)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	checkUser := make([]*CheckUser, 0)
	for rows.Next() {
		user := new(CheckUser)
		var count sql.NullString
		rows.Scan(&count)
		user.Count = count.String
		checkUser = append(checkUser, user)
	}
	return checkUser
}

func ModelsGetPatientOfArena(number_cart string) ([]*ModePatientFindOfArena)  {
		fmt.Println(number_cart)
		query := `select fam || ' ' || im || ' ' || OT as fio , pasp, DR, SEX from patient where UID = ?`
		var Fio sql.NullString
		var Pasport sql.NullString
		var DateRogd sql.NullString
		var Pol sql.NullString


	    rows, _ := db.Select(query, number_cart)
		bks := make([]*ModePatientFindOfArena, 0)

		for rows.Next() {
				bk := new(ModePatientFindOfArena)
				rows.Scan(&Fio, &Pasport, &DateRogd, &Pol)
				bk.Fio = Fio.String
				bk.Pasport = Pasport.String
				bk.DateRogd = utils.Cut(DateRogd.String, 0, 10)
				bk.Pol = Pol.String

				bks = append(bks, bk)
			}
			fmt.Println(bks)

			return bks
		}
