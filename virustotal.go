///if want you can add api of the virustotal here

package main

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

const apiKey = "YOUR_VIRUSTOTAL_API_KEY"

type VirusTotalResponse struct {
    Data struct {
        Attributes struct {
            LastAnalysisStats struct {
                Malicious int `json:"malicious"`
            } `json:"last_analysis_stats"`
            Reputation int `json:"reputation"`
        } `json:"attributes"`
    } `json:"data"`
}

func checkURLReputation(inputURL string, report URLReport) URLReport {
    base64URL := base64.URLEncoding.EncodeToString([]byte(inputURL))
    url := fmt.Sprintf("https://www.virustotal.com/api/v3/urls/%s", base64URL)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatalf("Failed to create request: %v", err)
    }
    req.Header.Add("x-apikey", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Request failed: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read response body: %v", err)
    }

    var vtResponse VirusTotalResponse
    if err := json.Unmarshal(body, &vtResponse); err != nil {
        log.Fatalf("Failed to unmarshal JSON: %v", err)
    }

    if vtResponse.Data.Attributes.LastAnalysisStats.Malicious > 0 {
        report.IsMalicious = true
        report.Reason = "Detected as malicious by VirusTotal"
    } else {
        report.IsMalicious = false
        report.Reason = "URL appears safe according to VirusTotal"
    }

    return report
}
