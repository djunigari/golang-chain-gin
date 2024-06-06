package chains

import (
	"os"
	"strconv"

	chain "github.com/djunigari/golang-chain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var printLog bool

func init() {
	value, err := strconv.ParseBool(os.Getenv("CHAIN_LOGGER"))
	if err != nil {
		printLog = false
		return
	}
	printLog = value
}

type Context struct {
	C  *gin.Context
	Tx *gorm.DB
}

type ChainExecutor struct {
	name    string
	actions *chain.Actions[Context]

	// chain *chain.Processor[Context]
}

func NewChain(name string) *ChainExecutor {
	return &ChainExecutor{
		name:    name,
		actions: nil,
	}
}

func (e *ChainExecutor) Actions(actions ...*chain.Action[Context]) *ChainExecutor {
	e.actions = (*chain.Actions[Context])(&actions)
	return e
}

func (e ChainExecutor) Run(ctx *gin.Context) {
	chain.New(e.name, e.actions, printLog).Run(
		&Context{
			C: ctx,
		},
	)
}
