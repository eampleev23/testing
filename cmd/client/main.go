package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://practicum.yandex.ru")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	//if _, err := io.CopyN(os.Stdout, response.Body, 512); err != nil {
	//	fmt.Println(err)
	//}
	/*другой вариант*/
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(body) > 512 {
		body = body[:512]
	}
	fmt.Print(string(body))
}
