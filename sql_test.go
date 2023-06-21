package database_go

import (
	"context"
	"fmt"
	"testing"
)

/*
----------------------------------------------------------------
untuk memasukkan data berupa context tanpa menampilkan / ambil data dari database
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
	rows, err := db.QueryContext(ctx, scriptsqlsaya) // dikirim dalam bentuk context

	if err != nil {
		panic(err)
	}

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

	defer rows.Close()

}
