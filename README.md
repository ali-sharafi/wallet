# wallet

## Introduction

The wallet is responsible for managing clients' wallets in a system.

## Installation

First of all to install wallet dependencies you should run:

```bash
go get .
```
Then for add your custom config file run:

```bash
cd conf && cp app.example.ini app.ini
```

At the end just run:
```bash
go run .
```
The wallet is ready to use.

Available endpoints:

## Authentication

### POST /auth

Get Authentication token.

+ Request (application/json)
```
{
    "username": "test@test.com"
    "password": "12345678"
}
```

+ Response 200 (application/json)
```
{
  "code": 0,
  "msg": "Success",
  "data": "JWT TOKEN"
}
```
## Wallets

### GET /api/v1/wallets
List of wallets

+ Request headers
```
{
    "Authorization": "JWT TOKEN"
}
```
+ Response 200 (application/json)
```
{
  "code": 0,
  "msg": "Success",
  "data": [
    {
      "id": 1,
      "balance": "1000",
      "created_at": "2022-05-17 07:41:38",
      "updated_at": "2022-05-17 07:41:38"
    },
    .
    .
    .
}
```
### GET /api/v1/wallets/:id/balance
Get a wallet balance

+ Request headers
```
{
    "Authorization": "JWT TOKEN"
}
```
+ Response 200 (application/json)
```
{
  "code": 0,
  "msg": "Success",
  "data": "1000"
}
```

### POST /api/v1/wallets/:id/credit
Add credit to a wallet

+ Request headers
```
{
    "Authorization": "JWT TOKEN"
}
```
+ Request body
```
{
    "amount": "10.01"
}
```
+ Response 200 (application/json)
```
{
  "code": 0,
  "msg": "Success",
  "data": "1010.01"
}
```
### POST /api/v1/wallets/:id/debit
Add debit to a wallet

+ Request headers
```
{
    "Authorization": "JWT TOKEN"
}
```
+ Request body
```
{
    "amount": "0.01"
}
```
+ Response 200 (application/json)
```
{
  "code": 0,
  "msg": "Success",
  "data": "1010"
}
```