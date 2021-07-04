package mysql

import "github.com/adebmbng/bcg-test/entities"

// GetPromoBySKUs to get promos data by skus
func (r *repository) GetPromoBySKUs(q []string) ([]*entities.Promotion, error) {
	var res []*entities.Promotion
	err := r.db.Model(entities.Promotion{}).Where(`item IN (?)`, res).Where(`is_active = ?`, 1).Error
	return res, err
}
