package smartag

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

// UnicomOrderConfirm invokes the smartag.UnicomOrderConfirm API synchronously
// api document: https://help.aliyun.com/api/smartag/unicomorderconfirm.html
func (client *Client) UnicomOrderConfirm(request *UnicomOrderConfirmRequest) (response *UnicomOrderConfirmResponse, err error) {
	response = CreateUnicomOrderConfirmResponse()
	err = client.DoAction(request, response)
	return
}

// UnicomOrderConfirmWithChan invokes the smartag.UnicomOrderConfirm API asynchronously
// api document: https://help.aliyun.com/api/smartag/unicomorderconfirm.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnicomOrderConfirmWithChan(request *UnicomOrderConfirmRequest) (<-chan *UnicomOrderConfirmResponse, <-chan error) {
	responseChan := make(chan *UnicomOrderConfirmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnicomOrderConfirm(request)
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

// UnicomOrderConfirmWithCallback invokes the smartag.UnicomOrderConfirm API asynchronously
// api document: https://help.aliyun.com/api/smartag/unicomorderconfirm.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnicomOrderConfirmWithCallback(request *UnicomOrderConfirmRequest, callback func(response *UnicomOrderConfirmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnicomOrderConfirmResponse
		var err error
		defer close(result)
		response, err = client.UnicomOrderConfirm(request)
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

// UnicomOrderConfirmRequest is the request struct for api UnicomOrderConfirm
type UnicomOrderConfirmRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer               `position:"Query" name:"ResourceOwnerId"`
	TmsCode              string                         `position:"Query" name:"TmsCode"`
	OrderItem            *[]UnicomOrderConfirmOrderItem `position:"Query" name:"OrderItem"  type:"Repeated"`
	OrderPostFee         requests.Integer               `position:"Query" name:"OrderPostFee"`
	TradeId              string                         `position:"Query" name:"TradeId"`
	OwnerUserId          string                         `position:"Query" name:"OwnerUserId"`
	ResourceOwnerAccount string                         `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                         `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer               `position:"Query" name:"OwnerId"`
	TmsOrderCode         string                         `position:"Query" name:"TmsOrderCode"`
}

// UnicomOrderConfirmOrderItem is a repeated param struct in UnicomOrderConfirmRequest
type UnicomOrderConfirmOrderItem struct {
	ScItemName   string    `name:"ScItemName"`
	ItemAmount   string    `name:"ItemAmount"`
	SnList       *[]string `name:"SnList" type:"Repeated"`
	OrderItemId  string    `name:"OrderItemId"`
	ScItemCode   string    `name:"ScItemCode"`
	ItemQuantity string    `name:"ItemQuantity"`
	TradeId      string    `name:"TradeId"`
	TradeItemId  string    `name:"TradeItemId"`
}

// UnicomOrderConfirmResponse is the response struct for api UnicomOrderConfirm
type UnicomOrderConfirmResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUnicomOrderConfirmRequest creates a request to invoke UnicomOrderConfirm API
func CreateUnicomOrderConfirmRequest() (request *UnicomOrderConfirmRequest) {
	request = &UnicomOrderConfirmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "UnicomOrderConfirm", "smartag", "openAPI")
	return
}

// CreateUnicomOrderConfirmResponse creates a response to parse from UnicomOrderConfirm response
func CreateUnicomOrderConfirmResponse() (response *UnicomOrderConfirmResponse) {
	response = &UnicomOrderConfirmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
