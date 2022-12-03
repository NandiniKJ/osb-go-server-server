/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"bytes"
)

func CreateTenantNew() {
    url := "https://abimedical-application-d6.vmr7h1rf0o4.us-east.codeengine.appdomain.cloud/api/create-tenant/"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"name": "demo1234"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Authorization", "Basic cnBvbGVwZWRkaTphZG1pbjEyMw==")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

func CreateTenant() {
	client := &http.Client{}
	var data = strings.NewReader(`{
	"name": "demo1234"
}`)
	req, err := http.NewRequest("POST", "https://abimedical-application-d6.vmr7h1rf0o4.us-east.codeengine.appdomain.cloud/api/create-tenant/", data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Create tenant")
	req.Header.Set("Authorization", "Basic cnBvbGVwZWRkaTphZG1pbjEyMw==")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func CreateAdminTenant() {
	fmt.Println("Create admin tenant")
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://abimedical-application-d6.vmr7h1rf0o4.us-east.codeengine.appdomain.cloud/clients/demo1234/api/v1/core/administrator/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func ServiceInstanceDeprovision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ServiceInstanceGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ServiceInstanceLastOperationGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ServiceInstanceProvision(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp := make(map[string]string)
	resp["dashboard_url"] = "http://www.ibm.com"
	resp["description"] = "Service is created"
	resp["operation"] = "provision45"
	log.Printf("Lets Go...")
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	CreateTenantNew()
	CreateAdminTenant()
	w.Write(jsonResp)
}

func ServiceInstanceUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
