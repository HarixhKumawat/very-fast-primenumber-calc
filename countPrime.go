package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"time"
)

func worker(nums chan int, results chan int) {
	for num := range nums {
		prime := 1
		till := int(math.Sqrt(float64(num)))
		if num%2 == 0 || num%3 == 0 {
			results <- 0
		} else {
			// for id := range primes {
			// 	fmt.Print()
			// 	if (num % id) == 0 {
			// 		prime = 0
			// 		break
			// 	}
			// }
			for loop := 2; loop < till; loop++ {
				if (num % loop) == 0 {
					prime = 0
					break
				}
			}
			if prime == 1 {
				//fmt.Print("prime")
				results <- num
			} else {
				results <- 0
			}
		}

	}
}

func main() {
	csvFile, err := os.Create("employee.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()
	csvwriter := csv.NewWriter(csvFile)

	start := time.Now()

	nums := make(chan int, 1000)
	results := make(chan int)

	var prime []int
	var write [][]string

	fmt.Printf("program srtarted")
	for i := 0; i < cap(nums); i++ {
		go worker(nums, results)
	}
	go func() {
		for i := 1; i <= 10000000; i++ {
			//fmt.Print(ports)
			nums <- i
		}
	}()

	for i := 1; i <= 10000000; i++ {
		pri := <-results
		if pri != 0 {
			//fmt.Println(pri)
			prime = append(prime, pri)
		}
	}
	sort.Ints(prime)
	for id, sq := range prime {
		write = append(write, []string{1: strconv.Itoa(id), 2: strconv.Itoa(sq)})
		//fmt.Printf("the number iss... %d\n", sq)
	}

	err = csvwriter.WriteAll(write) // calls Flush internally
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
