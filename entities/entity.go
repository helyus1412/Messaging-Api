package entities

type Messages struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	UserId  int    `json:"user_id"`
	From    string `json:"from"`
	To      string `json:"to"`
}
