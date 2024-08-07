# E-commerce System Models with Midtrans Payment Gateway

This repository contains the models used in an e-commerce system, including integration with the Midtrans payment gateway. These models represent the essential entities in the system, such as users, products, carts, and transactions.

## Relation Database
![relasi ecommerce](https://github.com/flanker0507/fiber-e-commerce-system-API/assets/108620970/bc7a0c7a-0c1b-4357-9a7a-2edd0d96363a)

## Models

### User

The `User` model represents a user in the system.

#### Fields:
- `ID` (int): The unique identifier for the user.
- `Name` (string): The name of the user.
- `Email` (string): The email address of the user.
- `Password` (string): The hashed password of the user.
- `Role` (string): The role of the user (e.g., admin, customer).
- `Token` (string): The authentication token for the user.
- `CreatedAt` (time.Time): The timestamp when the user was created.
- `UpdatedAt` (time.Time): The timestamp when the user was last updated.

### Product

The `Product` model represents an item available for purchase.

#### Fields:
- `ID` (int): The unique identifier for the product.
- `Name` (string): The name of the product.
- `Quantity` (int): The quantity of the product in stock.
- `Price` (float64): The price of the product.

### Cart

The `Cart` model represents a shopping cart containing products.

#### Fields:
- `ID` (int): The unique identifier for the cart.
- `UserID` (int): The unique identifier for the user who owns the cart.
- `CreatedAt` (time.Time): The timestamp when the cart was created.
- `UpdatedAt` (time.Time): The timestamp when the cart was last updated.
- `Items` ([]CartItem): The list of items in the cart.
- `User` (User): The user who owns the cart.

### CartItem

The `CartItem` model represents an item in the cart.

#### Fields:
- `ID` (int): The unique identifier for the cart item.
- `CartID` (int): The unique identifier for the cart.
- `ProductID` (int): The unique identifier for the product.
- `Quantity` (int): The quantity of the product in the cart.
- `Product` (Product): The product associated with the cart item.
- `CreatedAt` (time.Time): The timestamp when the cart item was created.
- `UpdatedAt` (time.Time): The timestamp when the cart item was last updated.

### Transaction

The `Transaction` model represents a completed transaction.

#### Fields:
- `ID` (int): The unique identifier for the transaction.
- `CartID` (int): The unique identifier for the cart associated with the transaction.
- `UserID` (int): The unique identifier for the user who made the transaction.
- `Total` (float64): The total amount of the transaction.
- `Status` (string): The status of the transaction (e.g., pending, completed).
- `PaymentURL` (string): The URL for the payment gateway (Midtrans).
- `Items` ([]TransactionItem): The list of items in the transaction.
- `CreatedAt` (time.Time): The timestamp when the transaction was created.
- `UpdatedAt` (time.Time): The timestamp when the transaction was last updated.

## Features

### User Management

- **User Registration and Authentication**: Allows users to register and log in to the system securely.
- **User Roles**: Assign different roles to users, such as admin or customer, to control access and permissions.

### Product Management

- **Product Catalog**: Manage the list of products available for purchase, including their details such as name, quantity, and price.
- **Stock Management**: Keep track of the quantity of each product in stock.

### Shopping Cart

- **Cart Operations**: Users can add, update, and remove items from their shopping cart.
- **Persistent Cart**: The cart is saved in the database and associated with a specific user, allowing for a persistent shopping experience.

### Transactions

- **Order Processing**: Create and manage transactions, including calculating the total amount and maintaining the status of the transaction.
- **Transaction Items**: Keep track of each item within a transaction, including the product details and quantity.

### Payment Integration

- **Midtrans Payment Gateway**: Integrate with Midtrans to handle payment processing securely.
- **Payment URL**: Generate a payment URL for users to complete their transactions through Midtrans.
- **Transaction Status**: Update and manage the status of transactions based on payment results from Midtrans.

## Integration with Midtrans

This project integrates with the Midtrans payment gateway for processing transactions. To use Midtrans, you need to set up an account and obtain the necessary API keys.

### Installation

To install the Midtrans SDK, use the following command:

```bash
go get github.com/veritrans/go-midtrans
