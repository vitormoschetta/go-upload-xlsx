package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tealeg/xlsx"
)

func UploadExcel(w http.ResponseWriter, r *http.Request) {
	// Obtenha o arquivo enviado pelo cliente
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Falha ao obter o arquivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Abra o arquivo Excel
	xlFile, err := xlsx.OpenReaderAt(file, handler.Size)
	if err != nil {
		http.Error(w, "Falha ao abrir o arquivo Excel", http.StatusInternalServerError)
		return
	}

	// Percorra as linhas da segunda aba do arquivo Excel
	if len(xlFile.Sheets) > 0 {
		sheet := xlFile.Sheets[1]
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Print(cell.String())
			}
		}
	}

	// Responda ao cliente com uma mensagem de sucesso
	fmt.Print(w, "Arquivo %s enviado com sucesso", handler.Filename)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/upload", UploadExcel).Methods("POST")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
