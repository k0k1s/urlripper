//main file of the urlripper

package main

import (
    "fmt"
    "github.com/fatih/color"
)

func main() {
    displayTitle()

    for {
        choice := displayMenu()

        if choice == "1" {
            fmt.Println("Enter a URL to scan:")
            var inputURL string
            fmt.Scanln(&inputURL)

            report := scanURL(inputURL)
            report = checkURLReputation(inputURL, report)
            saveReportToFile(report)

            postScanMenu(report)
        } else if choice == "2" {
            fmt.Println("Exiting the application.")
            break
        } else {
            fmt.Println("Invalid choice, please try again.")
        }
    }
}

func displayTitle() {
    bigTitle := `
 _    _  ____    _          
| |  | ||  __ \ | |          
| |  | || |_| | | |
| |  | || |_  / | |
| |__| || | \ \ | |__
 \____/ |_|  \_\|____| RIPPER
    `
    smallTitle := "@github/k0k1s"
    color.New(color.FgRed).Add(color.Bold).Println(bigTitle)
    color.New(color.FgGreen).Println(smallTitle)
    fmt.Println("====================")
}

func displayMenu() string {
    fmt.Println("Menu:")
    fmt.Println("1. Enter a URL to scan")
    fmt.Println("2. Exit")
    fmt.Print("Choose an option: ")

    var choice string
    fmt.Scanln(&choice)
    return choice
}

func postScanMenu(report URLReport) {
    for {
        fmt.Println("Options:")
        fmt.Println("1. Save as text")
        fmt.Println("2. Enter another URL")
        fmt.Println("3. Exit")
        fmt.Print("Choose an option: ")

        var choice string
        fmt.Scanln(&choice)

        if choice == "1" {
            err := saveReportToFile(report)
            if err != nil {
                fmt.Println("Error saving report:", err)
            } else {
                fmt.Println("Report saved successfully.")
            }
        } else if choice == "2" {
            break
        } else if choice == "3" {
            fmt.Println("Exiting the application.")
            return
        } else {
            fmt.Println("Invalid choice, please try again.")
        }
    }
}
