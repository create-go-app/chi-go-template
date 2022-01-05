package healthcheck

import "github.com/go-chi/chi/v5"

const (
	groupURL  = "/hc"
	statusURL = "/status"
)

// Routes function to create router.
func Routes(m *chi.Mux) {
	// Create group.
	m.Route(groupURL, func(r chi.Router) {
		r.Get(statusURL, getStatus) // get status route
	})
}
