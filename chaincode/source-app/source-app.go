package main

import (
	"encoding/json"
    "fmt"
  //  "strconv"
  //  "strings"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

/* 
Function annotation:
1. json
    1.1 func Marshal(v interface{}) ([]byte, error)
        Marshal returns the JSON encoding of v.
    1.2 func Unmarshal(data []byte, v interface{}) error
        Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
2. shim
    2.1 func PutState(key string, value []byte) error
        func PutState puts the specified `key` and `value` into the transaction's writeset as a data-write proposal. PutState doesn't effect the ledger until the transaction is validated and successfully committed. Simple keys must not be an empty string and must not start with null character (0x00), in order to avoid range query collisions with composite keys, which internally get prefixed with 0x00 as composite key namespace.
    2.2 GetHistoryForKey(key string) (HistoryQueryIteratorInterface, error)
        GetHistoryForKey returns a history of key values across time. For each historic key update, the historic value and associated transaction id and timestamp are returned. The timestamp is the timestamp provided by the client in the proposal header. GetHistoryForKey requires peer configuration core.ledger.history.enableHistoryDatabase to be true. The query is NOT re-executed during validation phase, phantom reads are not detected. That is, other committed transactions may have updated the key concurrently, impacting the result set, and this would not be detected at validation/commit time. Applications susceptible to this should therefore not use GetHistoryForKey as part of transactions that update ledger, and should limit use to read-only chaincode operations.
*/
type ExpressChainCode struct{
}

//快递信息
type ExpressInfo struct{
    CoName string `json:CoName`                              //物流公司名称
    CoInfo string `json:CoInfo`                              //物流公司信息
    DeliverTime string `json:DeliverTime`                    //发件日期
    EstimatedDeliveryTime string `json:EstimatedDeliveryTime`//预计配送时间
    BatchNum string `json:BatchNum`                          //快递批次号
    Weight string `json:Weight`                              //快递重量
    Price string `json:Price`                                //快递费用
    Deliverer string `json:Deliverer`                        //投件人
    DelivererAdd string `json:DelivererAdd`                  //投件人所在地
}

//中转信息
type TransferInfo struct{
    ArrTime string `json:ArrTime`                       //到达时间
    TranferStationAdd string `json:TranferStationAdd`   //中转站地址
    HandlerInfo string `json:HandlerInfo`               //处理人员信息
    PkgStatus string `json:PkgStatus`                   //快递包状态
    DepartureTime string `json:DepartureTime`           //出发时间
    NextDestAdd string `json:NextDestAdd`               //下一个站点地址
    Mission string `json:Mission`                       //转运/派送
    VehicleType string `json:VehicleType`               //运送方式
    VehicleInfo string `json:VehicleInfo`               //交通工具信息
    DriverInfo string `json:DriverInfo`                 //司机信息
}
//初始化
func (a *ExpressChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}
//提供给外部的调用
func (a *ExpressChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    fn,args := stub.GetFunctionAndParameters()
    if fn == "addExpressInfo"{
        return a.addExpressInfo(stub,args)
    }else if fn == "addTransferInfo"{
        return a.addTransferInfo(stub,args)
    }else if fn == "getExpressInfo"{
        return a.getExpressInfo(stub,args)
    }else if fn == "getTransferInfo"{
        return a.getTransferInfo(stub,args)
    }

    return shim.Error("Recevied unkown function invocation")
}
//添加快递信息
func (a *ExpressChainCode) addExpressInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    if len(args)!=10{
        return shim.Error("Incorrect number of arguments.")
    }
    ExpressID := args[0]
    if ExpressID == ""{
        return shim.Error("ExpressID can not be empty.")
    }
    
    var expressInfos ExpressInfo
    expressInfos.CoName = args[1]
    expressInfos.CoInfo = args[2]
    expressInfos.DeliverTime = args[3]
    expressInfos.EstimatedDeliveryTime = args[4]
    expressInfos.BatchNum = args[5]
    expressInfos.Weight = args[6]
    expressInfos.Price = args[7]
    expressInfos.Deliverer = args[8]
    expressInfos.DelivererAdd = args[9]
    var err error 
    ExpressInfoJSONasBytes,err := json.Marshal(expressInfos)//将快递信息转为JSON格式
    if err != nil{
        return shim.Error(err.Error())
    }

    err = stub.PutState(ExpressID,ExpressInfoJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }

    return shim.Success(nil)
}

//添加中转信息
func(a *ExpressChainCode) addTransferInfo(stub shim.ChaincodeStubInterface,args []string) pb.Response{
    if len(args)!=11{
        return shim.Error("Incorrect number of arguments.")
    }
    ExpressID := args[0]
    if ExpressID == ""{
        return shim.Error("ExpressID can not be empty.")
    }

    var transferInfo TransferInfo
    transferInfo.ArrTime = args[1]
    transferInfo.TranferStationAdd = args[2]
    transferInfo.HandlerInfo = args[3]
    transferInfo.PkgStatus = args[4]
    transferInfo.DepartureTime = args[5]
    transferInfo.NextDestAdd = args[6]
    transferInfo.Mission = args[7]
    transferInfo.VehicleType = args[8]
    transferInfo.VehicleInfo = args[9]
    transferInfo.DriverInfo = args[10]
    
    var err error
    TransferInfoJSONasBytes,err := json.Marshal(transferInfo)
    if err != nil{
        return shim.Error(err.Error())
    } 
    err = stub.PutState(ExpressID,TransferInfoJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(nil)ExpressInfos
}

//获取快递信息
func(a *ExpressChainCode) getExpressInfo(stub shim.ChaincodeStubInterface,args []string) pb.Response{
    
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    ExpressID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(ExpressID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var expressInfo ExpressInfo

    for resultsIterator.HasNext(){
        var expressInfoGeted ExpressInfo
        response,err :=resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&expressInfoGeted)
        if expressInfoGeted.CoName != ""{
            expressInfo = expressInfoGeted
            continue
        }
    }
    jsonsAsBytes,err := json.Marshal(expressInfo)   
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

//获取中转信息
func(a *ExpressChainCode) getTransferInfo(stub shim.ChaincodeStubInterface,args []string) pb.Response{
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    
    ExpressID := args[0]
    resultsIterator,err :=stub.GetHistoryForKey(ExpressID)
    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    
    var transferInfos []TransferInfo
    for resultsIterator.HasNext(){
        var ExpressInfos TransferInfo
        response,err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&ExpressInfos)
        if ExpressInfos.HandlerInfo != ""{
            transferInfos = append(transferInfos,ExpressInfos)
        }
    }
    jsonsAsBytes,err := json.Marshal(transferInfos)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

func main(){
     err := shim.Start(new(ExpressChainCode))
     if err != nil {
         fmt.Printf("Error starting Express chaincode: %s ",err)
     }
}
