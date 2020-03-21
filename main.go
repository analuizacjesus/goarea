package goarea

import "math"

//Pi = 3.1416
const Pi = 3.1416

//Circ é responsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Nao eh visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) * 2
}
