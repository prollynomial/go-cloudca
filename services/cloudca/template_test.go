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
	TEST_TEMPLATE_ID              = "test_template_id"
	TEST_TEMPLATE_NAME            = "test_template"
	TEST_TEMPLATE_DESCRIPTION     = "test_template_description"
	TEST_TEMPLATE_SIZE            = 60
	TEST_TEMPLATE_IS_PUBLIC       = true
	TEST_TEMPLATE_IS_READY        = true
	TEST_TEMPLATE_SSH_KEY_ENABLED = false
	TEST_TEMPLATE_EXTRACTABLE     = true
	TEST_TEMPLATE_OS_TYPE         = "test_template_os_type"
	TEST_TEMPLATE_OS_TYPE_ID      = "test_template_os_type_id"
	TEST_TEMPLATE_HYPERVISOR      = "test_template_hypervisor"
	TEST_TEMPLATE_FORMAT          = "test_template_format"
	TEST_TEMPLATE_ZONE_NAME       = "test_template_zone_name"
	TEST_TEMPLATE_PROJECT_ID      = "test_template_project_id"
)

func buildTemplateJsonResponse(template *Template) []byte {
	return []byte(`{"id":"` + template.Id + `",` +
		` "name": "` + template.Name + `",` +
		` "description": "` + template.Description + `",` +
		` "size": ` + strconv.Itoa(template.Size) + `,` +
		` "isPublic": ` + strconv.FormatBool(template.IsPublic) + `,` +
		` "isReady": ` + strconv.FormatBool(template.IsReady) + `,` +
		` "sshKeyEnabled": ` + strconv.FormatBool(template.SSHKeyEnabled) + `,` +
		` "extractable": ` + strconv.FormatBool(template.Extractable) + `,` +
		` "osType": "` + template.OSType + `",` +
		` "osTypeId": "` + template.OSTypeId + `",` +
		` "hypervisor": "` + template.Hypervisor + `",` +
		` "format": "` + template.Format + `",` +
		` "zoneName": "` + template.ZoneName + `",` +
		` "projectId": "` + template.ProjectId + `"}`)
}

func buildListTemplateJsonResponse(templates []Template) []byte {
	resp := `[`
	for i, t := range templates {
		resp += string(buildTemplateJsonResponse(&t))
		if i != len(templates)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetTemplateReturnTemplateIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	templateService := TemplateApi{
		entityService: mockEntityService,
	}

	expectedTemplate := Template{Id: TEST_TEMPLATE_ID,
		Name:          TEST_TEMPLATE_NAME,
		Description:   TEST_TEMPLATE_DESCRIPTION,
		Size:          TEST_TEMPLATE_SIZE,
		IsPublic:      TEST_TEMPLATE_IS_PUBLIC,
		IsReady:       TEST_TEMPLATE_IS_READY,
		SSHKeyEnabled: TEST_TEMPLATE_SSH_KEY_ENABLED,
		Extractable:   TEST_TEMPLATE_EXTRACTABLE,
		OSType:        TEST_TEMPLATE_OS_TYPE,
		OSTypeId:      TEST_TEMPLATE_OS_TYPE_ID,
		Hypervisor:    TEST_TEMPLATE_HYPERVISOR,
		Format:        TEST_TEMPLATE_FORMAT,
		ZoneName:      TEST_TEMPLATE_ZONE_NAME,
		ProjectId:     TEST_TEMPLATE_PROJECT_ID,
	}

	mockEntityService.EXPECT().Get(TEST_TEMPLATE_ID, gomock.Any()).Return(buildTemplateJsonResponse(&expectedTemplate), nil)

	//when
	template, _ := templateService.Get(TEST_TEMPLATE_ID)

	//then
	if assert.NotNil(t, template) {
		assert.Equal(t, expectedTemplate, *template)
	}
}

func TestGetTemplateReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	templateService := TemplateApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_TEMPLATE_ID, gomock.Any()).Return(nil, mockError)

	//when
	template, err := templateService.Get(TEST_TEMPLATE_ID)

	//then
	assert.Nil(t, template)
	assert.Equal(t, mockError, err)

}

func TestListTemplateReturnTemplatesIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	templateService := TemplateApi{
		entityService: mockEntityService,
	}

	expectedTemplate1 := Template{Id: "list_id_1",
		Name:          "list_name_1",
		Description:   "list_description_1",
		Size:          132123,
		IsPublic:      true,
		IsReady:       true,
		SSHKeyEnabled: true,
		Extractable:   true,
		OSType:        "list_os_type_1",
		OSTypeId:      "list_os_type_id_1",
		Hypervisor:    "list_hypervisor_1",
		Format:        "list_format_1",
		ZoneName:      "list_zone_name_1",
		ProjectId:     "list_project_id_1",
	}
	expectedTemplate2 := Template{Id: "list_id_2",
		Name:          "list_name_2",
		Description:   "list_description_2",
		Size:          4525,
		IsPublic:      false,
		IsReady:       false,
		SSHKeyEnabled: false,
		Extractable:   false,
		OSType:        "list_os_type_2",
		OSTypeId:      "list_os_type_id_2",
		Hypervisor:    "list_hypervisor_2",
		Format:        "list_format_2",
		ZoneName:      "list_zone_name_2",
		ProjectId:     "list_project_id_2",
	}

	expectedTemplates := []Template{expectedTemplate1, expectedTemplate2}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTemplateJsonResponse(expectedTemplates), nil)

	//when
	templates, _ := templateService.List()

	//then
	if assert.NotNil(t, templates) {
		assert.Equal(t, expectedTemplates, templates)
	}
}

func TestListTemplateReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	templateService := TemplateApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	templates, err := templateService.List()

	//then
	assert.Nil(t, templates)
	assert.Equal(t, mockError, err)

}
