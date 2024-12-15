# Julo backend case study

## prerequisite
```
golang
doker compose
```
## Initiate The Project

```
go mod tidy
go mod vendor
```

## Running

To run the project, run the following command:

```
docker-compose up --build
```

You should be able to access the API at http://localhost:8080

If you change `database.sql` file, you need to reinitate the database by running:

```
docker-compose down --volumes
docker-compose up --build
```

## Postman collection
for all the api in this project you can import it to your postman using postman collections bellow
julo.postman_collection.json
