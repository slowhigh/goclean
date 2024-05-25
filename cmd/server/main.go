package main

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
