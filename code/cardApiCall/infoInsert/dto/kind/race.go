package kind

type RaceKind struct {
	Kind
}

var (
	RaceSpellcaster = RaceKind{Kind{ID: 1, NameJa: "魔法使い族", NameEn: "Spellcaster"}}
	RaceDragon      = RaceKind{Kind{ID: 2, NameJa: "ドラゴン族", NameEn: "Dragon"}}
	RaceZombie      = RaceKind{Kind{ID: 3, NameJa: "アンデット族", NameEn: "Zombie"}}
	RaceWarrior     = RaceKind{Kind{ID: 4, NameJa: "戦士族", NameEn: "Warrior"}}
	RaceBeastWarrior = RaceKind{Kind{ID: 5, NameJa: "獣戦士族", NameEn: "Beast-Warrior"}}
	RaceBeast       = RaceKind{Kind{ID: 6, NameJa: "獣族", NameEn: "Beast"}}
	RaceWingedBeast = RaceKind{Kind{ID: 7, NameJa: "鳥獣族", NameEn: "Winged Beast"}}
	RaceFiend       = RaceKind{Kind{ID: 8, NameJa: "悪魔族", NameEn: "Fiend"}}
	RaceFairy       = RaceKind{Kind{ID: 9, NameJa: "天使族", NameEn: "Fairy"}}
	RaceInsect      = RaceKind{Kind{ID: 10, NameJa: "昆虫族", NameEn: "Insect"}}
	RaceDinosaur    = RaceKind{Kind{ID: 11, NameJa: "恐竜族", NameEn: "Dinosaur"}}
	RaceReptile     = RaceKind{Kind{ID: 12, NameJa: "爬虫類族", NameEn: "Reptile"}}
	RaceFish        = RaceKind{Kind{ID: 13, NameJa: "魚族", NameEn: "Fish"}}
	RaceSeaSerpent  = RaceKind{Kind{ID: 14, NameJa: "海竜族", NameEn: "Sea Serpent"}}
	RaceAqua        = RaceKind{Kind{ID: 15, NameJa: "水族", NameEn: "Aqua"}}
	RacePyro        = RaceKind{Kind{ID: 16, NameJa: "炎族", NameEn: "Pyro"}}
	RaceThunder     = RaceKind{Kind{ID: 17, NameJa: "雷族", NameEn: "Thunder"}}
	RaceRock        = RaceKind{Kind{ID: 18, NameJa: "岩石族", NameEn: "Rock"}}
	RacePlant       = RaceKind{Kind{ID: 19, NameJa: "植物族", NameEn: "Plant"}}
	RaceMachine     = RaceKind{Kind{ID: 20, NameJa: "機械族", NameEn: "Machine"}}
	RacePsychic     = RaceKind{Kind{ID: 21, NameJa: "サイキック族", NameEn: "Psychic"}}
	RaceDivineBeast = RaceKind{Kind{ID: 22, NameJa: "幻神獣族", NameEn: "Divine-Beast"}}
	RaceWyrm        = RaceKind{Kind{ID: 23, NameJa: "幻竜族", NameEn: "Wyrm"}}
	RaceCyberse     = RaceKind{Kind{ID: 24, NameJa: "サイバース族", NameEn: "Cyberse"}}
	RaceIllusion    = RaceKind{Kind{ID: 25, NameJa: "幻想魔族", NameEn: "Illusion"}}
)
