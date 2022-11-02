func isPalindrome(x int) bool {
    convertedString := strconv.Itoa(x)
    splittedString := strings.Split(convertedString,"")
    
    fmt.Println("check!: ", splittedString)
    
    b := len(splittedString)
    
    for a := 0; a < len(splittedString); a++ {
        b = b - 1
        if splittedString[a] != splittedString[b] {
            return false
        }
    }
    return true
}