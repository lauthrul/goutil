package model

//import (
//	"fmt"
//	"fund/common"
//	"github.com/lauthrul/goutil/log"
//	"github.com/valyala/fastjson"
//	"strings"
//	"time"
//)
//
//func FundListUpdateCheck(cache *FundCache) (bool, error) {
//	log.Info("check fund list update ...")
//
//	var fc FundCache
//
//	url := fmt.Sprintf(fundListUrl, time.Now().Format("20060102150405"))
//
//	// HEAD
//	resp, err := common.Client.Head(url)
//	if err != nil {
//		log.Error(err)
//		return false, err
//	}
//
//	length := resp.Header.ContentLength()
//	if length > 0 {
//		if length == cache.Length {
//			log.Info("fund list no update")
//			return false, nil
//		}
//		fc.Length = length
//	}
//
//	// GET
//	resp, err = common.Client.Get(url)
//	if err != nil {
//		log.Error(err)
//		return false, err
//	}
//
//	data := string(resp.Body())
//	idx := strings.Index(data, "[")
//	if idx > 0 {
//		data = data[idx:]
//	}
//	data = strings.TrimRight(data, ";")
//
//	var parser fastjson.Parser
//	root, err := parser.Parse(data)
//	if err != nil {
//		log.Error(err)
//		return false, err
//	}
//	objs, err := root.Array()
//	if err != nil {
//		log.Error("invalid data")
//		return false, err
//	}
//	for _, obj := range objs {
//		var fund FundBasicInfo
//		if v, _ := obj.Array(); v != nil {
//			for i, vv := range v {
//				// "000001","HXCZHH","华夏成长混合","混合型","HUAXIACHENGZHANGHUNHE"
//				s := strings.Trim(vv.String(), `"`)
//				switch i {
//				case 0:
//					fund.Code = s
//				case 1:
//					fund.ShortName = s
//				case 2:
//					fund.Name = s
//				case 3:
//					fund.Type = s
//				case 4:
//					fund.FullName = s
//				}
//			}
//			fc.List = append(fc.List, fund)
//		}
//	}
//	*cache = fc
//
//	return true, nil
//}
