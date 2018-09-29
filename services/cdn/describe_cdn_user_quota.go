package cdn

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

// DescribeCdnUserQuota invokes the cdn.DescribeCdnUserQuota API synchronously
// api document: https://help.aliyun.com/api/cdn/describecdnuserquota.html
func (client *Client) DescribeCdnUserQuota(request *DescribeCdnUserQuotaRequest) (response *DescribeCdnUserQuotaResponse, err error) {
	response = CreateDescribeCdnUserQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnUserQuotaWithChan invokes the cdn.DescribeCdnUserQuota API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecdnuserquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnUserQuotaWithChan(request *DescribeCdnUserQuotaRequest) (<-chan *DescribeCdnUserQuotaResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnUserQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnUserQuota(request)
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

// DescribeCdnUserQuotaWithCallback invokes the cdn.DescribeCdnUserQuota API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecdnuserquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnUserQuotaWithCallback(request *DescribeCdnUserQuotaRequest, callback func(response *DescribeCdnUserQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnUserQuotaResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnUserQuota(request)
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

// DescribeCdnUserQuotaRequest is the request struct for api DescribeCdnUserQuota
type DescribeCdnUserQuotaRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCdnUserQuotaResponse is the response struct for api DescribeCdnUserQuota
type DescribeCdnUserQuotaResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DomainQuota      int    `json:"DomainQuota" xml:"DomainQuota"`
	RefreshUrlQuota  int    `json:"RefreshUrlQuota" xml:"RefreshUrlQuota"`
	RefreshDirQuota  int    `json:"RefreshDirQuota" xml:"RefreshDirQuota"`
	RefreshUrlRemain int    `json:"RefreshUrlRemain" xml:"RefreshUrlRemain"`
	RefreshDirRemain int    `json:"RefreshDirRemain" xml:"RefreshDirRemain"`
	PreloadQuota     int    `json:"PreloadQuota" xml:"PreloadQuota"`
	PreloadRemain    int    `json:"PreloadRemain" xml:"PreloadRemain"`
	BlockQuota       int    `json:"BlockQuota" xml:"BlockQuota"`
	BlockRemain      int    `json:"BlockRemain" xml:"BlockRemain"`
}

// CreateDescribeCdnUserQuotaRequest creates a request to invoke DescribeCdnUserQuota API
func CreateDescribeCdnUserQuotaRequest() (request *DescribeCdnUserQuotaRequest) {
	request = &DescribeCdnUserQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnUserQuota", "", "")
	return
}

// CreateDescribeCdnUserQuotaResponse creates a response to parse from DescribeCdnUserQuota response
func CreateDescribeCdnUserQuotaResponse() (response *DescribeCdnUserQuotaResponse) {
	response = &DescribeCdnUserQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
