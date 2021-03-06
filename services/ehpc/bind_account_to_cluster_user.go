package ehpc

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

// BindAccountToClusterUser invokes the ehpc.BindAccountToClusterUser API synchronously
// api document: https://help.aliyun.com/api/ehpc/bindaccounttoclusteruser.html
func (client *Client) BindAccountToClusterUser(request *BindAccountToClusterUserRequest) (response *BindAccountToClusterUserResponse, err error) {
	response = CreateBindAccountToClusterUserResponse()
	err = client.DoAction(request, response)
	return
}

// BindAccountToClusterUserWithChan invokes the ehpc.BindAccountToClusterUser API asynchronously
// api document: https://help.aliyun.com/api/ehpc/bindaccounttoclusteruser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindAccountToClusterUserWithChan(request *BindAccountToClusterUserRequest) (<-chan *BindAccountToClusterUserResponse, <-chan error) {
	responseChan := make(chan *BindAccountToClusterUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindAccountToClusterUser(request)
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

// BindAccountToClusterUserWithCallback invokes the ehpc.BindAccountToClusterUser API asynchronously
// api document: https://help.aliyun.com/api/ehpc/bindaccounttoclusteruser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindAccountToClusterUserWithCallback(request *BindAccountToClusterUserRequest, callback func(response *BindAccountToClusterUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindAccountToClusterUserResponse
		var err error
		defer close(result)
		response, err = client.BindAccountToClusterUser(request)
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

// BindAccountToClusterUserRequest is the request struct for api BindAccountToClusterUser
type BindAccountToClusterUserRequest struct {
	*requests.RpcRequest
	UserPwd     string `position:"Query" name:"UserPwd"`
	ClusterId   string `position:"Query" name:"ClusterId"`
	AccountName string `position:"Query" name:"AccountName"`
	AccountUid  string `position:"Query" name:"AccountUid"`
	UserName    string `position:"Query" name:"UserName"`
}

// BindAccountToClusterUserResponse is the response struct for api BindAccountToClusterUser
type BindAccountToClusterUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBindAccountToClusterUserRequest creates a request to invoke BindAccountToClusterUser API
func CreateBindAccountToClusterUserRequest() (request *BindAccountToClusterUserRequest) {
	request = &BindAccountToClusterUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "BindAccountToClusterUser", "ehs", "openAPI")
	return
}

// CreateBindAccountToClusterUserResponse creates a response to parse from BindAccountToClusterUser response
func CreateBindAccountToClusterUserResponse() (response *BindAccountToClusterUserResponse) {
	response = &BindAccountToClusterUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
