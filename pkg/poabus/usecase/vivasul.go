package usecase

import (
	"strings"

	"github.com/gocolly/colly"
	"github.com/nenitf/poabus/pkg/poabus/entity"
)

type useCase struct {
}

func NovaBuscaPeloVivaSul() *useCase {
	return &useCase{}
}

func (s *useCase) Execute(url string) (linha entity.Linha, err error) {
	c := colly.NewCollector()

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		split := strings.Split(e.Text, "Mapa da Rota")
		linha.Nome = strings.TrimSpace(split[0])
	})

	c.OnHTML("#uteisibc table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Úteis", Operacao: "Inverno", Rota: "Bairro/Centro", Horario: e.Text}
		linha.HorariosUteisInvernoBC = append(linha.HorariosUteisInvernoBC, h)
	})

	c.OnHTML("#sabadoibc table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Sábado", Operacao: "Inverno", Rota: "Bairro/Centro", Horario: e.Text}
		linha.HorariosSabadosInvernoBC = append(linha.HorariosSabadosInvernoBC, h)
	})

	c.OnHTML("#domingoibc table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Domingo", Operacao: "Inverno", Rota: "Bairro/Centro", Horario: e.Text}
		linha.HorariosDomingoInvernoBC = append(linha.HorariosDomingoInvernoBC, h)
	})

	c.OnHTML("#uteisvbc table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Úteis", Operacao: "Verão", Rota: "Bairro/Centro", Horario: e.Text}
		linha.HorariosUteisVeraoBC = append(linha.HorariosUteisVeraoBC, h)
	})

	c.OnHTML("#sabadovbc table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Sábado", Operacao: "Verão", Rota: "Bairro/Centro", Horario: e.Text}
		linha.HorariosSabadosVeraoBC = append(linha.HorariosSabadosVeraoBC, h)
	})

	c.OnHTML("#domingovbc table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Domingo", Operacao: "Verão", Rota: "Bairro/Centro", Horario: e.Text}
		linha.HorariosDomingoVeraoBC = append(linha.HorariosDomingoVeraoBC, h)
	})

	c.OnHTML("#uteisicb table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Úteis", Operacao: "Inverno", Rota: "Centro/Bairro", Horario: e.Text}
		linha.HorariosUteisInvernoCB = append(linha.HorariosUteisInvernoCB, h)
	})

	c.OnHTML("#sabadoicb table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Sábado", Operacao: "Inverno", Rota: "Centro/Bairro", Horario: e.Text}
		linha.HorariosSabadosInvernoCB = append(linha.HorariosSabadosInvernoCB, h)
	})

	c.OnHTML("#domingoicb table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Domingo", Operacao: "Inverno", Rota: "Centro/Bairro", Horario: e.Text}
		linha.HorariosDomingoInvernoCB = append(linha.HorariosDomingoInvernoCB, h)
	})

	c.OnHTML("#uteisvcb table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Úteis", Operacao: "Verão", Rota: "Centro/Bairro", Horario: e.Text}
		linha.HorariosUteisVeraoCB = append(linha.HorariosUteisVeraoCB, h)
	})

	c.OnHTML("#sabadovcb table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Sábado", Operacao: "Verão", Rota: "Centro/Bairro", Horario: e.Text}
		linha.HorariosSabadosVeraoCB = append(linha.HorariosSabadosVeraoCB, h)
	})

	c.OnHTML("#domingovcb table td", func(e *colly.HTMLElement) {
		h := entity.Horario{Dia: "Domingo", Operacao: "Verão", Rota: "Centro/Bairro", Horario: e.Text}
		linha.HorariosDomingoVeraoCB = append(linha.HorariosDomingoVeraoCB, h)
	})

	// c.OnRequest(func(r *colly.Request) {
	// fmt.Println("Visiting", r.URL)
	// })

	c.Visit(url)

	return
}
