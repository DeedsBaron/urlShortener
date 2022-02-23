package apiserver

import (
	"bytes"
	"fmt"
	"github.com/DeedsBaron/colors"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"shortener/internal/app/config"
	"strconv"
	"testing"
)

func init() {
	var _ = func() bool {
		testing.Init()
		return true
	}()
	//change test dir for correct default config parse
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../../")
	err := os.Chdir(dir)
	fmt.Println(os.Getwd())
	if err != nil {
		panic(err)
	}
}

func TestAPIServerInMem_CreateURL_and_GetFullURL(t *testing.T) {
	type ResponseURL struct {
		shortURL string
	}
	var Responses []ResponseURL
	fmt.Println(colors.Green + "Inmem solution tests" + colors.Res)
	tests := []struct {
		testDescription string
		testElem        []byte
	}{
		{
			testDescription: colors.Purple + "TEST1: valid URL" + colors.Res,
			testElem:        []byte(`{"longUrl": "https://www.google.com/"}`),
		},
		{
			testDescription: colors.Purple + "TEST2: another long valid URL" + colors.Res,
			testElem:        []byte(`{"longUrl": "https://en.wikipedia.org/wiki/Phrases_from_The_Hitchhiker%27s_Guide_to_the_Galaxy#The_Answer_to_the_Ultimate_Question_of_Life,_the_Universe,_and_Everything_is_42"}`),
		},
		{
			testDescription: colors.Purple + "TEST2: bad JSON" + colors.Res,
			testElem:        []byte(`123124safdasf`),
		},
		{
			testDescription: colors.Purple + "TEST3: Missing field 'longUrl' from JSON object" + colors.Res,
			testElem:        []byte(`{"thecake is a lie": "https://www.google.com/"}`),
		},
		{
			testDescription: colors.Purple + "TEST4: Extraneous data after JSON object" + colors.Res,
			testElem:        []byte(`{"longUrl": "https://www.google.com/"}, {"1234":"1234"}`),
		},
		{
			testDescription: colors.Purple + "TEST5: URL is already in base" + colors.Res,
			testElem:        []byte(`{"longUrl": "https://www.google.com/"}`),
		},
		{
			testDescription: colors.Purple + "TEST6: Not valid URL" + colors.Res,
			testElem:        []byte(`{"longUrl": "424242424424"}`),
		},
	}
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	s := New(config)
	fmt.Println(colors.Yellow + "Testing POST method \"CreateURL\"")
	s.logger.SetLevel(logrus.FatalLevel)
	for _, val := range tests {
		res := httptest.NewRecorder()
		fmt.Println(val.testDescription)
		fmt.Println(colors.Cyan+"Post req body: "+colors.Res, string(val.testElem))
		req, _ := http.NewRequest(http.MethodPost, "/encode/", bytes.NewBuffer(val.testElem))
		s.configureRouter()
		s.router.ServeHTTP(res, req)

		fmt.Println(colors.Cyan+"Response body: "+colors.Res, res.Body.String())
		if len(res.Body.String()) >= 10 {
			Responses = append(Responses, ResponseURL{shortURL: res.Body.String()[len(res.Body.String())-10 : len(res.Body.String())]})
		}
		fmt.Println(colors.Cyan+"Response code: "+colors.Res, res.Code)
	}
	fmt.Println(colors.Yellow + "Testing GET method \"GetFullURL\"" + colors.Res)
	for i, val := range Responses {
		if i == 3 {
			break
		}
		fmt.Println(colors.Purple + "TEST" + strconv.Itoa(i) + ": ")
		fmt.Println(colors.Cyan + "Get req: " + colors.Res + "\t" + config.Options.Schema + "//" + config.Options.Prefix + "/" + val.shortURL)
		res := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/"+val.shortURL, nil)
		s.router.ServeHTTP(res, req)
		fmt.Println(colors.Cyan+"Response body: "+colors.Res, res.Body.String())
		fmt.Println(colors.Cyan+"Response code: "+colors.Res, res.Code)
	}
}
