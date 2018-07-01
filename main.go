package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/cbroglie/mustache"
	"github.com/lionralfs/tal-framework"
)

func renderTemplate(template *mustache.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		framework := tal.New("node_modules/tal/config")
		// TODO: read queries from request
		deviceConfigRaw, err := framework.GetConfigurationFromFilesystem("default-webkit-default", "/devices")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		appID := "sampleapp"

		re, _ := regexp.Compile("%application%")
		deviceConfigRaw = re.ReplaceAllString(deviceConfigRaw, appID)

		var deviceConfigParsed tal.DeviceConfig
		errr := json.Unmarshal([]byte(deviceConfigRaw), &deviceConfigParsed)
		if errr != nil {
			panic(errr)
		}

		template.FRender(w, map[string]string{
			"root_html_tag":        framework.GetRootHTMLTag(deviceConfigParsed),
			"headers":              framework.GetDeviceHeaders(deviceConfigParsed),
			"application_id":       appID,
			"device_configuration": deviceConfigRaw,
			"extra_body":           framework.GetDeviceBody(deviceConfigParsed),
		})
	}
}

func main() {
	tmpl, _ := mustache.ParseFile("./views/index.mustache")
	http.HandleFunc("/", renderTemplate(tmpl))
	http.Handle("/tal/", http.StripPrefix("/tal/", http.FileServer(http.Dir("./node_modules/tal"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
