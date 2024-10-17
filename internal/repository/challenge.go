package repository

import (
	"context"
	"encoding/json"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

type ChallengeRepository interface {
	GetAllChallenges(ctx context.Context) ([]challenge.Challenge, error)
}

type challengeRepo struct {
	basedir string
}

type challengeInfo struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	ShortDesc      string `json:"shortDescription"`
	DescFile       string `json:"descriptionFile"`
	Difficulty     uint8  `json:"difficulty"`
	ReleaseDate    string `json:"releaseDate"`
	ContainerImage string `json:"containerImage"`
	ChallengeFile  string `json:"challengeFile"`
}

var _ ChallengeRepository = (*challengeRepo)(nil)

func (f *challengeRepo) GetAllChallenges(ctx context.Context) ([]challenge.Challenge, error) {
	entries, err := os.ReadDir(f.basedir)
	if err != nil {
		return nil, err
	}
	var challenges []challenge.Challenge
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

func readChallenge(basedir string, subDir string) (challenge.Challenge, error) {
	challengeDir := path.Join(basedir, subDir)
	jsonFile, err := os.ReadFile(path.Join(challengeDir, "challenge.json"))
	if err != nil {
		return challenge.Challenge{}, err
	}

	var challengeJson challengeInfo
	err = json.Unmarshal(jsonFile, &challengeJson)

	releaseDate, err := time.Parse(time.DateOnly, challengeJson.ReleaseDate)
	if err != nil {
		return challenge.Challenge{}, err
	}

	text, err := os.ReadFile(path.Join(challengeDir, challengeJson.DescFile))
	if err != nil {
		return challenge.Challenge{}, err
	}

	challengeFile, err := filepath.Abs(path.Join(challengeDir, challengeJson.ChallengeFile))
	if err != nil {
		return challenge.Challenge{}, err
	}

	return challenge.Challenge{
		Id:                challengeJson.Id,
		Name:              challengeJson.Title,
		Description:       challengeJson.ShortDesc,
		ChallengeMarkdown: string(text),
		ReleaseDate:       releaseDate,
		ContainerImage:    challengeJson.ContainerImage,
		ChallengeFile:     challengeFile,
	}, nil
}
