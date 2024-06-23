package instructioner

const (
	resultTryAgainLater = "tryAgainLater"
	resultComplete      = "complete"
	resultFailed        = "failed"
)

type instructionStepFunc = func(executor) (string, error)

type instDef struct {
	steps []instructionStepFunc
}

var instructionList = map[string]instDef{
	"Hello":                      helloInstruction,
	"instructionTestInstruction": instructionTestInstruction,
}
