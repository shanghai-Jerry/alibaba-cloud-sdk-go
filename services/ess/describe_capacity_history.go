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

// DescribeCapacityHistory invokes the ess.DescribeCapacityHistory API synchronously
// api document: https://help.aliyun.com/api/ess/describecapacityhistory.html
func (client *Client) DescribeCapacityHistory(request *DescribeCapacityHistoryRequest) (response *DescribeCapacityHistoryResponse, err error) {
	response = CreateDescribeCapacityHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCapacityHistoryWithChan invokes the ess.DescribeCapacityHistory API asynchronously
// api document: https://help.aliyun.com/api/ess/describecapacityhistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCapacityHistoryWithChan(request *DescribeCapacityHistoryRequest) (<-chan *DescribeCapacityHistoryResponse, <-chan error) {
	responseChan := make(chan *DescribeCapacityHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCapacityHistory(request)
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

// DescribeCapacityHistoryWithCallback invokes the ess.DescribeCapacityHistory API asynchronously
// api document: https://help.aliyun.com/api/ess/describecapacityhistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCapacityHistoryWithCallback(request *DescribeCapacityHistoryRequest, callback func(response *DescribeCapacityHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCapacityHistoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeCapacityHistory(request)
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

// DescribeCapacityHistoryRequest is the request struct for api DescribeCapacityHistory
type DescribeCapacityHistoryRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	EndTime              string           `position:"Query" name:"EndTime"`
	StartTime            string           `position:"Query" name:"StartTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeCapacityHistoryResponse is the response struct for api DescribeCapacityHistory
type DescribeCapacityHistoryResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	TotalCount           int                  `json:"TotalCount" xml:"TotalCount"`
	PageNumber           int                  `json:"PageNumber" xml:"PageNumber"`
	PageSize             int                  `json:"PageSize" xml:"PageSize"`
	CapacityHistoryItems CapacityHistoryItems `json:"CapacityHistoryItems" xml:"CapacityHistoryItems"`
}

// CreateDescribeCapacityHistoryRequest creates a request to invoke DescribeCapacityHistory API
func CreateDescribeCapacityHistoryRequest() (request *DescribeCapacityHistoryRequest) {
	request = &DescribeCapacityHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeCapacityHistory", "ESS", "openAPI")
	return
}

// CreateDescribeCapacityHistoryResponse creates a response to parse from DescribeCapacityHistory response
func CreateDescribeCapacityHistoryResponse() (response *DescribeCapacityHistoryResponse) {
	response = &DescribeCapacityHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
