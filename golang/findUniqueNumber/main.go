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

func main() {
	FindUniq([]float32{0, 0, 0, 0, 0.55})
	FindUniq([]float32{1, 2, 1, 1, 1})
	FindUniq([]float32{4.4, 4.4, 4.4, 5, 4.4})
}

// Best codewars solution
// func FindUniq(arr []float32) float32 {
//	 if arr[0] != arr[1] && arr[1] == arr[2] { return arr[0] }
//	 for i, v := range arr[1:] {
//		 if v != arr[i] { return v }
//	 }
// 	 return 0
// }
