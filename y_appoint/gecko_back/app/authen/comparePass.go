package authen

import (
	"gecko/crossLogging"
	"gecko/models/userrepo"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func ComparePass(storedUser userrepo.User, pass string) (*userrepo.User, error) {
	crossLogging.Logger.Info("login request", zap.String("username", storedUser.Email))

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.PassHash), []byte(pass))
	if err != nil {
		crossLogging.Logger.Error("error while comparing passwords", zap.Error(err))
		return nil, err
	}

	crossLogging.Logger.Info("Authenticated")
	return &storedUser, nil
}
