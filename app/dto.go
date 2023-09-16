package app

// Animal response for single animal
type animalResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	NumFinger     int    `json:"numFinger"`
	Skill         string `json:"skill"`
	SkillDesc     string `json:"skillDesc"`
	HandUrl       string `json:"handUrl"`
	UpFingerUrl   string `json:"upFingerUrl"`
	DownFingerUrl string `json:"downFingerUrl"`
}

// Animals response for list of animals
type animalsGetResponse struct {
	Animals []animalResponse `json:"animals"`
}

type AnimalPostRequest struct {
	ID string `json:"id"`
}
type animalSelectedResponse struct {
	IsPlayer      bool   `json:"isPlayer"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	NumFinger     int    `json:"numFinger"`
	Skill         string `json:"skill"`
	SkillDesc     string `json:"skillDesc"`
	HandUrl       string `json:"handUrl"`
	UpFingerUrl   string `json:"upFingerUrl"`
	DownFingerUrl string `json:"downFingerUrl"`
}
type AnimalPostResponse struct {
	MyAnimal  animalSelectedResponse `json:"myAnimal"`
	NpcAnimal animalSelectedResponse `json:"npcAnimal"`
}

type actionRequest struct {
	IsMyTurn          bool `json:"isMyTurn"`
	UseSkill          bool `json:"useSkill"`
	NumMyUpFinger     int  `json:"numMyUpFinger"`
	NumExpectedFinger *int `json:"numExpectedFinger"` // null許容
}

type actionResponse struct {
	IsFinished        bool `json:"isFinished"`
	IsMyTurn          bool `json:"isMyTurn"`
	NumMyFinger       int  `json:"numMyFinger"`
	NumNpcFinger      int  `json:"numNpcFinger"`
	NumMyUpFinger     int  `json:"numMyUpFinger"`
	NumNpcUpFinger    int  `json:"numNpcUpFinger"`
	NumExpectedFinger *int `json:"numExpectedFinger"`
}
