module github.com/Garage-Standard-Inc/roomap-golib/repository

go 1.19

require (
	github.com/Garage-Standard-Inc/roomap-golib/os v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
)

require github.com/joho/godotenv v1.4.0 // indirect

replace github.com/Garage-Standard-Inc/roomap-golib/os => ../os
