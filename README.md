# Julo-Task #
This repo is my submission for julo backend engineer skill test. Fingers crossed!
# Table of Contents
* [Prequisites](#prequisites)
* [Installation](#installation)
	* [DB Migration](#db-migration)
* [Postman Collection](#postman-collection)
* [App](#app)
	* [Environment Variables](#environment-variables)
	* [Redis Use Case](#redis-use-case)
	* [Endpoints](#api)
	
## Prequisites
* Docker
* DB Client (Optional) ([DBeaver](https://dbeaver.io/), [Tableplus](https://tableplus.com/), etc.) 
* go 1.18 (optional, used to run directly in local machine)

## Postman Collection
The postman collection json for this repo can be found at `/postman`. You can download it and import it on your local postman application.

## Installation
To run the app, you need to have `docker` installed in your machine or `go1.18` if you want to run the app directly in your local machine. Ensure you have **nothing** on port **5050** (App) and port **3306** (MySQL) and port **6379** (Redis) running because the app will be using those ports. Then go to root directory and run the following commands.

```
$ docker-compose up
```

If you are facing issue, try running docker-compose with sudo permission
```
$ sudo docker-compose up --build
```

### DB Migration
Running the docker compose up will create the default db `julo_db` as well as **automatically** creating the required tables.

### Viewing the database
The database we are using is `MySQL` and it will be running at PORT `3306`, you can use a DB client to connect to the DB on PORT 3306 and use the **username** `root` with the **password** `test_pass` as credentials. 

## App
The app is using using the `Echo` golang framework as well as `MySQL 8.0` for the database and `Redis` for caching

### Environment Variables
* `SERVER_PORT` defines which port the application will listen to 
	* The default port is `5050`
* `STORAGE_DSN` connection string for connecting to the db
	* The default value is `root:test_pass@tcp(julo-db:3306)/julo_db`
* `REDIS_HOST` redis host we will be connecting to
	* The default value is `julo-redis:6379`
* `JWT_SECRET` secret we will use to sign the jwt token with
	* The default value is `jwtseecreet`
	
### Redis Use Case
1. Get User Wallet Balance

We will be caching the user wallet balance and info with a ttl of `5 seconds`. The key prefix is `usr-wlt-bln-{{cutomer_xid}}`.

	
### API
  - [Health Check](#get-health---health-check)
  - [Init Wallet](#post-apiv1init---init-wallet)
  - [Enable Wallet](#post-apiv1wallet---enable-wallet)
  - [Disable Wallet](#patch-apiv1wallet---disable-wallet)
  - [View Wallet Balance](#get-apiv1wallet---get-wallet-balance)
  - [Deposit To Wallet](#post-apiv1walletdeposits---deposit-to-wallet)
  - [Withdraw From Wallet](#post-apiv1walletwithdrawals---withdraw-from-wallet)
  
### GET /health - Health Check
This endpoint can be used to verify that the app is running

### Response
```
ok
```

### POST /api/v1/init - Init Wallet
This endpoint will create a wallet using a customer_xid and return a jwt token. If a wallet based on a customer_xid already exists it will just return a jwt token.

### Request
```
//HTTP Request Body (JSON)
{
    "customer_xid" : "test" //Required
}
```
### Response
```
//HTTP Response (Application/JSON)
{
    "code": "AU0000",
    "status": "SUCCESS",
    "message": "Successfully initialized wallet",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjeGlkIjoidGVzdC0yIn0.NcwgB_WEaQi4nl3tLUVFf2w_a8LZhQXs5LbCvuPhTYo"
    }
}
```

### POST /api/v1/wallet - Enable Wallet
This endpoint will update wallet status to enabled to be able to view balance, make deposit and withdraw.

### Request

**Headers**

Authorization : Bearer {{token}}

### Response
```
//HTTP Response (Application/JSON)
{
    "code": "AU0000",
    "status": "SUCCESS",
    "message": "Successfully enabled wallet",
    "data": {
        "wallet": {
            "id": "9a8bf6de-808d-4266-a194-18b6e930e043",
            "owned_by": "test-2",
            "status": "enabled",
            "enabled_at": "2022-10-08T12:42:53.792Z",
            "balance": 0
        }
    }
}
```

If wallet is already enabled
```
{
    "message": "wallet needs to be disabled to use this resource"
}
```


### PATCH /api/v1/wallet - Disable Wallet
This endpoint will update wallet status to disabled. (Requires JWT & Wallet Enabled)

### Request

**Headers**

Authorization : Bearer {{token}}

### Response
```
//HTTP Response (Application/JSON)
{
    "code": "AU0000",
    "status": "SUCCESS",
    "message": "Successfully disabled wallet",
    "data": {
        "wallet": {
            "id": "9a8bf6de-808d-4266-a194-18b6e930e043",
            "owned_by": "test-2",
            "status": "disabled",
            "disabled_at": "2022-10-08T12:48:01.625Z",
            "balance": 0
        }
    }
}
```

If wallet is already disabled
```
{
    "message": "wallet needs to be enabled to use this resource"
}
```

### GET /api/v1/wallet - Get Wallet Balance
This endpoint will get wallet info and balance from the cache first. If it doesn't exist in the cached it will fetch from the db. (Requires JWT & Wallet Enabled)

### Request

**Headers**

Authorization : Bearer {{token}}

### Response
```
//HTTP Response (Application/JSON)
{
    "code": "AU0000",
    "status": "SUCCESS",
    "message": "Successfully fetched wallet",
    "data": {
        "wallet": {
            "id": "9a8bf6de-808d-4266-a194-18b6e930e043",
            "owned_by": "test-2",
            "status": "enabled",
            "enabled_at": "2022-10-08T12:49:52.442Z",
            "balance": 0
        }
    }
}
```

### POST /api/v1/wallet/deposits - Deposit to wallet
This endpoint will get a update the wallet balance as well as create a record in the transaction table. (Requires JWT & Wallet Enabled)

### Request

**Headers**

Authorization : Bearer {{token}}

```
//HTTP Request Body (JSON)
{
    "amount": 30000, //Required
    "reference_id": "ref-desdede" //Required
}
```

### Response
```
//HTTP Response (Application/JSON)
{
    "code": "AU0000",
    "status": "SUCCESS",
    "message": "Successfully deposited to wallet",
    "data": {
        "deposit": {
            "id": "1ad5668a-b405-4adb-9af8-4525fe245729",
            "deposited_by": "test-2",
            "status": "success",
            "deposited_at": "2022-10-08T12:53:54.25029195Z",
            "amount": 30000,
            "reference_id": "ref-desdede"
        }
    }
}
```

### POST /api/v1/wallet/withdrawals - Withdraw from wallet
This endpoint will get a decrement the wallet balance based on the request amount as well as create a record in the transaction table. It will also return an error if the wallet balance is not enough. (Requires JWT & Wallet Enabled)

### Request

**Headers**

Authorization : Bearer {{token}}

```
//HTTP Request Body (JSON)
{
    "amount": 30000, //Required
    "reference_id": "ref-desdede" //Required
}
```

### Response
```
//HTTP Response (Application/JSON)
{
    "code": "AU0000",
    "status": "SUCCESS",
    "message": "Successfully withdrawn from wallet",
    "data": {
        "withdrawal": {
            "id": "ebf25e1a-fcc1-4faa-99e3-60b423d31fc7",
            "withdrawn_by": "test-2",
            "status": "success",
            "withdrawn_at": "2022-10-08T12:55:20.435068016Z",
            "amount": 500,
            "reference_id": "ref-12122d1"
        }
    }
}
```
