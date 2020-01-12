# mailer
My Email sender helper package

## Installation
```bash
go get github.com/dkuye/mailer
```

## Usage
It require the following configuration in you .env file
```
MAIL_HOST=smpt.somedomain.com
MAIL_PORT=587
MAIL_USERNAME=noreply@somedomain.com
MAIL_PASSWORD=password
MAIL_FROM_ADDRESS=noreply@somedomain.com
MAIL_FROM_ALIAS="Company Name"
```


```go
package main

import (
    "get github.com/dkuye/mailer"
    "github.com/joho/godotenv"
)

func main(){
    // Open .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
    // Connect to you database
    mailer.Send(
        mailer.Template{
            Path:    "./public/emailTemplates",
            Layout:  "layout.html",
            File:    "password-reset.html",
            Email:   "som@domain.com",
            Subject: "Mail subject...",
            Data: map[string]string{"SomeKey": "SomeValue"),},
        })
}
```
