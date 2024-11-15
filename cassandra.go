package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	gocqlastra "github.com/datastax/gocql-astra"
	"github.com/gocql/gocql"
	"github.com/jaswdr/faker"
)

type Aluno struct {
	Name string
	Date string
}

type AlunosAprovados struct {
	aluno_id      gocql.UUID
	nome          string
	disciplina_id gocql.UUID
	semestre      int
	ano           int
	nota_final    float64
}

var cursos = []string{
	"Administracao",
	"Ciencia da Computacao",
	"Engenharia Civil",
	"Engenharia Eletrica",
}

var cursos_id = []gocql.UUID{}
var departamentos_id = []gocql.UUID{}
var disciplinas_adm_id = []gocql.UUID{}
var disciplinas_cc_id = []gocql.UUID{}
var disciplinas_engc_id = []gocql.UUID{}
var disciplinas_enge_id = []gocql.UUID{}
var alunos_id = []gocql.UUID{}
var professores_id = []gocql.UUID{}

func pickRandom(arr []gocql.UUID) gocql.UUID {
	return arr[rand.Intn(len(arr))]
}

func randYear() int {
	min := time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Year()
}

func main() {

	if len(os.Getenv("ASTRA_DB_APPLICATION_TOKEN")) > 0 {

		if len(os.Getenv("ASTRA_DB_ID")) == 0 {
			panic("database ID is required when using a token")
		}
	}

	cluster, _ := gocqlastra.NewClusterFromURL("https://api.astra.datastax.com", os.Getenv("ASTRA_DB_ID"), os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), 10*time.Second)
	cluster.Keyspace = " " // Especifique o keyspace aqui
	cluster.Consistency = gocql.Quorum
	session, _ := gocql.NewSession(*cluster)

	m := make(map[string][]string)

	m["Administracao"] = []string{
		"Matematica aplicada a administracao",
		"Fundamentos da administracao",
		"Estudos em macroeconomia",
		"Sociologia I",
		"Sociologia II",
		"Linguagens e generos textuais",
		"Etica nas organizacoes",
		"Ensino social cristao",
		"Calculo basico I",
		"Calculo basico II",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia I",
		"Filosofia II",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Gramatica",
		"Estudos em microeconomia",
		"Sustentabilidade",
		"Probabilidade estatistica",
		"Estudos do MEI",
		"Gestao de negocios",
		"Gestao de pessoas",
		"Gesta Financeira",
		"Matematica Financeira",
		"Gestao de recusos humanos",
		"PowerBI",
		"Excel Basico",
		"Excel intermediario",
		"Excel avancado",
		"Redacao",
		"Comunicacao para empresas",
		"Estrategia de marketing",
		"Direito Tributario",
		"Logistica",
		"Implementacao de negocios",
	}

	m["Ciencia da Computacao"] = []string{
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Sociologia",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"Circuitos eletricos",
		"Ecologia e sustentabilidade",
		"Sistemas embarcados",
		"Redes de computadores",
		"Calculo numerico",
		"Termodinamica",
		"Cinematica e dinamica",
		"Algebra Linear",
		"Automatos",
		"Redes moveis",
		"Algoritmos",
		"Sistema Operacional",
		"Computacao grafica",
		"Compiladores",
		"Empreendedorismo",
		"Inovacao",
		"Teste de software",
		"IA",
		"Orientacao a objetos",
		"Desenvolvimento de jogos",
		"Robotica",
		"TCC I",
		"TCC II",
	}

	m["Engenharia Civil"] = []string{
		"Ensino social cristao",
		"Calculo basico",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Desenho tecnico",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"Arquitetura e representacao grafica",
		"Mecanica geral",
		"Topografia",
		"Eletricidade geral",
		"Instalacoes eletricas",
		"Economia",
		"Geotecnia I",
		"Geotecnia II",
		"Transportes I",
		"Transportes II",
		"Optativa I",
		"Optativa II",
		"Custos",
		"Tecnologia das construcoes",
		"Metodos estatisticos",
		"Termodinamica",
		"Gestao de obras",
		"Planejamento de obras",
	}

	m["Engenharia Eletrica"] = []string{
		"Ensino social cristao",
		"Calculo basico",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Desenho tecnico",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"sinais e Sistemas",
		"Praticas de inovacao I",
		"Praticas de inovacao II",
		"Praticas de inovacao III",
		"Circuitos integrados",
		"Mecanica Geral",
		"Gestao organizacional",
		"TCC 1",
		"TCC 2",
		"Redes",
		"Mecanica dos solidos",
		"Mecanica dos fluidos",
		"Sistemas eletricos",
		"Seguranca do trabalho",
		"Economia",
		"Conversao de energia I",
		"Conversao de energia II",
		"Conversao de energia III",
	}

	departamentos := []string{
		"Matematica",
		"Ciencia da Computacao",
		"Quimica",
		"Fisica",
		"Engenharia",
	}

	fakeName := faker.New().Person()
	fakeDate := faker.New().Time()

	var pk gocql.UUID
	for i := 0; i < 20; i++ {
		pk = insertAluno(session, Aluno{fakeName.Name(), fakeDate.UnixDate(time.Now())})
		alunos_id = append(alunos_id, pk)
	}

	for i := 0; i < 10; i++ {
		pk = insertProfessor(session, fakeName.Name())
		professores_id = append(professores_id, pk)
	}

	for _, curso := range cursos {
		pk = insertCurso(session, curso)
		cursos_id = append(cursos_id, pk)
	}

	for _, departamento := range departamentos {
		pk = insertDepartamento(session, departamento)
		departamentos_id = append(departamentos_id, pk)
	}

	for _, disc := range m["Administracao"] {
		pk = insertDisciplina(session, disc)
		disciplinas_adm_id = append(disciplinas_adm_id, pk)
	}

	for _, disc := range m["Ciencia da Computacao"] {
		pk = insertDisciplina(session, disc)
		disciplinas_cc_id = append(disciplinas_cc_id, pk)
	}

	for _, disc := range m["Engenharia Civil"] {
		pk = insertDisciplina(session, disc)
		disciplinas_engc_id = append(disciplinas_engc_id, pk)

	}

	for _, disc := range m["Engenharia Eletrica"] {
		pk = insertDisciplina(session, disc)
		disciplinas_enge_id = append(disciplinas_enge_id, pk)

	}

	atualizarFKAluno(session)
	atualizarFKCurso(session)
	atualizarFKDepartamento(session)

	atualizarFKDisciplina(session, disciplinas_adm_id[:], 0)
	atualizarFKDisciplina(session, disciplinas_cc_id[:], 1)
	atualizarFKDisciplina(session, disciplinas_engc_id[:], 2)
	atualizarFKDisciplina(session, disciplinas_enge_id[:], 3)

	atualizarFKProfessor(session)
	insertDisciplinasM(session, professores_id, disciplinas_adm_id)
	insertDisciplinasM(session, professores_id, disciplinas_cc_id)
	insertDisciplinasM(session, professores_id, disciplinas_engc_id)
	insertDisciplinasM(session, professores_id, disciplinas_enge_id)

	for i := 0; i < 6; i++ {
		query := `INSERT INTO grupo_tcc (tcc_id, aluno1_id, aluno2_id, aluno3_id, orientador_id) VALUES (?, ?, ?, ?, ?)`

		tccID := gocql.TimeUUID()
		aluno1 := pickRandom(alunos_id)
		aluno2 := pickRandom(alunos_id)
		aluno3 := pickRandom(alunos_id)
		orientador := pickRandom(professores_id)

		session.Query(query, tccID, aluno1, aluno2, aluno3, orientador).Exec()

		session.Query(query, aluno1, aluno2, aluno3, orientador).Exec()

	}

	for _, aluno := range alunos_id {
		query := `INSERT INTO historico_escolar (aluno_id, disciplina_id, semestre, ano, nota_final, disciplina_nome) VALUES (?, ?, ?, ?, ?, ?)`

		var dataEntrada string
		var curso gocql.UUID
		session.Query(`SELECT data_ingresso FROM aluno WHERE aluno_id = ?`, aluno).Scan(&dataEntrada)
		session.Query(`SELECT curso_id FROM aluno WHERE aluno_id = ?`, aluno).Scan(&curso)

		date := strings.Split(dataEntrada, " ")
		ano, _ := strconv.ParseInt(date[len(date)-1], 10, 64)

		var d []gocql.UUID
		switch curso {
		case cursos_id[0]:
			d = disciplinas_adm_id
		case cursos_id[1]:
			d = disciplinas_cc_id
		case cursos_id[2]:
			d = disciplinas_engc_id
		case cursos_id[3]:
			d = disciplinas_enge_id
		}

		for i, disc := range d {

			var nome string

			query1 := `SELECT nome FROM disciplina WHERE disciplina_id = ?`

			session.Query(query1, disc).Scan(&nome)
			session.Query(query,
				aluno,
				disc,
				int(i/5)+1,
				ano+int64(i/10),
				rand.Float32()*10,
				nome).Exec()
		}
	}

	for _, curso := range cursos_id {
		query := `INSERT INTO matriz_curricular (curso_id, disciplina_id, semestre) VALUES (?, ?, ?)`

		var d []gocql.UUID
		switch curso {
		case cursos_id[0]:
			d = disciplinas_adm_id
		case cursos_id[1]:
			d = disciplinas_cc_id
		case cursos_id[2]:
			d = disciplinas_engc_id
		case cursos_id[3]:
			d = disciplinas_enge_id
		}

		for i, disc := range d {
			session.Query(query,
				curso,
				disc,
				int(i/5)+1).Exec()
		}
	}
}

