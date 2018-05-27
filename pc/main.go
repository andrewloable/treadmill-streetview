package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/tarm/serial"
)

// SerialConfig ...
type SerialConfig struct {
	Name string
	Baud int
}

var config SerialConfig
var loops int64

func readConfigFromFile(ps *SerialConfig) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(file, &ps)
}

func processSerial(c *serial.Config) {
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
	}
	for {
		buf := make([]byte, 9)
		n, err := s.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < n; i++ {
			//fmt.Println(buf[i], ":", string(buf[i]))
			if buf[i] == 45 {
				fmt.Println("LOOP Detected")
				loops += 1
			}
		}
	}
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	config = SerialConfig{}
	readConfigFromFile(&config)

	c := &serial.Config{
		Name: config.Name,
		Baud: config.Baud,
	}

	go processSerial(c)

	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          60 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	api := r.Group("/api")

	// Serve Static Page
	r.Static("/app", "./public")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/app")
	})

	api.GET("/get_loops", func(c *gin.Context) {
		c.JSON(200, loops)
		return
	})

	api.GET("/reset_loops", func(c *gin.Context) {
		c.JSON(200, loops)
		loops = 0
		return
	})
	open("http://localhost:1513/")
	r.Run(":1513")
}
