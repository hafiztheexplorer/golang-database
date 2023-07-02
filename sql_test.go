package database_go

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

/*
----------------------------------------------------------------
ExecSQLContext untuk memasukkan data berupa context tanpa menampilkan / ambil data dari database
----------------------------------------------------------------
*/
func TestExecSql(t *testing.T) {
	// kita coba masukkan data ke database / tabel database
	// tanpa ada feedback / hasil / result balik ke kita
	db := GetConnection() // koneksi ke database kita
	defer db.Close()

	ctx := context.Background() // something something go context, lupa nanti pelajari lagi

	scriptsqlsaya := "insert into daftar_pelanggan(id,nama) values('P0003','contoh nama pelanggan 3') "
	_, err := db.ExecContext(ctx, scriptsqlsaya) // dikirim dalam bentuk context

	if err != nil {
		panic(err)
	}

	fmt.Println("sukses menambah data ke tabel database - pelanggan baru")
	// pastikan data sudah benar2 terkirim dengan select * "namatable", di mysql workbench
}

/*
----------------------------------------------------------------
untuk menampilkan data dengan context yang didapat dari database dan lewat pointer ke header rownya
tipe data rows
----------------------------------------------------------------
*/
func TestQuerySql(t *testing.T) {
	// kita coba masukkan data ke database / tabel database
	// tanpa ada feedback / hasil / result balik ke kita
	db := GetConnection() // koneksi ke database kita
	defer db.Close()

	ctx := context.Background() // something something go context, lupa nanti pelajari lagi

	scriptsqlsaya := "select * from daftar_pelanggan"
	rows, err := db.QueryContext(ctx, scriptsqlsaya) // dikirim dalam bentuk query context

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		// cursor ke database dengan rows.Next() jadi menggunakan for
		//selama bernilai true rows.Nextnya maka akan terus teriterasi
		var id, nama string
		err = rows.Scan(&id, &nama) // kenapa pointer "&"", nanti ini hasil data sesungguhnya yang akan ditangkap, bukan data yang diduplikasi karena ini go
		if err != nil {
			panic(err)
		}
		fmt.Println("id = ", id)
		fmt.Println("nama = ", nama)
	}

}

/*
----------------------------------------------------------------
untuk menampilkan data dengan context yang didapat dari database dan lewat pointer ke database yang lebih komples
----------------------------------------------------------------
*/

func TestQuerySqlKompleks(t *testing.T) {
	// kita coba masukkan data ke database / tabel database
	// tanpa ada feedback / hasil / result balik ke kita
	db := GetConnection() // koneksi ke database kita
	defer db.Close()

	ctx := context.Background() // something something go context, lupa nanti pelajari lagi

	scriptsqlsaya := "select id,nama,email,balance,rating,created_at,DOB,sex from daftar_pelanggan_2"
	rows, err := db.QueryContext(ctx, scriptsqlsaya) // dikirim dalam bentuk context

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		// cursor ke database dengan rows.Next() jadi menggunakan for
		//selama bernilai true rows.Nextnya maka akan terus teriterasi
		var id, nama, email, sex sql.NullString
		var balance sql.NullInt32
		var rating sql.NullFloat64
		var DOB, created_at sql.NullTime
		// sql.blabla dia nullable, nilai bisa null kalo diambil dari sql spt itu
		// coba dilihat sp.blabla nya ada 2 keluaran, ada tipe data
		// dan booleannya true / false menandakan kosong atau tidaknya
		// var DOB []uint8
		// var created_at []uint8

		err := rows.Scan(&id, &nama, &email, &balance, &rating, &created_at, &DOB, &sex) // kenapa pointer "&"", nanti ini hasil data sesungguhnya yang akan ditangkap, bukan data yang diduplikasi karena ini go
		if err != nil {
			panic(err)
		}

		fmt.Println("--------------------------------------------------------")
		if id.Valid { // nanti yang di print nilainya
			fmt.Println("id = ", id.String)
		}
		if nama.Valid {
			fmt.Println("nama = ", nama.String)
		}
		if nama.Valid {
			fmt.Println("Tanggal Lahir = ", DOB.Time)
		}
		if nama.Valid {
			fmt.Println("Jenis Kelamin = ", sex.String)
		}
		if nama.Valid {
			fmt.Println("E-Mail = ", email.String)
		}
		if nama.Valid {
			fmt.Println("Balance = ", balance.Int32)
		}
		if nama.Valid {
			fmt.Println("Rating = ", rating.Float64)
		}
		if nama.Valid {
			fmt.Println("dibuat pada = ", created_at.Time)
		}
		fmt.Println("---------------------------------------------------------")

	}

}

/*
----------------------------------------------------------------
untuk SQL injection by Go, query sql with parameter
----------------------------------------------------------------
*/
func TestQuerySqlWParam(t *testing.T) {
	// kita coba masukkan data ke database / tabel database
	// tanpa ada feedback / hasil / result balik ke kita
	db := GetConnection() // koneksi ke database kita
	defer db.Close()

	ctx := context.Background() // something something go context, lupa nanti pelajari lagi

	// simulasi input dari user
	username := "username1'; #" // markicob sql injection pada ini dengan "username1'; #"
	userpassword := "password"

	scriptsqlsaya := "select * from daftar_user where username= '" + username + "' and userpassword = '" + userpassword + "' limit 1;" //sql script ini rawan kena inject dari inputan, data stringnya disambung, perhatikan tanda petiknya
	rows, err := db.QueryContext(ctx, scriptsqlsaya)                                                                                   // dikirim dalam bentuk context

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() { // kalau misal ada datanya
		var username string
		err := rows.Scan(&username, &userpassword)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login dengan username: ", username)
		fmt.Println("dan dengan menggunakan password: ", userpassword)
		fmt.Println(scriptsqlsaya)
	} else {
		fmt.Println("gagal login")

	}

}

