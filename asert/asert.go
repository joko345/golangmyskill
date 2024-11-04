package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Misal kubus adalah struct yang memiliki metode Volume, Luas, dan Keliling
type Kubus struct {
	sisi float64
}

// Method Volume untuk menghitung volume kubus
func (k Kubus) Volume() float64 {
	return k.sisi * k.sisi * k.sisi
}

// Method Luas untuk menghitung luas kubus
func (k Kubus) Luas() float64 {
	return 6 * (k.sisi * k.sisi)
}

// Method Keliling untuk menghitung keliling kubus
func (k Kubus) Keliling() float64 {
	return 12 * k.sisi
}

func TestVolume(t *testing.T) {
	kubus := Kubus{sisi: 3}
	volumeAsli := 27.0 // Nilai volume yang diharapkan
	assert.Equal(t, kubus.Volume(), volumeAsli, "Perhitungan volume salah")
}

func TestLuas(t *testing.T) {
	kubus := Kubus{sisi: 3}
	luasAsli := 54.0 // Nilai luas yang diharapkan
	assert.Equal(t, kubus.Luas(), luasAsli, "Perhitungan luas salah")
}

func TestKeliling(t *testing.T) {
	kubus := Kubus{sisi: 3}
	kelilingAsli := 36.0 // Nilai keliling yang diharapkan
	assert.Equal(t, kubus.Keliling(), kelilingAsli, "Perhitungan keliling salah")
}
