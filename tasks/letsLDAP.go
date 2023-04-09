package tasks

import (
	"fmt"
	"log"
)

type Employee struct {
	hsbc_ad_Grade string
	displayName   string
	uid           string
	memberOf      []string
}

func QueryAndSave(employeeID string) error {
	emp, err1 := QueryEmployeeInfoFromLDAP(employeeID)
	if err1 != nil {
		log.Fatal(err1)
		return fmt.Errorf("Could not find employee : %f from LDAP.", employeeID)
	}
	err2 := SaveEmployeeInfoToBQ(emp)

	if err2 != nil {
		log.Fatal(err2)
		return fmt.Errorf("Could not find employee : %f from LDAP.", employeeID)
	} else {
		return nil
	}
}

func QueryEmployeeInfoFromLDAP(employeeID string) (Employee, error) {
	println("query LDAP by employeeID:" + employeeID)
	var emptyList = []string{}
	emp := Employee{
		hsbc_ad_Grade: "4",
		displayName:   "John Duo",
		uid:           "43780000",
		memberOf:      emptyList,
	}
	return emp, nil
}

func SaveEmployeeInfoToBQ(emp Employee) error {
	println("save employeeID:" + emp.uid)
	return nil
}
