package main

import (
	"fmt"
	"typing-speed/typingtest"
)

func main() {
	fmt.Println("=== Typing Speed Test ===")
	fmt.Println("")

	sentence := typingtest.PickSentence()
	result := typingtest.StartTest(sentence)

	fmt.Println("\n=== Hasil ===")
	fmt.Printf("Target   : %s\n", result.Target)
	fmt.Printf("Diketik  : %s\n", result.Typed)
	fmt.Printf("Waktu    : %.2f detik\n", result.Duration)
	fmt.Printf("Kecepatan: %.2f WPM\n", result.WPM)
	fmt.Printf("Akurasi  : %.2f%%\n", result.Accuracy)
}
