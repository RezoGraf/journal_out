package models

import (
	"database/sql"
	"../db"
	"../utils"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"github.com/lib/pq"
)

const programname = "vaccinations"

type ModelUser struct {
	Id  string `json:"id"`
	Fam string `json:"fam"`
	Im  string `json:"im"`
	Ot  string `json:"ot"'`
}

type ModePatientFindOfArena struct {
	Fio      string `json:"Fio"`
	Pasport  string `json:"Pasport"`
	DateRogd string `json:"DateRogd"`
	Pol      string `json:"Pol"`
	Uid      string `json:"NumberKart"`
}

type ModelPrivivka struct {
	Date        string `json:"date"`
	Preparat    string `json:"preparat"`
	Seria       string `json:"seria"`
	Doza        string `json:"doza"`
}

type CheckUser struct {
	Count string
}

func ModelsAddPrivivka(userId, vaccination, date, preparat, seria, doza, numberkart string) {
	db.Exec(`INSERT INTO vac_vaccinations (vrach_id, type_vaccinations, date, preparat, seria, doza, numberkart)
	VALUES ($1, $2, $3::timestamp, $4, $5, $6, $7)`,
		utils.NullableInt(userId), utils.NullableString(vaccination), utils.NullableTime(date), utils.NullableString(preparat),
		utils.NullableString(seria), utils.NullableString(doza), utils.NullableInt(numberkart))
}

func ModelsGetUserInfo(name, pass string) []*ModelUser {
	rows, err := db.Select("postgres", `SELECT id, fam, im, ot FROM users WHERE login = $1 and pass = $2 and access_program = $3`,
		name, pass, programname)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	AuthUser := make([]*ModelUser, 0)
	for rows.Next() {
		user := new(ModelUser)
		var id sql.NullString
		var fam sql.NullString
		var im sql.NullString
		var ot sql.NullString
		rows.Scan(&id, &fam, &im, &ot)
		user.Id = id.String
		user.Fam = fam.String
		user.Im = im.String
		user.Ot = ot.String
		AuthUser = append(AuthUser, user)
	}
	return AuthUser
}

func ModelsGetPrivivka(type_vaccinations, numberKart string) []*ModelPrivivka {
	rows, err := db.Select("postgres", `SELECT date, preparat, seria, doza
	FROM vac_vaccinations
	WHERE type_vaccinations = $1 and numberKart = $2`,
		type_vaccinations, numberKart)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	getVaccinations := make([]*ModelPrivivka, 0)
	for rows.Next() {
		vacc := new(ModelPrivivka)
		var date pq.NullTime
		var preparat sql.NullString
		var seria sql.NullString
		var doza sql.NullString
		rows.Scan(&date, &preparat, &seria, &doza)
		vacc.Date = utils.FormatDatePqNullTime(date)
		vacc.Preparat = preparat.String
		vacc.Doza = doza.String
		vacc.Seria = seria.String
		getVaccinations = append(getVaccinations, vacc)
	}
	return getVaccinations
}

func ModelsAuth(name, pass string) []*CheckUser {
	rows, err := db.Select("postgres", `SELECT count(*) FROM users WHERE login = $1 and pass = $2 and access_program = $3`,
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

func ModelsGetPatientOfArena(number_cart string) ([]*ModePatientFindOfArena) {
	fmt.Println(number_cart)
	query := `select fam || ' ' || im || ' ' || OT as fio , pasp, DR, SEX, uid from patient where UID = ?`
	var Fio sql.NullString
	var Pasport sql.NullString
	var DateRogd pq.NullTime
	var Pol sql.NullString
	var Uid sql.NullString

	rows, err := db.Select("firebirdsl", query, number_cart)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	bks := make([]*ModePatientFindOfArena, 0)

	for rows.Next() {
		bk := new(ModePatientFindOfArena)
		rows.Scan(&Fio, &Pasport, &DateRogd, &Pol, &Uid)
		bk.Fio = Fio.String
		bk.Pasport = Pasport.String
		bk.DateRogd = utils.FormatDatePqNullTime(DateRogd)
		bk.Pol = Pol.String
		bk.Uid = Uid.String

		bks = append(bks, bk)
	}
	fmt.Println(bks)

	return bks
}
