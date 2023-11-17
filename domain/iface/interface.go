package iface

import "ddd_demo/domain/model"

type UserRepo interface {
	GetUserByID(id string) (*model.User, error)
}
