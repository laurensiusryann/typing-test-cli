package typingtest

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type TestResult struct {
	Target   string
	Typed    string
	Duration float64
	WPM      float64
	Accuracy float64
}

var sentences = []string{
	"Go adalah bahasa pemrograman yang menyenangkan",
	"Aplikasi typing speed test bisa berguna untuk latihan mengetik",
	"The only way to do great work is to love what you do",
	"Practice makes perfect when learning to code in Go",
	"It always seems impossible until it's done",
}

func PickSentence() string {
	rand.Seed(time.Now().UnixNano())
	return sentences[rand.Intn(len(sentences))]
}

func StartTest(target string) TestResult {
	fmt.Println("Kalimat:", target)
	fmt.Println("\nTekan Enter untuk mulai...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	start := time.Now()

	fmt.Print("Ketik ulang: ")
	reader := bufio.NewReader(os.Stdin)
	typed, _ := reader.ReadString('\n')
	typed = strings.TrimSpace(typed)

	duration := time.Since(start).Seconds()
	wpm := calculateWPM(target, duration)
	accuracy := calculateAccuracy(target, typed)

	return TestResult{
		Target:   target,
		Typed:    typed,
		Duration: duration,
		WPM:      wpm,
		Accuracy: accuracy,
	}
}

func calculateWPM(sentence string, seconds float64) float64 {
	words := len(strings.Split(sentence, " "))
	return float64(words) / (seconds / 60)
}

func calculateAccuracy(target, typed string) float64 {
	targetWords := strings.Split(target, " ")
	typedWords := strings.Split(typed, " ")

	correct := 0
	for i := 0; i < len(targetWords) && i < len(typedWords); i++ {
		if targetWords[i] == typedWords[i] {
			correct++
		}
	}

	return (float64(correct) / float64(len(targetWords))) * 100
}
