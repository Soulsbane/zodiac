package zodiac

var (
	animalString = []string{"Rat", "Ox", "Tiger", "Rabbit", "Dragon", "Snake",
		"Horse", "Goat", "Monkey", "Rooster", "Dog", "Pig"}
	stemYYString  = []string{"Yang", "Yin"}
	elementString = []string{"Wood", "Fire", "Earth", "Metal", "Water"}
	stemCh        = []rune("甲乙丙丁戊己庚辛壬癸")
	branchCh      = []rune("子丑寅卯辰巳午未申酉戌亥")
)

type Sign struct {
	Animal     string
	YinYang    string
	Element    string
	StemBranch string
	Year       int
}

func GetSign(year int) Sign {
	var sign Sign

	year -= 4
	stem := year % 10
	branch := year % 12

	sign.Animal = animalString[branch]
	sign.YinYang = stemYYString[stem%2]
	sign.Element = elementString[stem/2]
	sign.StemBranch = string([]rune{stemCh[stem], branchCh[branch]})
	sign.Year = year%60 + 1

	return sign
}
