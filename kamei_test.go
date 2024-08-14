package kamei

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		kanji string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "kanji_1_charactor",
			args: args{
				kanji: "叶",
			},
			want: []string{
				"かなえ",
				"かのう",
			},
		},
		{
			name: "kanji_2_charactor",
			args: args{
				kanji: "中田",
			},
			want: []string{
				"なかた",
				"なかだ",
			},
		},
		{
			name: "not_kanji",
			args: args{
				kanji: "さとう",
			},
			want: nil,
		},
		{
			name: "empty",
			args: args{
				kanji: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.kanji); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		substr string
	}
	tests := []struct {
		name string
		args args
		want FamilyNames
	}{
		{
			name: "kanji_1_charactor",
			args: args{
				substr: "佐",
			},
			want: FamilyNames{
				2: {
					"佐伯": {
						"さいき",
						"さえき",
						"さはく",
						"さへき",
					},
					"佐倉": {
						"さくら",
						"さのくら",
					},
					"佐原": {
						"さはら",
						"さわら",
					},
					"佐山": {
						"さやま",
					},
					"佐川": {
						"さかわ",
						"さがわ",
					},
					"佐田": {
						"さた",
						"さだ",
					},
					"佐藤": {
						"さとう",
					},
					"佐賀": {
						"さか",
						"さが",
					},
					"佐野": {
						"さの",
					},
				},
				3: {
					"亜佐倉": {
						"あさくら",
					},
					"佐々山": {
						"ささやま",
					},
					"佐々木": {
						"ささき",
					},
					"佐久間": {
						"さくま",
					},
					"宇佐美": {
						"うさみ",
					},
				},
			},
		},
		{
			name: "kanji_2_charactor",
			args: args{
				substr: "佐々",
			},
			want: FamilyNames{
				3: {
					"佐々山": {
						"ささやま",
					},
					"佐々木": {
						"ささき",
					},
				},
			},
		},
		{
			name: "unmatch",
			args: args{
				substr: "家名",
			},
			want: nil,
		},
		{
			name: "not_kanji",
			args: args{
				substr: "ささ",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.substr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPrefix(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		args args
		want FamilyNames
	}{
		{
			name: "kanji_1_charactor",
			args: args{
				prefix: "佐",
			},
			want: FamilyNames{
				2: {
					"佐伯": {
						"さいき",
						"さえき",
						"さはく",
						"さへき",
					},
					"佐倉": {
						"さくら",
						"さのくら",
					},
					"佐原": {
						"さはら",
						"さわら",
					},
					"佐山": {
						"さやま",
					},
					"佐川": {
						"さかわ",
						"さがわ",
					},
					"佐田": {
						"さた",
						"さだ",
					},
					"佐藤": {
						"さとう",
					},
					"佐賀": {
						"さか",
						"さが",
					},
					"佐野": {
						"さの",
					},
				},
				3: {
					"佐々山": {
						"ささやま",
					},
					"佐々木": {
						"ささき",
					},
					"佐久間": {
						"さくま",
					},
				},
			},
		},
		{
			name: "kanji_2_charactor",
			args: args{
				prefix: "佐々",
			},
			want: FamilyNames{
				3: {
					"佐々山": {
						"ささやま",
					},
					"佐々木": {
						"ささき",
					},
				},
			},
		},
		{
			name: "unmatch",
			args: args{
				prefix: "家名",
			},
			want: nil,
		},
		{
			name: "not_kanji",
			args: args{
				prefix: "ささ",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPrefix(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasSuffix(t *testing.T) {
	type args struct {
		suffix string
	}
	tests := []struct {
		name string
		args args
		want FamilyNames
	}{
		{
			name: "kanji_1_charactor",
			args: args{
				suffix: "条",
			},
			want: FamilyNames{
				2: {
					"一条": {
						"いちじょう",
					},
					"三条": {
						"さんじょう",
					},
					"上条": {
						"うえじょう",
						"かみしょう",
						"かみじょう",
						"じょうじょう",
					},
					"九条": {
						"くじょう",
					},
					"五条": {
						"ごじょう",
					},
					"北条": {
						"ほじょう",
						"きたじょう",
						"ほうしょう",
						"ほうじょう",
						"ほくじゅう",
						"ほくじょう",
					},
					"南条": {
						"なんじょう",
					},
					"斉条": {
						"さいじょう",
					},
					"斎条": {
						"さいじょう",
					},
					"新条": {
						"しんじょう",
					},
					"東条": {
						"とうじょう",
					},
					"西条": {
						"さいじょう",
						"にしじょう",
					},
					"齊条": {
						"さいじょう",
					},
					"齋条": {
						"さいじょう",
					},
				},
				3: {
					"東十条": {
						"ひがしじゅうじょう",
					},
				},
			},
		},
		{
			name: "kanji_2_charactor",
			args: args{
				suffix: "五条",
			},
			want: FamilyNames{
				2: {
					"五条": {
						"ごじょう",
					},
				},
			},
		},
		{
			name: "unmatch",
			args: args{
				suffix: "家名",
			},
			want: nil,
		},
		{
			name: "not_kanji",
			args: args{
				suffix: "じょう",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasSuffix(tt.args.suffix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPickup(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
	}{
		{
			name: "0",
			args: args{
				n: 0,
			},
			wantLength: 0,
		},
		{
			name: "10",
			args: args{
				n: 10,
			},
			wantLength: 10,
		},
		{
			name: "over",
			args: args{
				n: 1000000000,
			},
			wantLength: len(FlatAll()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pickup(tt.args.n); len(got) != tt.wantLength {
				t.Errorf("Pickup() length = %d, want %d", len(got), tt.wantLength)
			}
		})
	}
}
