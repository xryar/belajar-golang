package main

import (
	f "fmt"
	m "math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return m.Phi * m.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return m.Phi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return m.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// embedded interface
type hitung2d interface {
	luas() float64
	kelilingDimensi() float64
}

type hitung3d interface {
	volume() float64
}

type hitungDimensi interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return m.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return m.Pow(k.sisi, 2) * 6
}

func (k *kubus) kelilingDimensi() float64 {
	return k.sisi * 12
}

func main() {
	var bangunRuang hitungDimensi = &kubus{4}
	var bangunDatar hitung = lingkaran{14.0}
	var bangunLingkaran lingkaran = bangunDatar.(lingkaran)

	f.Println("==== kubus")
	f.Println("luas		  :", bangunRuang.luas())
	f.Println("keliling   :", bangunRuang.kelilingDimensi())
	f.Println("volume 	  :", bangunRuang.volume())

	bangunDatar = persegi{10.0}
	f.Println("==== persegi")
	f.Println("luas 	  :", bangunDatar.luas())
	f.Println("keliling   :", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	f.Println("==== Keliling")
	f.Println("luas 	  :", bangunDatar.luas())
	f.Println("keliling   :", bangunDatar.keliling())
	f.Println("jari-jari  :", bangunLingkaran.jariJari())
}
