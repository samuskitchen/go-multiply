package config

func Start(port string) {
	server := newServer(port)

	// start the server.
	server.Start()
}
