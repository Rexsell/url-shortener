package server

import (
	"log"
	"net/http"
)

type Server struct {
	Port    string
	Handler func(http.ResponseWriter, *http.Request)
	//Сюда бы еще что-то закинул, если вообще надо
	//Мб встроить сюда структурку с БД, но хз или просто порт/путь к бд

	//Может быть мапу map[string]func(), где будут URI и функция ручки
	//В методе Start() их перебирать и запускать в http.HandleFunc(string, func())
}

func New(port string) *Server {
	return &Server{
		Port:    port,
		Handler: ShortenerHandler,
	}
}

func (s *Server) Start() error {
	log.Printf("Server started at port %s", s.Port)
	http.HandleFunc("/api/short", s.Handler)
	err := http.ListenAndServe(s.Port, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
