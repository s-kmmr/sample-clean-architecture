package injector

import (
	"github.com/s-kmmr/sample-clean-architecture/domain/repository"
	"github.com/s-kmmr/sample-clean-architecture/infrastructure/client"
	pd "github.com/s-kmmr/sample-clean-architecture/infrastructure/persistence/database"
	"github.com/s-kmmr/sample-clean-architecture/interfaces/controllers"
	"github.com/s-kmmr/sample-clean-architecture/interfaces/gateway/database"
	"github.com/s-kmmr/sample-clean-architecture/usecases"
)

type Injector struct {
	memberController controllers.MemberController
}

func NewInjector() Injector {
	i := Injector{}
	i.memberController = i.newMemberController()
	return i
}
func (i *Injector) MemberController() controllers.MemberController {
	return i.memberController
}

func (i *Injector) newMemberController() controllers.MemberController {
	m := controllers.NewMemberController(
		i.newMemberUseCase(),
	)
	return m
}

func (i *Injector) newMemberUseCase() usecases.MemberUseCase {
	u := usecases.NewMemberUseCase(
		i.newMemberRepository(),
		i.newTxRepository(),
	)
	return u
}

func (i *Injector) newMemberRepository() repository.MemberRepository {
	// frameworks
	m := pd.NewMemberHandler(i.newFwSqlHandler())
	// gateways
	r := database.NewMemberRepository(m)
	return r
}

func (i *Injector) newTxRepository() repository.TransactionRepository {
	// frameworks
	tr := pd.NewTransactionHandler(i.newFwSqlHandler())
	// gateways
	r := database.NewTransactionRepository(tr)
	return r
}

// frameworks
func (i *Injector) newFwSqlHandler() *client.SqlHandler {
	return client.NewSqlHandler()
}
