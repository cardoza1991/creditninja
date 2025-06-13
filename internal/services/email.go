package services

import "log"

func SendEmail(to, sub, body string) {
    // Integrate real email provider later
    log.Printf("Email to %s: %s", to, sub)
}
