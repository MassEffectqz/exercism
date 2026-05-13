package blackjack

func ParseCard(card string) int {
	switch card{
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
func FirstTurn(card1, card2, dealerCard string) string {
	cardValue1 := ParseCard(card1)
	cardValue2 := ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	sum := cardValue1 + cardValue2
	if card1 == "ace" && card2 == "ace" { 
		return "P"
	}
	if sum == 21 { 
		if dealerValue == 11 || dealerValue == 10 {
			return "S"
		}
		return "W"
	}
	if sum >= 17 && sum <= 20 {
		return "S"
	}
	if sum >= 12 && sum <= 16 { 
		if dealerValue >= 7 { 
			return "H"
		}
		return "S"
	}
	return "H"
}
