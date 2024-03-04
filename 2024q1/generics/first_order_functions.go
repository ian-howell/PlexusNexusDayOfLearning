package generics

func Map[T, U any](f func(T) U, a []T) []U {
	res := make([]U, len(a))
	for i, v := range a {
		res[i] = f(v)
	}
	return res
}

func Filter[T any](f func(T) bool, a []T) []T {
	var res []T
	for _, v := range a {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func Reduce[T, U any](f func(U, T) U, a []T, init U) U {
	res := init
	for _, v := range a {
		res = f(res, v)
	}
	return res
}
