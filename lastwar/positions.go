package lastwar

import (
	"lastwar/notifier/states"
	"os"
)

func GetAlertPositions() map[string]struct{ x, y, w, h int } {
	if os.Getenv("APPLICATION_SATE") == states.CONQUERED {
		return map[string]struct{ x, y, w, h int }{
			"science":     {262, 850, 30, 30},
			"development": {40, 850, 30, 30},
			"interior":    {480, 850, 52, 50},
			"security":    {470, 580, 52, 50},
			"strategy":    {255, 580, 52, 50},
		}
	}

	if os.Getenv("APPLICATION_SATE") == states.CONQUERER {
		return map[string]struct{ x, y, w, h int }{
			"science":            {264, 852, 24, 32},
			"development":        {50, 850, 22, 31},
			"interior":           {481, 853, 22, 33},
			"security":           {480, 577, 22, 32},
			"strategy":           {265, 577, 21, 31},
			"admin_commander":    {391, 302, 20, 30},
			"military_commander": {140, 299, 23, 33},
		}
	}

	return map[string]struct{ x, y, w, h int }{
		"science":     {251, 750, 52, 50},
		"development": {36, 750, 52, 50},
		"interior":    {466, 750, 52, 50},
		"security":    {466, 472, 52, 50},
		"strategy":    {250, 472, 52, 50},
	}
}

func GetDevelopmentPosition() (string, string) {
	if os.Getenv("APPLICATION_SATE") == states.CONQUERED {
		return "150", "1000"
	}

	if os.Getenv("APPLICATION_SATE") == states.CONQUERER {
		return "133", "970"
	}

	return "100", "800"
}

func GetSciencePosition() (string, string) {
	if os.Getenv("APPLICATION_SATE") == states.CONQUERED {
		return "400", "1000"
	}

	if os.Getenv("APPLICATION_SATE") == states.CONQUERER {
		return "353", "999"
	}

	return "300", "800"
}

func GetInteriorPosition() (string, string) {
	if os.Getenv("APPLICATION_SATE") == states.CONQUERED {
		return "600", "1000"
	}

	if os.Getenv("APPLICATION_SATE") == states.CONQUERER {
		return "570", "967"
	}

	return "500", "800"
}

func GetSecurityPosition() (string, string) {
	if os.Getenv("APPLICATION_SATE") == states.CONQUERED {
		return "575", "698"
	}

	if os.Getenv("APPLICATION_SATE") == states.CONQUERER {
		return "576", "693"
	}
	return "500", "500"
}

func GetStategyPosition() (string, string) {
	if os.Getenv("APPLICATION_SATE") == states.CONQUERED {
		return "359", "704"
	}

	if os.Getenv("APPLICATION_SATE") == states.CONQUERER {
		return "351", "678"
	}

	return "350", "500"
}

func GetAdministrativeCommandPosition() (string, string) {
	return "474", "411"
}

func GetMilitaryCommandPosition() (string, string) {
	return "228", "409"
}
