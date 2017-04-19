package door

// Door is a struct containing all the data concerning a Door
type Door struct {
	ID          string `json:"id"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Price       string `json:"price"`
	OwnerName   string `json:"ownerName"`
	OwnerPhone  string `json:"ownerPhone"`
}
