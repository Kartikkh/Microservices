# Microservices

> This is a flow for Order Management system which involves 3 Microservices
 
- [Product Inventory](https://github.com/Kartikkh/Microservices/tree/master/Product-Inventory) Product. 
- [Cart](https://github.com/Kartikkh/Microservices/tree/master/Cart) cart   
- [Orders](https://github.com/Kartikkh/Microservices/tree/master/Orders) Order Manager
## Requirements

To run OMS you'll need:

- [Golang](https://golang.org/) v1.10+.

## Usage

### Getting started

You will need to fork the project and clone it locally.

```sh
$ git clone git@github.com:YOUR_USERNAME/Microservices.git
$ cd Microservices
```

Move all the folders from Microservices to 
```sh
~/go/src/github.com/ 
```

### Installation

Install all the dependencies

```sh
$ go mod
```

#### Now run the Product server 

```sh
$ cd /Product-Inventory
go run app.go
```

#### Now run the Cart server 

```sh
$ cd /Cart
go run app.go
```

#### Now run the Product server 

```sh
$ cd /Orders
go run app.go
```

You can use postman to check the API endpoints.

## Contribute

Feel free to raise an issue about a bug or new feature you would like to see in this project.
 
If you are willing to do some work, we will be glad to accept your PR.

## License

This project is [Licensed](LICENSE) under the MIT License (MIT).