/*
----------------------------------------------------------------
untuk SQL  with parameter anti sql injection
----------------------------------------------------------------
*/
func TestQuerySqlWParam2(t *testing.T) {
	// kita coba masukkan data ke database / tabel database
	// tanpa ada feedback / hasil / result balik ke kita
	db := GetConnection() // koneksi ke database kita
	defer db.Close()

	ctx := context.Background() // something something go context, lupa nanti pelajari lagi

	// simulasi input dari user
	username := "username1'; #" // markicob sql injection pada ini dengan "username1'; #"
	userpassword := "password"

	scriptsqlsaya := "select * from daftar_user where username = ? and userpassword = ? limit 1;" //data stringnya disambung, perhatikan tanda petiknya
	rows, err := db.QueryContext(ctx, scriptsqlsaya, username, userpassword)
	// tambahkan arguments setelah string script, di case ini username dan password
	// dikirim dalam bentuk context

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() { // kalau misal ada datanya
		var username string
		err := rows.Scan(&username, &userpassword)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login dengan username: ", username)
		fmt.Println("dan dengan menggunakan password: ", userpassword)
		fmt.Println(scriptsqlsaya)
	} else {
		fmt.Println("username / password salah, gagal login")

	}

}

/*
----------------------------------------------------------------
ExecSQLContext untuk memasukkan data berupa context tanpa menampilkan / ambil data
dari database dan tanpa beresiko kena SQL injection
----------------------------------------------------------------
*/
func TestExecSqlParam2(t *testing.T) {
	// kita coba masukkan data ke database / tabel database
	// tanpa ada feedback / hasil / result balik ke kita
	db := GetConnection() // koneksi ke database kita
	defer db.Close()

	ctx := context.Background() // something something go context, lupa nanti pelajari lagi

	id := "P0005; drop table user; #"
	nama := "contoh nama pelanggan 4"
	scriptsqlsaya := "insert into daftar_pelanggan(id,nama) values( ? , ? ) "
	// biar tidak ada sql injection, jangan hardcode inputan di sini
	_, err := db.ExecContext(ctx, scriptsqlsaya, id, nama)
	// dikirim dalam bentuk context

	if err != nil {
		panic(err)
	}

	fmt.Println("sukses menambah data ke tabel database - pelanggan baru")
	// pastikan data sudah benar2 terkirim dengan select * "namatable", di mysql workbench
}

/*
----------------------------------------------------------------
input data dengan auto increment
----------------------------------------------------------------
*/
func TestExecSqlParamAutoIncrement(t *testing.T) {
	// kita coba masukkan data ke database / tabel database
	// tanpa ada feedback / hasil / result balik ke kita
	db := GetConnection() // koneksi ke database kita
	defer db.Close()

	ctx := context.Background()

	email := "contoh mail 3"
	isi_komentar := "no komen"
	scriptsqlsaya := "insert into komentar_pelanggan(email,isi_komentar) values( ? , ? ) "
	result, err := db.ExecContext(ctx, scriptsqlsaya, email, isi_komentar)
	// dikirim dalam bentuk context

	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("ID terakhir yang sudah dimasukkan :", insertId)

}

/*
----------------------------------------------------------------
TestExecSql SQL Query Prepare Statement
----------------------------------------------------------------
*/
func TestExecSqlParamPrepareStmt(t *testing.T) {
	// Prepare statement
	// kita coba masukkan data ke database / tabel database
	// tanpa ada feedback / hasil / result balik ke kita
	db := GetConnection() // koneksi ke database kita
	defer db.Close()

	ctx := context.Background()
	scriptsqlsaya := "insert into komentar_pelanggan(email,isi_komentar) values( ? , ? ) "
	stmt, err := db.PrepareContext(ctx, scriptsqlsaya)
	// dikirim dalam bentuk context
	if err != nil {
		panic(err)
	}
	defer stmt.Close() // wajib dikasih untuk "prepare statement"

	for i := 0; i < 10; i++ {
		email := "contoh email " + strconv.Itoa(i) + "@mail.com"
		isi_komentar := "contoh komentar ke-" + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, isi_komentar)
		if err != nil {
			panic(err)
		}
		id, _ := result.LastInsertId()
		fmt.Println("Comment ID ke-", id)
	}

}

/*
----------------------------------------------------------------
Database Transaction di Go - ketika commit dan rollback
----------------------------------------------------------------
*/

func TestDbTrsctn(t *testing.T) {
	// di bawah ini syambung ke function dari module yang syambung ke database local
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin() // db.Begin ada 2 keluaran, tx dan err
	if err != nil {
		panic(err)
	}

	// kalau tidak, jalankan transaction di bawah ini
	// sqlnya
	scriptsqlsaya := "INSERT INTO komentar_pelanggan(email, isi_komentar) VALUES( ? , ? )"
	// perulangannya
	for i := 0; i < 10; i++ {
		email := "contoh email" + strconv.Itoa(i) + "@mail.com"
		isi_komentar := "contoh komentar ke-" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, scriptsqlsaya, email, isi_komentar)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment ID ke-", id)
	}

	// lalu commit/rollback (pilih salah satu)
	// tx.commit & tx.Rollback return valuenya bisa error juga jadi buat seperti ini
	// kalau ok kan yaudah commit gt aja
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
