package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"io"
	"net/http"
	"net/url"

	_ "github.com/lib/pq"
)

type Company struct {
	ID      string `json:"company_id"`
	Name    string `json:"company_name"`
	Sector  string `json:"company_sector"`
	Address string `json:"company_address"`
	ZipCode string `json:"company_area_code"`
	City    string `json:"company_city"`
	Country string `json:"company_country"`
	Phone   string `json:"company_phone"`
	Email   string `json:"company_email"`
}

type Location struct {
    Address string
    City    string
    ZipCode string
}

type CompaniesWrapper struct {
	Companies []Company `json:"items"`
}

type GeoApifyResponse struct {
	Features []struct {
		Properties struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"properties"`
	} `json:"features"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "xhuet"
	password = "1234"
	dbname   = "internship"
)

const apiKey = "8b7904aac90941b893df97c45046e288"

func getCoordinates(loc Location) (float64, float64, error) {
	fullAddress := fmt.Sprintf("%s, %s, %s, %s", loc.Address, loc.City, loc.ZipCode, "France")

	encodedAddress := url.QueryEscape(fullAddress)

	url := fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&apiKey=%s", encodedAddress, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var geoResp GeoApifyResponse
	err = json.Unmarshal(body, &geoResp)
	if err != nil {
		return 0, 0, err
	}

	if len(geoResp.Features) == 0 {
		return 0, 0, fmt.Errorf("aucune coordonnée trouvée pour l'adresse: %s", fullAddress)
	}

	return geoResp.Features[0].Properties.Lat, geoResp.Features[0].Properties.Lon, nil
}

func getAllAdresses(db *sql.DB) ([]Location, error) {
	rows, err := db.Query("SELECT address, city, zip_code FROM companies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []Location
	for rows.Next() {
		var loc Location
		if err := rows.Scan(&loc.Address, &loc.City, &loc.ZipCode); err != nil {
			return nil, err
		}
		locations = append(locations, loc)
	}
	return locations, nil
}

func insertCompany(db *sql.DB, c Company) error {
	query := `
		INSERT INTO companies (id, name, sector, address, zip_code, city, country, phone, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO NOTHING;
	`
	_, err := db.Exec(query, c.ID, c.Name, c.Sector, c.Address, c.ZipCode, c.City, c.Country, c.Phone, c.Email)
	return err
}


func main() {
	data, err := os.ReadFile("internship.json")
	if err != nil {
		fmt.Println("Errorr while reading the file:", err)
		return
	}

	var wrapper CompaniesWrapper
	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		fmt.Println("Errorr parsing JSON:", err)
		return
	}
	fmt.Printf("Companies found: %d\n", len(wrapper.Companies))

	companies := wrapper.Companies

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the db successfuly")
	for _, c := range companies {
		err := insertCompany(db, c)
		if err != nil {
			fmt.Println("Errorr insertion:", err)
		} else {
			fmt.Println("Company inserted:", c.Name)
		}
	}
	locations, err := getAllAdresses(db)
	if err != nil {
        fmt.Println("Erreur récupération adresses:", err)
        return
	}
	for _, loc := range locations {
    lat, lon, err := getCoordinates(loc)
    if err != nil {
        fmt.Println("Erreur Geoapify:", err)
        continue
    }
    fmt.Printf("%s, %s, %s -> Lat: %f, Lon: %f\n", loc.Address, loc.ZipCode, loc.City, lat, lon)
}
}
