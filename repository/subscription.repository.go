package repository

type SubscriptionRepository struct {
	ConnectionString string
}

func NewSubscriptionRepository(connString string) *SubscriptionRepository {
	repo := SubscriptionRepository{
		ConnectionString: connString,
	}

	return &repo
}
