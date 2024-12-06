package commands

type Add struct {
	A int
	B int

	Result struct {
		C     int
		Error error
	}
}

type Subtraction struct {
	A int
	B int

	Result struct {
		C     int
		Error error
	}
}

type Multiplication struct {
	A int
	B int

	Result struct {
		C     int
		Error error
	}
}

type Division struct {
	A int
	B int

	Result struct {
		C     int
		Error error
	}
}
