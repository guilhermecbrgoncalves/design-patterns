// The objective of the Template Method is to allow override methods in subclasses when the structure is the same.
// Template Method is a design pattern that defines the skeleton of an algorithm in the superclass but allows subclasses to override specific steps of the algorithm without changing its structure.
// In this example we create 2 differente templates: email and sms. Each of these has the same the same methods (name) but
// different outputs. Then to use these templates according with the need we create an object and inject which template we would like to use.

package main

import "fmt"

func main() {
	smsOTP := &sms{}
	o := otp{
		iOtp: smsOTP, // Here we inject the "Template"
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &email{}
	o = otp{
		iOtp: emailOTP, // Here we inject the "Template"
	}
	o.genAndSendOTP(4)

}
