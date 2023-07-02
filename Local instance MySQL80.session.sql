/* 
untuk contoh pertama menggunakan context dari golang, 
keluaran data berupa rows dan error , lebih jelasnya lihat package database_go
*/
USE belajar_golang_database;
create table daftar_pelanggan
(
id		VARCHAR(100) not null,
nama	VARCHAR(100) not null,
primary key(id)
)engine=innoDB;

select * from daftar_pelanggan;
desc daftar_pelanggan;

/* 
untuk tipe data lainm buat tabel baru
*/
create table daftar_pelanggan_2
(
id		VARCHAR(100) not null,
nama	VARCHAR(100) not null,
email 	VARCHAR(100) not null,
balance integer	default 0,
rating	double default 0.0,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
DOB	DATE,
sex SET ('L','P','') DEFAULT NULL,
primary key(id)
)engine=innoDB;

desc daftar_pelanggan_2;

insert into daftar_pelanggan_2 (id, nama, email, balance, rating, DOB, sex) 
VALUES
('P0001','contoh nama 1','nama1@email.com',1000000,4.8,'1991-1-1','L'),
('P0002','contoh nama 2','nama2@email.com',1000000,4.3,'1992-2-2','P'),
('P0003','contoh nama 3','nama3@email.com',1000000,4.9,'1993-3-3','P'),
('P0004','contoh nama 4','nama4@email.com',1000000,3.6,'1994-4-4','L'),
('P0005','contoh nama 5','nama5@email.com',1000000,2.8,'1995-5-5','P');

select * from daftar_/* 
untuk contoh pertama menggunakan context dari golang, 
keluaran data berupa rows dan error , lebih jelasnya lihat package database_go
*/
USE belajar_golang_database;
create table daftar_pelanggan
(
id		VARCHAR(100) not null,
nama	VARCHAR(100) not null,
primary key(id)
)engine=innoDB;

select * from daftar_pelanggan;
desc daftar_pelanggan;

/* 
untuk tipe data lainm buat tabel baru
*/
create table daftar_pelanggan_2
(
id		VARCHAR(100) not null,
nama	VARCHAR(100) not null,
email 	VARCHAR(100) not null,
balance integer	default 0,
rating	double default 0.0,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
DOB	DATE,
sex SET ('L','P','') DEFAULT NULL,
primary key(id)
)engine=innoDB;

desc daftar_pelanggan_2;

insert into daftar_pelanggan_2 (id, nama, email, balance, rating, DOB, sex) 
VALUES
('P0001','contoh nama 1','nama1@email.com',1000000,4.8,'1991-1-1','L'),
('P0002','contoh nama 2','nama2@email.com',1000000,4.3,'1992-2-2','P'),
('P0003','contoh nama 3','nama3@email.com',1000000,4.9,'1993-3-3','P'),
('P0004','contoh nama 4','nama4@email.com',1000000,3.6,'1994-4-4','L'),
('P0005','contoh nama 5','nama5@email.com',1000000,2.8,'1995-5-5','P');

select * from daftar_pelanggan_2;

drop table daftar_pelanggan_2;

CREATE TABLE komentar_pelanggan
(
    id INT NOT NULL auto_increment,
    email VARCHAR(100) NOT NULL,
    isi_komentar TEXT,
    primary key (id)
) engine=InnoDB;

select * from komentar_pelanggan;
DESCRIBE komentar_pelanggan;