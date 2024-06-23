package instructioner

import (
	"db"
	"fmt"
	"sharedconfig"
	"sqseventdispatcher"
)

type executor interface {
	processNextStep() error
	getInstruction() inst
	statusUpdate(string) error
}

type inst struct {
	Id                int64
	Name              string
	CurrentStep       int
	CurrentStepResult string
	Data              string
}

type execute struct {
	instruction inst
	config      sharedconfig.SharedConfiger
}

func NewInstructionExecutor(config sharedconfig.SharedConfiger, i inst) executor {
	return &execute{
		instruction: i,
		config:      config,
	}
}

func (t *execute) processNextStep() error {
	instructionToProcess, ok := instructionList[t.instruction.Name]
	if !ok {
		return fmt.Errorf("instruction %s not defined", t.instruction.Name)
	}

	if t.instruction.CurrentStep > len(instructionToProcess.steps) {
		return fmt.Errorf("cannot process next step, instruction should be complete id: %d", t.instruction.Id)
	}

	stepResult, err := instructionToProcess.steps[t.instruction.CurrentStep-1](t)
	if err != nil {
		err2 := t.updateCurrentStepResult(resultFailed)
		if err2 != nil {
			return err2
		}
		fmt.Println(err)
		return err
	}

	if stepResult == resultComplete {
		err = t.updateCurrentStepResult(resultComplete)
		if err != nil {
			return err
		}

		return nil
	}

	err = t.moveToNextStep(stepResult)
	if err != nil {
		return err
	}

	event := &sQSMessage{
		EventName: "processInstruction",
		Id:        t.instruction.Id,
	}

	dispatcher := sqseventdispatcher.NewDispatcher(t.config.GetSQSConfig())
	err = dispatcher.Send(*event)
	if err != nil {
		return err
	}

	return nil
}

func (t *execute) moveToNextStep(newStepResult string) error {
	conn, err := db.Db(t.config)
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := "UPDATE instructions set current_step = current_step + 1, current_step_result = ? where id = ?"
	stmt, err := conn.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newStepResult, t.instruction.Id)
	if err != nil {
		return err
	}

	return nil
}

func (t *execute) updateCurrentStepResult(newStepResult string) error {
	conn, err := db.Db(t.config)
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := "UPDATE instructions set current_step_result = ? where id = ?"
	stmt, err := conn.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newStepResult, t.instruction.Id)
	if err != nil {
		return err
	}

	return nil
}

func (t *execute) getInstruction() inst {
	return t.instruction
}

func (t *execute) statusUpdate(message string) error {
	conn, err := db.Db(t.config)
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := "INSERT INTO status_updates (`instruction_id`, `message`) values(?,?)"
	stmt, err := conn.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.instruction.Id, message)
	if err != nil {
		return err
	}

	return nil
}
