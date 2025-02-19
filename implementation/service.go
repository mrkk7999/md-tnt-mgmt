package implementation

import "md-tnt-mgmt/iface"

type service struct {
	repo iface.Repository
}

func New(repo iface.Repository) iface.Service {
	return &service{
		repo: repo,
	}
}
