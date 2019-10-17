package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-http-utils/cors"
)

type HouseDAO struct {
	Id         int
	Name       string
	Region     string
	CoatOfArms string
	Words      string
}

var houses = []HouseDAO{
	HouseDAO{
		Id:         1,
		Name:       "House Algood",
		Region:     "The Westerlands",
		CoatOfArms: "A golden wreath, on a blue field with a gold border(Azure, a garland of laurel within a bordure or)",
		Words:      "",
	},
	HouseDAO{
		Id:         2,
		Name:       "House Allyrion of Godsgrace",
		Region:     "Dorne",
		CoatOfArms: "Gyronny Gules and Sable, a hand couped Or",
		Words:      "No Foe May Pass",
	},
	HouseDAO{
		Id:         3,
		Name:       "House Amber",
		Region:     "The North",
		CoatOfArms: "",
		Words:      "",
	},
}

type MyArgs struct {
	Id int `json:"Id"`
}

type MaRequest struct {
	Method string `json:"method"`
	Params MyArgs `json:"params"`
	Id     string `json:"id"`
}

type MyResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Id      string      `json:"id"`
}

func getHouse(id int) HouseDAO {
	var house HouseDAO

	for index := 0; index < len(houses); index++ {
		if houses[index].Id == id {
			house = houses[index]
		}
	}

	return house
}

func main() {
	log.Println("Server")
	mux := http.NewServeMux()

	mux.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		var request MaRequest
		var response MyResponse

		json.NewDecoder(r.Body).Decode(&request)

		if request.Method == "house.GetHouses" {
			response = MyResponse{
				Jsonrpc: "2.0",
				Result:  houses,
				Id:      request.Id,
			}
		} else if request.Method == "house.GetHouse" {
			response = MyResponse{
				Jsonrpc: "2.0",
				Result:  getHouse(request.Params.Id),
				Id:      request.Id,
			}
		} else {
			response = MyResponse{
				Jsonrpc: "2.0",
				Result:  nil,
				Id:      request.Id,
			}
		}

		json.NewEncoder(w).Encode(response)
	})

	log.Fatal(http.ListenAndServe(":8080", cors.Handler(mux)))
}
