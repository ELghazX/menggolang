package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (v *NotFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"Validation Error"}
	}

	if id != "ELghaz" {
		return &NotFoundError{"Data not found"}
	}

	// oke
	return nil
}

func main() {
	err := SaveData("Bukan Elghaz", nil)

	if err != nil {
		// eror
		if validationErr, ok := err.(*ValidationError); ok {
			fmt.Println("Validation error", validationErr.Error())
		} else if notFoundErr, ok := err.(*NotFoundError); ok {
			fmt.Println("not found error ", notFoundErr.Error())
		} else {
			fmt.Println("Unknown Error", err.Error())
		}

		// atau bisa paka switch case
		// switch finalError := err.(type) {
		// case *ValidationError:
		// 	fmt.Println("Validation Error", finalError.Error())
		// case *NotFoundError:
		// 	fmt.Println("not found error", finalError.Error())
		// default:
		// 	fmt.Println("unknown error", finalError.Error())
		// }
	} else {
		// sukses
		fmt.Println("Sukses")
	}
}
