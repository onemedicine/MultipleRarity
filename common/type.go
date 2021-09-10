package common

import "math/big"

type Attributes struct {
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
}

type Summoner struct {
	Xp    *big.Int
	Log   *big.Int
	Class *big.Int
	Level *big.Int
}
