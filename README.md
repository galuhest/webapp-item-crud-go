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