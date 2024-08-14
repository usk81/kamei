package kamei

import (
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

type tempPickupName struct {
	Kanji string
	Rubys []string
}

// All gets all the family names registered in this package
func All() FamilyNames {
	return familyNames
}

// Find rubys by kanji
func Find(kanji string) []string {
	l := utf8.RuneCountInString(kanji)
	all := All()
	kanjis, ok := all[l]
	if !ok {
		return nil
	}
	rubys, ok := kanjis[kanji]
	if !ok {
		return nil
	}
	return rubys
}

// Contains returns family names containing the string substr exist in the list of family names registered in this package.
func Contains(substr string) FamilyNames {
	return searchByKanji(substr, strings.Contains)
}

// HasPrefix returns family names begin with prefix exist in the list of family names registered in the package.
func HasPrefix(prefix string) FamilyNames {
	return searchByKanji(prefix, strings.HasPrefix)
}

// HasSuffix returns family names end with suffix exist in the list of family names registered in the package.
func HasSuffix(suffix string) FamilyNames {
	return searchByKanji(suffix, strings.HasSuffix)
}

func searchByKanji(kanji string, fn func(s string, search string) bool) FamilyNames {
	lk := utf8.RuneCountInString(kanji)
	all := All()
	result := FamilyNames{}
	for n, kanjis := range all {
		if n < lk {
			continue
		}
		for kj, rubys := range kanjis {
			if fn(kj, kanji) {
				if _, ok := result[n]; ok {
					result[n][kj] = rubys
				} else {
					result[n] = map[string][]string{
						kj: rubys,
					}
				}
			}
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func FlatAll() FlatFamilyNames {
	fs := All()
	result := FlatFamilyNames{}
	for _, fn := range fs {
		for kanji, rubys := range fn {
			result[kanji] = rubys
		}
	}
	return result
}

// Pickup randomly selects from all the family names registered in the package.
func Pickup(n int) FlatFamilyNames {
	if n == 0 {
		return nil
	}
	fs := All()
	var list []tempPickupName
	for _, fn := range fs {
		for kanji, rubys := range fn {
			list = append(list, tempPickupName{
				Kanji: kanji,
				Rubys: rubys,
			})
		}
	}
	l := len(list)
	if n >= l {
		return FlatAll()
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := FlatFamilyNames{}
	for i := 0; i < n; i++ {
		i := r.Intn(l)
		r := list[i]
		result[r.Kanji] = r.Rubys
	}
	return result
}
