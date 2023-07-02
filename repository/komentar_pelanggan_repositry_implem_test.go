package repository

import (
	"context"
	"fmt"
	database_go "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMemasukkanKomentar_pelanggan(t *testing.T) {
	komentar_pelanggan_repository := Komentar_pelangganRepositoryBaru(database_go.GetConnection())

	ctx := context.Background()
	komentar_pelanggan := entity.Komentar_pelanggan{
		Email:        "contohmai3l@mail.com",
		Isi_komentar: "contoh komentar 1 dari repo",
	}
	hasil, err := komentar_pelanggan_repository.Insert(ctx, komentar_pelanggan)

	if err != nil {
		panic(err)
	}

	fmt.Println(hasil)
}

func TestFindByIDKomentar_pelanggan(t *testing.T) {
	komentar_pelanggan_repository := Komentar_pelangganRepositoryBaru(database_go.GetConnection())
	komentar_pelanggan, err := komentar_pelanggan_repository.FindById(context.Background(), 116)

	if err != nil {
		panic(err)
	}

	fmt.Println(komentar_pelanggan)
}

func TestFindAllKomentar_pelanggan(t *testing.T) {
	komentar_pelanggan_repository := Komentar_pelangganRepositoryBaru(database_go.GetConnection())
	komentar_pelanggans, err := komentar_pelanggan_repository.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _, komentar_pelanggan := range komentar_pelanggans {
		fmt.Println(komentar_pelanggan)
	}
}

func TestUpdateByIdKomentar_pelanggan(t *testing.T) {
	komentar_pelanggan_repository := Komentar_pelangganRepositoryBaru(database_go.GetConnection())

	ctx := context.Background()
	komentar_pelanggan := entity.Komentar_pelanggan{
		Email:        "sample update 2",
		Isi_komentar: "sample update 2",
	}
	hasil, err := komentar_pelanggan_repository.UpdateById(ctx, 69, komentar_pelanggan)

	if err != nil {
		panic(err)
	}

	fmt.Println(hasil)

}

func TestDeleteByIDKomentar_pelanggan(t *testing.T) {
	komentar_pelanggan_repository := Komentar_pelangganRepositoryBaru(database_go.GetConnection())

	ctx := context.Background()
	komentar_pelanggan, _ := komentar_pelanggan_repository.DeleteById(ctx, 116)

	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(komentar_pelanggan)
}
