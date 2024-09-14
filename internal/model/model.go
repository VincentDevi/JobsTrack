package model

type OrganisationsPageElement struct {
	Name           string
	Country        string
	Language       string
	Tag            string
	Message_number int64
	Offer_number   int64
}

func DummiesOrganisationsPageElement() []OrganisationsPageElement {
	return []OrganisationsPageElement{
		{
			Name:           "Odoo",
			Country:        "Belgium",
			Language:       "French",
			Tag:            "Dev",
			Message_number: 1,
			Offer_number:   1,
		},
		{
			Name:           "Aerospacelab",
			Country:        "Belgium",
			Language:       "French",
			Tag:            "Dev",
			Message_number: 1,
			Offer_number:   1,
		},
		{
			Name:           "Komon",
			Country:        "Belgium",
			Language:       "French",
			Tag:            "Dev",
			Message_number: 2,
			Offer_number:   1,
		},
	}
}

type MessagesPageElement struct {
	Message_type string
	Offer        string
	Organisation string
	Channel      string
	Created_at   string
}

func DummiesMessagesPageElement() []MessagesPageElement {
	return []MessagesPageElement{
		{
			Message_type: "first",
			Offer:        "fullstack dev",
			Organisation: "Odoo",
			Channel:      "def@gmail.com",
			Created_at:   "14/09/2024",
		},
		{
			Message_type: "last",
			Offer:        "frontend",
			Organisation: "Aerospacelab",
			Channel:      "def@gmail.com",
			Created_at:   "10/04/2024",
		},
		{
			Message_type: "continue",
			Offer:        "backend dev",
			Organisation: "Komon",
			Channel:      "def@gmail.com",
			Created_at:   "20/02/2024",
		},
		{
			Message_type: "continuet",
			Offer:        "fullstack dev",
			Organisation: "Komon",
			Channel:      "def@gmail.com",
			Created_at:   "02/02/2023",
		},
	}
}

type OffersPageElement struct {
	Organisation string
	Remote       bool
	Salary       float64
	Messages     int64
	Created_at   string
	Updated_at   string
}

func DummiesOffersPageElement() []OffersPageElement {
	return []OffersPageElement{
		{
			Organisation: "Odoo",
			Remote:       false,
			Salary:       2500.20,
			Messages:     1,
			Created_at:   "02/02/2020",
			Updated_at:   "02/02/2020",
		},
		{
			Organisation: "Aerospacelab",
			Remote:       false,
			Salary:       2500.20,
			Messages:     1,
			Created_at:   "02/02/2020",
			Updated_at:   "02/02/2020",
		},
		{
			Organisation: "Komon",
			Remote:       true,
			Salary:       2500.20,
			Messages:     2,
			Created_at:   "02/02/2020",
			Updated_at:   "02/02/2020",
		},
	}
}
