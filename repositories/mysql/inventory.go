package mysql

import "github.com/adebmbng/bcg-test/entities"

// GetInventoriesBySKUs get list of inventory by SKUs
func (r *repository) GetInventoriesBySKUs(q []string) ([]*entities.Inventory, error) {
	var res []*entities.Inventory
	err := r.db.Model(entities.Inventory{}).Where(`sku IN (?)`, res).Where(`is_active = ?`, 1).Error
	return res, err
}
