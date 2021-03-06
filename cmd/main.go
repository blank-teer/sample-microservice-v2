package main

import (
	"log"

	"sample-microservice-v2/internal/config"
	userService "sample-microservice-v2/internal/domain/user"
	"sample-microservice-v2/internal/repository/postgres"
	userStorage "sample-microservice-v2/internal/repository/postgres/user"
	"sample-microservice-v2/internal/transport/http"
	userHandler "sample-microservice-v2/internal/transport/http/user"
	userUsecase "sample-microservice-v2/internal/usecase/user"
)

func main() {
	c, err := config.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	p, err := postgres.NewConn(c.Postgres)
	if err != nil {
		log.Fatalln(err)
	}

	rUser := userStorage.NewRepository(p)
	sUser := userService.NewService(rUser, nil)
	uUser := userUsecase.NewUsecase(sUser)
	hUser := userHandler.NewHandler(uUser)

	s := http.NewServer(c.HTTP)
	s.MountRoutes(hUser)
	s.Run()
}
