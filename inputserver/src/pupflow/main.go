package main

import (
	"github.com/banthar/Go-SDL/sdl"
	"flag"
	"time"
	"os"
	"encoding/json"
	"log"
	"strconv"
)

const (
	UPDATES_PER_SECOND = 50
	SLEEP_TIME_US = 1000000/UPDATES_PER_SECOND
	SAVE_SLEEP_SEC = 10
)

var (
	dataFlag = flag.String("data", "localhost:13370", "UDP port to broadcast to")
	webFlag = flag.String("web", "localhost:8181", "Address to spawn webserver on")
	fileFlag = flag.String("save", "conf", "Save file")
)

var (
	config = map[uint16]SceneObject {}
)

func main() {
	flag.Parse()

	load_config(*fileFlag)

	sdl.Init(sdl.INIT_JOYSTICK)
	c := openJoystickStream()
	go startNetworking(*dataFlag, c)
	startWebserver(*webFlag)
}

func openJoystickStream() (<-chan []byte) {
	c := make(chan []byte)
	j := openAllJoysticks()
	go func() {
		for {
			sdl.JoystickUpdate()
			state := getJoystickState(j)
			for k, v := range config {
				if int(k) >= len(state) {
					continue
				}
				c <- v.MarshalWithValue(state[k])
			}
			time.Sleep(SLEEP_TIME_US * time.Microsecond)
		}
	}()
	return c
}


func load_config(file string) {
	f, e := os.Open(file)
	if e == nil {
		defer f.Close()
		strmap := make(map[string]SceneObject)
		d := json.NewDecoder(f)
		e := d.Decode(&strmap)
		if e != nil {
			log.Printf("Could not load config file: %s\n", e)
		}
		for k, v := range strmap {
			id, e := strconv.Atoi(k)
			if e != nil {
				continue
			}
			config[uint16(id)] = v
		}
	} else {
		log.Printf("Could not open config file: %s\n", e)
	}

	go func() {
		for {
			time.Sleep(SAVE_SLEEP_SEC * time.Second)
			f, e := os.Create(file)
			if e != nil {
				log.Printf("Could not save config: %s\n", e)
				continue
			}
			strmap := make(map[string]SceneObject)
			for k, v := range config {
				strmap[strconv.Itoa(int(k))] = v
			}
			d, _ := json.Marshal(strmap)
			f.Write(d)
			f.Close()
			log.Printf("Saved.")
		}
	}()
}
