package common

type Pagination struct {
	CurrentPage  int `form:"currentPage" binding:"required,min=1"`
	ItemsPerPage int `form:"itemsPerPage" binding:"required,min=1,max=100"`
}

func (p *Pagination) CalculateOffset() Pagination {
	p.CurrentPage = p.ItemsPerPage * (p.CurrentPage - 1)
	return *p
}
