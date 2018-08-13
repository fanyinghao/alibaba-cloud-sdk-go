package linkface

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

// Data is a nested struct in linkface response
type Data struct {
	Rate             float64                `json:"Rate" xml:"Rate"`
	Groups           []string               `json:"Groups" xml:"Groups"`
	GroupIds         []string               `json:"GroupIds" xml:"GroupIds"`
	CurrentFaceInfos []CurrentFaceInfosItem `json:"CurrentFaceInfos" xml:"CurrentFaceInfos"`
	FailedFaceInfos  []FailedFaceInfosItem  `json:"FailedFaceInfos" xml:"FailedFaceInfos"`
	UserFaceMetas    []UserFaceMetasItem    `json:"UserFaceMetas" xml:"UserFaceMetas"`
	TopUserItem      []TopUserItemItem      `json:"TopUserItem" xml:"TopUserItem"`
}
