package student

import "fmt"

func Split(str, charset string) []string {
	l := 0
	l1 := 0

	for range str {
		l++
	}
	for range charset {
		l1++
	}
	if l1 > l {
		arr := make([]string, 1)
		arr[0] = str
		return arr
	}

	a := []rune(str)
	p := ""

	count := 0
	count2 := 1
	//count3 := 0
	ch := false
	ch2 := false
	ch3 := false
	for i := 0; i < l-l1+1; i++ {

		for j := i; j < i+l1; j++ {

			p = p + string(a[j])

		}

		if p == charset {
			if ch3 == false {
				count++
				ch3 = true
			}

			ch = true
			// if ch != ch2 && i != 0 {
			// 	count++
			// }

			if ch == ch2 {
				count++

			}
			if ch != ch2 {
				count--
			}
			if i == 0 {
				count--
			}
			i = i + l1 - 1
			if l-i-1 < l1 {
				fmt.Println("LEI", l-i-1)
				count = count + 2
				fmt.Println("LEI", count)
			}
		} else {

			ch = false
			if i == 0 {
				count--
			}

		}
		if i == l-l1 && p != charset {
			count++
		}
		//	fmt.Printf(p)

		p = ""
		if ch2 != ch {
			count++
			//	fmt.Println("LEIp", count)
			ch2 = ch
		}
		//fmt.Println(count)
	}

	arr := make([]string, count)
	//------------------------------------
	ch = false
	ch2 = false
	//	ch3 = false
	count = 0
	for i := 0; i < l-l1+1; i++ {

		for j := i; j < i+l1; j++ {

			p = p + string(a[j])

		}

		//	fmt.Println(count2)
		if p == charset {
			// if ch3 == false {
			// 	count++
			// 	ch3 = true
			// }
			ch = true
			if ch != ch2 && i != 0 {

				arr[count+1] = str[i-count2+1 : i]

			}
			if ch != ch2 || ch == ch2 && i == l-l1 {
				count--
			}
			count2 = 0
			if ch == ch2 {
				count++
			}
			if i == 0 {
				count--
			}
			i = i + l1 - 1

			if l-i-1 < l1 {

				arr[count+3] = str[i+1 : l]
			}

		} else {

			ch = false
			if i == 0 {
				count--
			}

		}
		if i == l-l1 && p != charset {
			if ch2 != ch {
				arr[count+2] = str[i-count2+1 : l]
			} else {
				arr[count+1] = str[i-count2+1 : l]
				fmt.Println(count, "------", l)
			}

		}
		count2++

		//	fmt.Printf(p)

		p = ""
		if ch2 != ch {
			count++
			ch2 = ch
		}
		//	fmt.Println(count)
	}

	//----------------------------------------------

	return arr
}