func insertProfessor(session *gocql.Session, nome string) gocql.UUID {

	professorID := gocql.TimeUUID()
	query := `INSERT INTO professor (professor_id, nome) VALUES (?, ?) IF NOT EXISTS`
	if err := session.Query(query, professorID, nome).Exec(); err != nil {
		log.Fatal(err)
	}

	return professorID
}

func atualizarFKProfessor(session *gocql.Session) {
	query := `UPDATE professor SET departamento_id=? WHERE professor_id=?`
	for _, id := range professores_id {
		session.Query(query,
			pickRandom(departamentos_id[:]),
			id).Exec()
	}
}

func insertAluno(session *gocql.Session, aluno Aluno) gocql.UUID {
	alunoID := gocql.TimeUUID()

	query := `INSERT INTO aluno (aluno_id, nome, data_ingresso) VALUES (?, ?, ?) IF NOT EXISTS`
	if err := session.Query(query, alunoID, aluno.Name, aluno.Date).Exec(); err != nil {
		log.Fatal(err)
	}

	return alunoID
}

func atualizarFKAluno(session *gocql.Session) {
	query := `UPDATE aluno SET curso_id=? WHERE aluno_id=?`
	for _, id := range alunos_id {
		session.Query(query,
			pickRandom(cursos_id[:]),
			id).Exec()
	}
}

