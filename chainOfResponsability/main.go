// The objective is to create a chain of responsabilities each one of them with
// validations that allows the chain to continue executing

package main

func main() {

	cashier := &cashier{}

	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{}
	//Patient visiting
	reception.execute(patient)
}
