package algo

import (
	"cmp"
	"slices"

	"github.com/KKogaa/things-todo-backend/core/entity"
)

type PrioritySchedule struct {
}

func NewPrioritySchedule() PrioritySchedule {
	return PrioritySchedule{}
}

// order tasks based on priorty, OR also addd by duration?
func (p PrioritySchedule) Schedule(taskManager TaskManager) {
	slices.SortFunc(taskManager.Tasks, func(a, b *entity.Task) int {
		return cmp.Compare(b.Priority, a.Priority)
	})

}
