package todo 

import (
	"fmt"
)

type item struct { //a struct type format to help us store our values
	Task string
	Done bool
	CreatedAt time.time
	CompletedAt time.Time
}

//declare a slice called Todos of the data-type item created above.
type Todos []item

func (t *Todos) Add(task string) { //function to add task to our todo-list,t is a variable pointing to our slice Todos
	//which is of the struct data type "item".
	todo := item{
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}
	*t= append(t, todo)
}
func (t *Todos) Complete(index int) error {
	ls := *t 
	if index <= 0 || > len(ls) {
		return fmt.Errorf("invalid index")
	}
	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true
	return nil
}

func (*t Todos) Delete(index int) error {
	ls := *t 
	if index <= 0 || > len(ls) {
		return fmt.Errorf("invalid index")
	}
	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (*t Todos) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
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