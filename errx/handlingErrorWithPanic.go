package errx

func ErrorX(err error) {
	if err != nil {
		panic(err)
	}
}
