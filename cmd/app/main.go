package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"itk/internal/handler"
	"itk/internal/repository"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	var db *repository.Database
	var err error

	for i := 0; i < 10; i++ {
		db, err = repository.NewDatabaseFromEnv()
		if err == nil {
			break
		}
		log.Println("Ожидание БД...")
		time.Sleep(time.Second)
	}

	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.DB.Close()

	log.Println("Подключение к БД успешно")

	walletRepo := repository.NewWalletRepository(db.DB)
	walletHandler := handler.WalletHandler{Repo: walletRepo}

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/wallet", walletHandler.HandleOperation).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{id}", walletHandler.HandleGetBalance).Methods("GET")
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Println("Сервер запущен на порту :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
