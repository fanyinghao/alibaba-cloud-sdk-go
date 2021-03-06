package scdn

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

// DescribeScdnDomainPvData invokes the scdn.DescribeScdnDomainPvData API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainpvdata.html
func (client *Client) DescribeScdnDomainPvData(request *DescribeScdnDomainPvDataRequest) (response *DescribeScdnDomainPvDataResponse, err error) {
	response = CreateDescribeScdnDomainPvDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainPvDataWithChan invokes the scdn.DescribeScdnDomainPvData API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainpvdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainPvDataWithChan(request *DescribeScdnDomainPvDataRequest) (<-chan *DescribeScdnDomainPvDataResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainPvDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainPvData(request)
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

// DescribeScdnDomainPvDataWithCallback invokes the scdn.DescribeScdnDomainPvData API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainpvdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainPvDataWithCallback(request *DescribeScdnDomainPvDataRequest, callback func(response *DescribeScdnDomainPvDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainPvDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainPvData(request)
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

// DescribeScdnDomainPvDataRequest is the request struct for api DescribeScdnDomainPvData
type DescribeScdnDomainPvDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnDomainPvDataResponse is the response struct for api DescribeScdnDomainPvData
type DescribeScdnDomainPvDataResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	DomainName     string         `json:"DomainName" xml:"DomainName"`
	StartTime      string         `json:"StartTime" xml:"StartTime"`
	PvDataInterval PvDataInterval `json:"PvDataInterval" xml:"PvDataInterval"`
}

// CreateDescribeScdnDomainPvDataRequest creates a request to invoke DescribeScdnDomainPvData API
func CreateDescribeScdnDomainPvDataRequest() (request *DescribeScdnDomainPvDataRequest) {
	request = &DescribeScdnDomainPvDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainPvData", "", "")
	return
}

// CreateDescribeScdnDomainPvDataResponse creates a response to parse from DescribeScdnDomainPvData response
func CreateDescribeScdnDomainPvDataResponse() (response *DescribeScdnDomainPvDataResponse) {
	response = &DescribeScdnDomainPvDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
