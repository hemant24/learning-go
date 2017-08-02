package main

import "time"

type requestBody struct {
	ID string `json:"id"`
	Skills []string `json:"skills"`
}

type skillScore struct {
	SkillName string `json:"skill_name"`
	Score float32 `json:"score"`
	LastScored time.Time `json:"last_scored"`
}

type mockDataStore struct {
	data map[string] map[string]skillScore
}

func (mds *mockDataStore) WriteData(profileID string, score skillScore) {
	if mds.data == nil {
		mds.data = make(map[string]map[string]skillScore)
	}

	var scores map[string]skillScore
	scores, ok := mds.data[profileID]

	if !ok {
		scores = make(map[string]skillScore)
	}

	scores[score.SkillName] = score

	mds.data[profileID] = scores
}

func (mds *mockDataStore) ReadData(profileID string, skillName string) (skillScore, bool) {
	profileSkills, ok := mds.data[profileID]
	if !ok {
		return skillScore{}, false
	}

	skill, ok := profileSkills[skillName]

	if !ok {
		return skillScore{}, false
	}

	return skill, true
}

type skillScoreMessage struct {
	ProfileID string `json:"profile_id"`
	SkillName string `json:"skill_name"`
}
