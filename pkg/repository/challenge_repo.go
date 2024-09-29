package repository

import (
	"context"
	"encoding/json"
	"github.com/secfault-org/hacktober/pkg/model"
	"log"
	"os"
	"path"
	"time"
)

type challengeRepo struct {
	basedir string
}

type challengeInfo struct {
	Title          string `json:"title"`
	ShortDesc      string `json:"shortDescription"`
	DescFile       string `json:"descriptionFile"`
	Difficulty     uint8  `json:"difficulty"`
	ReleaseDate    string `json:"releaseDate"`
	ContainerImage string `json:"containerImage"`
}

var _ ChallengeRepository = (*challengeRepo)(nil)

func (f *challengeRepo) GetAllChallenges(ctx context.Context) ([]model.Challenge, error) {
	entries, err := os.ReadDir(f.basedir)
	if err != nil {
		return nil, err
	}
	var challenges []model.Challenge
	for _, entry := range entries {
		if entry.IsDir() {
			if entry.IsDir() {
				challenge, err := readChallenge(f.basedir, entry.Name())
				if err != nil {
					log.Fatal(err)
				}
				challenges = append(challenges, challenge)
			}
		}
	}
	return challenges, nil
}

func readChallenge(basedir string, challengeDir string) (model.Challenge, error) {

	jsonFile, err := os.ReadFile(path.Join(basedir, challengeDir, "challenge.json"))
	if err != nil {
		return model.Challenge{}, err
	}

	var challengeJson challengeInfo
	err = json.Unmarshal(jsonFile, &challengeJson)

	releaseDate, err := time.Parse(time.DateOnly, challengeJson.ReleaseDate)
	if err != nil {
		return model.Challenge{}, err
	}

	text, err := os.ReadFile(path.Join(basedir, challengeDir, challengeJson.DescFile))

	if err != nil {
		return model.Challenge{}, err
	}

	return model.Challenge{
		Name:              challengeJson.Title,
		Description:       challengeJson.ShortDesc,
		ChallengeMarkdown: string(text),
		ReleaseDate:       releaseDate,
		ContainerImage:    challengeJson.ContainerImage,
	}, nil
}
