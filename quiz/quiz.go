package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type Question struct {
	text string
}

type Answer struct {
	text string
}

func main() {
	fmt.Println("Welcome to the quiz!")

	var filename = flag.String("csv", "problems", "a csv file in the format of 'question,answer'")
	var limit = flag.Int("limit", 30, "time limit (seconds) for the quiz")

	flag.Parse()
	fmt.Printf("Loading quiz from file: %s.csv\n", *filename)
	fmt.Println("Time limit: ", *limit)

	content, err := ioutil.ReadFile(*filename + ".csv")

	if err != nil {
		fmt.Printf("Failed to open the file %s\n", *filename)
		os.Exit(1)
	}

	r := csv.NewReader(strings.NewReader(string(content)))

	data := make(map[string]string)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(r)
		}

		data[record[0]] = strings.Trim(record[1], " ")
	}

	counter := 1
	correct := 0

	input := bufio.NewReader(os.Stdin)

	fmt.Println("Ready to start? Press a key!")
	_, err = input.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	for k, v := range data {
		fmt.Printf("Problem %d: %s:", counter, k)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(data))
			return
		case answer := <-answerCh:
			counter++
			if strings.Compare(strings.Trim(answer, "\n "), v) == 0 {
				correct++
			}
		}
	}
	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(data))

}
