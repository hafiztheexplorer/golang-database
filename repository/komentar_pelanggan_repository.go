package repository

import (
	"context"
	"golang-database/entity"
)

// implementasi function/method yang berhubungan ke database
type Komentar_pelangganRepository interface {
	// semuanya pake context untuk interaksi dengan database
	// untuk memasukkan data ke database lewat file listperintah.go disebut juga (struct model/entity)
	// return ada dua, dari listperintah dan error
	Insert(ctx context.Context, komentar_pelanggan entity.Komentar_pelanggan) (entity.Komentar_pelanggan, error)
	// untuk query data by ID data ke database lewat file listperintah.go disebut juga (struct model/entity)
	FindById(ctx context.Context, id int32) (entity.Komentar_pelanggan, error)
	// untuk query data semuanya ke database lewat file listperintah.go disebut juga (struct model/entity)
	FindAll(ctx context.Context) ([]entity.Komentar_pelanggan, error)
	// - update by id under construction
	UpdateById(ctx context.Context, id int32, komentar_pelanggan entity.Komentar_pelanggan) (entity.Komentar_pelanggan, error)
	//  - delete by id under construction, return ada dua, dari listperintah dan error
	DeleteById(ctx context.Context, id int32) (entity.Komentar_pelanggan, error)
}
