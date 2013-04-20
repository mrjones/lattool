package main

import (
	"../latvis"

	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello world")

	http := http.Client{}


	engine := latvis.NewRenderEngine(
		latvis.NewLocalFSBlobStore("/tmp/lattool"),
		http.Transport)

	fmt.Println(engine.GetOAuthUrl("urn:ietf:wg:oauth:2.0:oob", ""))
	fmt.Print("Enter verification code: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	handle := latvis.GenerateHandle();

	bounds, err := latvis.NewBoundingBox(
		latvis.Coordinate{Lat: 40.69834018178774, Lng: -74.02381896972656},
		latvis.Coordinate{Lat: 40.783660996197945, Lng: -73.96064758300781})

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	start = start.Add(-1 * 24 * 7 * time.Hour)

	end := time.Now()

	request := &latvis.RenderRequest{
		Bounds: bounds,
		Start: start,
		End: end,
	}

	if err != nil {
		log.Fatal(err)
	}

	engine.Execute(request, verificationCode, "urn:ietf:wg:oauth:2.0:oob", handle)

	blob, err := engine.FetchImage(handle)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got: %d bytes\n", len(blob.Data))

	ioutil.WriteFile("/var/www/lattool/latest.png", blob.Data, 0777)
}
