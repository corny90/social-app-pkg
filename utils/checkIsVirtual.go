package utils

func CheckIsVirtual(userType string) bool {

	if userType == "v" || userType == "user-virtual-old" || userType == "user-virtual" || userType == "user-v" {
		return true
	}

	return false
}
