package chat

import (
	"bufio"
	"chatbot/internal/gemini"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// typeEffect simulates a typing effect for a realistic AI feel
func typeEffect(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
	fmt.Println()
}

// centerText centers text based on a given width
func centerText(text string, width int) string {
	spaces := (width - len(text)) / 2
	if spaces > 0 {
		return strings.Repeat(" ", spaces) + text
	}
	return text
}

// StartChat runs the hacker-style chatbot
func StartChat() {
	width := 60 // Adjust width for centering text
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	red := color.New(color.FgRed, color.Bold)
	cyan := color.New(color.FgCyan, color.Bold)

	fmt.Println(green.Sprintf("%s", centerText(`
	████████╗ █████╗ ███╗   ███╗███████╗███████╗███╗   ███╗
	╚══██╔══╝██╔══██╗████╗ ████║██╔════╝██╔════╝████╗ ████║
	██║   ███████║██╔████╔██║█████╗  █████╗  ██╔████╔██║
	██║   ██╔══██║██║╚██╔╝██║██╔══╝  ██╔══╝  ██║╚██╔╝██║
	██║   ██║  ██║██║ ╚═╝ ██║███████╗███████╗██║ ╚═╝ ██║
	╚═╝   ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝╚══════╝╚═╝     ╚═╝
	`, width)))

	fmt.Println(green.Sprintf("%s", centerText("🤖 GAI Chatbot - Created by Tameem Ahmad Shahzad", width)))
	fmt.Println(yellow.Sprintf("%s", centerText("=============================================", width)))
	fmt.Println(yellow.Sprintf("%s", centerText("Type 'exit' to quit", width)))
	fmt.Println(yellow.Sprintf("%s", centerText("=============================================", width)))

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(cyan.Sprintf("\nYou> "))
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Println(red.Sprintf("%s", centerText("⚠️ ERROR: Invalid input! Try again.", width)))
			continue
		}

		if input == "exit" {
			fmt.Println(green.Sprintf("%s", centerText("👋 Chatbot: Terminating session... Goodbye!", width)))
			break
		}

		// AI Thinking effect
		thinkingAnim(width, yellow)

		// Get AI response
		response, err := gemini.GetChatResponse(input)
		if err != nil {
			fmt.Println(red.Sprintf("%s", centerText("⚠️ SYSTEM ERROR: "+err.Error(), width)))
			continue
		}

		// Typing effect for AI response
		fmt.Print(green.Sprintf("\n🤖 Chatbot: "))
		typeEffect(response, 50*time.Millisecond)
	}
}

// thinkingAnim creates a flickering "AI is thinking" effect
func thinkingAnim(width int, yellow *color.Color) {
	loadingTexts := []string{"🤔 AI is processing.", "🤔 AI is processing..", "🤔 AI is processing..."}
	for i := 0; i < 3; i++ {
		fmt.Print("\r" + yellow.Sprintf("%s", centerText(loadingTexts[i], width)))
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Print("\r" + strings.Repeat(" ", width) + "\r") // Clear line
}
