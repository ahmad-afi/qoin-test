---
--- PostgreSQL database dump
---
CREATE TABLE pelanggan (
    id_pelanggan SERIAL PRIMARY KEY,
    nama_pelanggan varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    no_telepon varchar(255) NOT NULL UNIQUE
);

CREATE TABLE menu (
    id_menu SERIAL PRIMARY KEY,
    nama_menu varchar(255) NOT NULL,
    harga_menu INTEGER NOT NULL,
    deskripsi_menu text NOT NULL
);


CREATE TABLE rekomendasi_menu (
    id_rekomendasi SERIAL PRIMARY KEY,
    id_menu INTEGER REFERENCES menu(id_menu),
    jumlah_dipesan INTEGER NOT NULL
);

CREATE TABLE stok_menu (
    id_stok_menu SERIAL PRIMARY KEY,
    id_menu INTEGER REFERENCES menu(id_menu),
    stok INTEGER NOT NULL
);

CREATE TABLE bahan_baku (
    id_bahan_baku SERIAL PRIMARY KEY,
    nama_bahan_baku varchar(255) NOT NULL,
    stok_bahan_baku INTEGER NOT NULL
);

CREATE TABLE diskon (
    kode_diskon SERIAL PRIMARY KEY,
    jenis_diskon varchar(255) NOT NULL,
    nominal_diskon INTEGER NOT NULL,
    berlaku_hingga DATE NOT NULL
);

CREATE TABLE ulasan (
    id_ulasan SERIAL PRIMARY KEY,
    id_pelanggan INTEGER REFERENCES pelanggan(id_pelanggan),
    id_menu INTEGER REFERENCES menu(id_menu),
    ulasan varchar(255) NOT NULL,
    rating INTEGER NOT NULL
);


CREATE TABLE pesanan (
    id_pesanan SERIAL PRIMARY KEY,
    id_pelanggan INTEGER REFERENCES pelanggan(id_pelanggan),
    tanggal_pesanan DATE NOT NULL,
    jenis_pesanan varchar(255) NOT NULL,
    status_pesanan varchar(255) NOT NULL,
    kode_diskon INTEGER REFERENCES diskon(kode_diskon)
);


CREATE TABLE detail_pesanan (
    id__detail_pesanan SERIAL PRIMARY KEY,
    id_pesanan INTEGER REFERENCES pesanan(id_pesanan),
    id_menu INTEGER REFERENCES menu(id_menu),
    jumlah_pesanan INTEGER NOT NULL,
    harga_menu INTEGER NOT NULL
);

CREATE TABLE laporan_penghasilan (
    id_laporan SERIAL PRIMARY KEY,
    jenis_laporan varchar(255) NOT NULL,
    tanggal_laporan DATE NOT NULL,
    total_penghasilan INTEGER NOT NULL
);

CREATE TABLE pembayaran (
    id_pembayaran SERIAL PRIMARY KEY,
    id_pesanan INTEGER REFERENCES pesanan(id_pesanan),
    jumlah_bayar INTEGER NOT NULL,
    metode_pembayaran varchar(255) NOT NULL
);