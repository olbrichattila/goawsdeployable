package instructioner

var instructionTestInstruction = instDef{
	steps: []instructionStepFunc{
		instructionTestInstructionStep1,
		instructionTestInstructionStep2,
		instructionTestInstructionStep3,
		instructionTestInstructionStep4,
		instructionTestInstructionStep5,
		instructionTestInstructionStep6,
		instructionTestInstructionStep7,
		instructionTestInstructionStep8,
		instructionTestInstructionStep9,
		instructionTestInstructionStep10,
	},
}

func instructionTestInstructionStep1(e executor) (string, error) {
	err := e.statusUpdate("instructionTestInstructionStep1 executed")
	if err != nil {
		return "", err
	}
	return resultTryAgainLater, nil
}

func instructionTestInstructionStep2(e executor) (string, error) {
	err := e.statusUpdate("instructionTestInstructionStep2 executed")
	if err != nil {
		return "", err
	}
	return resultTryAgainLater, nil
}

func instructionTestInstructionStep3(e executor) (string, error) {
	err := e.statusUpdate("instructionTestInstructionStep3 executed")
	if err != nil {
		return "", err
	}
	return resultTryAgainLater, nil
}

func instructionTestInstructionStep4(e executor) (string, error) {
	err := e.statusUpdate("instructionTestInstructionStep4 executed")
	if err != nil {
		return "", err
	}
	return resultTryAgainLater, nil
}

func instructionTestInstructionStep5(e executor) (string, error) {
	err := e.statusUpdate("instructionTestInstructionStep5 executed")
	if err != nil {
		return "", err
	}
	return resultTryAgainLater, nil
}

func instructionTestInstructionStep6(e executor) (string, error) {
	err := e.statusUpdate("instructionTestInstructionStep6 executed")
	if err != nil {
		return "", err
	}
	return resultTryAgainLater, nil
}

func instructionTestInstructionStep7(e executor) (string, error) {
	err := e.statusUpdate("instructionTestInstructionStep7 executed")
	if err != nil {
		return "", err
	}
	return resultTryAgainLater, nil
}

func instructionTestInstructionStep8(e executor) (string, error) {
	err := e.statusUpdate("instructionTestInstructionStep8 executed")
	if err != nil {
		return "", err
	}
	return resultTryAgainLater, nil
}

func instructionTestInstructionStep9(e executor) (string, error) {
	err := e.statusUpdate("instructionTestInstructionStep3 executed")
	if err != nil {
		return "", err
	}
	return resultTryAgainLater, nil
}

func instructionTestInstructionStep10(e executor) (string, error) {
	err := e.statusUpdate("instructionTestInstructionStep10 executed")
	if err != nil {
		return "", err
	}
	return resultComplete, nil
}
