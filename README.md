# E-commerce System Models

This repository contains the models used in an e-commerce system. These models represent the essential entities in the system, such as users, products, carts, and transactions.

## Relation Database
![relasi ecommerce](https://github.com/flanker0507/fiber-e-commerce-system-API/assets/108620970/507417e6-63d2-4957-8d60-e540925c41d7)

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
- `PaymentURL` (string): The URL for the payment gateway.
- `Items` ([]TransactionItem): The list of items in the transaction.
- `CreatedAt` (time.Time): The timestamp when the transaction was created.
- `UpdatedAt` (time.Time): The timestamp when the transaction was last updated.

## Installation

To use these models, clone the repository and import the package into your Go project:

```bash
git clone https://github.com/yourusername/ecommerce-system-models.git
