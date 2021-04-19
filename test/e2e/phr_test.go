package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.com/alexgtn/buildit/pkg/plantHireRequest/delivery/http/dto"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestPhrCreateAndRead(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	startTime := time.Now()
	endTime := time.Now()
	url := fmt.Sprintf("%s/phr", getHost())
	phr := &dto.PlantHireRequest{
		SiteEngineerName:      "siteEngineerName",
		ConstructionSiteName:  "constructionSiteName",
		Comment:               "comment",
		Cost:                  123,
		PlantInventoryEntryID: 1,
		StartDate:             &startTime,
		EndDate:               &endTime,
	}
	m, err := json.Marshal(phr)
	if err != nil {
		t.Fatalf("Could not marshal json, err %v", err)
	}
	client := http.Client{
		Timeout: 1 * time.Minute,
	}

	// Execute POST
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(m))
	if err != nil {
		t.Fatalf("Could not create post request, err %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Could not send get request, err %v", err)
	}
	if resp.StatusCode != 201 {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Unexpected http code=%d body=%s", resp.StatusCode, body)
	}

	// Execute GET
	url = fmt.Sprintf("%s/phr", getHost())
	req2, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatalf("Could not create get request, err %v", err)
	}
	req2.Header.Set("Content-Type", "application/json")

	resp2, err := client.Do(req2)
	if err != nil {
		t.Fatalf("Could not send get request, err %v", err)
	}

	if resp2.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp2.Body)
		t.Fatalf("api error code=%d body=%s", resp2.StatusCode, body)
	}
	responseGetPhrs := []*dto.PlantHireRequest{}
	err = json.NewDecoder(resp2.Body).Decode(&responseGetPhrs)
	if err != nil {
		t.Fatalf("Could not decode json, err %v", err)
	}

	// Assert that the length of the array with PHRs is 1
	if len(responseGetPhrs) != 1 {
		t.Errorf("Number of PHRs returned is wrong .\nCurrent: '%d'\nExpected: '%d'", len(responseGetPhrs), 1)
	}
}

func getHost() string {
	p := os.Getenv("TEST_MAIN_SERVICE")
	if p != "" {
		return p
	}
	return "http://localhost:8080"
}
