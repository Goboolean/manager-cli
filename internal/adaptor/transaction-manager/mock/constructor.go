package transactionManagerMock

type TransactionManagerMock struct {
}

func New() (*TransactionManagerMock, error) {
	return &TransactionManagerMock{}, nil
}
