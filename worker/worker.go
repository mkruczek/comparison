package main

import (
	"bytes"
	"comparasion/common"
	"comparasion/resources"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	go work(common.GinPointers, common.GinPointersPort)
	go work(common.GinCallback, common.GinCallbackPort)
	go work(common.EchoCallback, common.EchoCallbackPort)
	go work(common.EchoPointers, common.EchoPointersPort)

	select {}
}

func work(version, port string) {

	url := fmt.Sprintf("http://localhost%s/api/%s/resources", port, version)

	var body, newBody []byte
	var resp *http.Response
	var resource resources.Resources

	start := time.Now()
	var total int
	for i := 0; i < 100000; i++ {

		c := http.Client{Timeout: time.Duration(1) * time.Second}

		//POST request
		body = []byte(`{"Name":"test","Value":"test"}`)
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		if err != nil {
			log.Printf("Worker %s Error at 41 %s", version, err)
			goto FINISH
		}
		resp, err = c.Do(req)
		if err != nil {
			log.Printf("Worker %s Error at 46 %s", version, err)
			goto FINISH
		}
		body, err = io.ReadAll(resp.Body)
		req.Body.Close()

		resource = resources.Resources{}
		err = json.Unmarshal(body, &resource)
		if err != nil {
			log.Printf("Worker %s Error at 55 %s", version, err)
			goto FINISH
		}

		//PUT request
		newBody = []byte(`{"ID":"` + resource.ID.String() + `","Name":"newTest","Value":"newTest"}`)
		req, err = http.NewRequest(http.MethodPut, url, bytes.NewBuffer(newBody))
		if err != nil {
			log.Printf("Worker %s Error at 63 %s", version, err)
			goto FINISH
		}
		resp, err = c.Do(req)
		if err != nil {
			log.Printf("Worker %s Error at 68 %s", version, err)
			goto FINISH
		}

		//GET request
		req, err = http.NewRequest(http.MethodGet, url+"/"+resource.ID.String(), nil)
		if err != nil {
			log.Printf("Worker %s Error at 75 %s", version, err)
			goto FINISH
		}
		resp, err = c.Do(req)
		if err != nil {
			log.Printf("Worker %s Error at 80 %s", version, err)
			goto FINISH
		}

		//GET all request
		req, err = http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Printf("Worker %s Error at 87 %s", version, err)
			goto FINISH
		}
		resp, err = c.Do(req)
		if err != nil {
			log.Printf("Worker %s Error at 92 %s", version, err)
			goto FINISH
		}

		//DELETE request
		req, err = http.NewRequest(http.MethodDelete, url+"/"+resource.ID.String(), nil)
		if err != nil {
			log.Printf("Worker %s Error at 99 %s", version, err)
			goto FINISH
		}
		resp, err = c.Do(req)
		if err != nil {
			log.Printf("Worker %s Error at 104 %s", version, err)
			goto FINISH
		}
		total++
	FINISH:
		if i%10000 == 0 {
			fmt.Printf("Worker %s: %d requests.\n", version, i)
		}

	}

	fmt.Printf("Worker %s finished: %d requests in %s.\n", version, total, time.Since(start))
}
