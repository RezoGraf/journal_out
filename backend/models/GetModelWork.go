package models

import (
	"database/sql"
	_ "github.com/nakagami/firebirdsql"
	"../db"
	"../utils"
	"fmt"
)

type ModePatientFindOfArena struct {
	Fio string `json:"Fio"`
	Pasport string `json:"Pasport"`
	DateRogd string `json:"DateRogd"`
	Pol string `json:"Pol"`
}

func GetPatientOfArena(number_cart string) ([]*ModePatientFindOfArena)  {
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
