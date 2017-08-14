package barrier

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var timeoutMilliseconds = 5000

type barrierResp struct {
	Err error
	Resp string
}

func barrier(endpoints ...string) {

	requestNumber := len(endpoints)
	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)
	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		//block the execution waiting for data form the in channel
		resp := <-in
		if resp.Err != nil {
			fmt.Println("Error : ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}

	if hasError{
		for _, resp := range response {
			fmt.Println(resp.Resp)
		}
	}

}

// makeRequest peform http get request to an url and accepts a channel to ouptput barrierResp values
func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}

	// Create the http client and set the timeout

	client := http.Client{
		Timeout: timeDuration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}

	// peform http get request
	resp, err := clint.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	// send it throught the channel
	res.Resp = string(byt)
	out <- res
}

//captureBarrierOutput will capture the outputs in stdout
func captureBarrierOutput(endpoints ...string) string {
	// create reader input to a bytes buffer before sending
	// the contents of the buffer through a channel
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)

	//close the writter to the signal the gorutine that
	// no more input is going to come it
	write.Close()
	temp := <-out

	return temp
}
