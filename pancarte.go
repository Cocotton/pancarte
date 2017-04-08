package main

type owner struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type door struct {
	Address     string  `json:"address"`
	Description string  `json:"descrition"`
	Price       string  `json:"price"`
	Owners      []Owner `json:"owner"`
}

func main() {

}
