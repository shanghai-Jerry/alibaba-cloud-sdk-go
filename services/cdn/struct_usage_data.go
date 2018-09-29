package cdn

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

// UsageData is a nested struct in cdn response
type UsageData struct {
	Path      string                                    `json:"Path" xml:"Path"`
	Time      string                                    `json:"Time" xml:"Time"`
	Traffic   int                                       `json:"Traffic" xml:"Traffic"`
	TimeStamp string                                    `json:"TimeStamp" xml:"TimeStamp"`
	Acc       int                                       `json:"Acc" xml:"Acc"`
	Value     ValueInDescribeDomainRealTimeHttpCodeData `json:"Value" xml:"Value"`
}
