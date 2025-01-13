package presentation

import (
	"context"
	"github.com/danielchiovitti/ballot-box/pkg/domain/service"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redis"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redisbloom"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/route"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"strconv"
	"sync"
)

var handlerLock sync.Mutex
var handlerInstance *Handler

type Handler struct {
	healthRoute                    *route.HealthRoute
	votingRoute                    *route.VotingRoute
	viper                          *viper.Viper
	config                         shared.ConfigInterface
	reserveUseCaseFactory          redisbloom.ReserveUseCaseFactoryInterface
	createSteamGroupUseCaseFactory redis.CreateStreamGroupUseCaseFactoryInterface
	consumeOltpService             service.ConsumeOltpServiceInterface
	consumeOlapService             service.ConsumeOlapServiceInterface
}

func NewHandler(
	healthRoute *route.HealthRoute,
	votingRoute *route.VotingRoute,
	viper *viper.Viper,
	config shared.ConfigInterface,
	reserveUseCaseFactory redisbloom.ReserveUseCaseFactoryInterface,
	createSteamGroupUseCaseFactory redis.CreateStreamGroupUseCaseFactoryInterface,
	consumeOltpService service.ConsumeOltpServiceInterface,
	consumeOlapService service.ConsumeOlapServiceInterface,
) *Handler {
	if handlerInstance == nil {
		handlerLock.Lock()
		defer handlerLock.Unlock()
		if handlerInstance == nil {
			handlerInstance = &Handler{
				healthRoute:                    healthRoute,
				votingRoute:                    votingRoute,
				viper:                          viper,
				config:                         config,
				reserveUseCaseFactory:          reserveUseCaseFactory,
				createSteamGroupUseCaseFactory: createSteamGroupUseCaseFactory,
				consumeOltpService:             consumeOltpService,
				consumeOlapService:             consumeOlapService,
			}
		}
	}

	return handlerInstance
}

func (h *Handler) GetRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/health", h.healthRoute.HealthRoutes())
	r.Mount("/voting", h.votingRoute.VotingRoutes())
	return r
}

func (h *Handler) GetViper() *viper.Viper {
	return h.viper
}

func (h *Handler) SetBloomFilter() {
	reserveUseCase := h.reserveUseCaseFactory.Build()
	err := reserveUseCase.Execute(h.config.GetBloomName(), h.config.GetBloomPrecision(), uint64(h.config.GetBloomInitial()))
	if err != nil && err.Error() != "ERR item exists" {
		panic(err)
	}
}

func (h *Handler) CreateStreamGroup() {
	createStreamGroupUseCase := h.createSteamGroupUseCaseFactory.Build()
	ctx := context.Background()
	err := createStreamGroupUseCase.Execute(ctx, h.config.GetOltpStreamName(), h.config.GetOltpStreamGroupName())
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		panic(err)
	}

	err = createStreamGroupUseCase.Execute(ctx, h.config.GetOlapStreamName(), h.config.GetOlapStreamGroupName())
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		panic(err)
	}
}

func (h *Handler) StartConsumers() {

	for i := 0; i < h.config.GetOltpConsumersQty(); i++ {
		go h.consumeOltpService.Run(strconv.Itoa(i))
	}

	for i := 0; i < h.config.GetOlapConsumersQty(); i++ {
		go h.consumeOlapService.Run(strconv.Itoa(i))
	}
}
