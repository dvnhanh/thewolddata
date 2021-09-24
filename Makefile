# run test
test-integration:
	go test internal/repos/mysqlrepo/mysql.go internal/repos/mysqlrepo/intergration_test.go
test-sqlmock:
	go test internal/repos/mysqlrepo/mysql.go internal/repos/mysqlrepo/sqlmock_test.go
# build project
build:
	echo "****** THE WORLD DATA ******"
	go build -o api.exe cmd/main.go
# run program (file.exe)
run:
	./api
