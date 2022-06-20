package helper

func PanicIfError(err Error) {
	if err != nil {
		panic(err)
	}
}
