package main

import (
	"encoding/json"
    "fmt"
  //  "strconv"
  //  "strings"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)
type ExpressChainCode struct{	
}

//express数据结构体
type ExpressInfo struct{
    ExpressID string `json:ExpressID`                             //食品ID
    ExpressProInfo ProInfo `json:ExpressProInfo`                  //生产信息
    //ExpressIngInfo []IngInfo `json:ExpressIngInfo`                  //配料信息
    ExpressLogInfo LogInfo `json:ExpressLogInfo`                  //物流信息
}

type ExpressAllInfo struct{
    ExpressID string `json:ExpressId`
    ExpressProInfo ProInfo `json:ExpressProInfo`
    //ExpressIngInfo []IngInfo `json:ExpressIngInfo`
    ExpressLogInfo []LogInfo `json:ExpressLogInfo`
}

//生产信息
type ProInfo struct{
    CoName string `json:CoName`                         //食品名称
    CoInfo string `json:CoInfo`                         //食品规格
    DeliverTime string `json:DeliverTime`                   //食品出产日期
    EstimatedDeliveryTime string `json:EstimatedDeliveryTime`                   //食品保质期
    BatchNum string `json:BatchNum`                           //食品批次号
    Weight string `json:Weight`                         //食品生产许可证编号
    Price string `json:Price`                 //食品生产商名称
    Deliverer string `json:Deliverer`                 //食品生产价格
    DelivererAdd string `json:DelivererAdd`                 //食品生产所在地
}
/* type IngInfo struct{
    IngID string `json:IngID`                               //配料ID
    IngName string `json:IngName`                           //配料名称
} */

type LogInfo struct{
    ArrTime string `json:ArrTime`             //出发时间
    TranferStationAdd string `json:TranferStationAdd`                 //到达时间
    HandlerInfo string `json:HandlerInfo`                     //处理业务(储存or运输)
    PkgStatus string `json:PkgStatus`             //出发地
    DepartureTime string `json:DepartureTime`                           //目的地
    NextDestAdd string `json:NextDestAdd`                   //销售商
    Mission string `json:Mission`                 //存储时间
    VehicleType string `json:VehicleType`                             //运送方式
    VehicleInfo string `json:VehicleInfo`                     //物流公司名称
    DriverInfo string `json:DriverInfo`                           //费用
}

func (a *ExpressChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

func (a *ExpressChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    fn,args := stub.GetFunctionAndParameters()
    if fn == "addProInfo"{
        return a.addProInfo(stub,args)
    }else if fn == "getExpressInfo"{
        return a.getExpressInfo(stub,args)
    }else if fn == "addLogInfo"{
        return a.addLogInfo(stub,args)
    }else if fn == "getProInfo"{
        return a.getProInfo(stub,args)
    }else if fn == "getLogInfo"{
        return a.getLogInfo(stub,args)
    }else if fn == "getLogInfo_l"{
        return a.getLogInfo_l(stub,args)
    }

    return shim.Error("Recevied unkown function invocation")
}

func (a *ExpressChainCode) addProInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error 
    var ExpressInfos ExpressInfo

    if len(args)!=10{
        return shim.Error("Incorrect number of arguments.")
    }
    ExpressInfos.ExpressID = args[0]
    if ExpressInfos.ExpressID == ""{
        return shim.Error("ExpressID can not be empty.")
    }
    
    
    ExpressInfos.ExpressProInfo.CoName = args[1]
    ExpressInfos.ExpressProInfo.CoInfo = args[2]
    ExpressInfos.ExpressProInfo.DeliverTime = args[3]
    ExpressInfos.ExpressProInfo.EstimatedDeliveryTime = args[4]
    ExpressInfos.ExpressProInfo.BatchNum = args[5]
    ExpressInfos.ExpressProInfo.Weight = args[6]
    ExpressInfos.ExpressProInfo.Price = args[7]
    ExpressInfos.ExpressProInfo.Deliverer = args[8]
    ExpressInfos.ExpressProInfo.DelivererAdd = args[9]
    ProInfosJSONasBytes,err := json.Marshal(ExpressInfos)
    if err != nil{
        return shim.Error(err.Error())
    }

    err = stub.PutState(ExpressInfos.ExpressID,ProInfosJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }

    return shim.Success(nil)
}

/* func(a *ExpressChainCode) addIngInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
        
    var ExpressInfos ExpressInfo
    var IngInfoitem IngInfo

    if  (len(args)-1)%2 != 0 || len(args) == 1{
        return shim.Error("Incorrect number of arguments")
    }

    ExpressID := args[0]
    for i :=1;i < len(args);{   
        IngInfoitem.IngID = args[i]
        IngInfoitem.IngName = args[i+1]
        ExpressInfos.ExpressIngInfo = append(ExpressInfos.ExpressIngInfo,IngInfoitem)
        i = i+2
    }
    
    
    ExpressInfos.ExpressID = ExpressID
    //ExpressInfos.ExpressIngInfo = expressIngInfo
    IngInfoJsonAsBytes,err := json.Marshal(ExpressInfos)
    if err != nil {
    return shim.Error(err.Error())
    }

    err = stub.PutState(ExpressInfos.ExpressID,IngInfoJsonAsBytes)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
        
} */

func(a *ExpressChainCode) addLogInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
 
    var err error
    var ExpressInfos ExpressInfo

    if len(args)!=11{
        return shim.Error("Incorrect number of arguments.")
    }
    ExpressInfos.ExpressID = args[0]
    if ExpressInfos.ExpressID == ""{
        return shim.Error("ExpressID can not be empty.")
    }
    ExpressInfos.ExpressLogInfo.ArrTime = args[1]
    ExpressInfos.ExpressLogInfo.TranferStationAdd = args[2]
    ExpressInfos.ExpressLogInfo.HandlerInfo = args[3]
    ExpressInfos.ExpressLogInfo.PkgStatus = args[4]
    ExpressInfos.ExpressLogInfo.DepartureTime = args[5]
    ExpressInfos.ExpressLogInfo.NextDestAdd = args[6]
    ExpressInfos.ExpressLogInfo.Mission = args[7]
    ExpressInfos.ExpressLogInfo.VehicleType = args[8]
    ExpressInfos.ExpressLogInfo.VehicleInfo = args[9]
    ExpressInfos.ExpressLogInfo.DriverInfo = args[10]
    
    LogInfosJSONasBytes,err := json.Marshal(ExpressInfos)
    if err != nil{
        return shim.Error(err.Error())
    } 
    err = stub.PutState(ExpressInfos.ExpressID,LogInfosJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
}



func(a *ExpressChainCode) getExpressInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    ExpressID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(ExpressID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var expressAllinfo ExpressAllInfo

    for resultsIterator.HasNext(){
        var ExpressInfos ExpressInfo
        response,err :=resultsIterator.Next()
        if err != nil {
             return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&ExpressInfos)
        if ExpressInfos.ExpressProInfo.CoName !=""{
            expressAllinfo.ExpressProInfo = ExpressInfos.ExpressProInfo
        }else if ExpressInfos.ExpressLogInfo.HandlerInfo !=""{
            expressAllinfo.ExpressLogInfo = append(expressAllinfo.ExpressLogInfo,ExpressInfos.ExpressLogInfo)
        }

    }
    
    jsonsAsBytes,err := json.Marshal(expressAllinfo)
    if err != nil{
        return shim.Error(err.Error())
    }

    return shim.Success(jsonsAsBytes)
}
 

func(a *ExpressChainCode) getProInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
    
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    ExpressID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(ExpressID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var expressProInfo ProInfo

    for resultsIterator.HasNext(){
        var ExpressInfos ExpressInfo
        response,err :=resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&ExpressInfos)
        if ExpressInfos.ExpressProInfo.CoName != ""{
            expressProInfo = ExpressInfos.ExpressProInfo
            continue
        }
    }
    jsonsAsBytes,err := json.Marshal(expressProInfo)   
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

/* func(a *ExpressChainCode) getIngInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
 
    if len(args) !=1{
        return shim.Error("Incorrect number of arguments.")
    }
    ExpressID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(ExpressID)

    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var expressIngInfo []IngInfo
    for resultsIterator.HasNext(){
        var ExpressInfos ExpressInfo
        response,err := resultsIterator.Next()
        if err != nil{
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&ExpressInfos)
        if ExpressInfos.ExpressIngInfo != nil{
            expressIngInfo = ExpressInfos.ExpressIngInfo
            continue
        }
    }
    jsonsAsBytes,err := json.Marshal(expressIngInfo)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
} */

func(a *ExpressChainCode) getLogInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{

    var LogInfos []LogInfo

    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }

    ExpressID := args[0]
    resultsIterator,err :=stub.GetHistoryForKey(ExpressID)
    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

   
    for resultsIterator.HasNext(){
        var ExpressInfos ExpressInfo
        response,err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&ExpressInfos)
        if ExpressInfos.ExpressLogInfo.HandlerInfo != ""{
            LogInfos = append(LogInfos,ExpressInfos.ExpressLogInfo)
        }
    }
    jsonsAsBytes,err := json.Marshal(LogInfos)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

func(a *ExpressChainCode) getLogInfo_l(stub shim.ChaincodeStubInterface,args []string) pb.Response{
    var Loginfo LogInfo

    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }

    ExpressID := args[0]
    resultsIterator,err :=stub.GetHistoryForKey(ExpressID)
    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

   
    for resultsIterator.HasNext(){
        var ExpressInfos ExpressInfo
        response,err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&ExpressInfos)
        if ExpressInfos.ExpressLogInfo.HandlerInfo != ""{
           Loginfo = ExpressInfos.ExpressLogInfo
           continue 
       }
    }
    jsonsAsBytes,err := json.Marshal(Loginfo)
    if err != nil{
        return shim.Error(err.Error ())
    }
    return shim.Success(jsonsAsBytes)
}


func main(){
     err := shim.Start(new(ExpressChainCode))
     if err != nil {
         fmt.Printf("Error starting Express chaincode: %s ",err)
     }
}
