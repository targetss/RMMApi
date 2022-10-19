package main

import (
	"time"
)

type Requests interface {
}

// Информация по конкретному компьютеру
type ComputerInfo struct {
	ID              int    `json:"id"`
	VersionAgent    string `json:"version"`
	Description     string `json:"description"`
	OperatingSystem string `json:"operating_system"`
	Hostname        string `json:"hostname"`
	WMI             string `json:"wmi_detail"`
	SiteID          int    `json:"site_id"`
}

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

// Информация по сайтам(компаниям)
type Site struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ClientID  int    `json:"client_id"`
	CreatedBy string `json:"created_by"`
}
