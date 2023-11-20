// link to kata https://www.codewars.com/kata/585d7d5adb20cf33cb000235/train/go
package main

func FindUniq(arr []float32) float32 {
	var unique float32
	m := make(map[float32]int)

	for _, num := range arr {
		m[num] = m[num] + 1
	}

	for k, e := range m {
		if e == 1 {
			unique = k
		}
	}

	return unique
}
