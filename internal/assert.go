package internal

func Assert(context interface {
	Helper()
	Errorf(string, ...interface{})
}) func(bool, interface{}) bool {
	return func(test bool, subject interface{}) bool {
		if test {
			return test
		}

		context.Helper()
		context.Errorf("\n\nunexpected %#v\n\n", subject)

		return test
	}
}
