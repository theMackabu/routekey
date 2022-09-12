package models

type QRCode struct {
	ID      uint64 `json:"id" gorm:"primary_key"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
	Image   []byte `json:"image"`
}

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Error ServiceError `json:"error"`
}

type LinkBody struct {
	CustomURL   string `json:"customurl"`
	Target      string `json:"target"`
	Description string `json:"description"`
	Reusable    *bool  `json:"reusable"`
	Password    string `json:"password"`
	ExpireIn    string `json:"expire_in"`
	Domain      string `json:"domain"`
}
