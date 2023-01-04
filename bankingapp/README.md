# Project: banking-app
Golang API for a banking application. It describes create/view a customer, open a bank account for a customer, withdraw money from bank account, deposit money to bank account, view balance for any account, view all the fixed accounts for a customer etc.

# Getting started with the project

## Prerequisite:

1. [Go](https://go.dev) should be installed.
1. [MySQL](https://dev.mysql.com/downloads/mysql/) should be installed. You can download community version from mysql website.

Please follow the steps bellow to run this project.

1. Execute sql statements provided in [ddl.sql](https://github.com/chauhansantosh/go-training/blob/dev/bankingapp/util/ddl.sql "ddl").
1. Execute ```go run main.go``` in the directory containing main.go file.


# 📁 Collection: Customer 

## End-point: Create customer success
Create cutomer successfully for the given paylod.
### Method: PUT
>```
>http://localhost:8080/customer
>```
### Body (**raw**)

```json
{
    "customer_id": 1255,
    "customer_name": "Amit",
    "customer_type": "INDIVIDUAL"
}
```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": false,
    "customer": {
        "customer_id": 123,
        "customer_name": "Santosh",
        "customer_type": "INDIVIDUAL"
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Create customer error - validation
Throw error if required fields are missing from the payload.
### Method: PUT
>```
>http://localhost:8080/customer
>```
### Body (**raw**)

```json
{
    "customer_id": 0,
    "customer_name": "",
    "customer_type": ""
}
```

### Response: 400
```json
{
    "statusCode": 400,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": true,
    "errorResponse": [
        {
            "errorMessage": "customer_id is a required field",
            "errorCode": "1002"
        },
        {
            "errorMessage": "customer_name is a required field",
            "errorCode": "1002"
        },
        {
            "errorMessage": "customer_type is a required field",
            "errorCode": "1002"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get customers success
Get all the customers created in database.
### Method: GET
>```
>http://localhost:8080/customers
>```
### Body (**raw**)

```json

```

### Response: 200
```json
{
    "statusCode": 200,
    "hasErrorResponse": false,
    "customers": [
        {
            "customer_id": 123,
            "customer_name": "Santosh",
            "customer_type": "INDIVIDUAL",
            "created_at": "2022-12-20T10:13:27Z",
            "updated_at": "2022-12-20T10:13:27Z"
        },
        {
            "customer_id": 234,
            "customer_name": "Deepika",
            "customer_type": "INDIVIDUAL",
            "created_at": "2022-12-20T10:24:09Z",
            "updated_at": "2022-12-20T10:24:09Z"
        },
        {
            "customer_id": 345,
            "customer_name": "HCL",
            "customer_type": "COMPANY",
            "created_at": "2022-12-20T10:24:47Z",
            "updated_at": "2022-12-20T10:24:47Z"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get customer by customer_id
Get customer by customer_id passed in path variable.
### Method: GET
>```
>http://localhost:8080/customers/1249
>```
### Body (**raw**)

```json

```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 1,
    "hasErrorResponse": false,
    "customers": [
        {
            "customer_id": 123,
            "customer_name": "Santosh",
            "customer_type": "INDIVIDUAL",
            "created_at": "2022-12-20T10:13:27Z",
            "updated_at": "2022-12-20T10:13:27Z"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: Accounts - Common 


## End-point: Create account success
Create an account successfully for the given payload.
### Method: PUT
>```
>http://localhost:8080/account
>```
### Body (**raw**)

```json
{
    "customer_id": 123,
    "account_id": 51691852,
    "account_type": "SAVINGS",
    "balance": 10000.5
}
```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": false,
    "bank_account": {
        "customer_id": 123,
        "account_id": 51691852,
        "account_type": "SAVINGS",
        "balance": 10000.5
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Create account error - customer does not exist
Throw an error if customer does not exist in database while creating bank account for the customer.
### Method: PUT
>```
>http://localhost:8080/account
>```
### Body (**raw**)

```json
{
    "customer_id": 234,
    "account_id": 51691853,
    "account_type": "SAVINGS",
    "balance": 20000
}
```

### Response: 500
```json
{
    "statusCode": 500,
    "timeElapsed": 51989,
    "hasErrorResponse": true,
    "errorResponse": [
        {
            "errorMessage": "Error 1644 (47000): Customer 234 does not exist. Create customer before creating account.",
            "errorCode": "1006"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Create account error - missing required fields
Throw an error if missing required fields while creating bank account.
### Method: PUT
>```
>http://localhost:8080/account
>```
### Body (**raw**)

```json
{
    "customer_id": 34567,
    "account_id": 0,
    "account_type": "",
    "balance": 0
}
```

### Response: 400
```json
{
    "statusCode": 400,
    "timeElapsed": 908374,
    "hasErrorResponse": true,
    "errorResponse": [
        {
            "errorMessage": "account_id is a required field",
            "errorCode": "1005"
        },
        {
            "errorMessage": "account_type is a required field",
            "errorCode": "1005"
        },
        {
            "errorMessage": "balance is a required field",
            "errorCode": "1005"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Deposit into account
Deposit money into the account passed in path variable and display the balance in response along with customer id and account id.
### Method: PUT
>```
>http://localhost:8080/accounts/51691852/deposit
>```
### Body (**raw**)

```json
{
   "amount": 100
}
```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": false,
    "bank_account": {
        "customer_id": 123,
        "account_id": 51691852,
        "balance": 10100.5
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Withdraw from account
Withdraw from the account passed in path variable and display the balance in response.
### Method: PUT
>```
>http://localhost:8080/accounts/51691852/withdraw
>```
### Body (**raw**)

```json
{
   "amount": 100
}
```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": false,
    "bank_account": {
        "customer_id": 123,
        "account_id": 51691852,
        "balance": 10000.5
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Withdraw from FD account
Withdraw from the account passed in path variable and display the balance in response.
### Method: PUT
>```
>http://localhost:8080/accounts/51691852/withdraw
>```
### Body (**raw**)

```json
{
   "amount": 100
}
```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": false,
    "bank_account": {
        "customer_id": 123,
        "account_id": 51691852,
        "balance": 10000.5
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get all accounts
Get all the accounts in the database.
### Method: GET
>```
>http://localhost:8080/accounts
>```
### Body (**raw**)

```json

```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 3,
    "hasErrorResponse": false,
    "bank_accounts": [
        {
            "customer_id": 123,
            "account_id": 51691852,
            "account_type": "SAVINGS",
            "balance": 10000.5,
            "created_at": "2022-12-20T10:17:31Z",
            "updated_at": "2022-12-20T10:17:31Z",
            "is_active": true
        },
        {
            "customer_id": 234,
            "account_id": 51691853,
            "account_type": "SAVINGS",
            "balance": 20000,
            "created_at": "2022-12-20T10:26:59Z",
            "updated_at": "2022-12-20T10:26:59Z",
            "is_active": true
        },
        {
            "customer_id": 345,
            "account_id": 51691854,
            "account_type": "CURRENT",
            "balance": 20000,
            "created_at": "2022-12-20T10:27:34Z",
            "updated_at": "2022-12-20T10:27:34Z",
            "is_active": true
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get account by account_id
Get account details for the account id passed in path variable.
### Method: GET
>```
>http://localhost:8080/accounts/51691852
>```
### Body (**raw**)

```json

```

### Response: 200
```json
{
    "statusCode": 200,
    "hasErrorResponse": false,
    "bank_accounts": [
        {
            "customer_id": 123,
            "account_id": 51691852,
            "account_type": "SAVINGS",
            "balance": 10000.5,
            "created_at": "2022-12-20T10:17:31Z",
            "updated_at": "2022-12-20T10:17:31Z",
            "is_active": true
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get all accounts for customer
Get all accounts for a customer. A customer can have more than one FD account.
### Method: GET
>```
>http://localhost:8080/accounts/customer/234/getallaccounts
>```
### Body (**raw**)

```json

```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 16,
    "hasErrorResponse": false,
    "bank_accounts": [
        {
            "customer_id": 234,
            "account_id": 51691853,
            "account_type": "SAVINGS",
            "balance": 20000,
            "created_at": "2022-12-20T10:26:59Z",
            "updated_at": "2022-12-20T10:26:59Z",
            "is_active": true
        },
        {
            "customer_id": 234,
            "account_id": 51691856,
            "account_type": "FIXED",
            "balance": 20000,
            "created_at": "2022-12-20T10:34:35Z",
            "updated_at": "2022-12-20T10:34:35Z",
            "lock_period_fd": 2,
            "is_locked": true,
            "is_active": true,
            "penalty_fd": 0.1,
            "locked_until": "2024-12-20T15:34:36Z"
        },
        {
            "customer_id": 234,
            "account_id": 51691857,
            "account_type": "FIXED",
            "balance": 100000,
            "created_at": "2022-12-20T10:32:23Z",
            "updated_at": "2022-12-20T10:32:23Z",
            "lock_period_fd": 1,
            "is_locked": true,
            "is_active": true,
            "penalty_fd": 0.1,
            "locked_until": "2023-12-20T15:32:23Z"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: Accounts - Savings 


## End-point: Deposit into savings account more than 50K
Throw error if pan number not provided while depositing more than 50K in savings account.
### Method: PUT
>```
>http://localhost:8080/accounts/5169201/deposit
>```
### Body (**raw**)

```json
{
   "amount": 50001,
   "account_pan": "ASHJK1653E"
}
```

### Response: 400
```json
{
    "statusCode": 400,
    "timeElapsed": 43243,
    "hasErrorResponse": true,
    "errorResponse": [
        {
            "errorMessage": "Error 1644 (46000): Pan number is mandatory for depositing more than 50000.",
            "errorCode": "1014"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Create SAVINGS Account
### Method: PUT
>```
>http://localhost:8080/account
>```
### Body (**raw**)

```json
{
    "customer_id": 1255,
    "account_id": 5169201,
    "account_type": "SAVINGS",
    "balance": 100000
}
```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": false,
    "bank_account": {
        "customer_id": 1255,
        "account_id": 5169201,
        "account_type": "SAVINGS",
        "balance": 100000
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: Accounts - Current 


## End-point: Create CURRENT Account
### Method: PUT
>```
>http://localhost:8080/account
>```
### Body (**raw**)

```json
{
    "customer_id": 1256,
    "account_id": 5169202,
    "account_type": "CURRENT",
    "balance": 350000,
    "odallowed": true,
    "odamount": 3500
}
```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": false,
    "bank_account": {
        "customer_id": 1256,
        "account_id": 5169202,
        "account_type": "CURRENT",
        "balance": 350000,
        "odallowed": true,
        "odamount": 3500
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Deposit in CURRENT Account
### Method: PUT
>```
>http://localhost:8080/accounts/5169202/deposit
>```
### Body (**raw**)

```json
{
    "amount": 3500
}
```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": false,
    "bank_account": {
        "customer_id": 1256,
        "account_id": 5169202,
        "balance": 353300
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Err: Deposit in CURRENT Account
### Method: PUT
>```
>http://localhost:8080/accounts/5169202/deposit
>```
### Body (**raw**)

```json
{
    "amount": 250001
}
```

### Response: 400
```json
{
    "statusCode": 400,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": true,
    "errorResponse": [
        {
            "errorMessage": "PAN is mandatory to deposit the amount more than 250000",
            "errorCode": "1014"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Withdraw from CURRENT Account
### Method: PUT
>```
>http://localhost:8080/accounts/5169202/withdraw
>```
### Body (**raw**)

```json
{
    "amount":100
}
```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": false,
    "bank_account": {
        "customer_id": 1256,
        "account_id": 5169202,
        "balance": 349800,
        "odallowed": true,
        "odamount": 3500
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Err: Withdraw from CURRENT Account
### Method: PUT
>```
>http://localhost:8080/accounts/5169202/withdraw
>```
### Body (**raw**)

```json
{
    "amount":500000
}
```

### Response: 400
```json
{
    "statusCode": 400,
    "timeElapsed": 9223372036854,
    "hasErrorResponse": true,
    "errorResponse": [
        {
            "errorMessage": "Insufficient fund. Withdraw amount cannot be covered by overdraft amount, please enter again",
            "errorCode": "1011"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: Accounts - Fixed 


## End-point: Create FIXED account error
Throw error if locking period was not provided while creating a fixed account.
### Method: PUT
>```
>http://localhost:8080/account
>```
### Body (**raw**)

```json
{
    "customer_id": 345,
    "account_id": 51691857,
    "account_type": "FIXED",
    "balance": 100000
}
```

### Response: 400
```json
{
    "statusCode": 400,
    "timeElapsed": 93389,
    "hasErrorResponse": true,
    "errorResponse": [
        {
            "errorMessage": "lock_period_fd is required for AccountType FIXED",
            "errorCode": "1005"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Create FIXED account  success
Create a fixed account successfully for the given valid payload.
### Method: PUT
>```
>http://localhost:8080/account
>```
### Body (**raw**)

```json
{
    "customer_id": 234,
    "account_id": 51691857,
    "account_type": "FIXED",
    "balance": 100000,
    "lock_period_fd": 1
}
```

### Response: 200
```json
{
    "statusCode": 200,
    "timeElapsed": 233802,
    "hasErrorResponse": false,
    "bank_account": {
        "customer_id": 234,
        "account_id": 51691857,
        "account_type": "FIXED",
        "balance": 100000,
        "lock_period_fd": 1
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Deposit into fixed account
Throw error if deposing into fixed account after account was created. Deposit in FD can be done only once while creating the account.
### Method: PUT
>```
>http://localhost:8080/accounts/51691857/deposit
>```
### Body (**raw**)

```json
{
   "amount": 100
}
```

### Response: 400
```json
{
    "statusCode": 400,
    "timeElapsed": 308928,
    "hasErrorResponse": true,
    "errorResponse": [
        {
            "errorMessage": "Deposit in fixed account is not allowed more than once.",
            "errorCode": "1014"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Err: Withdraw from fixed account
Throw error if withdrawing before maturity date from fixed account.
### Method: PUT
>```
>http://localhost:8080/accounts/51691204/withdraw
>```
### Body (**raw**)

```json
{
   "amount": 100
}
```

### Response: 400
```json
{
    "statusCode": 400,
    "timeElapsed": 231871,
    "hasErrorResponse": true,
    "errorResponse": [
        {
            "errorMessage": "Locking period of your FD is still not complete. You need to pass preMatureWithdrawal as true to withdraw. Penalty will be applied in case of premature withdrawal.",
            "errorCode": "1011"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Err2: Withdraw from fixed account
Throw error if withdrawing less amount than deposted from fixed account. All amount has to be withdrawn from FD.
### Method: PUT
>```
>http://localhost:8080/accounts/51691204/withdraw
>```
### Body (**raw**)

```json
{
   "amount": 100,
   "preMatureWithdrawal": true
}
```

### Response: 400
```json
{
    "statusCode": 400,
    "timeElapsed": 332045,
    "hasErrorResponse": true,
    "errorResponse": [
        {
            "errorMessage": "All amount has to be withdrawn from FD account. Your account balance is 100000.000000. Penalty will be applied in case of premature withdrawal.",
            "errorCode": "1011"
        }
    ]
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃


**Final Project:**

Write a program for Banking transaction :

Three Accounts to be created:

1. Savings Account
1. Current Account
1. Fixed Account.

Savings Account: Features

1. Each customer should have only one Savings Account in a Bank
1. Customer can deposit any amount In Savings Account.
1. If the Amount deposition is more than 50 thousand then application should ask for PAN card entry while deposition.
1. After deposition there should be display of current balance in his account.
1. There should be option to with draw the amount from the Account of corresponding customer.
1. Withdrawing is possible only if With draw amount is less than or equal to Balance amount.
1. If with draw is more than Balance then proper message should be displayed. And then should give option for customer to re enter once again to provide withdraw amount.
1. There should be Limit of amount to be present in Savings Account, since this varies between banks , So decision is left to the developer to decide on upper limit.  You can also provide default message to customer during his Account creation, specifying what is the upper limit of amount that be deposited in Savings account

Current Account : Features

1. Each Company  should have only one Current Account in a Bank.
1. Company can deposit any amount in Current Account .
1. If the deposit Amount is more than 2.5 lakhs then application should ask for PAN card of the company.
1. After deposition , at any time Balance should be displayed.
1. Company can with draw the amount if the withdraw amount is Less than balance.
1. There should be Overdraft option , Company can select whether it wanted to have Overdraft option for it.
1. If company chooses Overdraft then while withdrawing the amount , if the withdraw amount is more than the balance available , then the remaining amount should be withdrawn from Overdraft limit
1. There should be some upper limitation for Over draft and this varies from bank to bank , so choice is given to developer to provide the overdraft amount.
1. At any point the withdraw amount should not cross the Balance+Overdraft amount. If crossed then appropriate message should be displayed.
1. At any point display of current balance should be provided.

Fixed Account:

1. Customer can create any number of fixed accounts and for any amount
1. In Fixed account ,only once deposition is allowed and it should be locked for the period specified.
1. In Fixed account only once withdrawal is allowed and it is allowed by closing the account. After closing the fixed account all the amount should be withdrawn at once , there should not be any amount left in this account.
1. The closing of fixed account can be done only after the completion of locking period.
1. If it is closed before locking period then penalty should be imposed and 0.1 percent of the deposit amount should be deducted and remaining should be paid to the customer.
1. So there should be option for displaying what is the amount the customer is getting when he closed the fixed Account.
1. There should be option to display the list of all the fixed accounts opened by  each of the customer  along with total amount of all the fixed accounts.
1. So if any one account is closed then the list and the amount should be displayed accordingly.

Important Points before executing:

1. Use Concurrency , channels features . Functions related to transactions like Withdrawing , Deposit should be created as goroutines. 
1. For creating the Account and capturing the details of the account holder , create as endpoints , using these endpoints allow user to enter account holder details and save them into database when clicked on submit button.
1. All the account holders details created should be saved in database .(use any database e.g.sqlite3)
1. Create static pages using which user can enter account holder details.
1. Use Interface , structs , Slices ,Methods and Functions during development.
1. Try to implement atleast for 3 or more customers .
1. There should not be any race conditions occurring.
1. Appropriate Positive and Negative messages or information should be displayed to the user.
