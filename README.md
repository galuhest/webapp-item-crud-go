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