//package main
//
//import (
//	"log"
//	"net/smtp"
//)
//
//func main() {
//	// Set up authentication information.
//	auth := smtp.PlainAuth("", "bdauren06@gmail.com", "lrhcvmwjvdkbjrkc", "smtp.gmail.com")
//
//	// Connect to the server, authenticate, set the sender and recipient,
//	// and send the email all in one step.
//	to := []string{"kukareku_147@mail.ru"}
//	msg := []byte("To: kukareku_147@mail.ru\r\n" +
//		"Subject: discount Gophers!\r\n" +
//		"\r\n" +
//		"This is the email body.\r\n")
//	err := smtp.SendMail("smtp.gmail.com:587", auth, "bdauren06@gmail.com", to, msg)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("ui/html/*.tmpl"))

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/application", application)
	//http.HandleFunc("/register", registerHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
