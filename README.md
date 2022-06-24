# First API: dengan Golang

Cara menjalankan:
1. [Install Golang](https://go.dev/doc/install)
1. Pastikan Golang 1.18
1. `go get ./...`
1. Siapkan database MySQL (saya menggunakan versi 8)
1. Eksekusi file-file `*.sql` dalam folder `database/migrations/*.sql` ke database MySQL
1. Siapkan environment variable dibawah ini:
   - MYSQL_PORT
   - MYSQL_PASSWORD
   - MYSQL_HOST
   - MYSQL_USER
   - MYSQL_DATABASE
   - SERVER_PORT
   - SERVER_HOST
   - API_KEY
   - LOG_LEVEL [List](https://github.com/rs/zerolog/blob/master/log.go#L112-L133)
1. `go run .`

Referensi:
- https://documenter.getpostman.com/view/7918444/UVJhEabr
- https://blog.gethired.id/devcode-challenge-2nd/
