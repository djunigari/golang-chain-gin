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
	chain *chain.Processor[Context]
}

func NewChain(actions ...*chain.Action[Context]) *ChainExecutor {
	return &ChainExecutor{chain: chain.New((*chain.Actions[Context])(&actions), printLog)}
}

func (e ChainExecutor) Run(ctx *gin.Context) {
	e.chain.Run(
		&Context{
			C: ctx,
		},
	)
}
