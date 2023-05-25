package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rajpatil7322/MicroService-in-Golang/entity"
)

func GetProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile("./data/data.json")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write(data)
		return
	}
}

func CreateProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// Check if data is proper JSON (data validation)
		var product entity.Transaction
		err = json.Unmarshal(data, &product)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid Data Format"))
			fmt.Println(err)
			return
		}
		// Load existing products and append the data to product list
		var products []entity.Transaction
		data, err = ioutil.ReadFile("./data/data.json")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(data, &products)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		products = append(products, product)
		updatedData, err := json.Marshal(products)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = ioutil.WriteFile("./data/data.json", updatedData, os.ModePerm)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Added New Product"))
		return
	}
}
