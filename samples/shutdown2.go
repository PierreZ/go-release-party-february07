http.HandleFunc("/", handler)
err := srv.ListenAndServe()
if err != http.ErrServerClosed { // HL
		log.Fatalf("listen: %s\n", err)
}
log.Println("Server gracefully stopped")
