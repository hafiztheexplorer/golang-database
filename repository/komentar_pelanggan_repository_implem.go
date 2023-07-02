package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type komentar_pelangganRepositoryImplem struct {
	DB *sql.DB
}

func Komentar_pelangganRepositoryBaru(db *sql.DB) Komentar_pelangganRepository {
	return &komentar_pelangganRepositoryImplem{DB: db}
}

// logic insert data
func (repository *komentar_pelangganRepositoryImplem) Insert(ctx context.Context, komentar_pelanggan entity.Komentar_pelanggan) (entity.Komentar_pelanggan, error) {
	script := "INSERT INTO komentar_pelanggan(email, isi_komentar) VALUES(? , ? )"
	hasil, err := repository.DB.ExecContext(ctx, script, komentar_pelanggan.Email, komentar_pelanggan.Isi_komentar)

	if err != nil {
		return komentar_pelanggan, err
	}

	id, err := hasil.LastInsertId()
	if err != nil {
		return komentar_pelanggan, err
	}

	komentar_pelanggan.Id = int32(id)
	return komentar_pelanggan, nil

}

// logic view data by id
func (repository *komentar_pelangganRepositoryImplem) FindById(ctx context.Context, id int32) (entity.Komentar_pelanggan, error) {
	script := "SELECT * FROM komentar_pelanggan WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	komentar_pelanggan := entity.Komentar_pelanggan{}
	if err != nil {
		return komentar_pelanggan, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&komentar_pelanggan.Id, &komentar_pelanggan.Email, &komentar_pelanggan.Isi_komentar)
		return komentar_pelanggan, nil
	} else {
		// tidak ada
		return komentar_pelanggan, errors.New("ID " + strconv.Itoa(int(id)) + " IS NOT FOUND")
	}
}

// logic view data all
func (repository *komentar_pelangganRepositoryImplem) FindAll(ctx context.Context) ([]entity.Komentar_pelanggan, error) {
	script := "SELECT * FROM komentar_pelanggan"
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var komentar_pelanggans []entity.Komentar_pelanggan
	for rows.Next() {
		// ada
		komentar_pelanggan := entity.Komentar_pelanggan{}
		rows.Scan(&komentar_pelanggan.Id, &komentar_pelanggan.Email, &komentar_pelanggan.Isi_komentar)
		komentar_pelanggans = append(komentar_pelanggans, komentar_pelanggan)

	}
	return komentar_pelanggans, nil
}

// logic update data by id
func (repository *komentar_pelangganRepositoryImplem) UpdateById(ctx context.Context, id int32, komentar_pelanggan entity.Komentar_pelanggan) (entity.Komentar_pelanggan, error) {
	script := "UPDATE komentar_pelanggan SET email = ? , isi_komentar = ? WHERE id = ? LIMIT 1"
	hasil, err := repository.DB.QueryContext(ctx, script, komentar_pelanggan.Email, komentar_pelanggan.Isi_komentar, id)
	komentar_pelanggan2 := entity.Komentar_pelanggan{}
	if err != nil {
		return komentar_pelanggan2, err
	}
	defer hasil.Close()
	if hasil.Next() {
		// ada
		hasil.Scan(&komentar_pelanggan2.Id, &komentar_pelanggan2.Email, &komentar_pelanggan2.Isi_komentar)
	}
	return komentar_pelanggan2, nil
	// else {
	// 	// tidak ada
	// 	return komentar_pelanggan2, err
	// 	// errors.New("ID" + strconv.Itoa(int(id)) + "IS NOT FOUND")
	// }
}

// logic delete data by id
func (repository *komentar_pelangganRepositoryImplem) DeleteById(ctx context.Context, id int32) (entity.Komentar_pelanggan, error) {
	script := "DELETE FROM komentar_pelanggan WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	komentar_pelanggan := entity.Komentar_pelanggan{}
	if err != nil {
		return komentar_pelanggan, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&komentar_pelanggan.Id, &komentar_pelanggan.Email, &komentar_pelanggan.Isi_komentar)
		return komentar_pelanggan, nil
	} else {
		// tidak ada
		return komentar_pelanggan, errors.New("ID " + strconv.Itoa(int(id)) + " IS NOT FOUND")
	}
}
