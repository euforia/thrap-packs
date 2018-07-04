package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	_version   string
	_buildtime string
)

// Version returns the full version of the app. This is populated at
// build time
func Version() string {
	return _version + " " + _buildtime
}

var (
	showVersion = flag.Bool("version", false, "show version")
	bindAddr    = flag.String("bind-addr", ":8080", "Bind address")
)

func init() {

	if len(_version) == 0 {
		_version = "unknown"
	}
	if len(_buildtime) == 0 {
		_buildtime = time.Now().UTC().String()
	}

}

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Println(Version())
		os.Exit(0)
	}

	//
	// TODO: Code below is a sample.  Add code here.
	//

	http.HandleFunc("/v1/status", func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(map[string]string{
			"version":   _version,
			"buildtime": _buildtime,
		})
		w.Write(b)
	})

	err := http.ListenAndServe(*bindAddr, nil)
	if err != nil {
		log.Fatal(err)
	}

}
