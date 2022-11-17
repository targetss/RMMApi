package main

type Requests interface {
}

type InfoClient struct {
	IDClient     int      `json:"id"`
	Client       string   `json:"client"`
	Sites        []string `json:"site"`
	CountAgent   []int    `json:"count_agent"`
	AgentOnline  []int    `json:"agent_online"`
	AgentOffline []int    `json:"agent_offline"`
	CountNotes   []string `json:"notes"`
}

type InfoObject struct {
	CountSite    int `json:"count_site"`
	CountAgent   int `json:"count_agent"`
	AgentOffline int `json:"agent_offline"`
	AgentOnline  int `json:"agent_online"`
}

type FullComputerInfo struct {
	ID               int                      `json:"id"`
	VersionAgent     string                   `json:"version"`
	Description      string                   `json:"description"`
	OperatingSystem  string                   `json:"operating_system"`
	Disks            []map[string]interface{} `json:"disks"`
	PublicIP         string                   `json:"public_ip"`
	TotalRAM         int                      `json:"total_ram"`
	LoggedInUsername string                   `json:"logged_in_username"`
	Goarch           string                   `json:"goarch"`
	Software         []map[string]interface{} `json:"software"`
	Hostname         string                   `json:"hostname"`
	WMIInfo          []map[string]interface{} `json:"wmi_detail"`
	SiteID           int                      `json:"site_id"`
}

// Информация по конкретному компьютеру
type ComputerInfo struct {
	ID              int                      `json:"id"`
	VersionAgent    string                   `json:"version"`
	Description     string                   `json:"description"`
	OperatingSystem string                   `json:"operating_system"`
	Hostname        string                   `json:"hostname"`
	WMIInfo         []map[string]interface{} `json:"wmi_detail"`
	SiteID          int                      `json:"site_id"`
}

// Упрощенная структура административных учетных записей TacticalRMM
type AccountsUser struct {
	LastLogin   string `json:"last_login"`
	IsSuperUser bool   `json:"is_superuser"`
	UserName    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateJoined  string `json:"date_joined"` //Дата создания учетной записи
	LastLoginIP string `json:"last_login_ip"`
	RoleName    string `json:"role"`
}

// Информация по сайтам(компаниям)
type Site struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ClientID  int    `json:"client_id"`
	CreatedBy string `json:"created_by"`
}
