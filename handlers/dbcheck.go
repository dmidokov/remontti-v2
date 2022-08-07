package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func dbCheck(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		type Student struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		}

		type Students []Student

		var students Students

		sql := "select * FROM students"
		data, err := conn.Query(sql)
		if err != nil {
			panic(err)
		}

		for data.Next() {
			var student Student
			err = data.Scan(&student.ID, &student.Name)
			if err != nil {
				fmt.Fprintf(w, "Error Looping data "+err.Error())
			}
			students = append(students, student)
			fmt.Println(student.ID, student.Name)
		}

		out, _ := json.Marshal(students)

		fmt.Fprint(w, string(out))

	}
}
