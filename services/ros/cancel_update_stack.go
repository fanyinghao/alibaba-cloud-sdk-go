package ros

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

// CancelUpdateStack invokes the ros.CancelUpdateStack API synchronously
// api document: https://help.aliyun.com/api/ros/cancelupdatestack.html
func (client *Client) CancelUpdateStack(request *CancelUpdateStackRequest) (response *CancelUpdateStackResponse, err error) {
	response = CreateCancelUpdateStackResponse()
	err = client.DoAction(request, response)
	return
}

// CancelUpdateStackWithChan invokes the ros.CancelUpdateStack API asynchronously
// api document: https://help.aliyun.com/api/ros/cancelupdatestack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelUpdateStackWithChan(request *CancelUpdateStackRequest) (<-chan *CancelUpdateStackResponse, <-chan error) {
	responseChan := make(chan *CancelUpdateStackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelUpdateStack(request)
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

// CancelUpdateStackWithCallback invokes the ros.CancelUpdateStack API asynchronously
// api document: https://help.aliyun.com/api/ros/cancelupdatestack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelUpdateStackWithCallback(request *CancelUpdateStackRequest, callback func(response *CancelUpdateStackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelUpdateStackResponse
		var err error
		defer close(result)
		response, err = client.CancelUpdateStack(request)
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

// CancelUpdateStackRequest is the request struct for api CancelUpdateStack
type CancelUpdateStackRequest struct {
	*requests.RpcRequest
	CancelType string `position:"Query" name:"CancelType"`
	StackId    string `position:"Query" name:"StackId"`
}

// CancelUpdateStackResponse is the response struct for api CancelUpdateStack
type CancelUpdateStackResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelUpdateStackRequest creates a request to invoke CancelUpdateStack API
func CreateCancelUpdateStackRequest() (request *CancelUpdateStackRequest) {
	request = &CancelUpdateStackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ROS", "2019-09-10", "CancelUpdateStack", "ROS", "openAPI")
	return
}

// CreateCancelUpdateStackResponse creates a response to parse from CancelUpdateStack response
func CreateCancelUpdateStackResponse() (response *CancelUpdateStackResponse) {
	response = &CancelUpdateStackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
