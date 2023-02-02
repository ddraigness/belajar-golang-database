package belajargolangdatabase

/**
- Pengenalan Package Database
    1. bahasa pemrograman go-lang secara default memiliki sebuah package bernama database
    2. package database adalah package yang berisikan kumpulan standard interface yang menjadi standard untuk berkomunikasi ke database
    3. hal ini menjadikan kode program yang kita buat untuk mengakses jenis database apapun bisa menggunakan kode yang sama
    4. yang berbeda hanya kode SQL yang perlu kita gunakan sesuai dengan database yang kita gunakan

- cara kerja package database
aplikasi, call -> database interface, call -> database driver, call -> DBMS

- database driver
    1. sebelum kita membuat kode program menggunakan database di go-lang, terlebih dahulu kita wajib menambahkan driver database nya
    2. tanpa driver database, maka package database di go-lang tidak mengerti apapun, karena hanya berisi kontrak interface saja
    3. https://golang.org/s/sqldrivers (Package : github.com/go-sql-driver/mysql)

- membuat koneksi ke database
    1. hal yang pertama akan kita lakukan ketika membuat aplikasi yang akan menggunakan database adalah melakukan koneksi ke database nya
    2. untuk melakukan koneksi ke database di golang, kita bisa membuat object sql.DB menggunakan function sql.Open(driver, dataSourceName)
    3. untuk menggunakan database MySQL, kita bisa menggunakan driver mysql
    4. sedangkan untuk dataSourceName, tiap database biasanya punya cara penulisan masing-masing, misal di mysql, kita bisa menggunakan dataSourceName seperti ini : 
    ~ username:password@tcp(host:port)/database_name
    5. jika object sql.DB sudah tidak digunakan lagi, disarankan untuk menutupnya menggunakan function Close()

- database pooling
    1. sql.DB di golang sebenarnya bukanlah sebuah koneksi ke database
    2. melainkan sebuah pool ke database, atau dikenal dengan konsep Database Pooling
    3. di dalam sql.DB, golang melakukan management koneksi ke database secara otomatis. Hal ini menjadikan kita tidak perlu melakukan management koneksi database secara manual
    4. dengan kemampuan database pooling ini, kita bisa menentukan jumlah minimal dan maksimal koneksi yang dibuat oleh golang, sehingga tidak membanjiri koneksi database, karena biasanya ada batas maksimal koneksi yang bisa ditangani oleh database yang kita gunakan

- pengaturan database pooling
    1. (DB) SetMaxIdleConns(number) : Pengaturan berapa jumlah koneksi minimal yang dibuat
    2. (DB) SetMaxOpenConns(number) : Pegaturan berapa jumlah koneksi maksimal yang dibuat
    3. (DB) SetConnMaxIdleTime(duration) : Pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus
    4. (DB) SetConnMaxLifetime(duration) : Pengaturan berapa lama koneksi boleh digunakan

- eksekusi perintal SQL
    1. saat membuat aplikasi menggunakan database, sudah pasti kita ingin berkomunikasi dengan database menggunakan perintah SQL
    2. di golang juga menyediakan function yang bisa kita gunakan untuk mengirim perintah SQL ke database menggunakan function (DB) ExecContext(context, sql, params)
    3. ketika mengirim perintah SQL, kita butuh mengirimkan context, dan seperti yang sudah pernah kita pelajari di course golang context, dengan context, kita bisa mengirim sinyal cancel jika kita ingin membatalkan pengiriman perintah SQL nya

- Query SQL
    1. untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan perintah Exec, namun jika kita membutuhkan result, seperti SELECT SQL, kita bisa menggunakan function yang berbeda
    2. function untuk melakukan query ke database, bisa menggunakan function (DB) QueryContext(context, sql, params)

- Rows
    1. hasil Query function adalah sebuah data struct sql.Rows
    2. Rows digunakan untuk melakukan iterasi terhadap hasil dari query
    3. kita bisa menggunakan function (Rows) Next() (boolean) untuk melakukan iterasi terhadap data hasil query, jika return data false, artinya sudah tidak ada data lagi di dalam result
    4. untuk membaca tiap data, kita bisa menggunakan (Rows) Scan(columns...)
    5. setelah menggunakan Rows, jangan lupa untuk menutupnya menggunakan (Rows) Close()

- Tipe Data Column
    1. Sebelumnya kita hanya membuat table dengan tipe data di kolom nya berubah VARCHAR
    2. untuk VARCHAR di database, biasanya kita gunakan String di Golang
    3. bagaimana dengan tipe data yang lain?
    4. apa representasinya di golang, misal tipe data timestamp, date dan lain-lain

- mapping tipe data
    tipe data database(tipe data golang) :
    1. varchar, char (string)
    2. int, bigint (int32,int64)
    3. float, double (float32, float64)
    4. boolean (bool)
    5. date, datetime, time, timestamp (time.Time)

- error tipe data Date
    1. secara default, Driver MySQL untuk golang akan melakukan query tipe data DATE, DATETIME, TIMESTAMP menjadi []byte / []uint8. Dimana ini bisa dikonversi menjadi String, lalu diparsing menjadi time.Time
    2. namun hal ini merepotkan jika dilakukan manual, kita bisa meminta Driver MySQL untuk Golang secara otomatis melakukan parsing dengan menambahkan parameter parseDate=true

- Nullable Type
    1. Golang database tidak mengerti dengan tipe data NULL di database
    2. oleh karena itu, khusus untuk kolom yang bisa NULL di database, akan jadi masalah jika kita melakukan secara bulat-bulat menggunakan tipe data representasinya di Golang

- error data null
    1. konversi secara otomatis NULL tidak didukung oleh Driver MySQL golang
    2. oleh karena itu, khusus tipe kolom yang bisa NULL, kita perlu menggunakan tipe data yang ada di dalam package sql

- tipe data nullable
    tipe data golang, (tipe data nullable) :
    1. string, (database/sql.NullString)
    2. bool, (database/sql.NullBool)
    3. float64, (database/sql.NullFloat64)
    4. int32, (database/sql.NullInt32)
    5. int64, (database/sql.NullInt64)
    6. time.Time, (database/sql.NullTime)

- SQL dengan Parameter
    1. Saat membuat aplikasi, kita tidak mungkin akan melakukan hardcode perintah SQL di kode golang kita
    2. biasanya kita akan menerima input data dari user, lalu membuat perintah SQL dari input user, dan mengirimnya menggunakan SQL

- SQL Injection
    1. SQL Injection adalah sebuah teknik yang menyalahgunakan sebuah celah keamanan yang terjadi dalam lapisan basis data sebuah aplikasi
    2. Biasa, SQL Injection dilakukan dengan mengirim input dari user dengan perintah yang salah, sehingga menyebabkan hasil SQL yang kita buat menjadi tidak valid
    3. SQL Injection sangat berbahaya, jika sampai kita salah membuat SQL, bisa jadi data kita tidak aman

- Solusinya?
    1. Jangan membuat query SQL secara manual dengan menggabungkan String secara bulat-bulat
    2. jika kita membutuhkan parameter ketika membuat SQL, kita bisa menggunakan function Execute atau Query dengan parameter yang akan kita bahas di chapter selanjutnya

- SQL dengan Parameter
    1. sekarang kita sudah tahu bahaya nya SQL Injection jika menggabungkan string ketika membuat query
    2. jika ada kebutuhan seperti itu, sebenarnya function Exec dan Query memiliki parameter tambahan yang bisa kita gunakan untuk mensubtitusi parameter dari function tersebut ke SQL query yang kita buat
    3. untuk menandai sebuah SQL membutuhkan parameter, kita bisa gunakan karakter '?' (tanda tanya)

- Contoh SQL
    1. SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1
    2. INSERT INTO user(username, password) VALUES (?, ?)
    3. Dan lain-lain

- Auto Increment
    1. kadang kita membuat sebuah table dengan id auto increment
    2. dan kadang pula, kita ingin mengambil data id yang sudah kita insert ke dalam MySQL
    3. sebenarnya kita bisa melakukan query ulang ke database menggunakan SELECT LAST_INSERT_ID()
    4. Tapi untungnya di golang ada cara yang lebih mudah
    5. kita bisa menggunakan function (Result) LastInsertId() untuk mendapatkan ID terkahir yang dibuat secara auto increment
    6. Result adalah object yang dikembalikan ketika kita menggunakan function Exec

- Prepare Statement
    - Query atau Exec dengan Parameter
        1. saat kita menggunakan Function Query atau Exec yang menggunakan parameter, sebenarnya implementasi dibawahnya menggunakan Prepare Statement
        2. jadi tahapan pertama statement nya disiapkan terlebih dahulu, setelah itu baru di isi dengan parameter
        3. kadang ada kasus kita ingin melakukan beberapa hal yang sama sekaligus, hanya berbeda parameternya, Misal insert data langsung banyak
        4. pembuatan Prepare Statement bisa dilakukan dengan manual, tanpa harus menggunakan Query atau Exec dengan parameter

    - Prepare Statement part 2
        1. saat kita membuat Prepare Statement, secara otomatis akan mengenali koneksi database yang digunakan
        2. sehingga ketika kita mengeksekusi Prepare Statement berkali-kali, maka akan menggunakan koneksi yang sama dan lebih efisien karena pembuatan prepare statementnya hanya sekali di awal saja
        3. jika menggunakan Query dan Exec dengan parameter, kita tidak bisa menjamin bahwa koneksi yang digunakan akan sama, oleh karena itu, bisa jadi prepare statement akan selalu dibuat berkali-kali walaupun kita menggunakan SQL yang sama
        4. untuk membuat Prepare Statement, kita bisa menggunakan function (DB) Prepare(context, sql)
        5. Prepare Statement di representasikan dalam struct database/sql.Stmt
        6. sama seperti resource sql lainnya, Stmt hatus di Close() jika suda tidak digunakan lagi

    - database transaction
        1. salah satu fitur andalan di database adalah transaction
        2. materi database transaction sudah saya bahas tuntas di materi MySQL database, jadi silahkan pelajari di course tersebut
        3. di course ini kita akan fokus bagaimana menggunakan database transaction di golang

    - transaction di golang
        1. secara default, semua perintah SQL yang kita kirim menggunakan Golang akan otomatis di commit, atau istilahnya auto commit
        2. namun kita bisa menggunakan fitur transaksi sehingga SQL yang kita kirim tidak secara otomatis di commit ke database
        3. untuk memulai transaksi, kita bisa menggunakan function (DB) Begin(), dimana akan menghasilkan struct Tx yang merupakan representasi Transaction
        4. struct Tx ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana hampir semua function di DB ada di Tx, seperti Exec, Query atau Prepare
        5. setelah selesai proses transaksi, kita bisa gunakan function (Tx) Commit() untuk melakukan commit atau Rollback()

    - Repository Pattern
        1. dalam buku Domain-Driven Design, Eric Evans menjelaskan bahwa "repository is a mechanism for encapsulating storage, retrieval, and search behavior, which emulates a collection of object"
        2. Pattern Repository ini biasanya digunakan sebagai jembatan antar business logic aplikasi kita dengan semua perintah SQL ke database
        3. jadi semua perintah SQL akan ditulis di Repository, sedangkan business logic kode program kita hanya cukup menggunakan Repository tersebut

- Entity / Model
    1. dalam pemrograman berorientasi object, biasanya sebuah tabel di database akan selalu dibuat representasinya sebagai class Entity atau Model, namun di Golang, karena tidak mengenal Class, jadi kita akan representasikan data dalam bentuk struct
    2. ini bisa mempermudah ketika membuat kode program
    3. misal ketika kita query ke repository, dibanding mengembalikan array, alangkah baiknya repository melakukan konversi terlebih dahulu ke struct Entity / Model, sehingga kita tinggal menggunakan objectnya saja

*/