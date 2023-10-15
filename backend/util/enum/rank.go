package enum

type Rank string

const (
	Silver   Rank = "silver"
	Gold     Rank = "gold"
	Platinum Rank = "platinum"
)

func GetRankByTotalOrder(total int) Rank {
	if total > 100000 {
		return Platinum
	} else if total >= 30000 {
		return Gold
	}
	return Silver
}
