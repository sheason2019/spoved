package utils

func MapTo[Item any, Result any](arr []Item, fn func(item Item, index int) Result) []Result {
	res := make([]Result, len(arr))
	for i, v := range arr {
		res[i] = fn(v, i)
	}

	return res
}
