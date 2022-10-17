package main

import "time"

// Упрощенная структура административных учетных записей TacticalRMM
type AccountsUser struct {
	lastLogin   time.Time `json:"last_login"`
	isSuperUser bool      `json:"is_superuser"`
	userName    string    `json:"username"`
	firstName   string    `json:"first_name"`
	lastName    string    `json:"last_name"`
	email       string    `json:"email"`
	dateJoined  time.Time `json:"date_joined"` //Дата создания учетной записи
	lastLoginIP string    `json:"last_login_ip"`
}
