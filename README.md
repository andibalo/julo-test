# julo-Task #
This repo is my submission for julo backend engineer skill test. Fingers crossed!
# Table of Contents
* [Prequisites](#prequisites)
* [Installation](#installation)
	* [DB Migration](#db-migration)
* [Postman Collection](#postman-collection)
* [App](#app)
	* [Environment Variables](#environment-variables)
	* [Endpoints](#api)
	
## Prequisites
* Docker
* DB Client ([DBeaver](https://dbeaver.io/), [Tableplus](https://tableplus.com/), etc.)
* go 1.18 (optional, used to run directly in local machine)

## Postman Collection
The postman collection json for this monorepo can be found at `/postman`. You can download it and import it on your local postman application.

## Installation
To run the app, you need to have `docker` installed in your machine or `go1.18` if you want to run the app directly in your local machine. Ensure you have **nothing** on port **5050** (App) and port **3306** (MySQL) running because the app will be using that port. Then go to root directory and run the following commands.

```
$ docker-compose up
```

If you are facing issue, try running docker-compose with sudo permission
```
$ sudo docker-compose up --build
```

Before testing out the APIs you will need to set up the tables in the DB

### DB Migration
Running the docker compose up will create the default db `julo_db`. To setup the DB tables you can use a DB client.

The database we are using is `MySQL` and it will be running at PORT `3306`, you can use a DB client to connect to the DB on PORT 3306 and use the **username** `root` with the **password** `test_pass` as credentials. 

After connecting to the DB, run the following sql to create the initial table.

```
CREATE TABLE `cake` (
    `id` VARCHAR(255) NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `rating` FLOAT(20) NOT NULL,
    `image` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
```

(Optional) Run the following SQL to create dummy data
```
INSERT INTO cake (id,title,description,rating,image) VALUES
 ("73c30425-1af5-4095-bf88-4cf21841dfd5","Cheesecake","Best cheesecake in the world", 9.2, "http://cakeimage.com"),
 ("23c30425-13f5-4295-df88-4cf21841dfd5","Mooncake","Best mooncake in the world", 8.2, "http://cakeimage.com"),
 ("d3c30425-1335-1095-ce38-3cf21841dfd5","Kue Lapis","Best kue lapis in the world", 9.8, "http://cakeimage.com"),
 ("e3c30425-a335-1295-ce38-3cf21e41dfd5","Red velvet","Best red velvet in the world", 8.4, "http://cakeimage.com"),
 ("77e6db98-12c5-4d4f-bebb-f92ef686f1ea","Birthday Cake","Best birthday cake in the world", 7.8, "http://cakeimage.com"),
 ("036d041a-d998-4492-9894-b5bd59a348d3","Sweet cake","Best sweet cake in the world", 8.8, "http://cakeimage.com"),
 ("421db05a-5edc-4c43-bd31-ae0744501888","Chocolate cake","Best chocolate cake in the world", 7.5, "http://cakeimage.com");
```

And you're done! You can now start testing out the APIs

## App
The app is using using the echo golang framework as well as MySQL 8.0 for the database.

### Environment Variables
* `SERVER_PORT` defines which port the application will listen to 
	* The default port is `5050`
* `STORAGE_DSN` connection string for connecting to the db
	* The default value is `root:test_pass@tcp(julo-db:3306)/julo_db`
	
### API
  - [Health Check](#get-health---health-check)
  - [Create Cake](#post-cakes---create-cake)
  - [Get All Cake](#get-cakespage1perpage5---get-cakes)
  - [Get Cake Detail](#get-cakesid---get-cake-detail)
  - [Delete Cake](#delete-cakesid---delete-cake)
  - [Update Cake](#patch-cakesid---update-cake)
  
  
### GET /health - Health Check
This endpoint can be used to verify that the app is running

### Response
```
ok
```

### POST /cakes - Create Cake
This endpoint will validate the request and insert cake data into the database. 

### Request
```
//HTTP Request Body (JSON)
{
    "title": "best-titel",  	// Required, Max 50 Characters
    "description" : "cdcddcdc", // Required, Max 255 Characters
    "rating": 3, 		// Required, Must be between 1 and 10
    "image" :"test-crece" 	// Required, Max 255 Characters
}
```
### Response
```
//HTTP Response (Application/JSON)
{
    "code": "AU0000",
    "status": "SUCCESS",
    "message": "Successfully created cake",
    "data": null
}
```

### PATCH /cakes/:id - Update Cake
This endpoint will validate the request and update cake data in database based on id passed to the route param. 

### Request
```
//HTTP Request Body (JSON)
{
    "title": "best-titel",  	// Required, Max 50 Characters
    "description" : "cdcddcdc", // Required, Max 255 Characters
    "rating": 3, 		// Required, Must be between 1 and 10
    "image" :"test-crece" 	// Required, Max 255 Characters
}
```
### Response
```
//HTTP Response (Application/JSON)
{
    "code": "AU0000",
    "status": "SUCCESS",
    "message": "Successfully updated cake",
    "data": null
}
```

### DELETE /cakes/:id - Delete Cake
This endpoint will delete cake data in database based on id passed to the route param. 

### Response
```
//HTTP Response (Application/JSON)
{
    "code": "AU0000",
    "status": "SUCCESS",
    "message": "Successfully deleted cake",
    "data": null
}
```

### GET /cakes/:id - Get Cake Detail
This endpoint will get cake detail data in database based on id passed to the route param. 

### Response
```
//HTTP Response (Application/JSON)
{
    "code": "AU0000",
    "status": "SUCCESS",
    "message": "Successfully deleted cake",
    "data": {
        "id": "uuid",
        "title": "best-titel",
        "description": "cdcddcdc",
        "rating": 3,
        "image": "test-crece",
        "created_at": "2022-10-02T19:19:34Z",
        "updated_at": "2022-10-02T19:19:34Z"
    }
}
```

### GET /cakes?page=1&perPage=5 - Get Cakes
This endpoint will get a list of cakes **sorted by** *rating ascending* and *created at date descending*. It will also take query params for pagination purposes.
* `page` represents the current page of data. Default value is 1
* `perPage` represents the limit of data per page. Default value is 5


### Response
```
//HTTP Response (Application/JSON)
{
    "code": "AU0000",
    "status": "SUCCESS",
    "message": "Successfully deleted cake",
     "data": {
        "cakes": [
	  {
		"id": "uuid",
		"title": "best-titel",
		"description": "cdcddcdc",
		"rating": 3,
		"image": "test-crece",
		"created_at": "2022-10-02T19:19:34Z",
		"updated_at": "2022-10-02T19:19:34Z"
	 },
	 {
		"id": "uuid",
		"title": "best-titel",
		"description": "cdcddcdc",
		"rating": 3,
		"image": "test-crece",
		"created_at": "2022-10-02T19:19:34Z",
		"updated_at": "2022-10-02T19:19:34Z"
	 }
	],
        "total_record": 7,
        "current_page": 5,
        "total_page": 3
    }
}
```
