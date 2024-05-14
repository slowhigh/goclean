package main

// @title   goclean
// @version  1.0
// @description goclean

// @host   localhost:5000

// @schemes http
type Server interface {
	Run() error
}

func main() {
	var server Server

	server, err := InitServer()
	if err != nil {
		panic(err)
	}

	if err := server.Run(); err != nil {
		panic(err)
	}
}
