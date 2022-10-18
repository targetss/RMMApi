package main

import (
	"time"
)

// Упрощенная структура административных учетных записей TacticalRMM
type AccountsUser struct {
	LastLogin   time.Time `json:"last_login"`
	IsSuperUser bool      `json:"is_superuser"`
	UserName    string    `json:"username"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateJoined  time.Time `json:"date_joined"` //Дата создания учетной записи
	LastLoginIP string    `json:"last_login_ip"`
	RoleName    string    `json:"role"`
}
