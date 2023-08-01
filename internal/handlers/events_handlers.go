package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"mini-project-new/internal/domain"
	"net/http"
	"strconv"
)

var events = make(map[uuid.UUID]*domain.Event)

// CreateEvents creates a new event.
// @Summary Create a new event
// @Description Create a new event
// @Tags Events
// @Accept json
// @Produce json
// @Param eventData body CreateEventRequest true "Event data"
// @Success 201 {object} domain.Event "Created event"
// @Router /v1/events [post]
func CreateEvents(w http.ResponseWriter, r *http.Request) {
	var newEvent domain.Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request body")
		return
	}

	newEvent.EventsID = uuid.New()
	domain.Events = append(domain.Events, newEvent)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}

// UpdateEvents updates an existing event by ID.
// @Summary Update an event
// @Description Update an existing event by ID
// @Tags Events
// @Accept json
// @Produce json
// @Param events_id path string true "Event ID"
// @Param eventData body UpdateEventRequest true "Event data"
// @Success 200 {object} domain.Event "Updated event"
// @Router /v1/events/{events_id} [patch]
func UpdateEvents(w http.ResponseWriter, r *http.Request) {
	eventID := chi.URLParam(r, "events_id")

	id, err := uuid.Parse(eventID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid event ID")
		return
	}

	index := -999
	for i, event := range domain.Events {
		if event.EventsID == id {
			index = i
			break
		}
	}

	if index == -999 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Event not found")
		return
	}

	var updatedEvent domain.Event
	err = json.NewDecoder(r.Body).Decode(&updatedEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request body")
		return
	}
	updatedEvent.EventsID = id
	domain.Events[index] = updatedEvent

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEvent)
}

// DeleteEvents deletes an event by ID.
// @Summary Delete an event
// @Description Delete an event by ID
// @Tags Events
// @Accept json
// @Produce json
// @Param events_id path string true "Event ID"
// @Success 200 {object} SuccessResponse "Delete success"
// @Router /v1/events/{events_id} [delete]
func DeleteEvents(w http.ResponseWriter, r *http.Request) {
	eventID := chi.URLParam(r, "events_id")

	id, err := uuid.Parse(eventID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid event")
		return
	}

	index := -1
	for i, event := range domain.Events {
		if event.EventsID == id {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Event not found")
		return
	}

	domain.Events = append(domain.Events[:index], domain.Events[index+1:]...)

	response := map[string]string{"message": "delete success!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetEvent gets an event by ID.
// @Summary Get an event
// @Description Get an event by ID
// @Tags Events
// @Accept json
// @Produce json
// @Param events_id path string true "Event ID"
// @Success 200 {object} domain.Event "Event"
// @Router /v1/events/{events_id} [get]
func GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := chi.URLParam(r, "events_id")

	id, err := uuid.Parse(eventID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid event ID")
		return
	}

	event, found := events[id]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Event not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

// GetAllEvents gets all events with pagination and category filter.
// @Summary Get all events with pagination and category filter
// @Description Retrieve a paginated list of events filtered by category
// @Tags Events
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of events per page"
// @Param category query string false "Filter events by category"
// @Success 200 {array} domain.Event "List of events"
// @Router /v1/events [get]
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	category := r.URL.Query().Get("category")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	var filteredEvents []domain.Event
	if category != "" {
		for _, event := range domain.Events {
			if event.Category == category {
				filteredEvents = append(filteredEvents, event)
			}
		}
	} else {

		filteredEvents = domain.Events
	}

	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > len(filteredEvents) {
		endIndex = len(filteredEvents)
	}

	pagedEvents := filteredEvents[startIndex:endIndex]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pagedEvents)
}

type CreateEventRequest struct {
	NameEvents  string          `json:"name_events"`
	VenueEvents string          `json:"venue_events"`
	MaxCapacity int             `json:"max_capacity"`
	Category    string          `json:"category"`
	GuestStar   []domain.Artist `json:"guest_star"`
}

type UpdateEventRequest struct {
	NameEvents  string          `json:"name_events"`
	VenueEvents string          `json:"venue_events"`
	MaxCapacity int             `json:"max_capacity"`
	Category    string          `json:"category"`
	GuestStar   []domain.Artist `json:"guest_star"`
}
type SuccessResponse struct {
	Message string `json:"message"`
}
