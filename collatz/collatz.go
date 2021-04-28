package collatz

func CollatzSteps(n int) int {
	return len(CollatzNumbers(n))
}

func CollatzNumbers(n int) []int {
	ns := []int{}
	for {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		ns = append(ns, n)
		if n == 1 {
			break
		}
	}
	return ns
}
