package ivision

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

// CreateTag invokes the ivision.CreateTag API synchronously
// api document: https://help.aliyun.com/api/ivision/createtag.html
func (client *Client) CreateTag(request *CreateTagRequest) (response *CreateTagResponse, err error) {
	response = CreateCreateTagResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTagWithChan invokes the ivision.CreateTag API asynchronously
// api document: https://help.aliyun.com/api/ivision/createtag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTagWithChan(request *CreateTagRequest) (<-chan *CreateTagResponse, <-chan error) {
	responseChan := make(chan *CreateTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTag(request)
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

// CreateTagWithCallback invokes the ivision.CreateTag API asynchronously
// api document: https://help.aliyun.com/api/ivision/createtag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTagWithCallback(request *CreateTagRequest, callback func(response *CreateTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTagResponse
		var err error
		defer close(result)
		response, err = client.CreateTag(request)
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

// CreateTagRequest is the request struct for api CreateTag
type CreateTagRequest struct {
	*requests.RpcRequest
	Description string           `position:"Query" name:"Description"`
	ProjectId   string           `position:"Query" name:"ProjectId"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	Name        string           `position:"Query" name:"Name"`
}

// CreateTagResponse is the response struct for api CreateTag
type CreateTagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Tag       Tag    `json:"Tag" xml:"Tag"`
}

// CreateCreateTagRequest creates a request to invoke CreateTag API
func CreateCreateTagRequest() (request *CreateTagRequest) {
	request = &CreateTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivision", "2019-03-08", "CreateTag", "ivision", "openAPI")
	return
}

// CreateCreateTagResponse creates a response to parse from CreateTag response
func CreateCreateTagResponse() (response *CreateTagResponse) {
	response = &CreateTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}