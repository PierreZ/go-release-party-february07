import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	// Unlock HTTP/2 server push
	p, err := w.(http.Pusher) // HL
	if err != nil {
		p.Push("/my_awesome_logo.png", nil) // HL
	}
}
