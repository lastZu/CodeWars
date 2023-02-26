package main

func DigitalRoot(n int) int {
	temp := n
	root := 0

	for temp != 0 {
		root += temp % 10
		temp /= 10
		if temp == 0 && root/10 != 0 {
			temp = root
			root = 0
		}
	}

	return root
}

func DigitalRoot2(n int) int {
	temp := n

	for temp >= 10 {
		temp = temp/10 + temp%10
	}

	return temp
}
