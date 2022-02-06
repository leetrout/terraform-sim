package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TestResp struct {
	Id   int
	Name string
}

// func (*TestResp) MarshalJSON() ([]byte, error) {
// 	m := map[string]string{
// 		"foo": "bar",
// 	}
// 	return json.Marshal(m)
// }

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		w.Write([]byte(`{"Method": "Not supported"}`))
		return
	}
	var t = &TestResp{Id: 1, Name: "foo"}
	fmt.Printf("%+v\n", t)
	jsonOut, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		// TODO
		return
	}
	fmt.Println("no errors")
	fmt.Println(string(jsonOut))
	w.Write(jsonOut)
}
