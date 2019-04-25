package rsw

import (
	"bytes"
	"fmt"
	"testing"
)

func createSolveTest2048A2(time uint64, t *testing.T) {
	key := make([]byte, 32)
	copy(key[:], []byte(fmt.Sprintf("opencxcreatesolve%d", time)))
	rswTimelock, err := New2048A2(key)
	if err != nil {
		t.Fatalf("There was an error creating a new timelock puzzle: %s", err)
	}
	puzzle, expectedAns, err := rswTimelock.SetupTimelockPuzzle(time)
	if err != nil {
		t.Fatalf("There was an error setting up the timelock puzzle: %s\n", err)
	}
	puzzleAns, err := puzzle.Solve()
	if err != nil {
		t.Fatalf("Error solving puzzle: %s\n", err)
	}
	if !bytes.Equal(puzzleAns, expectedAns) {
		t.Fatalf("Answer did not equal puzzle for time = %d. Expected %x, got %x\n", time, expectedAns, puzzleAns)
	}
	return
}

func createTest2048A2(time uint64, t *testing.T) {
	key := make([]byte, 32)
	copy(key[:], []byte(fmt.Sprintf("opencx%d", time)))
	rswTimelock, err := New2048A2(key)
	if err != nil {
		t.Fatalf("There was an error creating a new timelock puzzle: %s", err)
	}
	if _, _, err = rswTimelock.SetupTimelockPuzzle(time); err != nil {
		t.Fatalf("There was an error setting up the timelock puzzle: %s\n", err)
	}
	return
}

// TestCreate tests are to show that it's not the creation/setup steps that we're waiting for, it's the solve step.
// Solving whatever is created by TestCreateQuintrillion2048A2 is going to take a long time, much past our 10 minute
// testing limit
func TestCreateZero2048A2(t *testing.T) {
	createTest2048A2(0, t)
	return
}

func TestCreateOne2048A2(t *testing.T) {
	createTest2048A2(1, t)
	return
}

func TestCreateTen2048A2(t *testing.T) {
	createTest2048A2(10, t)
	return
}

func TestCreateHundred2048A2(t *testing.T) {
	createTest2048A2(100, t)
	return
}

func TestCreateThousand2048A2(t *testing.T) {
	createTest2048A2(1000, t)
	return
}

func TestCreateTenThousand2048A2(t *testing.T) {
	createTest2048A2(10000, t)
	return
}

func TestCreateHundredThousand2048A2(t *testing.T) {
	createTest2048A2(100000, t)
	return
}

func TestCreateMillion2048A2(t *testing.T) {
	createTest2048A2(1000000, t)
	return
}

func TestCreateTenMillion2048A2(t *testing.T) {
	createTest2048A2(10000000, t)
	return
}

func TestCreateHundredMillion2048A2(t *testing.T) {
	createTest2048A2(100000000, t)
	return
}

func TestCreateBillion2048A2(t *testing.T) {
	createTest2048A2(1000000000, t)
	return
}

func TestCreateTenBillion2048A2(t *testing.T) {
	createTest2048A2(10000000000, t)
	return
}

func TestCreateHundredBillion2048A2(t *testing.T) {
	createTest2048A2(100000000000, t)
	return
}

func TestCreateTrillion2048A2(t *testing.T) {
	createTest2048A2(1000000000000, t)
	return
}

func TestCreateTenTrillion2048A2(t *testing.T) {
	createTest2048A2(10000000000000, t)
	return
}

func TestCreateHundredTrillion2048A2(t *testing.T) {
	createTest2048A2(100000000000000, t)
	return
}

func TestCreateQuadrillion2048A2(t *testing.T) {
	createTest2048A2(1000000000000000, t)
	return
}

func TestCreateTenQuadrillion2048A2(t *testing.T) {
	createTest2048A2(10000000000000000, t)
	return
}

func TestCreateHundredQuadrillion2048A2(t *testing.T) {
	createTest2048A2(100000000000000000, t)
	return
}

func TestCreateQuintrillion2048A2(t *testing.T) {
	createTest2048A2(1000000000000000000, t)
	return
}

func TestCreateLCSTime2048A2(t *testing.T) {
	createTest2048A2(79685186856218, t)
	return
}

func TestZero2048A2(t *testing.T) {
	createSolveTest2048A2(0, t)
	return
}

func TestOne2048A2(t *testing.T) {
	createSolveTest2048A2(1, t)
	return
}

func TestTen2048A2(t *testing.T) {
	createSolveTest2048A2(10, t)
	return
}

func TestHundred2048A2(t *testing.T) {
	createSolveTest2048A2(100, t)
	return
}

func TestThousand2048A2(t *testing.T) {
	createSolveTest2048A2(1000, t)
	return
}

func TestTenThousand2048A2(t *testing.T) {
	createSolveTest2048A2(10000, t)
	return
}

func TestHundredThousand2048A2(t *testing.T) {
	createSolveTest2048A2(100000, t)
	return
}

func TestMillion2048A2(t *testing.T) {
	createSolveTest2048A2(1000000, t)
	return
}

func TestTenMillion2048A2(t *testing.T) {
	createSolveTest2048A2(10000000, t)
	return
}

func TestHundredMillion2048A2(t *testing.T) {
	createSolveTest2048A2(100000000, t)
	return
}