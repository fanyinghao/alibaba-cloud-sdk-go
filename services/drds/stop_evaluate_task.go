package drds

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

// StopEvaluateTask invokes the drds.StopEvaluateTask API synchronously
// api document: https://help.aliyun.com/api/drds/stopevaluatetask.html
func (client *Client) StopEvaluateTask(request *StopEvaluateTaskRequest) (response *StopEvaluateTaskResponse, err error) {
	response = CreateStopEvaluateTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StopEvaluateTaskWithChan invokes the drds.StopEvaluateTask API asynchronously
// api document: https://help.aliyun.com/api/drds/stopevaluatetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopEvaluateTaskWithChan(request *StopEvaluateTaskRequest) (<-chan *StopEvaluateTaskResponse, <-chan error) {
	responseChan := make(chan *StopEvaluateTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopEvaluateTask(request)
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

// StopEvaluateTaskWithCallback invokes the drds.StopEvaluateTask API asynchronously
// api document: https://help.aliyun.com/api/drds/stopevaluatetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopEvaluateTaskWithCallback(request *StopEvaluateTaskRequest, callback func(response *StopEvaluateTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopEvaluateTaskResponse
		var err error
		defer close(result)
		response, err = client.StopEvaluateTask(request)
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

// StopEvaluateTaskRequest is the request struct for api StopEvaluateTask
type StopEvaluateTaskRequest struct {
	*requests.RpcRequest
	TaskId requests.Integer `position:"Query" name:"TaskId"`
}

// StopEvaluateTaskResponse is the response struct for api StopEvaluateTask
type StopEvaluateTaskResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Success          bool             `json:"Success" xml:"Success"`
	TaskManageResult TaskManageResult `json:"TaskManageResult" xml:"TaskManageResult"`
}

// CreateStopEvaluateTaskRequest creates a request to invoke StopEvaluateTask API
func CreateStopEvaluateTaskRequest() (request *StopEvaluateTaskRequest) {
	request = &StopEvaluateTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "StopEvaluateTask", "Drds", "openAPI")
	return
}

// CreateStopEvaluateTaskResponse creates a response to parse from StopEvaluateTask response
func CreateStopEvaluateTaskResponse() (response *StopEvaluateTaskResponse) {
	response = &StopEvaluateTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
