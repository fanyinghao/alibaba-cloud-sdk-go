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

package dts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyMigrationObject invokes the dts.ModifyMigrationObject API synchronously
// api document: https://help.aliyun.com/api/dts/modifymigrationobject.html
func (client *Client) ModifyMigrationObject(request *ModifyMigrationObjectRequest) (response *ModifyMigrationObjectResponse, err error) {
	response = CreateModifyMigrationObjectResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMigrationObjectWithChan invokes the dts.ModifyMigrationObject API asynchronously
// api document: https://help.aliyun.com/api/dts/modifymigrationobject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMigrationObjectWithChan(request *ModifyMigrationObjectRequest) (<-chan *ModifyMigrationObjectResponse, <-chan error) {
	responseChan := make(chan *ModifyMigrationObjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMigrationObject(request)
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

// ModifyMigrationObjectWithCallback invokes the dts.ModifyMigrationObject API asynchronously
// api document: https://help.aliyun.com/api/dts/modifymigrationobject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMigrationObjectWithCallback(request *ModifyMigrationObjectRequest, callback func(response *ModifyMigrationObjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMigrationObjectResponse
		var err error
		defer close(result)
		response, err = client.ModifyMigrationObject(request)
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

// ModifyMigrationObjectRequest is the request struct for api ModifyMigrationObject
type ModifyMigrationObjectRequest struct {
	*requests.RpcRequest
	MigrationJobId  string `position:"Query" name:"MigrationJobId"`
	MigrationObject string `position:"Query" name:"MigrationObject"`
	ClientToken     string `position:"Query" name:"ClientToken"`
	OwnerId         string `position:"Query" name:"OwnerId"`
}

// ModifyMigrationObjectResponse is the response struct for api ModifyMigrationObject
type ModifyMigrationObjectResponse struct {
	*responses.BaseResponse
	Success    string `json:"Success" xml:"Success"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyMigrationObjectRequest creates a request to invoke ModifyMigrationObject API
func CreateModifyMigrationObjectRequest() (request *ModifyMigrationObjectRequest) {
	request = &ModifyMigrationObjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "ModifyMigrationObject", "dts", "openAPI")
	return
}

// CreateModifyMigrationObjectResponse creates a response to parse from ModifyMigrationObject response
func CreateModifyMigrationObjectResponse() (response *ModifyMigrationObjectResponse) {
	response = &ModifyMigrationObjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}