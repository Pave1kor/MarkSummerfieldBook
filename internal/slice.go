package internal

type SafeSlicer interface {
	Append(interface{})     //добавляет элемент в конец среза
	At(int) interface{}     //Возвращает элемент с указанным индексом
	Close() []interface{}   //Закрывает канал и возвращает срез
	Delete(int)             //Удаляет элемент с указанным индексом
	Len() int               //Возвращает количество элементов в срезе
	Update(int, UpdateFunc) //обновляет элемент с указанным индексом
}
type UpdateFunc func(interface{}) interface{}
type safeSlice chan commandData

type commandData struct {
	action  coomandAction
	index   int
	value   interface{}
	data    chan<- []interface{}
	result  chan<- interface{}
	updater UpdateFunc
}
type coomandAction int

const (
	appendVal coomandAction = iota
	at
	end
	delete
	lenght
	update
)

func (ss safeSlice) Append(val interface{}) {
	ss <- commandData{action: appendVal, value: val}
}
func (ss safeSlice) At(index int) interface{} {
	result := make(chan interface{})
	ss <- commandData{action: at, index: index, result: result}
	return <-result
}
func (ss safeSlice) Close() []interface{} {
	data := make(chan []interface{})
	ss <- commandData{action: end, data: data}
	return <-data
}
func (ss safeSlice) Delete(index int) {
	ss <- commandData{action: delete, index: index}
}
func (ss safeSlice) Len() int {
	result := make(chan interface{})
	ss <- commandData{action: lenght, result: result}
	return (<-result).(int)
}
func (ss safeSlice) Update(index int, updater UpdateFunc) {
	ss <- commandData{action: update, index: index, updater: updater}
}
func New() safeSlice {
	sm := make(safeSlice)
	go sm.run()
	return sm
}

func (ss safeSlice) run() {
	store := make([]interface{}, 0)
	for cmd := range ss {
		switch cmd.action {
		case appendVal:
			store = append(store, cmd.value)
		case at:
			if cmd.index < len(store) && cmd.index >= 0 {
				cmd.result <- store[cmd.index]
			} else {
				cmd.result <- nil
			}
		case end:
			cmd.data <- store
			close(ss)
		case delete:
			if cmd.index < len(store) && cmd.index >= 0 {
				store = append(store[:cmd.index], store[cmd.index+1:]...)
			}
		case lenght:
			cmd.result <- len(store)
		case update:
			if cmd.index <= len(store) && cmd.index >= 0 {
				store[cmd.index] = cmd.updater(store[cmd.index])
			}
		}
	}
}
