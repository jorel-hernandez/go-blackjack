package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	//panic("Please implement the ParseCard function")
	switch card {
		case "ace": return 11
		case "two": return 2
		case "three": return 3
		case "four": return 4
		case "five": return 5
		case "six": return 6
		case "seven": return 7
		case "eight": return 8
		case "nine": return 9
		case "ten", "jack", "queen", "king": return 10
		default: return 0
	}
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	//panic("Please implement the IsBlackjack function")
	switch card1Value, card2Value := ParseCard(card1), ParseCard(card2); {
		//case card1Value == 11 && card2Value == 11: return false
		case card1Value + card2Value == 21: return true
		default: return false 
	}

}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	//panic("Please implement the LargeHand function")
	switch {
		case isBlackjack == true: 
			switch {
				case dealerScore >= 10: return "S"
				default: return "W"
			}
	}
	return "P"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	//panic("Please implement the SmallHand function")
	switch {
		case handScore >= 17: return "S"
		case handScore <= 11: return "H"
		case handScore >= 12 && handScore <= 16:
			switch {
				case dealerScore >= 7: return "H"
				default: return "S"
			}
	}
	return "0"
}
