package awesomeProject

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetPublicAPI() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("Test new")
		fmt.Printf("%s\n", string(contents))
	}
}
