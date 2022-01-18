package interfaces

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"key-value-store/domain"
	"key-value-store/infrastructure"
)

func HandlerMain() {
	http.HandleFunc("/keys/all", func(rw http.ResponseWriter, req *http.Request) {
		enableCors(&rw)
		switch req.Method {
		case http.MethodGet:
			{
				allData := infrastructure.GetAllData()
				if allData != "err" {
					fmt.Fprint(rw, allData)
					log.Println("All data is printed.")
				} else {
					fmt.Fprintf(rw, "X000 No data in memory.")
				}

			}
		default:
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/keys", func(rw http.ResponseWriter, req *http.Request) {
		enableCors(&rw)
		switch req.Method {
		case http.MethodGet:
			{

				keys, ok := req.URL.Query()["key"]

				if !ok || len(keys[0]) < 1 {
					log.Println("X001 Url Param 'key' is missing.")
					fmt.Fprint(rw, "X001 Url Param 'key' is missing.")
					rw.WriteHeader(http.StatusBadRequest)
					return
				}
				key := keys[0]

				incomeData := infrastructure.ReadData(key)
				if incomeData == "" {
					log.Println("Url Param 'key' is: " + string(key))
					fmt.Fprint(rw, "X002 There is no value for this key => ", key)
					rw.WriteHeader(http.StatusNoContent)
				}
				log.Println("Incoming value => " + incomeData)
				fmt.Fprint(rw, incomeData)
			}
		case http.MethodPut:
			{

				keys, ok := req.URL.Query()["key"]
				values, ok := req.URL.Query()["value"]
				if !ok || len(keys[0]) < 1 {
					log.Println("X001 Url Param 'key' is missing")
					fmt.Fprint(rw, "X001 Parameter Required")
					rw.WriteHeader(http.StatusNoContent)
					return
				}
				key := keys[0]
				value := values[0]
				log.Println("Url Param 'key' is: " + key)
				log.Println("Url Param 'value' is: " + value)
				infrastructure.WriteData(domain.NewMemory(key, value))
				fmt.Fprint(rw, "Key Added => ", key)

			}
		case http.MethodDelete:
			{
				keys, ok := req.URL.Query()["key"]

				if !ok || len(keys[0]) < 1 {
					infrastructure.DeleteAll()
					fmt.Fprint(rw, "All datas was deleted.")
					rw.WriteHeader(http.StatusOK)

					return
				} else {
					key := keys[0]

					infrastructure.DeleteKey(key)
					log.Println(" Deleted key is: " + key)
					fmt.Fprint(rw, " Deleted key is: "+key)
					rw.WriteHeader(http.StatusOK)
				}

			}
		default:
			enableCors(&rw)
			http.Error(rw, "X003 Method not allowed", http.StatusMethodNotAllowed)
		}

	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	(*w).Header().Set("Content-Type", "application/x-www-form-urlencoded")
}
