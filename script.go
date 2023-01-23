package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"net/http"

	
)

func main() {



	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        userInput := r.FormValue("user-input")
        fmt.Println("User input is:", userInput)
        var1 := userInput

			// var textingStr = "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."

		
			// (FUNC 1) find "(hex)" then convert the previous element from 16 to 32
			firstSoS := strings.Split(string(var1), " ") // split the input into a slice of strings (file casted from byte to string)
		
			for i := 0; i < len(firstSoS); i++ {
				if firstSoS[i] == "(hex)" {
					s2, _ := strconv.ParseInt(firstSoS[i-1], 16, 32) // func that converts from hex to dec
					firstSoS[i-1] = strconv.Itoa(int(s2))            // func that turns an int to a string. int was casted from 32 base int to 10 base int for itoa to work
				}
			}
		
			// (FUNC 2) find "(bin)" then convert the previous element from 2 to 32
			for i := 0; i < len(firstSoS); i++ {
				if firstSoS[i] == "(bin)" {
					s2, _ := strconv.ParseInt(firstSoS[i-1], 2, 32)
					firstSoS[i-1] = strconv.Itoa(int(s2))
				}
			}
		
			// (FUNC 3) find "(up)" then convert the previous element from lowercase to uppercase
			for i := 0; i < len(firstSoS); i++ {
				if firstSoS[i] == "(up)" {
					firstSoS[i-1] = strings.ToUpper(firstSoS[i-1])
				}
			}
		
			// (FUNC 4) find "(low)" then convert the previous element from uppercase to lowercase
			for i := 0; i < len(firstSoS); i++ {
				if firstSoS[i] == "(low)" {
					firstSoS[i-1] = strings.ToLower(firstSoS[i-1])
				}
			}
		
			// (FUNC 5) find "(cap)" then convert the previous element from a lowercase word to an uppercase word
			for i := 0; i < len(firstSoS); i++ {
				if firstSoS[i] == "(cap)" {
					firstSoS[i-1] = strings.Title(firstSoS[i-1])
				}
			}
		
			// (FUNC 3,4,5 sub-task) perform the conversion to the previous element(s) by the defined number in the string
			for i := 0; i < len(firstSoS); i++ { 
				if firstSoS[i] == "(up," {
					theStringNum := strings.Replace(firstSoS[i+1], ")", "", -1) // using replace to isolate number
					turnedToNum, _ := strconv.Atoi(theStringNum)                // atoi converts a string to a number
		
					for j := 0; j < turnedToNum; j++ {
						firstSoS[i-j-1] = strings.ToUpper(firstSoS[i-j-1]) 
					}
				}
			}
		
			for i := 0; i < len(firstSoS); i++ {
				if firstSoS[i] == "(low," {
					theStringNum := strings.Replace(firstSoS[i+1], ")", "", -1)
					turnedToNum, _ := strconv.Atoi(theStringNum)
		
					for j := 0; j < turnedToNum; j++ {
						firstSoS[i-j-1] = strings.ToLower(firstSoS[i-j-1])
					}
				}
			}
		
			for i := 0; i < len(firstSoS); i++ {
				if firstSoS[i] == "(cap," {
					theStringNum := strings.Replace(firstSoS[i+1], ")", "", -1)
					turnedToNum, _ := strconv.Atoi(theStringNum)
		
					for j := 0; j < turnedToNum; j++ {
						firstSoS[i-j-1] = strings.Title(firstSoS[i-j-1])
					}
				}
			}
		
			/*(FUNC 6)
			*/
			for i := 0; i < len(firstSoS); i++ { // ADD SPACE BETWEEN PUNC & JOINED WORD
				if firstSoS[i][0] == '!' || firstSoS[i][0] == '?' || firstSoS[i][0] == '.' || firstSoS[i][0] == ',' || firstSoS[i][0] == ':' || firstSoS[i][0] == ';' {
					firstSoS[i] = strings.Replace(firstSoS[i], "!", "! ", -1)
					firstSoS[i] = strings.Replace(firstSoS[i], "?", "? ", -1)
					firstSoS[i] = strings.Replace(firstSoS[i], ".", ". ", -1)
					firstSoS[i] = strings.Replace(firstSoS[i], ",", ", ", -1)
					firstSoS[i] = strings.Replace(firstSoS[i], ":", ": ", -1)
					firstSoS[i] = strings.Replace(firstSoS[i], ";", "; ", -1)
				}
			}
		
			firstSoSToString := strings.Join(firstSoS, " ") // change input from slice of strings to strings
			firstRuned := []rune(firstSoSToString)
		
			for i := 0; i < len(firstRuned); i++ { // REMOVES SPACE BEFORE PUNC
				if firstRuned[i] == '.' || firstRuned[i] == ',' || firstRuned[i] == '!' || firstRuned[i] == '?' || firstRuned[i] == ':' || firstRuned[i] == ';' {
					if firstRuned[i-1] == ' ' {
						firstRuned[i-1] = rune(0)
					}
					if firstRuned[i-2] == ')' {
						firstRuned[i-1] = ' '
					}
				}
			}
		
			// (FUNC 7) close the gap between the two single apostrophe's, for a single word or multiple words
			switchVar := false
		
			for i := 0; i < len(firstRuned); i++ { // explain all this better here below
				if firstRuned[i] == rune(39) && firstRuned[i-1] == ' ' {
					// if switchVar == false {
					if !switchVar {

						firstRuned[i+1] = rune(0)
						switchVar = true
					} else {
						firstRuned[i-1] = rune(0)
						switchVar = false
					}
				}
		
			}
		
			// (FUNC 8) instance of "a" turn to "an" if next word begins with vowel or "h"
			// notes: There are six vowels in the English language: a, e, i, o, u and sometimes y.
			firstRunedToSoS := strings.Split(string(firstRuned), " ") // convert result from SoR to SoS
		
			for i := 0; i < len(firstRunedToSoS); i++ {
				if firstRunedToSoS[i] == "A" {
					if firstRunedToSoS[i+1][0] == 'a' || firstRunedToSoS[i+1][0] == 'e' || firstRunedToSoS[i+1][0] == 'i' || firstRunedToSoS[i+1][0] == 'o' || firstRunedToSoS[i+1][0] == 'u' || firstRunedToSoS[i+1][0] == 'y' || firstRunedToSoS[i+1][0] == 'h' {
						firstRunedToSoS[i] = strings.Replace(firstRunedToSoS[i], "A", "An", -1)
		
					}
				}
				if firstRunedToSoS[i] == "a" {
					if firstRunedToSoS[i+1][0] == 'a' || firstRunedToSoS[i+1][0] == 'e' || firstRunedToSoS[i+1][0] == 'i' || firstRunedToSoS[i+1][0] == 'o' || firstRunedToSoS[i+1][0] == 'u' || firstRunedToSoS[i+1][0] == 'y' || firstRunedToSoS[i+1][0] == 'h' {
						firstRunedToSoS[i] = strings.Replace(firstRunedToSoS[i], "a", "an", -1)
		
					}
				}
			}
		
			// (FUNC TO REMOVE ALL INSTRUCTIONS)
			for i := 0; i < len(firstRunedToSoS); i++ {
				if firstRunedToSoS[i] == "(bin)" || firstRunedToSoS[i] == "(hex)" || firstRunedToSoS[i] == "(up)" || firstRunedToSoS[i] == "(cap)" {
					firstRunedToSoS[i] = ""
				}
		
				if firstRunedToSoS[i] == "(low)" {
					firstRunedToSoS[i] = ""
				}
		
				if firstRunedToSoS[i] == "(cap," || firstRunedToSoS[i] == "(up," || firstRunedToSoS[i] == "(low," {
					if firstRunedToSoS[i+1][0] >= '0' && firstRunedToSoS[i+1][0] <= '9' {
						firstRunedToSoS[i] = ""
						firstRunedToSoS[i+1] = ""
					}
				}
			}
		
			secondSoSToString := strings.Join(firstRunedToSoS, " ")
			// res := strings.Join(strings.Fields(secondSoSToString), " ") // another way to remove double spaces
		
			needToLearnRegexp := regexp.MustCompile(`\s+`) // RUN THROUGH ENTIRE STRING, & REMOVE DOUBLE WHITESPACE
			doubleSpacedRemoved := needToLearnRegexp.ReplaceAllLiteralString(secondSoSToString, " ")
		
			// REMOVES SPACE BEFORE PUNC (DUPLICATE)
			// I wanted to get rid of "(FUNC TO REMOVE ALL INSTRUCTIONS)" and write it into the specific func's but then when I did, "(FUNC 6)" was not working because I called the elements"
			// Therefore had to repeat this func to just move on
			secondRuned := []rune(doubleSpacedRemoved)
			for i := 0; i < len(secondRuned); i++ {
				if secondRuned[i] == '.' || secondRuned[i] == ',' || secondRuned[i] == '!' || secondRuned[i] == '?' || secondRuned[i] == ':' || secondRuned[i] == ';' {
					if secondRuned[i-1] == ' ' {
						secondRuned[i-1] = rune(0)
					}
		
					if secondRuned[len(secondRuned)-1] == '.' {
						continue
					}
		
					if secondRuned[len(secondRuned)-1] == ' ' {
						secondRuned[len(secondRuned)-1] = rune(0)
					}
		
					if secondRuned[i] == '.' && secondRuned[i+3] == rune(39) { // deals with "." only before closing "'" - not dynamic for other punc's
						secondRuned[i+1] = rune(0)
					}
				}
			}
		
		
		
			// FUNC to fix the NULL e.g Rune(0) bytes from the string
			for i := 0; i < len(secondRuned); i++ {
				if secondRuned[i] == rune(0) {
					secondRuned[i] = '{'
				}
			}
		
			thirdSoS := strings.Split(string(secondRuned), " ")
		
			for i := 0; i < len(thirdSoS); i++ {
				thirdSoS[i] = strings.Replace(thirdSoS[i], "{", "", -1)
			}
		
			// Getting the programme ready to print
		
			lastVar := strings.Join(thirdSoS, " ")

		
			// convertStringToByte := []byte(lastVar)
		
			// ioutil.WriteFile("result.txt", convertStringToByte, 0777)
		
			fmt.Println(lastVar)










    })

	http.ListenAndServe(":8000", nil)



	
 }