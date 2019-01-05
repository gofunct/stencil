// Copyright © 2019 Coleman Word <coleman.word@gofunct.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd


import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A start a web server to edit templates",
	Run: func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()

		r.Handle("/", http.FileServer(http.Dir("static")))
		r.HandleFunc("/generate", generate)
		addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
		if addr == ":" {
			addr = ":8080"
		}

		fmt.Printf("Listening on %s...\n", addr)
		h := handlers.LoggingHandler(os.Stderr, r)
		h = handlers.CompressHandler(h)
		h = handlers.RecoveryHandler()(h)
		if err := http.ListenAndServe(addr, h); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func generate(w http.ResponseWriter, r *http.Request) {
	// read input
	decoder := json.NewDecoder(r.Body)
	type Input struct {
		Protobuf string `json:"protobuf"`
		Template string `json:"template"`
	}
	var input Input
	if err := decoder.Decode(&input); err != nil {
		returnError(w, err)
		return
	}

	// create workspace
	dir, err := ioutil.TempDir("", "pggt")
	if err != nil {
		returnError(w, err)
	}
	// clean up
	defer func() {
		if err = os.RemoveAll(dir); err != nil {
			log.Printf("error: failed to remove temporary directory: %v", err)
		}
	}()
	if err = ioutil.WriteFile(filepath.Join(dir, "example.proto"), []byte(input.Protobuf), 0644); err != nil {
		returnError(w, err)
		return
	}
	if err = ioutil.WriteFile(filepath.Join(dir, "example.output.tmpl"), []byte(input.Template), 0644); err != nil {
		returnError(w, err)
		return
	}

	// generate
	cmd := exec.Command("protoc", "-I"+dir, "--gotemplate_out=template_dir="+dir+",debug=true:"+dir, filepath.Join(dir, "example.proto")) // #nosec
	out, err := cmd.CombinedOutput()
	if err != nil {
		returnError(w, errors.New(string(out)))
		return
	}

	// read output
	content, err := ioutil.ReadFile(filepath.Join(dir, "example.output")) // #nosec
	if err != nil {
		returnError(w, err)
		return
	}

	returnContent(w, content)
}

func returnContent(w http.ResponseWriter, output interface{}) {
	payload := map[string]interface{}{
		"output": fmt.Sprintf("%s", output),
	}
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func returnError(w http.ResponseWriter, err error) {
	payload := map[string]interface{}{
		"error": fmt.Sprintf("%v", err),
	}
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}