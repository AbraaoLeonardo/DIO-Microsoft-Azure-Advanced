package tarefas

type Tarefa struct {
	ID        int64  `json:"id"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Status    string `json:"status"`
}
