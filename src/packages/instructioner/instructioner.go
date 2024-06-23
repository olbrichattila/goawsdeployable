// Package instruction managing instructions
package instructioner

import (
	"db"
	"encoding/json"
	"sharedconfig"
)

func newInstructioner() instructioner {
	return &instruction{}
}

type createParams map[string]any

type instructioner interface {
	create(sharedconfig.SharedConfiger, string, createParams) (int64, error)
	process(sharedconfig.SharedConfiger, int64) error
}

type instruction struct {
}

func (t *instruction) create(config sharedconfig.SharedConfiger, name string, params createParams) (int64, error) {
	conn, err := db.Db(config)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	insertSQL := "INSERT INTO instructions (`name`, `current_step`, `current_step_result`, `data`) values (?,?,?,?)"

	data := "{}"
	if params != nil {
		r, err := json.Marshal(params)
		if err != nil {
			return 0, err
		}
		data = string(r)

	}
	stmt, err := conn.Prepare(insertSQL)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, 1, resultTryAgainLater, data)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (t *instruction) process(config sharedconfig.SharedConfiger, id int64) error {
	conn, err := db.Db(config)
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := "select `id`, `name`, `current_step`, `current_step_result`, `data` from instructions i WHERE i.id = ?"
	stmt, err := conn.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var i inst
	row := stmt.QueryRow(id)
	err = row.Scan(&i.Id, &i.Name, &i.CurrentStep, &i.CurrentStepResult, &i.Data)
	if err != nil {
		return err
	}

	executer := NewInstructionExecutor(config, i)
	executer.processNextStep()

	return nil
}
