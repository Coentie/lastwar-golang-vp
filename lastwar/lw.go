package lastwar

import (
	"fmt"
	"lastwar/notifier/hdadb"
	"log"
	"os"
	"strconv"
)

var (
	amountOfTaps    int
	amountOfTapsSet bool
)

func ApproveMemberSequence() error {
	SleepBetweenSequenceAction()
	if err := List(); err != nil {
		return fmt.Errorf("List failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := hdadb.Swipe(); err != nil {
		return fmt.Errorf("hdadb.Swipe failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := hdadb.Swipe(); err != nil {
		return fmt.Errorf("hdadb.Swipe failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := hdadb.Swipe(); err != nil {
		return fmt.Errorf("hdadb.Swipe failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := hdadb.Swipe(); err != nil {
		return fmt.Errorf("hdadb.Swipe failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := hdadb.Swipe(); err != nil {
		return fmt.Errorf("hdadb.Swipe failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := hdadb.Swipe(); err != nil {
		return fmt.Errorf("hdadb.Swipe failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := ApprovalTap(); err != nil {
		return fmt.Errorf("Failed Tapping: %w", err)
	}

	if err := Close(); err != nil {
		return fmt.Errorf("List failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := Close(); err != nil {
		return fmt.Errorf("List failed: %w", err)
	}
	SleepBetweenSequenceAction()
	return nil
}

func List() error {
	return hdadb.Tap("600", "1100")
}

func OpenDevelopment() error {
	return hdadb.Tap(GetDevelopmentPosition())
}

func OpenScience() error {
	return hdadb.Tap(GetSciencePosition())
}

func OpenInterior() error {
	return hdadb.Tap(GetInteriorPosition())
}

func OpenSecurity() error {
	return hdadb.Tap(GetSecurityPosition())
}

func OpenStrategy() error {
	return hdadb.Tap(GetStategyPosition())
}

func OpenMilitaryCommander() error {
	return hdadb.Tap(GetMilitaryCommandPosition())
}

func OpenAdministrativeCommander() error {
	return hdadb.Tap(GetAdministrativeCommandPosition())
}

func Close() error {
	log.Println("Closing menu...")
	return hdadb.Tap("650", "100")
}

func GetAmountOfTaps() int {
	if amountOfTapsSet {
		return amountOfTaps
	}

	TapsStr := os.Getenv("AMOUNT_OF_TapS")

	Taps, err := strconv.Atoi(TapsStr)
	if err != nil {
		fmt.Printf("ERROR: AMOUNT_OF_TapS must be an integer (got '%s'). DEFAULTING TO 10", TapsStr)
		Taps = 10
	}

	if Taps < 0 {
		fmt.Printf("ERROR: AMOUNT_OF_TapS cannot be negative (got %d). DEFAULTING TO 10", Taps)
		Taps = 10
	}

	amountOfTaps = Taps
	amountOfTapsSet = true
	return amountOfTaps
}
