package model

import (
	"fmt"
)

const (
	orderFilename = "receipt-%d.md"

	markdownTemplate = `# Order: %d

| Created At      | Drink ID | Amount |
|-----------------|----------|--------|
| %s | %d | %d |

Thanks for drinking with us!`
)

type Order struct {
	Base
	Amount  uint64 `json:"amount"`
	DrinkID uint   `json:"drink_id" gorm:"not null"`
	Drink   Drink  `json:"drink"`
}

func (o *Order) ToMarkdown() string {
	return fmt.Sprintf(
		markdownTemplate,
		o.ID,
		o.CreatedAt.Format("Jan 02 15:04:05"),
		o.DrinkID,
		o.Amount,
	)
}

func (o *Order) Filename() string {
	return fmt.Sprintf(orderFilename, o.ID)
}
