package rdc

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

// CreateEnterprise invokes the rdc.CreateEnterprise API synchronously
// api document: https://help.aliyun.com/api/rdc/createenterprise.html
func (client *Client) CreateEnterprise(request *CreateEnterpriseRequest) (response *CreateEnterpriseResponse, err error) {
	response = CreateCreateEnterpriseResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEnterpriseWithChan invokes the rdc.CreateEnterprise API asynchronously
// api document: https://help.aliyun.com/api/rdc/createenterprise.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateEnterpriseWithChan(request *CreateEnterpriseRequest) (<-chan *CreateEnterpriseResponse, <-chan error) {
	responseChan := make(chan *CreateEnterpriseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEnterprise(request)
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

// CreateEnterpriseWithCallback invokes the rdc.CreateEnterprise API asynchronously
// api document: https://help.aliyun.com/api/rdc/createenterprise.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateEnterpriseWithCallback(request *CreateEnterpriseRequest, callback func(response *CreateEnterpriseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEnterpriseResponse
		var err error
		defer close(result)
		response, err = client.CreateEnterprise(request)
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

// CreateEnterpriseRequest is the request struct for api CreateEnterprise
type CreateEnterpriseRequest struct {
	*requests.RpcRequest
	Emails         string `position:"Query" name:"Emails"`
	CreatorStaffId string `position:"Query" name:"CreatorStaffId"`
	Domain         string `position:"Query" name:"Domain"`
	Name           string `position:"Query" name:"Name"`
	Description    string `position:"Query" name:"Description"`
}

// CreateEnterpriseResponse is the response struct for api CreateEnterprise
type CreateEnterpriseResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateEnterpriseRequest creates a request to invoke CreateEnterprise API
func CreateCreateEnterpriseRequest() (request *CreateEnterpriseRequest) {
	request = &CreateEnterpriseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rdc", "2018-08-21", "CreateEnterprise", "rdc", "openAPI")
	return
}

// CreateCreateEnterpriseResponse creates a response to parse from CreateEnterprise response
func CreateCreateEnterpriseResponse() (response *CreateEnterpriseResponse) {
	response = &CreateEnterpriseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}