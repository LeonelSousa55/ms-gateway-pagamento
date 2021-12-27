package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreditCardNumber(t *testing.T) {

	_, err := NewCreditCard(
		"40000000000000000",
		"Jose da Silva",
		12,
		2024,
		123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard(
		"5102497878839657",
		"Jose da Silva",
		12,
		2024,
		123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMoth(t *testing.T) {

	_, err := NewCreditCard(
		"5102497878839657",
		"Jose da Silva",
		13,
		2024,
		123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard(
		"5102497878839657",
		"Jose da Silva",
		0,
		2024,
		123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard(
		"5102497878839657",
		"Jose da Silva",
		11,
		2024,
		123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {

	lastYear := time.Now().AddDate(-1, 0, 0)

	_, err := NewCreditCard(
		"5102497878839657",
		"Jose da Silva",
		12,
		lastYear.Year(),
		123)
	assert.Equal(t, "invalid expiration year", err.Error())
}
