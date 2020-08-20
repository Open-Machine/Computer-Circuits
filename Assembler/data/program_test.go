package data

import "testing"

func TestAddCommand(t *testing.T) {
	program := NewProgram(5)

	if len(program.commands) != 0 {
		t.Errorf("Expected length 0, got: %d", len(program.commands))
	}

	program.AddCommand(Command{0, NewIntParam(0)})

	if len(program.commands) != 1 {
		t.Errorf("Expected length 1, got: %d", len(program.commands))
	}
}

func TestToExecuterSuccess(t *testing.T) {
	program := NewProgram(3)
	program.AddCommand(Command{1, NewIntParam(2)})
	program.AddCommand(Command{15, NewIntParam(7)})
	program.AddCommand(Command{0, NewIntParam(0)})

	got, errors := program.ToExecuter()
	expected := "01020f070000"

	if !(len(errors) == 0 && got == expected) {
		t.Errorf("Expected: '%s', got: '%s'", expected, got)
	}
}

func TestToExecuterFail(t *testing.T) {
	program := NewProgram(3)
	program.AddCommand(Command{1, NewIntParam(2)})
	program.AddCommand(Command{1200, NewIntParam(7)})
	program.AddCommand(Command{0, NewIntParam(0)})

	execCode, errors := program.ToExecuter()

	if len(errors) != 1 {
		t.Errorf("Should result in error because of overflow. Executer code: %s // Errors: %v", execCode, errors)
	}
}