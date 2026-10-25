package session

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"

	"github.com/secfault-org/hacktober/internal/config"
	"github.com/secfault-org/hacktober/internal/container"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	cm "github.com/secfault-org/hacktober/internal/model/container"
)

type Session struct {
	mu               sync.Mutex
	active           *challenge.ActiveChallenge
	containerService container.Service
	cfg              *config.Config
}

var ErrStartCancelled = errors.New("challenge start cancelled")

func New(service container.Service, cfg *config.Config) *Session {
	return &Session{
		containerService: service,
		cfg:              cfg,
	}
}

func (s *Session) ActiveChallenge() *challenge.ActiveChallenge {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.active
}

func (s *Session) IsChallengeRunning() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.active != nil && s.active.IsRunning()
}

func (s *Session) CanStartChallenge() bool { return s.ActiveChallenge() == nil }
func (s *Session) CanStopChallenge() bool  { return s.IsChallengeRunning() }

func (s *Session) SetStarting(ch *challenge.Challenge) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.active != nil {
		return false
	}
	s.active = &challenge.ActiveChallenge{
		Challenge: ch,
		Container: &cm.Container{State: cm.Starting},
	}
	return true
}

func (s *Session) Start(ctx context.Context, ch *challenge.Challenge) (*challenge.ActiveChallenge, error) {
	flag, err := newFlag()
	if err != nil {
		s.clear()
		return nil, err
	}

	image := container.ResolveImage(s.cfg.Container.RegistryPrefix, ch.ContainerImage)
	c, err := s.containerService.StartContainer(ctx, image, flag, s.cfg.Container.ContainerPort)
	if err != nil {
		s.clear()
		return nil, err
	}

	s.mu.Lock()
	if s.active == nil {
		s.mu.Unlock()
		_ = s.containerService.StopContainer(ctx, c.ID)
		return nil, ErrStartCancelled
	}
	s.active = &challenge.ActiveChallenge{Challenge: ch, Container: c, Flag: flag}
	ac := s.active
	s.mu.Unlock()

	return ac, nil
}

func (s *Session) Stop(ctx context.Context) error {
	s.mu.Lock()
	ac := s.active
	s.active = nil
	s.mu.Unlock()

	if ac == nil || ac.Container == nil || ac.Container.ID == "" {
		return nil
	}
	return s.containerService.StopContainer(ctx, ac.Container.ID)
}

func (s *Session) clear() {
	s.mu.Lock()
	s.active = nil
	s.mu.Unlock()
}

func (s *Session) CheckFlag(flag string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.active != nil && flag == s.active.Flag
}

//func (s *Session) IsSolved(challengeID string) bool {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//	_, ok := s.solved[challengeID]
//	return ok
//}
//
//func (s *Session) MarkSolved(ctx context.Context, challengeID string) error {
//	if err := s.app.Users().RecordSolve(ctx, s.User.Fingerprint, challengeID); err != nil {
//		return err
//	}
//	s.mu.Lock()
//	s.solved[challengeID] = time.Now()
//	s.mu.Unlock()
//	return nil
//}

func newFlag() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
