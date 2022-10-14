package main

import (
	"fmt"
	"os"

	"github.com/jwalton/gchalk"
)

func main() {
	fmt.Println(gchalk.Blue("This line is blue"))

	// Combine styled and normal strings
	fmt.Println(gchalk.Blue("Hello") + " World" + gchalk.Red("!"))

	// Compose multiple styles using the chainable API
	fmt.Println(gchalk.WithBlue().WithBgRed().Bold("Hello world!"))

	// Pass in multiple arguments
	fmt.Println(gchalk.Blue("Hello", "World!", "Foo", "bar", "biz", "baz"))

	// Nest styles
	fmt.Println(gchalk.Green(
        "I am a green line " +
        gchalk.WithBlue().WithUnderline().Bold("with a blue substring") +
        " that becomes green again!",
	))

	// Use RGB colors in terminal emulators that support it.
	fmt.Println(gchalk.WithRGB(123, 45, 67).Underline("Underlined reddish color"))
	fmt.Println(gchalk.WithHex("#DEADED").Bold("Bold gray!"))

	// Use color name strings
	fmt.Println(gchalk.StyleMust("blue")("Hello World!"))


	// Write to stderr:
	os.Stderr.WriteString(gchalk.Stderr.Red("Ohs noes!\n"))
}
