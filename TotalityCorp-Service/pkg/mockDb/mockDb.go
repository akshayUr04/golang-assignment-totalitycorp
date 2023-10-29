package mockdb

type User struct {
	ID      int32
	Fname   string
	City    string
	Phone   int64
	Height  float32
	Married bool
}

var Users = map[int32]User{
	1: {
		ID:      1,
		Fname:   "Akshay",
		City:    "Kozhikode",
		Phone:   1234567890,
		Height:  5.8,
		Married: false,
	},
	2: {
		ID:      2,
		Fname:   "Amal",
		City:    "Pathanamthitta",
		Phone:   1234567890,
		Height:  5.4,
		Married: true,
	},
	3: {
		ID:      3,
		Fname:   "Amar",
		City:    "Palakkad",
		Phone:   1234567890,
		Height:  5.3,
		Married: true,
	},
	4: {
		ID:      4,
		Fname:   "Aswathi",
		City:    "Wayanad",
		Phone:   1234567890,
		Height:  5,
		Married: false,
	},
	5: {
		ID:      5,
		Fname:   "Abhi",
		City:    "Kottayam",
		Phone:   1234567890,
		Height:  5.7,
		Married: true,
	},
	6: {
		ID:      6,
		Fname:   "Amar",
		City:    "Kollam",
		Phone:   1234567890,
		Height:  5.9,
		Married: true,
	},
}
