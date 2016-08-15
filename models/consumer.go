package models

type Consumer struct {
    Uuid     string
	Username string
	Keys     []string
}

func NewConsumer(uuid, username string) (*Consumer) {
	if uuid == "" {
		uuid = newUuid()
	}
	return &Consumer{
		Uuid:     uuid,
		Username: username,
		Keys:     nil,
	}
}
