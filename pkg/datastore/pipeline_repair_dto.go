package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PipelineRepair struct {
	ID                          primitive.ObjectID  `bson:"_id,omitempty"`
	PipelineId                  *string             `bson:"pipelineId,omitempty"`
	ExteriorRepairDescription1  *string             `bson:"exteriorRepairDescription1,omitempty"`
	ExteriorRepairPrice1        *float64            `bson:"exteriorRepairPrice1,omitempty"`
	ExteriorRepairDescription2  *string             `bson:"exteriorRepairDescription2,omitempty"`
	ExteriorRepairPrice2        *float64            `bson:"exteriorRepairPrice2,omitempty"`
	ExteriorRepairDescription3  *string             `bson:"exteriorRepairDescription3,omitempty"`
	ExteriorRepairPrice3        *float64            `bson:"exteriorRepairPrice3,omitempty"`
	ExteriorRepairDescription4  *string             `bson:"exteriorRepairDescription4,omitempty"`
	ExteriorRepairPrice4        *float64            `bson:"exteriorRepairPrice4,omitempty"`
	ExteriorRepairDescription5  *string             `bson:"exteriorRepairDescription5,omitempty"`
	ExteriorRepairPrice5        *float64            `bson:"exteriorRepairPrice5,omitempty"`
	ExteriorRepairDescription6  *string             `bson:"exteriorRepairDescription6,omitempty"`
	ExteriorRepairPrice6        *float64            `bson:"exteriorRepairPrice6,omitempty"`
	ExteriorRepairDescription7  *string             `bson:"exteriorRepairDescription7,omitempty"`
	ExteriorRepairPrice7        *float64            `bson:"exteriorRepairPrice7,omitempty"`
	ExteriorRepairDescription8  *string             `bson:"exteriorRepairDescription8,omitempty"`
	ExteriorRepairPrice8        *float64            `bson:"exteriorRepairPrice8,omitempty"`
	ExteriorRepairDescription9  *string             `bson:"exteriorRepairDescription9,omitempty"`
	ExteriorRepairPrice9        *float64            `bson:"exteriorRepairPrice9,omitempty"`
	ExteriorRepairDescription10 *string             `bson:"exteriorRepairDescription10,omitempty"`
	ExteriorRepairPrice10       *float64            `bson:"exteriorRepairPrice10,omitempty"`
	ExteriorRepairPriceTotal    *float64            `bson:"exteriorRepairPriceTotal,omitempty"`
	InteriorRepairDescription1  *string             `bson:"interiorRepairDescription1,omitempty"`
	InteriorRepairPrice1        *float64            `bson:"interiorRepairPrice1,omitempty"`
	InteriorRepairDescription2  *string             `bson:"interiorRepairDescription2,omitempty"`
	InteriorRepairPrice2        *float64            `bson:"interiorRepairPrice2,omitempty"`
	InteriorRepairDescription3  *string             `bson:"interiorRepairDescription3,omitempty"`
	InteriorRepairPrice3        *float64            `bson:"interiorRepairPrice3,omitempty"`
	InteriorRepairDescription4  *string             `bson:"interiorRepairDescription4,omitempty"`
	InteriorRepairPrice4        *float64            `bson:"interiorRepairPrice4,omitempty"`
	InteriorRepairDescription5  *string             `bson:"interiorRepairDescription5,omitempty"`
	InteriorRepairPrice5        *float64            `bson:"interiorRepairPrice5,omitempty"`
	InteriorRepairDescription6  *string             `bson:"interiorRepairDescription6,omitempty"`
	InteriorRepairPrice6        *float64            `bson:"interiorRepairPrice6,omitempty"`
	InteriorRepairDescription7  *string             `bson:"interiorRepairDescription7,omitempty"`
	InteriorRepairPrice7        *float64            `bson:"interiorRepairPrice7,omitempty"`
	InteriorRepairDescription8  *string             `bson:"interiorRepairDescription8,omitempty"`
	InteriorRepairPrice8        *float64            `bson:"interiorRepairPrice8,omitempty"`
	InteriorRepairDescription9  *string             `bson:"interiorRepairDescription9,omitempty"`
	InteriorRepairPrice9        *float64            `bson:"interiorRepairPrice9,omitempty"`
	InteriorRepairDescription10 *string             `bson:"interiorRepairDescription10,omitempty"`
	InteriorRepairPrice10       *float64            `bson:"interiorRepairPrice10,omitempty"`
	InteriorRepairPriceTotal    *float64            `bson:"interiorRepairPriceTotal,omitempty"`
	CreatedDateTime             primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime              *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *PipelineRepair) ToModels() *models.PipelineRepair {
	return &models.PipelineRepair{
		ExteriorRepairDescription1:  u.ExteriorRepairDescription1,
		ExteriorRepairPrice1:        u.ExteriorRepairPrice1,
		ExteriorRepairDescription2:  u.ExteriorRepairDescription2,
		ExteriorRepairPrice2:        u.ExteriorRepairPrice2,
		ExteriorRepairDescription3:  u.ExteriorRepairDescription3,
		ExteriorRepairPrice3:        u.ExteriorRepairPrice3,
		ExteriorRepairDescription4:  u.ExteriorRepairDescription4,
		ExteriorRepairPrice4:        u.ExteriorRepairPrice4,
		ExteriorRepairDescription5:  u.ExteriorRepairDescription5,
		ExteriorRepairPrice5:        u.ExteriorRepairPrice5,
		ExteriorRepairDescription6:  u.ExteriorRepairDescription6,
		ExteriorRepairPrice6:        u.ExteriorRepairPrice6,
		ExteriorRepairDescription7:  u.ExteriorRepairDescription7,
		ExteriorRepairPrice7:        u.ExteriorRepairPrice7,
		ExteriorRepairDescription8:  u.ExteriorRepairDescription8,
		ExteriorRepairPrice8:        u.ExteriorRepairPrice8,
		ExteriorRepairDescription9:  u.ExteriorRepairDescription9,
		ExteriorRepairPrice9:        u.ExteriorRepairPrice9,
		ExteriorRepairDescription10: u.ExteriorRepairDescription10,
		ExteriorRepairPrice10:       u.ExteriorRepairPrice10,
		ExteriorRepairPriceTotal:    u.ExteriorRepairPriceTotal,
		InteriorRepairDescription1:  u.InteriorRepairDescription1,
		InteriorRepairPrice1:        u.InteriorRepairPrice1,
		InteriorRepairDescription2:  u.InteriorRepairDescription2,
		InteriorRepairPrice2:        u.InteriorRepairPrice2,
		InteriorRepairDescription3:  u.InteriorRepairDescription3,
		InteriorRepairPrice3:        u.InteriorRepairPrice3,
		InteriorRepairDescription4:  u.InteriorRepairDescription4,
		InteriorRepairPrice4:        u.InteriorRepairPrice4,
		InteriorRepairDescription5:  u.InteriorRepairDescription5,
		InteriorRepairPrice5:        u.InteriorRepairPrice5,
		InteriorRepairDescription6:  u.InteriorRepairDescription6,
		InteriorRepairPrice6:        u.InteriorRepairPrice6,
		InteriorRepairDescription7:  u.InteriorRepairDescription7,
		InteriorRepairPrice7:        u.InteriorRepairPrice7,
		InteriorRepairDescription8:  u.InteriorRepairDescription8,
		InteriorRepairPrice8:        u.InteriorRepairPrice8,
		InteriorRepairDescription9:  u.InteriorRepairDescription9,
		InteriorRepairPrice9:        u.InteriorRepairPrice9,
		InteriorRepairDescription10: u.InteriorRepairDescription10,
		InteriorRepairPrice10:       u.InteriorRepairPrice10,
		InteriorRepairPriceTotal:    u.InteriorRepairPriceTotal,
	}
}
