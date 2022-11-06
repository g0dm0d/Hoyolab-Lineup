package entity

type RetData struct {
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Npt  string `json:"next_page_token"`
	List []Team `json:"list"`
}

type Team struct {
	ID          string        `json:"id"`
	AvatarGroup []AvatarGroup `json:"avatar_group"`
}

type AvatarGroup struct {
	Groups []Group `json:"group"`
}

type Group struct {
	Name string `json:"name"`
}
