package main

/*import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//using ``, "",\"\"

	//len returns byte at position not the string character
	str := "t@r"
	fmt.Println(len(str))
	fmt.Println(str[2]) //returns byte at at position 2

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println("\n" + strings.Repeat("#", 20))
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}
	fmt.Println("\n" + strings.Repeat("#", 20))
	for _, r := range str {
		fmt.Printf("%c", r)
	}
	fmt.Println("\n" + strings.Repeat("#", 20))
	b := utf8.RuneCountInString(str)
	fmt.Println(b)

	// to see the string in rune form and as a byte
	rs := "special character"
	rn := []rune(rs)
	fmt.Println(rn)
	fmt.Println(string(rn[0:3]))

	//strings package fns
	oh := fmt.Println
	comparison := strings.Contains(rs, "pecc")
	oh(comparison)

	comparison2 := strings.ContainsAny(rs, "pecc")
	oh(comparison2)

	comparison3 := strings.ContainsRune(rs, 't')
	oh(comparison3)

	comparison4 := strings.Count(rs, "e")
	oh(comparison4)

	oh(strings.ToUpper(rs))

	comparison5 := strings.EqualFold(rs, string(rn))
	oh(comparison5)

	comparison6 := strings.ReplaceAll(rs, "a", "b")
	oh(comparison6)

	comparison7 := strings.Split(rs, "e")
	oh(comparison7)

}*/
