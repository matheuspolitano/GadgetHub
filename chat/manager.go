package chat

import db "github.com/matheuspolitano/GadgetHub/pkg/db/sqlc"

type ManagerChat struct {
	template Template
	messager Messager
	store    db.Store
}
