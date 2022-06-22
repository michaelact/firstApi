# First API: dengan Golang

Cara menjalankan:
1. [Install Golang](https://go.dev/doc/install)
2. Pastikan Golang 1.18
3. Siapkan database MySQL (saya menggunakan versi 8)
4. Eksekusi file-file `*.sql` dalam folder `database/migrations/*.sql` ke database MySQL
5. Siapkan environment variable dibawah ini:
   - MYSQL_PORT
   - MYSQL_PASSWORD
   - MYSQL_HOST
   - MYSQL_USER
   - MYSQL_DATABASE
   - SERVER_PORT
   - SERVER_HOST
   - SERVER_API_KEY
6. `go run main.go`

Referensi:
- https://documenter.getpostman.com/view/7918444/UVJhEabr
- https://blog.gethired.id/devcode-challenge-2nd/
