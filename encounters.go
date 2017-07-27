package pokeapigo

type EncounterMethod struct {
	id    int
	name  string
	order int
	names []Name
}

type EncounterCondition struct {
	id     int
	name   string
	names  []Name
	values []NamedAPIResource
}

type EncounterConditionValue struct {
	id        int
	name      string
	condition []NamedAPIResource
	names     []Name
}
