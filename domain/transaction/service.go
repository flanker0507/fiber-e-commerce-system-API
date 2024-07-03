package transaction

import (
	"fiber-e-commerce-system-API/domain/models"
	"fiber-e-commerce-system-API/payment"
	"strconv"
)

type service struct {
	repository     Repository
	paymentService payment.Service
}

type Service interface {
	GetTransactionsByUserID(userID int) ([]models.Transaction, error)
	CreateTransaction(input CreateTransactionInput) (models.Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
	GetAllTransactions() ([]models.Transaction, error)
}

func NewService(repository Repository, paymentService payment.Service) *service {
	return &service{repository, paymentService}
}

func (s *service) GetTransactionsByUserID(userID int) ([]models.Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (models.Transaction, error) {
	// Dapatkan item dalam cart
	cartItems, err := s.repository.GetCartItems(input.CartID)
	if err != nil {
		return models.Transaction{}, err
	}

	// Hitung total transaksi
	var total float64
	for _, item := range cartItems {
		total += float64(item.Quantity) * item.Product.Price
	}

	transaction := models.Transaction{
		CartID: input.CartID,
		UserID: input.UserID,
		Total:  total,
		Status: "pending",
	}

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID:    newTransaction.ID,
		Total: int(newTransaction.Total),
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transactionID, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetByID(transactionID)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	_, err = s.repository.Update(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllTransactions() ([]models.Transaction, error) {
	transactions, err := s.repository.FindAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
