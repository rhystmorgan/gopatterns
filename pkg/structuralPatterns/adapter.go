package structuralpatterns

import "fmt"

// Adapter Pattern

// Convert an interface of a class into another interface clients expect. Adapter lets
// classes work together that couldn't otherwise because of incompatible surfaces

// Target
// Client
// Adaptee
// Adapter

// Adapter Interface
type Notification interface {
	Send()
}

// Adaptee
type SMS struct {
	Login   func()
	SetPort func()
	SendSms func()
}

func Login() bool {
	return true
}

func SetPort() int {
	return 404
}

func SendSms() {
	fmt.Println("Sending SMS ...")
}

func NotifyUsers(notifier Notification) {
	notifier.Send()
}

func SmsNotifier() {
	smsNotify := 
}
