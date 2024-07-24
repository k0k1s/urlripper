//scanning the url
//
package main

import (
    "crypto/tls"
    "fmt"
    "net/http"
    "net/url"
    "strings"
    "time"
)

type URLReport struct {
    URL           string
    IsMalicious   bool
    Reason        string
    HTTPStatus    int
    ContentType   string
    Headers       http.Header
    RedirectedURL string
    CertIssuer    string
    CertExpiry    time.Time
}

func scanURL(inputURL string) URLReport {
    parsedURL, err := url.Parse(inputURL)
    if err != nil {
        return URLReport{
            URL:         inputURL,
            IsMalicious: true,
            Reason:      "Invalid URL format",
        }
    }

    if !strings.HasPrefix(parsedURL.Scheme, "http") {
        return URLReport{
            URL:         inputURL,
            IsMalicious: true,
            Reason:      "Unsupported URL scheme",
        }
    }

    client := http.Client{
        Timeout:       10 * time.Second,
        CheckRedirect: redirectPolicyFunc,
    }

    resp, err := client.Get(inputURL)
    if err != nil {
        return URLReport{
            URL:         inputURL,
            IsMalicious: true,
            Reason:      "Unable to reach URL",
        }
    }
    defer resp.Body.Close()

    report := URLReport{
        URL:        inputURL,
        HTTPStatus: resp.StatusCode,
        Headers:    resp.Header,
    }

    if resp.StatusCode != http.StatusOK {
        report.IsMalicious = true
        report.Reason = fmt.Sprintf("Received non-OK HTTP status: %d", resp.StatusCode)
        return report
    }

    contentType := resp.Header.Get("Content-Type")
    report.ContentType = contentType
    if !strings.HasPrefix(contentType, "text/html") {
        report.IsMalicious = true
        report.Reason = "Unexpected content type"
        return report
    }

    if strings.Contains(inputURL, "malicious.com") {
        report.IsMalicious = true
        report.Reason = "URL contains known malicious domain"
        return report
    }

    if resp.Request.URL.String() != inputURL {
        report.RedirectedURL = resp.Request.URL.String()
    }

    if parsedURL.Scheme == "https" {
        conn, err := tls.Dial("tcp", parsedURL.Host+":443", nil)
        if err == nil {
            defer conn.Close()
            cert := conn.ConnectionState().PeerCertificates[0]
            report.CertIssuer = cert.Issuer.CommonName
            report.CertExpiry = cert.NotAfter
        }
    }

    report.IsMalicious = false
    report.Reason = "URL appears safe"
    return report
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
    if len(via) >= 10 {
        return fmt.Errorf("stopped after 10 redirects")
    }
    return nil
}
