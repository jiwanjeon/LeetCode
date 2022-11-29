func romanToInt(s string) int {
	sum := 0
	before := ""
	for _, c := range s {

		switch string(c) {

		case "I":
			sum += 1
			before = "I"
		case "V":
			sum += 5
			if before == "I" {
				sum -= 1 * 2
			}
			before = "V"
		case "X":
			sum += 10
			if before == "I" {
				sum -= 1 * 2
			}
			before = "X"
		case "L":
			sum += 50
			if before == "X" {
				sum -= 10 * 2
			}
			before = "L"
		case "C":
			sum += 100
			if before == "X" {
				sum -= 10 * 2
			}
			before = "C"
		case "D":
			sum += 500
			if before == "C" {
				sum -= 100 * 2
			}
			before = "D"
		case "M":
			sum += 1000
			if before == "C" {
				sum -= 100 * 2
			}
			before = "M"

		}

	}
	return sum  
}