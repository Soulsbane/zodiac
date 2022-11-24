package zodiac

var (
	animals = []string{"Rat", "Ox", "Tiger", "Rabbit", "Dragon", "Snake", "Horse",
		"Goat", "Monkey", "Rooster", "Dog", "Pig"}
	yinYang  = []string{"Yang", "Yin"}
	elements = []string{"Wood", "Fire", "Earth", "Metal", "Water"}
	stems    = []rune("甲乙丙丁戊己庚辛壬癸")
	branches = []rune("子丑寅卯辰巳午未申酉戌亥")
)

type Sign struct {
	Animal  string
	YinYang string
	Element string
	Stem    string
	Branch  string
	Year    int
}

func GetSign(year int) Sign {
	var sign Sign

	year -= 4
	stem := year % 10
	branch := year % 12

	sign.Animal = animals[branch]
	sign.YinYang = yinYang[stem%2]
	sign.Element = elements[stem/2]
	sign.Stem = string(stems[stem])
	sign.Branch = string(branches[branch])
	sign.Year = year%60 + 1

	return sign
}
