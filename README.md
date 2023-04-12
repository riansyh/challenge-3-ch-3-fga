# Challenge-2 Chapter 3 FGA Golang

## Soal

Buatlah Rest API product (create, read, update, delete) dengan fitur login dan register, serta memiliki 3 fitur middleware antara lain :

-   Authentication
-   Authorization multi level user
-   Authorization access product by id

Notes : buatlah authentication dengan JWT token golang, lalu gunakan token tersebut untuk setiap hit Rest API product.

Skema yang harus diterapkan:
![skema](https://i.ibb.co/ZY0LjcL/skema.png)

## List Endpoint

| HTTP Method | Endpoint        | Description                    | Authorization              |
| ----------- | --------------- | ------------------------------ | -------------------------- |
| POST        | /users/register | Register user with 'user' role | -                          |
| POST        | /users/login    | Login user                     | -                          |
| GET         | /products       | Get all products list          | Admin only                 |
| POST        | /products       | Create a product               | Admin & user               |
| GET         | /products/:id   | Get product by id              | Admin & corresponding user |
| PUT         | /products/:id   | Update product by id           | Admin & corresponding user |
| DELETE      | /products/:id   | Delete product by id           | Admin & corresponding user |
