package transaction

import (
	"fmt"
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
)

type (
	ResourcesHolder struct {
		resourceHandles map[string]*ResourceHandle
	}
)

func (h *ResourcesHolder) EnlistResource(key string, r transaction.Resource) {
	if h == nil {
		panic(fmt.Errorf("ResourcesHolder is nil"))
	}
	if h.resourceHandles == nil {
		h.resourceHandles = make(map[string]*ResourceHandle)
	}

	if current := h.resourceHandles[key]; current != nil {
		panic(fmt.Errorf("there is already this resource=[%v] enlisted with key=[%s]", current.Resource, key))
	}
	h.resourceHandles[key] = &ResourceHandle{Resource: r}
}

func (h *ResourcesHolder) GetResource(key string) transaction.Resource {
	if h == nil {
		return nil
	}
	if handle, found := h.resourceHandles[key]; found {
		return handle.Resource
	}
	return nil
}

func (h *ResourcesHolder) foreach(from int, do func(handle *ResourceHandle) error) (index int, err error) {
	if h == nil {
		return -1, nil
	}
	index = 0
	for _, handle := range h.resourceHandles {
		if index < from || handle == nil {
			continue
		}
		index++
		if err = do(handle); err != nil {
			return
		}
	}
	return
}
