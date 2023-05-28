package errorUtils

import logUtils "github.com/AhmedEnnaime/GoQl/utils/logger"

func HandleError(err error, msg ...string) {
	if err == nil {
		return
	}

	var errorMessage string

	if len(msg) == 0 {
		errorMessage = "Something went wrong"
	} else {
		errorMessage = msg[0]
	}

	logUtils.Error(err, errorMessage)

}