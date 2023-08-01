package domain

import "github.com/google/uuid"

type Event struct {
	EventsID    uuid.UUID `json:"events_id"`
	NameEvents  string    `json:"name_events"`
	VenueEvents string    `json:"venue_events"`
	MaxCapacity int       `json:"max_capacity"`
	Category    string    `json:"category"`
	GuestStar   []Artist  `json:"guest_star"`
}

type Artist struct {
	Artists string `json:"artists"`
}

var Events = []Event{
	Event{
		EventsID:    uuid.MustParse("adbaf183-a1e6-4986-8c1f-80837af95e77"),
		NameEvents:  "Festival Gong",
		VenueEvents: "Stadion Si Jalak Harupat",
		MaxCapacity: 5000,
		Category:    "Festival",
		GuestStar: []Artist{
			{
				Artists: "Ariel",
			},
		},
	},
}