func insertCurso(session *gocql.Session, nome string) gocql.UUID {

	cursoID := gocql.TimeUUID()
	query := `INSERT INTO curso (curso_id, nome) VALUES (?, ?) IF NOT EXISTS`
	if err := session.Query(query, cursoID, nome).Exec(); err != nil {
		log.Fatal(err)
	}

	return cursoID
}

func atualizarFKCurso(session *gocql.Session) {
	query := `UPDATE curso SET departamento_id=? WHERE curso_id=?`
	for _, id := range cursos_id {
		session.Query(query,
			pickRandom(departamentos_id[:]),
			id).Exec()
	}
}

func insertDepartamento(session *gocql.Session, nome string) gocql.UUID {

	departamentoID := gocql.TimeUUID()
	query := `INSERT INTO departamento (departamento_id, nome) VALUES (?, ?) IF NOT EXISTS`
	if err := session.Query(query, departamentoID, nome).Exec(); err != nil {
		log.Fatal(err)
	}

	return departamentoID
}

func atualizarFKDepartamento(session *gocql.Session) {
	query := `UPDATE departamento SET chefe_id=? WHERE departamento_id=?`

	for _, id := range departamentos_id {
		session.Query(query, pickRandom(professores_id[:]), id).Exec()
	}
}

func insertDisciplina(session *gocql.Session, nome string) gocql.UUID {

	disciplinaID := gocql.TimeUUID()
	query := `INSERT INTO disciplina (disciplina_id, nome) VALUES (?, ?) IF NOT EXISTS`
	if err := session.Query(query, disciplinaID, nome).Exec(); err != nil {
		log.Fatal(err)
	}
	return disciplinaID
}

func atualizarFKDisciplina(session *gocql.Session, d []gocql.UUID, index int) {
	query := `UPDATE disciplina SET professor_id=?, curso_id=? WHERE disciplina_id=?`

	for _, id := range d {
		session.Query(query, pickRandom(professores_id[:]), pickRandom(cursos_id[:]), id).Exec()
	}
}

func insertDisciplinasM(session *gocql.Session, professoresID []gocql.UUID, disciplinasID []gocql.UUID) {
	for i, disc := range disciplinasID {
		var nome string

		query1 := `SELECT nome FROM disciplina WHERE disciplina_id = ?`

		session.Query(query1, disc).Scan(&nome)

		query := `INSERT INTO disciplina_ministrada (professor_id, disciplina_id, semestre, ano, disciplina_nome) VALUES (?, ?, ?, ?, ?)`

		professorID := pickRandom(professoresID[:])
		semestre := int(i/5) + 1
		ano := randYear()

		if err := session.Query(query, professorID, disc, semestre, ano, nome).Exec(); err != nil {
			fmt.Println("Erro ao inserir disciplina:", err)
		}
	}
}
