package todo

import (
	"time"
)

type item struct { // list to serve our  todo values.
	Task string
	Done bool
	CreatedAt time.time
	CompletedAt time.time
}

type Todos []item // an empty slice to reference the todo values.

// method to handle todo
func (t *Todos) Add(task string){
	todo := item{
		Task: task,
		Done: false,
		CreatedAt: time.Now()
		CompletedAt: time.Time{}
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return mil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func (*t Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}