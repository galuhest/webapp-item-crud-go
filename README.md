## Description
Simple webapp built using go

## SLO and SLI

## Architecture Diagram

## Owner
Estya

## Contact and On-Call Information
call 
* 1
* 2
* 3

## Links
www.google.com

## Onboarding and Development Guide
Bagian ini dibuat agar developer baru dapat dengan mudah mulai berkontribusi pada service ini. Terdapat dua bagian penting yang setidaknya harus dibuat. Pertama adalah langkah-langkah untuk menjalankan service, mulai dari cara checkout code dari github, konfigurasi environment, menjalankan service, hingga verfikasi apakah service tersebut berjalan dengan benar (termasuk semua command atau script yang perlu dijalankan untuk mencapai hal tersebut).  Bagian kedua adalah langkah-langkah untuk mulai berkontribusi pada service, mulai dari checkout code, bagaimana melakukan perubahan code, bagaimana membuat unit test, bagaimana menjalankan test, bagaimana commit perubahan tersebut, bagaimana meminta review code, bagaimana menjamin bahwa service berhasil dibuat dan di release dengan benar, hingga bagaimana caranya deploy service tersebut.

## Request Flows, Endpoints, and Dependencies

## On-Call Runbooks

## FAQ

## REQUIREMENTS
* github.com/joho/godotenv
* github.com/go-sql-driver/mysql
* github.com/julienschmidt/httprouter

## Route
### root
	GET /
### Get item name
	GET /item/{id}
### Create Item
	POST /item
	Body :
	name
### Change item name
	POST /item/{id}
	Body :
	name
### Delete item
	POST /item/{id}/delete

## Crudmd
### Command structure
	crudmd -env=path/to/.env crud param1 param2
Default value for -env is current folder
### Crud command
	get id
	create name
	update id name
	delete id
id must be in int form