package auth

import "context"

type Respository interface{
	CreateAuth(ctx context.Context, model AuthEntity) (err error)
}

type service struct {
	repo Respository
}

func newService(repo Respository) service {
	return service{
		repo: repo,
	}
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {
	authEntity := NewFromRegisterRequest(req)
	if err = authEntity.Validate(); err != nil {
		return
	}

	return s.repo.CreateAuth(ctx, authEntity)
}