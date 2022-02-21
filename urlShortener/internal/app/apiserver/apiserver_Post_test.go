package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"shortener/internal/app/config"
	"testing"
)

func init() {
	//change test dir for correct default config parse
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../../")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestAPIServer_CreateURL(t *testing.T) {
	tests := map[string]string{
		"longUrl": "https://vk.com/letsffucktheworld",
	}
	jsonData, err := json.Marshal(tests)
	if err != nil {
		log.Fatal(err)
	}
	config, _ := config.NewConfig()
	s := New(config)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//s := New(config)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/encode/", bytes.NewBuffer(jsonData))
	s.CreateURL().ServeHTTP(rec, req)
	fmt.Println(rec.Body.String())
}
