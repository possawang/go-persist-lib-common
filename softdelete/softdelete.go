package softdelete

import "time"

type SoftDeleteModel struct {
	Deleted   bool       `gorm:"column:deleted"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	CreatedBy string     `gorm:"column:created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	UpdatedBy *string    `gorm:"column:updated_by"`
}
