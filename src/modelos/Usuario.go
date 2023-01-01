package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id"`
	Nome     string    `json:"nome"`
	Nick     string    `json:"nick"`
	Email    string    `json:"email"`
	Senha    string    `json:"senha"`
	CriadoEm time.Time `json:"CriadoEm"`
}

// Preparar vai chamar os metodos para validar e formatar o usuario recebido
func (usuario *Usuario) Preparar() error{
	if erro := usuario.Validar(); erro != nil{
		return erro
	}

	usuario.formatar()
	return nil

}

func (usuario *Usuario) Validar() error{
	if usuario.Nome == ""{
		return errors.New("O nome é obrigatório e não pode ficar em branco!")
	}
	if usuario.Nick == ""{
		return errors.New("O nick é obrigatório e não pode ficar em branco!")
	}
	if usuario.Email == ""{
		return errors.New("O email é obrigatório e não pode ficar em branco!")
	}
	if usuario.Senha == ""{
		return errors.New("A senha é obrigatório e não pode ficar em branca!")
	}

	return nil
}

func (usuario *Usuario) formatar(){
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}