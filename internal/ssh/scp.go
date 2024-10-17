package ssh

import (
	"github.com/charmbracelet/ssh"
	"github.com/pkg/sftp"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	"io"
	"io/fs"
	"os"
	"strings"
)

type challengeHandler struct {
	challenges []challenge.Challenge
}

type listerAt []fs.FileInfo

type SCPHandler interface {
	sftp.FileLister
	sftp.FileReader
}

var (
	_ sftp.FileLister = &challengeHandler{}
	_ sftp.FileReader = &challengeHandler{}
)

func NewScpChallengeHandler(challenges []challenge.Challenge) SCPHandler {
	return &challengeHandler{challenges: challenges}
}

func (s *SSHServer) sftpSubsystem(handler SCPHandler) ssh.SubsystemHandler {
	return func(session ssh.Session) {
		srv := sftp.NewRequestServer(session, sftp.Handlers{
			FileList: handler,
			FileGet:  handler,
		})
		if err := srv.Serve(); err == io.EOF {
			if err := srv.Close(); err != nil {
				s.logger.Error("sftpSubsystem", "error", err)
			}
		} else if err != nil {
			s.logger.Error("sftpSubsystem", "error", err)
		}
	}
}

func (l listerAt) ListAt(ls []fs.FileInfo, offset int64) (int, error) {
	if offset >= int64(len(l)) {
		return 0, io.EOF
	}
	n := copy(ls, l[offset:])
	if n < len(ls) {
		return n, io.EOF
	}
	return n, nil
}

func (c challengeHandler) findChallengeByPath(path string) *challenge.Challenge {
	requestedId := strings.TrimPrefix(path, "/")

	for _, chall := range c.challenges {
		if chall.Id == requestedId && !chall.Locked() {
			return &chall
		}
	}
	return nil
}

func (c challengeHandler) Filelist(request *sftp.Request) (sftp.ListerAt, error) {
	chall := c.findChallengeByPath(request.Filepath)

	if chall == nil {
		return nil, sftp.ErrSshFxNoSuchFile
	}

	fi, err := os.Stat(chall.ChallengeFile)
	if err != nil {
		return nil, err
	}

	return listerAt([]fs.FileInfo{fi}), nil
}

func (c challengeHandler) Fileread(r *sftp.Request) (io.ReaderAt, error) {
	chall := c.findChallengeByPath(r.Filepath)

	if chall == nil {
		return nil, sftp.ErrSshFxNoSuchFile
	}

	var flags int
	pflags := r.Pflags()
	if pflags.Append {
		flags |= os.O_APPEND
	}
	if pflags.Creat {
		flags |= os.O_CREATE
	}
	if pflags.Excl {
		flags |= os.O_EXCL
	}
	if pflags.Trunc {
		flags |= os.O_TRUNC
	}

	if pflags.Read && pflags.Write {
		flags |= os.O_RDWR
	} else if pflags.Read {
		flags |= os.O_RDONLY
	} else if pflags.Write {
		flags |= os.O_WRONLY
	}

	f, err := os.OpenFile(chall.ChallengeFile, flags, 0600)
	if err != nil {
		return nil, err
	}

	return f, nil
}
