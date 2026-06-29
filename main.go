package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	{

		fmt.Println("Hello World")
		fmt.Println("Hello World")

	}
	countingLettersFunction()

	timeCalc()

	structuredLogging()
}

func structuredLogging() {

	fmt.Println("\nStructured logging")
	{
		slog.Debug("debug log message")
		slog.Info("info log message")
		slog.Warn("warning log message")
		slog.Error("error log message")
	}
	{
		fmt.Println("\nStructure logginf with custom fields")
		userID := 5
		userName := "Tony McClay"
		slog.Info("user ", "id", userID, "name", userName)
	}
	{
		fmt.Println("\nStructure logginf with custom json handler")

		options := &slog.HandlerOptions{Level: slog.LevelDebug}
		handler := slog.NewJSONHandler(os.Stderr, options)
		mySlog := slog.New(handler)
		now := time.Now()
		id := "Tony"
		mySlog.Debug("debug log",
			"id", id,
			"currentTime", now)
		mySlog.Debug("debug log message")
		mySlog.Info("info log message")
		mySlog.Warn("warning log message")
		mySlog.Error("error log message")
	}
}

func timeCalc() {

	{
		fmt.Println("\n2 hours aand 30 min in go time:")
		d := 2*time.Hour + 30*time.Minute
		fmt.Println(d)
		fmt.Printf("%v is of type %T\n", d, d)
	}
	{
		fmt.Println("\nparse time duration string 2h25m")
		d, err := time.ParseDuration("2h45m08s")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v is of type %T\n", d, d)
	}
	{
		//sleep for 30 sec
		fmt.Println("Sleeping for 3 sec")

		now := time.Now()
		fmt.Printf("%v  %T\n", now, now)

		time.Sleep(3 * time.Second)

		now = time.Now()
		fmt.Printf("%v  %T\n", now, now)
		fmt.Println(now.Format(time.RFC3339))
		fmt.Println(now.Format(time.DateOnly))
		fmt.Println(now.Format(time.TimeOnly))
		fmt.Println(now.Format(time.DateTime))
		fmt.Println(now.Format(time.Kitchen))
		fmt.Println(now.Format(time.Stamp))
		fmt.Println(now.Format("01/01/2006"))
	}
	{
		fmt.Println("Time Duration")
		startTime := time.Now()
		randSec := rand.Intn(10)
		var sleepSec int
		sleepSec = randSec
		time.Sleep(time.Duration(sleepSec) * time.Second)
		elapseTime := time.Since(startTime)
		fmt.Println(elapseTime)

	}
}

func countingLettersFunction() {

	fmt.Println("counting letters function")

	s := "The quick brown fox jumps over the lazy dog"
	sReader := strings.NewReader(s)
	counts, err := countLetters(sReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(counts)
}

func countLetters(r io.Reader) (map[string]int, error) {

	fmt.Println("Counting letters")
	buff := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buff)
		for _, b := range buff[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
