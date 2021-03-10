package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// strconv

	strNum := "42"
	num, err := strconv.ParseInt(strNum, 10, 64)
	// strconv.Atoi // "42" -> 42
	// strconv.Itoa // 42-> "42"

	if err != nil {
		panic(err)
	}

	fmt.Println("integer from string:", num)

	// strings

	fmt.Println("strings.Contains", strings.Contains("humanity", "man"))
	fmt.Println("strings.Count", strings.Count("hello", "l"))
	fmt.Println("strings.HasPrefix", strings.HasPrefix("username", "user"))
	fmt.Println("strings.HasSuffix", strings.HasSuffix("username", "name"))
	fmt.Println("strings.Index", strings.Index("abc", "b"))
	fmt.Println("strings.Join", strings.Join([]string{"admin", "user"}, "|"))
	fmt.Println("strings.Repeat", strings.Repeat("a", 3))
	fmt.Println("strings.Replace", strings.Replace("ga ga ga", "a", "o", -1))      // all entries
	fmt.Println("strings.Replace", strings.Replace("username", "name", "role", 1)) // one entry
	fmt.Println("strings.Split", strings.Split("one,two,three", ","))
	fmt.Println("strings.ToLower", strings.ToLower("UPPER"))
	fmt.Println("strings.ToUpper", strings.ToUpper("lower"))
	fmt.Println("strings.TrimSpace", strings.TrimSpace("   space   "))

	/*
		regexp

		you should compile your RegExp first
		using RegExp in runtime results bag performance
	*/

	reg, err := regexp.Compile("@")

	if err != nil {
		panic(err)
	}

	fmt.Println("regexp", reg.FindStringSubmatch("company@email.com"))

	var source = "Hello World"
	var re = regexp.MustCompile("World")
	fmt.Println(source, re.MatchString(source))
}
