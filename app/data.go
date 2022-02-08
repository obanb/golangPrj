package app

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
)

var LetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var Words2 = [20]string{"as",
	"so",
	"oh",
	"do",
	"no",
	"go",
	"if",
	"me",
	"am",
	"us",
	"by",
	"be",
	"to",
	"my",
	"we",
	"on",
	"is",
	"an",
	"of",
	"up"}
var Words3 = [20]string{"hut",
	"sip",
	"beg",
	"age",
	"cry",
	"nut",
	"few",
	"oak",
	"gas",
	"raw",
	"law",
	"hot",
	"fun",
	"rob",
	"pat",
	"fog",
	"joy",
	"hay",
	"shy",
	"flu",
}
var Words4 = [20]string{"boat",
	"sand",
	"raid",
	"riot",
	"belt",
	"beer",
	"soar",
	"vote",
	"ball",
	"cast",
	"rush",
	"risk",
	"like",
	"hurt",
	"rest",
	"axis",
	"camp",
	"pack",
	"copy",
	"foot",
}
var Words5 = [20]string{"award",
	"piece",
	"issue",
	"bland",
	"gloom",
	"greet",
	"death",
	"guide",
	"scale",
	"colon",
	"grand",
	"press",
	"debut",
	"angle",
	"noble",
	"troop",
	"essay",
	"child",
	"shark",
	"chart",
}
var Words6 = [20]string{"award",
	"piece",
	"issue",
	"bland",
	"gloom",
	"greet",
	"death",
	"guide",
	"scale",
	"colon",
	"grand",
	"press",
	"debut",
	"angle",
	"noble",
	"troop",
	"essay",
	"child",
	"shark",
	"chart",
}
var Words7 = [20]string{"deprive",
	"express",
	"distort",
	"passage",
	"stretch",
	"serious",
	"present",
	"penalty",
	"mixture",
	"contain",
	"confine",
	"fitness",
	"ecstasy",
	"perform",
	"enhance",
	"unaware",
	"climate",
	"default",
	"content",
	"outside",
}
var Words8 = [20]string{"misplace",
	"accurate",
	"affinity",
	"reliance",
	"material",
	"customer",
	"diameter",
	"remember",
	"activate",
	"appendix",
	"cylinder",
	"proclaim",
	"concrete",
	"retailer",
	"threaten",
	"physical",
	"sickness",
	"majority",
	"minimize",
	"displace",
}
var Words9 = [20]string{"favorable",
	"wisecrack",
	"situation",
	"empirical",
	"orchestra",
	"temporary",
	"treasurer",
	"potential",
	"sensation",
	"computing",
	"establish",
	"directory",
	"difficult",
	"quotation",
	"craftsman",
	"hierarchy",
	"defendant",
	"education",
	"cigarette",
	"depressed",
}
var Words10 = [20]string{"proportion",
	"earthquake",
	"functional",
	"wilderness",
	"connection",
	"tournament",
	"vegetarian",
	"presidency",
	"permission",
	"definition",
	"goalkeeper",
	"background",
	"possession",
	"gregarious",
	"engagement",
	"overcharge",
	"indication",
	"nomination",
	"foundation",
	"resolution",
}
var Words11 = [20]string{"legislature",
	"exploration",
	"performance",
	"entitlement",
	"color-blind",
	"manufacture",
	"operational",
	"disturbance",
	"unfortunate",
	"consolidate",
	"respectable",
	"mathematics",
	"replacement",
	"concentrate",
	"nonremittal",
	"integration",
	"demonstrate",
	"legislation",
	"electronics",
	"requirement",
}
var Words12 = [20]string{"organisation",
	"disagreement",
	"intermediate",
	"contemporary",
	"conservation",
	"conglomerate",
	"satisfaction",
	"refrigerator",
	"transmission",
	"acquaintance",
	"jurisdiction",
	"introduction",
	"reproduction",
	"presentation",
	"intervention",
	"interference",
	"headquarters",
	"discriminate",
	"neighborhood",
	"accumulation",
}

func getRandomWords(s int) [20]string {
	switch s {
	case 2:
		return Words2
	case 3:
		return Words3
	case 4:
		return Words4
	case 5:
		return Words5
	case 6:
		return Words6
	case 7:
		return Words7
	case 8:
		return Words8
	case 9:
		return Words10
	case 10:
		return Words10
	case 11:
		return Words11
	case 12:
		return Words12
	default:
		return Words6
	}
}

func getNumberBetween(min int, max int) int {
	num := rand.Intn((max - min) + min)
	return num
}

func getRandomWord(a [20]string) string {
	min := 0
	max := 20

	num := getNumberBetween(min, max)

	return a[num]
}

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = LetterRunes[rand.Intn(len(LetterRunes))]
	}
	s := string(b)
	return s
}

func getRandomLetter() string {
	return randomString(1)
}

func getRandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func getRandomIntBetween(s int) int {
	switch s {
	case 1:
		return getRandomInt(0, 9)
	case 2:
		return getRandomInt(0, 99)
	case 3:
		return getRandomInt(0, 999)
	case 4:
		return getRandomInt(0, 9999)
	case 5:
		return getRandomInt(0, 99999)
	case 6:
		return getRandomInt(0, 999999)
	case 7:
		return getRandomInt(0, 9999999)
	case 8:
		return getRandomInt(0, 99999999)
	case 9:
		return getRandomInt(0, 999999999)
	case 10:
		return getRandomInt(0, 9999999999)
	case 11:
		return getRandomInt(0, 99999999999)
	case 12:
		return getRandomInt(0, 999999999999)
	default:
		return getRandomInt(0, 999999)
	}
}

func getMongoId() primitive.ObjectID {
	return primitive.NewObjectID()
}
