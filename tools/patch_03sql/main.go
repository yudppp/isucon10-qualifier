package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("1_DummyEstateData.sql")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("File contents: %s", content)
	data := fmt.Sprintf("%s", content)
	data = strings.TrimPrefix(data, "INSERT INTO isuumo.estate (id, thumbnail, name, latitude, longitude, address, rent, door_height, door_width, popularity, description, features) VALUES (")
	data = strings.TrimSuffix(data, ");")
	// fmt.Println(data)
	rows := strings.Split(data, "), (")
	// fmt.Println(len(rows))

	// rows = rows[0:3]
	// fmt.Println(rows)

	patchedRows := make([]string, len(rows))
	for i, row := range rows {
		args := strings.Split(row, ",")
		args2 := strings.Split(row, ",")
		updates := append(args2[0:5], fmt.Sprintf("POINT(%v %v)", strings.TrimSpace(args2[3]), strings.TrimSpace(args2[4])))
		updates = append(updates, args[5:]...)
		patchedRows[i] = strings.Join(updates, ",")
	}

	query := "INSERT INTO isuumo.estate (id, thumbnail, name, latitude, longitude, geo, address, rent, door_height, door_width, popularity, description, features) VALUES ("
	query += strings.Join(patchedRows, "), (")
	query += ");"
	fmt.Println(query)
}
