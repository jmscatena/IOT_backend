package sensorrepository

import (
	"context"

	"backend/core/domain"
	"backend/core/dto"
	"github.com/booscaaa/go-paginate/paginate"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParms) (*domain.Pagination[[]domain.Sensor], error) {
	ctx := context.Background()
	Sensors := []domain.Sensor{}
	total := int32(0)
	pagin := paginate.Instance(paginate.Pagination{})
	query, queryCount := pagin.
		Query("SELECT s.* FROM Sensor s").
		Sort(pagination.Sort).
		Desc(pagination.Descending).
		Page(pagination.Page).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "device", "corporation", "description").
		Select()
	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			Sensor := domain.Sensor{}

			rows.Scan(
				&Sensor.ID,
				&Sensor.Device,
				&Sensor.Description,
			)

			Sensors = append(Sensors, Sensor)
		}
	}

	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination[[]domain.Sensor]{
		Items: Sensors,
		Total: total,
	}, nil
}
