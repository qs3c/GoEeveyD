package testDeferReturn

func DeferReturn() *MyStruct {
	a := &MyStruct{
		name: "Jerry",
	}

	defer func() {
		a.name = "Tom"
	}()
	return a
}

type MyStruct struct {
	name string
}
