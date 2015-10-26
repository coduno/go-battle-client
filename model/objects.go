package model

import "time"

type GameMap []Entity

type Entity struct {
	GameObject GameObject `json:"gameObject"`
	Type       string     `json:"type"`
}

type GameObject struct {
	Nick       string        `json:"nick"`
	HP         int           `json:"hp"`
	Deaths     int           `json:"deaths"`
	Kills      int           `json:"kills"`
	Level      int           `json:"level"`
	Pos        Position      `json:"position"`
	MoveTime   time.Time     `json:"moveTime"`
	MoveSpeed  time.Duration `json:"moveSpeed"`
	AttackTime time.Time     `json:"attackTime"`
	Spells     []Spell       `json:"attackSpeed"`
}

type Position struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type Spell struct {
	Name     string        `json:"name"`
	Cooldown time.Duration `json:"cooldown"`
	UseTime  time.Time     `json:"useTime"`
}

type BattleError struct {
	Message   string        `json:"message"`
	Behaviour string        `json:"behaviour"`
	Remaining time.Duration `json:"remaining"`
}

type TypedBattleError struct {
	BattleError BattleError `json:"battleError"`
	Type        string      `json:"type"`
}
