package ports

type InformRepository interface {
	Inform() (inform any, err error)
}

type InformService interface {
	Inform() (inform any, err error)
}