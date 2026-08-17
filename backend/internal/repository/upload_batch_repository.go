package repository

import (
	"errors"
	"time"

	"github.com/blueship581/gbinsureapi/internal/model"
	"github.com/blueship581/gbinsureapi/internal/util"
	"gorm.io/gorm"
)

// UploadBatchRepository 上传批次仓储。
type UploadBatchRepository struct{ db *gorm.DB }

// NewUploadBatchRepository 构造上传批次仓储。
func NewUploadBatchRepository(db *gorm.DB) *UploadBatchRepository {
	return &UploadBatchRepository{db: db}
}

// WithTx 使用事务连接构造仓储，便于 service 层在事务内编排多次写入。
func (r *UploadBatchRepository) WithTx(tx *gorm.DB) *UploadBatchRepository {
	return &UploadBatchRepository{db: tx}
}

// Transaction 在事务内执行 fn，任一步返回 error 则整体回滚。
func (r *UploadBatchRepository) Transaction(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

// Create 创建批次。
func (r *UploadBatchRepository) Create(batch *model.UploadBatch) error { return r.db.Create(batch).Error }

// FindByID 按 ID 查询。
func (r *UploadBatchRepository) FindByID(id uint) (*model.UploadBatch, error) {
	var batch model.UploadBatch
	if err := r.db.First(&batch, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util.ErrNotFound
		}
		return nil, err
	}
	return &batch, nil
}

// FindByNo 按批次号查询。
func (r *UploadBatchRepository) FindByNo(batchNo string) (*model.UploadBatch, error) {
	var batch model.UploadBatch
	if err := r.db.Where("batch_no = ?", batchNo).First(&batch).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util.ErrNotFound
		}
		return nil, err
	}
	return &batch, nil
}

// UpdateStatus 更新状态。
func (r *UploadBatchRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&model.UploadBatch{}).Where("id = ?", id).Update("upload_status", status).Error
}

// ExistsByClientInsuredDate 检查同一调用方/参保人当日是否已有批次（重复性检查）。
// 当日以 since 作为左闭边界（created_at >= since），由调用方传入 Asia/Shanghai 今日 00:00，
// 这样仅当日批次才视为重复，跨天批次不再误拦新上传。
func (r *UploadBatchRepository) ExistsByClientInsuredDate(clientID, insuredID uint, since time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&model.UploadBatch{}).
		Where("client_id = ? AND insured_person_id = ? AND created_at >= ?", clientID, insuredID, since).
		Count(&count).Error
	return count > 0, err
}

// ListByClient 分页查询（按调用方）。
func (r *UploadBatchRepository) ListByClient(clientID uint, page, pageSize int) ([]model.UploadBatch, int64, error) {
	q := r.db.Model(&model.UploadBatch{})
	if clientID > 0 {
		q = q.Where("client_id = ?", clientID)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var batches []model.UploadBatch
	err := r.db.Where("client_id = ?", clientID).Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&batches).Error
	return batches, total, err
}

// Count 统计。
func (r *UploadBatchRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&model.UploadBatch{}).Count(&count).Error
	return count, err
}

// Update 更新批次。
func (r *UploadBatchRepository) Update(batch *model.UploadBatch) error {
	return r.db.Save(batch).Error
}
