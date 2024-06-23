package instructioner

var helloInstruction = instDef{
	steps: []instructionStepFunc{
		helloInstructionStepFirstStep,
		helloInstructionStepSecondStep,
		helloInstructionStepThirdStep,
	},
}

func helloInstructionStepFirstStep(e executor) (string, error) {
	err := e.statusUpdate("Step 1 executed")
	if err != nil {
		return "", err
	}
	return resultTryAgainLater, nil
}

func helloInstructionStepSecondStep(e executor) (string, error) {
	err := e.statusUpdate("Step 2 executed")
	if err != nil {
		return "", err
	}
	return resultTryAgainLater, nil
}

func helloInstructionStepThirdStep(e executor) (string, error) {
	err := e.statusUpdate("Step 3 executed")
	if err != nil {
		return "", err
	}
	return resultComplete, nil
}
