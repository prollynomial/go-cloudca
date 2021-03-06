package cloudca

import (
	"github.com/cloud-ca/go-cloudca/mocks"
	"github.com/cloud-ca/go-cloudca/mocks/services_mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

const (
	TEST_COMPUTE_OFFERING_ID         = "some_id"
	TEST_COMPUTE_OFFERING_NAME       = "test_compute_offering"
	TEST_COMPUTE_OFFERING_MEMORY     = 4
	TEST_COMPUTE_OFFERING_CPU_NUMBER = 2
)

func buildComputeOfferingJsonResponse(computeOffering *ComputeOffering) []byte {
	return []byte(`{"id": "` + computeOffering.Id +
		`","name":"` + computeOffering.Name +
		`","memory":` + strconv.Itoa(computeOffering.Memory) +
		`,"cpuNumber": ` + strconv.Itoa(computeOffering.CpuNumber) + `}`)
}

func buildListComputeOfferingJsonResponse(computeOfferings []ComputeOffering) []byte {
	resp := `[`
	for i, d := range computeOfferings {
		resp += string(buildComputeOfferingJsonResponse(&d))
		if i != len(computeOfferings)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetComputeOfferingReturnComputeOfferingIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	computeOfferingService := ComputeOfferingApi{
		entityService: mockEntityService,
	}

	expectedComputeOffering := ComputeOffering{Id: TEST_COMPUTE_OFFERING_ID,
		Name:      TEST_COMPUTE_OFFERING_NAME,
		Memory:    TEST_COMPUTE_OFFERING_MEMORY,
		CpuNumber: TEST_COMPUTE_OFFERING_CPU_NUMBER}

	mockEntityService.EXPECT().Get(TEST_COMPUTE_OFFERING_ID, gomock.Any()).Return(buildComputeOfferingJsonResponse(&expectedComputeOffering), nil)

	//when
	computeOffering, _ := computeOfferingService.Get(TEST_COMPUTE_OFFERING_ID)

	//then
	if assert.NotNil(t, computeOffering) {
		assert.Equal(t, expectedComputeOffering, *computeOffering)
	}
}

func TestGetComputeOfferingReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	computeOfferingService := ComputeOfferingApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_COMPUTE_OFFERING_ID, gomock.Any()).Return(nil, mockError)

	//when
	computeOffering, err := computeOfferingService.Get(TEST_COMPUTE_OFFERING_ID)

	//then
	assert.Nil(t, computeOffering)
	assert.Equal(t, mockError, err)

}

func TestListComputeOfferingReturnComputeOfferingsIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	computeOfferingService := ComputeOfferingApi{
		entityService: mockEntityService,
	}

	expectedComputeOfferings := []ComputeOffering{
		ComputeOffering{
			Id:        "list_id_1",
			Name:      "list_name_1",
			Memory:    1,
			CpuNumber: 1,
		},
		ComputeOffering{
			Id:        "list_id_2",
			Name:      "list_name_2",
			Memory:    2,
			CpuNumber: 2,
		},
	}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListComputeOfferingJsonResponse(expectedComputeOfferings), nil)

	//when
	computeOfferings, _ := computeOfferingService.List()

	//then
	if assert.NotNil(t, computeOfferings) {
		assert.Equal(t, expectedComputeOfferings, computeOfferings)
	}
}

func TestListComputeOfferingReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	computeOfferingService := ComputeOfferingApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	computeOfferings, err := computeOfferingService.List()

	//then
	assert.Nil(t, computeOfferings)
	assert.Equal(t, mockError, err)

}
