module example/main

go 1.17

replace example/factory => ../factory

replace example/connection => ../connection

require example/factory v0.0.0-00010101000000-000000000000

require (
	example/connection v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/lib/pq v1.10.3 // indirect
)
