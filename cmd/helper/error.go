package helper

func PrintIfError(err error) {
	if err != nil {
		panic(err)
	}
}
