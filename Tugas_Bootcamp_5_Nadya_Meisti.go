package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/nadya-sheet"
)

func main() {
	var connStr = "user=username dbname=mydb sslmode=disable"
	var db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Pasien (
        ID serial PRIMARY KEY,
        Nama VARCHAR (255),
        Alamat TEXT,
        TanggalLahir DATE,
        JenisKelamin VARCHAR (10)
    );`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS CatatanMedis (
        ID serial PRIMARY KEY,
        PasienID INT,
        Tanggal DATE,
        Diagnosa TEXT,
        FOREIGN KEY (PasienID) REFERENCES Pasien(ID)
    );`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tabel berhasil dibuat.")
}
