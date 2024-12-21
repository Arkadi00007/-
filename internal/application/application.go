package application

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"LMS/pkg/calculation"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}
func (a *Application) Run() error {
	for {
		// читаем выражение для вычисления из командной строки
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}
		// убираем пробелы, чтобы оставить только вычислемое выражение
		text = strings.TrimSpace(text)
		// выходим, если ввели команду "exit"
		if text == "exit" {
			log.Println("aplication was successfully closed")
			return nil
		}
		//вычисляем выражение
		result, err := calculation.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed wit error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func RunServer(host string) {
	//ЗАПУСК СЕРВЕРА
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	port := ":8080"
	if host != "" {
		port = host
	}
	log.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	//ОБРАБОТКА ПОСТ-ЗАПРОСА
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	if req.Expression == "" {
		http.Error(w, `{"error": "Empty expression"}`, http.StatusUnprocessableEntity)
		return
	}

	result, err := calculation.Calc(req.Expression)
	if err != nil {
		if err.Error() == "incorrect symbols" {
			w.WriteHeader(422)
			json.NewEncoder(w).Encode(Response{
				Error: "Expression is not valid",
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Error: "Internal server error",
			})
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Result: result,
	})

}
