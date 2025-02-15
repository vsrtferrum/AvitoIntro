package module

import "github.com/vsrtferrum/AvitoIntro/internal/errors"

func (module *Module) authByToken(token string) (uint64, error) {
	id, ok := module.auth.Identify(token)
	if !ok {
		module.logger.WriteError(errors.ErrNoUserFound)
		return 0, errors.ErrNoUserFound
	}
	return id, nil
}
