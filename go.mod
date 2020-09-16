module github.com/unimal-jp/benkyoukai-go

go 1.14

replace local.packages/DB => ./DB

replace local.packages/Controllers => ./Controllers

replace local.packages/Models => ./Models

require (
	github.com/djimenez/iconv-go v0.0.0-20160305225143-8960e66bd3da // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.17
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/lib/pq v1.1.1
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	local.packages/Controllers v0.0.0-00010101000000-000000000000
	local.packages/DB v0.0.0-00010101000000-000000000000
	local.packages/Models v0.0.0-00010101000000-000000000000 // indirect

)
