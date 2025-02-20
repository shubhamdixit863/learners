package services

var DATA = 9

// if a function or a variable is in lower case its not available
// to the outside world ,outside packages
func getData() string {
	return "The service has been invoked"
}

func GetData() string {
	return "The service has been invoked"
}
