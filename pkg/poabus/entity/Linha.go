package entity

type Linha struct {
	Nome                                                                       string
	HorariosUteisInvernoBC, HorariosSabadosInvernoBC, HorariosDomingoInvernoBC []Horario
	HorariosUteisInvernoCB, HorariosSabadosInvernoCB, HorariosDomingoInvernoCB []Horario
	HorariosUteisVeraoBC, HorariosSabadosVeraoBC, HorariosDomingoVeraoBC       []Horario
	HorariosUteisVeraoCB, HorariosSabadosVeraoCB, HorariosDomingoVeraoCB       []Horario
}
