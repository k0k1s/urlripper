//saveas  a text

package main

import (
    "fmt"
    "os"
    "github.com/fatih/color"
)

func saveReportToFile(report URLReport) error {
    file, err := os.Create("report.txt")
    if err != nil {
        return fmt.Errorf("could not create file: %v", err)
    }
    defer file.Close()

    writer := fmt.Sprintf("URL Scanning Report:\n")
    writer += "--------------------\n"
    writer += fmt.Sprintf("URL               : %s\n", report.URL)
    writer += fmt.Sprintf("HTTP Status       : %d\n", report.HTTPStatus)
    writer += fmt.Sprintf("Content Type      : %s\n", report.ContentType)
    writer += fmt.Sprintf("Headers           : %v\n", report.Headers)
    if report.RedirectedURL != "" {
        writer += fmt.Sprintf("Redirected URL    : %s\n", report.RedirectedURL)
    }
    if report.CertIssuer != "" {
        writer += fmt.Sprintf("Certificate Issuer: %s\n", report.CertIssuer)
        writer += fmt.Sprintf("Certificate Expiry: %s\n", report.CertExpiry)
    }
    maliciousStatus := "No"
    if report.IsMalicious {
        maliciousStatus = "Yes"
    }
    writer += fmt.Sprintf("Is Malicious      : %s\n", maliciousStatus)
    writer += fmt.Sprintf("Reason            : %s\n", report.Reason)
    writer += "--------------------\n"

    _, err = file.WriteString(writer)
    if err != nil {
        return fmt.Errorf("could not write to file: %v", err)
    }

    return nil
}

func printReport(report URLReport) {
    color.New(color.FgCyan).Println("URL Scanning Report:")
    color.New(color.FgCyan).Println("--------------------")
    fmt.Printf("%-20s: %s\n", "URL", report.URL)
    fmt.Printf("%-20s: %d\n", "HTTP Status", report.HTTPStatus)
    fmt.Printf("%-20s: %s\n", "Content Type", report.ContentType)
    fmt.Printf("%-20s: %v\n", "Headers", report.Headers)
    if report.RedirectedURL != "" {
        fmt.Printf("%-20s: %s\n", "Redirected URL", report.RedirectedURL)
    }
    if report.CertIssuer != "" {
        fmt.Printf("%-20s: %s\n", "Certificate Issuer", report.CertIssuer)
        fmt.Printf("%-20s: %s\n", "Certificate Expiry", report.CertExpiry)
    }
    maliciousStatus := "No"
    statusColor := color.FgGreen
    if report.IsMalicious {
        maliciousStatus = "Yes"
        statusColor = color.FgRed
    }
    color.New(statusColor).Printf("%-20s: %v\n", "Is Malicious", maliciousStatus)
    fmt.Printf("%-20s: %s\n", "Reason", report.Reason)
    color.New(color.FgCyan).Println("--------------------")
}
