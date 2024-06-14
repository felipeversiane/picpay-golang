# PicPay Challenge API

API for managing users and transfers, developed for PicPay's challenge.

## Introduction

The API allows for the management of users and transfers. Below are described the main functionalities available, as well as instructions for setup and usage.

## Main Libraries Used

- **github.com/jackc/pgx/v5** - PostgreSQL database driver for Go.
- **github.com/joho/godotenv** - Loads environment variables from a .env file.
- **go.uber.org/zap** - Fast, structured logging for Go.
- **github.com/gin-gonic/gin** - Fast, and concurrency way to make a api.

For a full list of dependencies, please refer to the [go.mod](https://github.com/felipeversiane/picpay-golang/blob/main/go.mod) file.

## Functionalities

### Users

#### Create User

- **Description:** Creates a new user with the provided information.
- **Method:** `POST`
- **Endpoint:** `/api/v1/user`
- **Request Body:**

  ```json
  {
    "email": "pedro@xx.com",
    "password": "senha123!F",
    "first_name": "Pedro",
    "last_name": "Silva",
    "is_merchant": false,
    "document": "1234567220",
    "balance": 200.00
  }

#### Get User By ID

- **Description:** Find a user with the provided id.
- **Method:** `GET`
- **Endpoint:** `/api/v1/user/{id}`

  #### Get User By Email

- **Description:** Find a user with the provided email.
- **Method:** `GET`
- **Endpoint:** `/api/v1/user/find_user_by_email/{email}`

 #### Get User By Document

- **Description:** Find a user with the provided document.
- **Method:** `GET`
- **Endpoint:** `/api/v1/user/find_user_by_document/{document}`

 #### Get User By Document

- **Description:** Find a user with the provided document.
- **Method:** `GET`
- **Endpoint:** `/api/v1/user/find_user_by_document/{document}`

 #### Update User

- **Description:** Update a user with the provided ID.
- **Method:** `PUT`
- **Endpoint:** `/api/v1/user/{id}`
- **Request Body:**

  ```json
  {
    "first_name": "Pedro",
    "last_name": "Silva",
    "is_merchant": false,
    "balance": 200.00
  }

 #### Delete User

- **Description:** Delete a user with the provided id.
- **Method:** `DELETE`
- **Endpoint:** `/api/v1/user/{id}`

### Orders

#### Create Order

- **Description:** Creates a new order with the provided information.
- **Method:** `POST`
- **Endpoint:** `/api/v1/order`
- **Request Body:**

  ```json
  {
  "amount": 1000.00,
  "payee": "d260ff06-5369-4269-8d54-91bbf42fd26a",
  "payer": "103cdecf-6e1c-4ec9-9e54-2dc6a4f185f8"
  }   


#### Get Order By ID

- **Description:** Find a order with the provided id.
- **Method:** `GET`
- **Endpoint:** `/api/v1/order/{id}`

The API is deployed and accessible at [picpay-golang.onrender.com](https://picpay-golang.onrender.com/docs/index.html).


