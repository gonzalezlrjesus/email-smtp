package email

import (
	"fmt"
	"net/smtp"
	"reflect"
	"strings"
	"testing"
)

func TestSendEmail(t *testing.T) {
	assertThat := func(assumption, addr string, auth smtp.Auth, from string, to []string, msg []byte, expectedOutput string) {
		actualOutput := sendEmail(addr, auth, from, to, msg)
		if actualOutput != expectedOutput {
			t.Errorf("Error! Outputs are not equal Actual %v Expected %v", actualOutput, expectedOutput)
		}
	}
	addr := "smtp.gmail.com:587"

	auth := smtp.PlainAuth("", "your-email@gmail.com", "your-password", "smtp.gmail.com")

	from := "your-email@gmail.com"

	to := []string{"to-email-1@gmail.com", "to-email-2@hotmail.com"}
	concatenate := strings.Join(to, ", ")
	toMsg := fmt.Sprintf("To: %v\r\n", concatenate)

	subject := "Test Subject"
	subjectMsg := fmt.Sprintf("Subject: %v\r\n", subject)

	body := "Example body."
	bodyMsg := fmt.Sprintf("%v\r\n", body)

	msg := []byte(toMsg + subjectMsg + "\r\n" + bodyMsg)
	assertThat("Should send email from user's GMAIL account to any email and return a message of successful sent", addr, auth, from, to, msg, "email sent successfully")
}

func TestGetAddressSMTP(t *testing.T) {
	assertThat := func(assumption, expectedAddress string) {
		actualAddress := getAddressSMTP()
		if actualAddress != expectedAddress {
			t.Errorf("Error! Outputs are not equal Actual %v Expected %v", actualAddress, expectedAddress)
		}
	}
	assertThat("Should get from a environment variable the smtp address", "smtp.gmail.com:587")
}

func TestGetEmailFrom(t *testing.T) {
	assertThat := func(assumption, expectedEmail string) {
		actualEmail := getEmailFrom()
		if actualEmail != expectedEmail {
			t.Errorf("Error! Outputs are not equal Actual %v Expected %v", actualEmail, expectedEmail)
		}
	}
	assertThat("Should return the email from where the messages will be sent.", "your-email@gmail.com")
}

func TestPlainAuth(t *testing.T) {
	assertThat := func(assumption string, expectedAuth smtp.Auth) {
		actualAuth := plainAuth()
		if !reflect.DeepEqual(actualAuth, expectedAuth) {
			t.Errorf("Error! Outputs are not equal Actual %v Expected %v", actualAuth, expectedAuth)
		}
	}
	expectedAuth := smtp.PlainAuth("", "your-email@gmail.com", "your-password", "smtp.gmail.com")
	assertThat("Should returns an Auth that implements the PLAIN authentication mechanism", expectedAuth)
}

func TestGenerateMessage(t *testing.T) {
	assertThat := func(assumption string, emailList []string, subject string, body string, expectedMsg []byte) {
		actualMsg := joinMessageStructure(emailList, subject, body)
		if !reflect.DeepEqual(actualMsg, expectedMsg) {
			t.Errorf("Error! Outputs are not equal Actual %v Expected %v", actualMsg, expectedMsg)
		}
	}
	to := []string{"to-email-1@gmail.com", "to-email-2@hotmail.com"}
	concatenate := strings.Join(to, ", ")
	toMsg := fmt.Sprintf("To: %v\r\n", concatenate)

	subject := "Test Subject"
	subjectMsg := fmt.Sprintf("Subject: %v\r\n", subject)

	body := "Example body."
	bodyMsg := fmt.Sprintf("%v\r\n", body)

	expectedMsg := []byte(toMsg + subjectMsg + "\r\n" + bodyMsg)
	assertThat("Should receive the emails, subject, body and should return an array of bytes being the message", to, subject, body, expectedMsg)
}

func TestIntegrationSMTP(t *testing.T) {
	assertThat := func(assumption string, emailList []string, subject string, body string, expectedOutput string) {
		address := getAddressSMTP()
		auth := plainAuth()
		emailFrom := getEmailFrom()
		msg := joinMessageStructure(emailList, subject, body)
		actualOutput := sendEmail(address, auth, emailFrom, emailList, msg)
		if actualOutput != expectedOutput {
			t.Errorf("Error! Outputs are not equal Actual %v Expected %v", actualOutput, expectedOutput)
		}
	}
	to := []string{"to-email-1@gmail.com", "to-email-2@hotmail.com"}
	subject := "Test Subject"
	body := "Example body."
	assertThat("Should get data from different units and send email", to, subject, body, "email sent successfully")
}
