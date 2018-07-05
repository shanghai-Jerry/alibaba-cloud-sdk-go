package ess

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyNotificationConfiguration invokes the ess.ModifyNotificationConfiguration API synchronously
// api document: https://help.aliyun.com/api/ess/modifynotificationconfiguration.html
func (client *Client) ModifyNotificationConfiguration(request *ModifyNotificationConfigurationRequest) (response *ModifyNotificationConfigurationResponse, err error) {
	response = CreateModifyNotificationConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyNotificationConfigurationWithChan invokes the ess.ModifyNotificationConfiguration API asynchronously
// api document: https://help.aliyun.com/api/ess/modifynotificationconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyNotificationConfigurationWithChan(request *ModifyNotificationConfigurationRequest) (<-chan *ModifyNotificationConfigurationResponse, <-chan error) {
	responseChan := make(chan *ModifyNotificationConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyNotificationConfiguration(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyNotificationConfigurationWithCallback invokes the ess.ModifyNotificationConfiguration API asynchronously
// api document: https://help.aliyun.com/api/ess/modifynotificationconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyNotificationConfigurationWithCallback(request *ModifyNotificationConfigurationRequest, callback func(response *ModifyNotificationConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyNotificationConfigurationResponse
		var err error
		defer close(result)
		response, err = client.ModifyNotificationConfiguration(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyNotificationConfigurationRequest is the request struct for api ModifyNotificationConfiguration
type ModifyNotificationConfigurationRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	NotificationArn      string           `position:"Query" name:"NotificationArn"`
	NotificationType     *[]string        `position:"Query" name:"NotificationType"  type:"Repeated"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyNotificationConfigurationResponse is the response struct for api ModifyNotificationConfiguration
type ModifyNotificationConfigurationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyNotificationConfigurationRequest creates a request to invoke ModifyNotificationConfiguration API
func CreateModifyNotificationConfigurationRequest() (request *ModifyNotificationConfigurationRequest) {
	request = &ModifyNotificationConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ModifyNotificationConfiguration", "ESS", "openAPI")
	return
}

// CreateModifyNotificationConfigurationResponse creates a response to parse from ModifyNotificationConfiguration response
func CreateModifyNotificationConfigurationResponse() (response *ModifyNotificationConfigurationResponse) {
	response = &ModifyNotificationConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
