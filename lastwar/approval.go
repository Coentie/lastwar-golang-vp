package lastwar

import (
	"fmt"
	"github.com/fatih/color"
	"lastwar/notifier/hdadb"
	"lastwar/notifier/image"
	"math/rand"
	"os"
	"time"
)

func Approve(label string) error {
	if label == "development" {
		if err := approveDevelopment(); err != nil {
			return fmt.Errorf("Approving development failed: %w", err)
		}
		SleepBetweenSequenceAction()
	}

	if label == "science" {
		if err := approveScience(); err != nil {
			return fmt.Errorf("Approving development failed: %w", err)
		}
		SleepBetweenSequenceAction()
	}

	if label == "interior" {
		if err := approveInterior(); err != nil {
			return fmt.Errorf("Approving interior failed: %w", err)
		}
		SleepBetweenSequenceAction()
	}

	if label == "security" {
		if err := approveSecurity(); err != nil {
			return fmt.Errorf("Approving security failed: %w", err)
		}
		SleepBetweenSequenceAction()
	}

	if label == "strategy" {
		if err := approveStrategy(); err != nil {
			return fmt.Errorf("Approving stategy failed: %w", err)
		}
		SleepBetweenSequenceAction()
	}

	if label == "strategy" {
		if err := approveStrategy(); err != nil {
			return fmt.Errorf("Approving stategy failed: %w", err)
		}
		SleepBetweenSequenceAction()
	}

	if label == "admin_commander" {
		if err := approveAdministrativeCommander(); err != nil {
			return fmt.Errorf("Approving administrative commander failed: %w", err)
		}
		SleepBetweenSequenceAction()
	}

	if label == "military_commander" {
		if err := approveMilitaryCommander(); err != nil {
			return fmt.Errorf("Approving military commander failed: %w", err)
		}
		SleepBetweenSequenceAction()
	}

	return nil
}

func approveDevelopment() error {
	SleepBetweenSequenceAction()
	if err := OpenDevelopment(); err != nil {
		return fmt.Errorf("openDevelopment failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := ApproveMemberSequence(); err != nil {
		return fmt.Errorf("Approving members: %w", err)
	}

	return nil
}

func approveAdministrativeCommander() error {
	SleepBetweenSequenceAction()
	if err := OpenAdministrativeCommander(); err != nil {
		return fmt.Errorf("openDevelopment failed: %w", err)
	}
	if err := ApproveMemberSequence(); err != nil {
		return fmt.Errorf("Approving members: %w", err)
	}

	return nil
}

func approveMilitaryCommander() error {
	SleepBetweenSequenceAction()
	if err := OpenMilitaryCommander(); err != nil {
		return fmt.Errorf("openDevelopment failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := ApproveMemberSequence(); err != nil {
		return fmt.Errorf("Approving members: %w", err)
	}

	return nil
}

func approveScience() error {
	SleepBetweenSequenceAction()
	if err := OpenScience(); err != nil {
		return fmt.Errorf("openDevelopment failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := ApproveMemberSequence(); err != nil {
		return fmt.Errorf("Approving members: %w", err)
	}

	return nil
}

func approveInterior() error {
	SleepBetweenSequenceAction()
	if err := OpenInterior(); err != nil {
		return fmt.Errorf("openDevelopment failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := ApproveMemberSequence(); err != nil {
		return fmt.Errorf("Approving members: %w", err)
	}

	return nil
}

func approveSecurity() error {
	SleepBetweenSequenceAction()
	if err := OpenSecurity(); err != nil {
		return fmt.Errorf("openDevelopment failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := ApproveMemberSequence(); err != nil {
		return fmt.Errorf("Approving members: %w", err)
	}

	return nil
}

func approveStrategy() error {
	SleepBetweenSequenceAction()
	if err := OpenStrategy(); err != nil {
		return fmt.Errorf("openDevelopment failed: %w", err)
	}
	SleepBetweenSequenceAction()
	if err := ApproveMemberSequence(); err != nil {
		return fmt.Errorf("Approving members: %w", err)
	}

	return nil
}

func ApprovalTap() error {
	whitelistFilter := os.Getenv("ALLIANCE_FILTER")
	for i := 1; i <= GetAmountOfTaps(); i++ {
		if whitelistFilter == WHITELIST_FITLER_ALL {
			hdadb.Tap("517", "280")
			SleepBetweenTap()
			continue
		}
		name, err := image.ReadNameFromRegion()

		if name == "ee" || name == "Oe" || name == "" {
			return nil
		}

		if err != nil {
			fmt.Println("error reading name, approving")
			continue
		}

		if IsAllianceOfInterest(name) && whitelistFilter == WHITELIST_FITLER_WHITELIST {
			color.Green(fmt.Sprintf("Approved: %s", name))
			hdadb.Tap("517", "280")
			SleepBetweenTap()
			continue
		} else {
			color.Red(fmt.Sprintf("Rejected: %s", name))
			// Reject
			hdadb.Tap("611", "279")
			SleepBetweenTap()
			// Confirm
			hdadb.Tap("467", "870")
			SleepBetweenTap()
			continue
		}

		if IsAllianceOfInterest(name) && whitelistFilter == WHITELIST_FITLER_BLACKLIST {
			color.Red(fmt.Sprintf("Rejected: %s", name))
			hdadb.Tap("611", "279")
			SleepBetweenTap()
			// Confirm
			hdadb.Tap("467", "870")
			SleepBetweenTap()
			continue
		} else {
			color.Green(fmt.Sprintf("Approved: %s", name))
			// approve alliances out of the blacklist
			hdadb.Tap("517", "280")
			SleepBetweenTap()
			continue
		}
	}

	return nil
}

func SleepBetweenSequenceAction() {
	randomSleep()
}

func SleepBetweenTap() {
	randomSleep()
}

func randomSleep() {
	min := 1000
	max := 1500
	duration := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(duration) * time.Millisecond)
}
