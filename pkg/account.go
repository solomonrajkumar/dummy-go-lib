package dummylib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Tenant struct {
	ID                   int            `json:"id"`
	Name                 string         `json:"name"`
	Domain               string         `json:"domain"`
	AlternateDomain      interface{}    `json:"alternate_domain"`
	DxcTenantHomePagesID int            `json:"dxc_tenant_home_pages_id"`
	Subscriptions        []Subscription `json:"subscriptions"`
	ImageUrls            []string       `json:"image_urls"`
	ImageUrlsQualified   []string       `json:"image_urls_qualified"`
	TenantHomePage       TenantHomePage `json:"tenant_home_page"`
}

type TenantHomePage struct {
	ID        int       `json:"id"`
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Subscription struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c* Client) GetTenant(tenantId int) (*Tenant, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}

	fmt.Println(res)

	defer res.Body.Close()

	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading HTTP response body: %v", err)
	}
	fmt.Println(string(responseBytes))

	var tenant Tenant
	err = json.Unmarshal(responseBytes, &tenant)
	if err != nil {
		return nil, err
	}

	return &tenant, nil

	//res, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	log.Fatalf("error sending HTTP request: %v", err)
	//}
	//responseBytes, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatalf("error reading HTTP response body: %v", err)
	//}
	//
	//log.Println("Response:", string(responseBytes))
	//return nil
}