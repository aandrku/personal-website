package auth

import (
	"time"

	"github.com/aandrku/personal-website/pkg/services/email"
	"github.com/google/uuid"
)

var currentOTP uuid.UUID
var lastRefresh time.Time

func Refresh() error {
	currentOTP = uuid.New()
	lastRefresh = time.Now()
	if err := email.SendString("aandrku.dev OTP", currentOTP.String()); err != nil {
		return err
	}
	return nil
}

func GetCurrentOTP() (string, error) {
	if time.Since(lastRefresh) > time.Minute*5 {
		if err := Refresh(); err != nil {
			return "", err
		}
	}

	return currentOTP.String(), nil
}
