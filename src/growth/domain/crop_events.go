package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CropBatchCreated struct {
	UID             uuid.UUID
	BatchID         string
	Status          CropStatus
	Type            CropType
	Container       CropContainer
	ContainerType   string
	ContainerCell   int
	InventoryUID    uuid.UUID
	VarietyName     string
	PlantType       string
	FarmUID         uuid.UUID
	CreatedDate     time.Time
	InitialAreaUID  uuid.UUID
	InitialAreaName string
	Quantity        int
}

type CropBatchMoved struct {
	UID           uuid.UUID
	BatchID       string
	ContainerType string
	Quantity      int
	SrcAreaUID    uuid.UUID
	SrcAreaName   string
	DstAreaUID    uuid.UUID
	DstAreaName   string
	MovedDate     time.Time

	InitialArea InitialArea
	MovedArea   []MovedArea
}

type CropBatchWatered struct {
	UID          uuid.UUID
	AreaUID      uuid.UUID
	AreaName     string
	WateringDate time.Time

	InitialArea InitialArea
	MovedArea   []MovedArea
}
