package randomuser

import "time"

type Login struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Md5      string `json:"md5"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
}

type Result struct {
	Gender string `json:"gender"`
	Name   struct {
		Title string `json:"title"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Location struct {
		Street struct {
			Number int    `json:"number"`
			Name   string `json:"name"`
		} `json:"street"`
		City        string `json:"city"`
		State       string `json:"state"`
		Country     string `json:"country"`
		Postcode    any    `json:"postcode"` // field can return int or string
		Coordinates struct {
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
		} `json:"coordinates"`
		Timezone struct {
			Offset      string `json:"offset"`
			Description string `json:"description"`
		} `json:"timezone"`
	} `json:"location"`
	Email string `json:"email"`
	Login Login  `json:"login"`
	Dob   struct {
		Date time.Time `json:"date"`
		Age  int       `json:"age"`
	} `json:"dob"`
	Registered struct {
		Date time.Time `json:"date"`
		Age  int       `json:"age"`
	} `json:"registered"`
	Phone string `json:"phone"`
	Cell  string `json:"cell"`
	ID    struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"id"`
	Picture struct {
		Large     string `json:"large"`
		Medium    string `json:"medium"`
		Thumbnail string `json:"thumbnail"`
	} `json:"picture"`
	Nat string `json:"nat"`
}

type Response struct {
	Results []Result `json:"results"`
	Info    struct {
		Seed    string `json:"seed"`
		Results int    `json:"results"`
		Page    int    `json:"page"`
		Version string `json:"version"`
	} `json:"info"`
}
