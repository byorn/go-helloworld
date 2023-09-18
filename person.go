package main

type person struct {
	firstname string
	lastname  string
}

func newPerson(firstname string, lastname string) person {
	return person{
		firstname,
		lastname,
	}
}

/*
*
Go is pass by value

person - is a type of person struct value that can be passed
p - is the value of person, it is a copy of the original address location

once the method exists in the stack the address in memory is not updated. because what updated was the copy

You need to update the original address location, see next method
*/
func (p person) updateLastName(lastname string) {
	p.lastname = lastname
}

/*
*person will access only addresses that point to struct person - means pass the original address pointer in memory
*p - the address is then converted back to the value of the address.
 */
func (p *person) updateLastNameWithOriginalAddressLocationPointer(lastname string) {

	// *p is redundant
	p.lastname = lastname
}
