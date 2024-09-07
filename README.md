# kamei

"Kamei" is a package that manages Japanese family names, and is used to obtain the corresponding family names from part of a family name, or to obtain rubys.

Currently, it contains **2627** family names.

## Notice

This is a package that was extracted from a program I created for a different project, because the part that managed family names had become too large.
The collected data includes stage names, and therefore also includes names that are not normally used as family names.


## Install

```sh
go get -u github.com/usk81/kamei
```

## Usage

### All

All gets all the family names registered in this package

```Go
fmt.Printf("%#v", kamei.All())
```

### Find

Find rubys by kanji

```Go
fmt.Printf("%#v", kamei.Find("中田"))
// []string{"なかた", "なかだ"}
```

### Contains

Contains returns family names containing the pecified string exist in the list of family names registered in this package.

```Go
fmt.Printf("%#v", kamei.Contains("佐"))
// FamilyNames{2:map[string][]string{"佐伯":[]string{"さいき", "さえき", "さはく", "さへき"}, "佐倉":[]string{"さくら", "さのくら"}, "佐原":[]string{"さはら", "さわら"}, "佐山":[]string{"さやま"}, "佐川":[]string{"さかわ", "さがわ"}, "佐田":[]string{"さた", "さだ"}, "佐藤":[]string{"さとう"}, "佐賀":[]string{"さか", "さが"}, "佐野":[]string{"さの"}}, 3:map[string][]string{"亜佐倉":[]string{"あさくら"}, "佐々山":[]string{"ささやま"}, "佐々木":[]string{"ささき"}, "佐久間":[]string{"さくま"}, "宇佐美":[]string{"うさみ"}}}
```

### HasPrefix 

// HasPrefix returns family names begin with prefix exist in the list of family names registered in the package.

```Go
fmt.Printf("%#v", kamei.HasPrefix("佐々"))
// FamilyNames{3:map[string][]string{"佐々山":[]string{"ささやま"}, "佐々木":[]string{"ささき"}}}
```

### HasSuffix

HasSuffix returns family names end with suffix exist in the list of family names registered in the package.

```Go
fmt.Printf("%#v", kamei.HasSuffix("条"))
// FamilyNames{2:map[string][]string{"一条":[]string{"いちじょう"}, "三条":[]string{"さんじょう"}, "上条":[]string{"うえじょう", "かみしょう", "かみじょう", "じょうじょう"}, "九条":[]string{"くじょう"}, "五条":[]string{"ごじょう"}, "北条":[]string{"ほじょう", "きたじょう", "ほうしょう", "ほうじょう", "ほくじゅう", "ほくじょう"}, "南条":[]string{"なんじょう"}, "斉条":[]string{"さいじょう"}, "斎条":[]string{"さいじょう"}, "新条":[]string{"しんじょう"}, "東条":[]string{"とうじょう"}, "西条":[]string{"さいじょう", "にしじょう"}, "齊条":[]string{"さいじょう"}, "齋条":[]string{"さいじょう"}}, 3:map[string][]string{"東十条":[]string{"ひがしじゅうじょう"}}}
```

### Pickup

Pickup randomly selects from all the family names registered in the package.

> This is created assuming cases where multiple fake names are generated in test data, etc.

```Go
fmt.Printf("%#v", kamei.Pickup(10))
```
