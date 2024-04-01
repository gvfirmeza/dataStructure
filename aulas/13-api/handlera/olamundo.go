package handlers

import "net/http"

func OlaMundo(w http.ResponseWriter, r *http.Request) {
	// Define o código da resposta
	w.WriteHeader((http.StatusOK))

	// Define o tipo de conteúdo que sera enviado
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	// Enviar o conteúdo
	w.Write([]byte("Ola, Mudo!"))
}