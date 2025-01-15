package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11

	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	playerScore := ParseCard(card1) + ParseCard(card2)
	// Rule 1: Split Aces
	if playerScore == 22 {
		return "P"
	}

	// Rule 2: Blackjack automatic win or stand
	if playerScore == 21 {
		if ParseCard(dealerCard) == 10 || ParseCard(dealerCard) == 11 {
			return "S" // Stand and wait
		}
		return "W" // Automatically win
	}

	// Rule 3: Stand on [17, 20]
	if playerScore >= 17 && playerScore <= 20 {
		return "S"
	}

	// Rule 4: Stand or hit on [12, 16]
	if playerScore >= 12 && playerScore <= 16 {
		if ParseCard(dealerCard) >= 7 {
			return "H" // Hit
		}
		return "S" // Stand
	}

	// Rule 5: Hit on 11 or lower
	if playerScore <= 11 {
		return "H"
	}

	// Fallback (should not reach here)
	return ""
}
