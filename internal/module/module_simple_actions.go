package module

import (
	"github.com/vsrtferrum/AvitoIntro/internal/errors"
	"github.com/vsrtferrum/AvitoIntro/internal/model"
)

func (module *Module) authByToken(token string) (model.AuthAns, error) {
	val, ok := module.auth.Identify(token)
	if !ok {
		module.logger.WriteError(errors.ErrNoUserFound)
		return model.AuthAns{}, errors.ErrNoUserFound
	}
	return val, nil
}
