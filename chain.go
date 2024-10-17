package chains

import (
	"fmt"
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
	name      string
	processor *chain.Processor[Context]
}

func NewChain(name string) *ChainExecutor {
	processor := chain.New[Context](name, &chain.Actions[Context]{}, printLog)
	return &ChainExecutor{
		name:      name,
		processor: processor,
	}
}

func (e *ChainExecutor) Actions(actions ...interface{}) *ChainExecutor {
	if e.processor.Actions == nil {
		e.processor.Actions = &chain.Actions[Context]{}
	}

	for _, action := range actions {
		switch a := action.(type) {
		case *chain.Action[Context]:
			e.processor.AddAction(a)
		case *chain.Actions[Context]:
			e.processor.AddActions(a)
		default:
			fmt.Println("Type unknown")
		}
	}
	return e
}

func (e ChainExecutor) Run(ctx *gin.Context) {
	e.processor.Run(
		&Context{
			C: ctx,
		},
	)
}
