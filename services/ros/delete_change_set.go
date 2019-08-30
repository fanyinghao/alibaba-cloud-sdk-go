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

// DeleteChangeSet invokes the ros.DeleteChangeSet API synchronously
// api document: https://help.aliyun.com/api/ros/deletechangeset.html
func (client *Client) DeleteChangeSet(request *DeleteChangeSetRequest) (response *DeleteChangeSetResponse, err error) {
	response = CreateDeleteChangeSetResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteChangeSetWithChan invokes the ros.DeleteChangeSet API asynchronously
// api document: https://help.aliyun.com/api/ros/deletechangeset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteChangeSetWithChan(request *DeleteChangeSetRequest) (<-chan *DeleteChangeSetResponse, <-chan error) {
	responseChan := make(chan *DeleteChangeSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteChangeSet(request)
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

// DeleteChangeSetWithCallback invokes the ros.DeleteChangeSet API asynchronously
// api document: https://help.aliyun.com/api/ros/deletechangeset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteChangeSetWithCallback(request *DeleteChangeSetRequest, callback func(response *DeleteChangeSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteChangeSetResponse
		var err error
		defer close(result)
		response, err = client.DeleteChangeSet(request)
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

// DeleteChangeSetRequest is the request struct for api DeleteChangeSet
type DeleteChangeSetRequest struct {
	*requests.RpcRequest
	ChangeSetId string `position:"Query" name:"ChangeSetId"`
}

// DeleteChangeSetResponse is the response struct for api DeleteChangeSet
type DeleteChangeSetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteChangeSetRequest creates a request to invoke DeleteChangeSet API
func CreateDeleteChangeSetRequest() (request *DeleteChangeSetRequest) {
	request = &DeleteChangeSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ROS", "2019-09-10", "DeleteChangeSet", "ROS", "openAPI")
	return
}

// CreateDeleteChangeSetResponse creates a response to parse from DeleteChangeSet response
func CreateDeleteChangeSetResponse() (response *DeleteChangeSetResponse) {
	response = &DeleteChangeSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
