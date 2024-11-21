package main

import (
	"fmt"
)

func FirstOrLast(n, s, d, x int) int {
	if d > n {
		fmt.Println(n) // esli 1 ili posledmnij tomat
	} else {
		x = x + d
	}
	return x
}

func OK(n, s, d, x int) int {
	if s-d > 1 && s+d < n {
		x = x + d + d // ne dohodit do krajev
	}
	return x
}

func TillEndOr1(n, s, d, x int) int {
	var d2 int
	if s-d <= 1 {
		d2 := s - 1 //do 1 tomata
		return d2
	}
	if s+d >= n {
		d2 := n - s
		return d2
	}
	d3 := d - d2
	x = x + d2 + d2 + d3 //  dohodit do 1
	fmt.Println("Sarkanu tomatu skaits: ", x, " pēc ", d, "dienam.")
	return x
}
func main() {
	var n int     // tomātu skaits
	var s int     //location
	var d int     //dienas
	var x int = 1 //tomātu skaits pēc d dienam

	fmt.Println("Kopejais tomatu skaits: ")
	fmt.Scan(&n)
	fmt.Println("Sarkana tomātu atrašānas vieta: ")
	fmt.Scan(&s)
	fmt.Println("Dienu skaits: ")
	fmt.Scan(&d)

	if s > n {
		fmt.Println("Tomāts nav atrasts!")
	} else { //programma

		//
		if s == 1 || s == n {
			fmt.Println("Tomatu skaits: ", FirstOrLast(n, s, d, x))

		} else {
			if s-d > 1 && s+d < n {
				fmt.Println("Tomatu skaits: ", OK(n, s, d, x))
			} else {
				if s+d >= n || s-d <= 1 {
					fmt.Println("Tomatu skaits: ", TillEndOr1(n, s, d, x))
				}
			}
		}
	}
}
