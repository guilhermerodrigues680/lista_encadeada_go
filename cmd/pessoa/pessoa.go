package pessoa

// Pessoaer - caracteristicas de  uma pessoa
type Pessoaer interface {
	FezAniversario()
	GetNome() string
}

// Pessoa é uma pessoa
type pessoa struct {
	Nome, Sobrenome string
	Idade           int
}

// Incrementa 1 ano na idade da pessoa
func (p *pessoa) FezAniversario() {
	// p.idade++
	p.Idade++
}

func (p *pessoa) GetNome() string {
	return p.Nome
}

// func (p *pessoa) GetNome() string {
// 	return p.nome
// }

// NewPessoa retorna uma nova pessoa
func NewPessoa(nome, sobrenome string, idade int) Pessoaer {
	return &pessoa{nome, sobrenome, idade}
}
